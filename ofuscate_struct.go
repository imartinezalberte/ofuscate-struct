package ofuscatestruct

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
)

var (
	arrRegexp      = regexp.MustCompile("([a-zA-Z0-9]+)\\[([0-9]*)\\]")
	indexArrRegexp = regexp.MustCompile("\\[([0-9]{0,9})\\](.*)")
)

func OfuscateArr(input any, filter string) any {
	if input == nil {
		return nil
	}

	v := reflect.ValueOf(input)
	switch reflect.TypeOf(input).Kind() {
	case reflect.Array, reflect.Slice:
		// Create the array
		tmp := make([]any, v.Len(), v.Len())

		// Check if we have to something more about that
		b, a, _ := strings.Cut(filter, ".")
		if !strings.ContainsAny(b, "[]") && !indexArrRegexp.MatchString(b) {
			return input
		}

		// Just ofuscate the following index
		options := indexArrRegexp.FindStringSubmatch(b)
		inside, next := options[1], join(".", options[2], a)
		if index, err := strconv.Atoi(inside); err == nil {
			for i := 0; i < len(tmp); i++ {
				tmp[i] = v.Index(i).Interface()
			}

			if index < len(tmp) {
				if et := reflect.TypeOf(input).Elem().Kind(); (et == reflect.Struct || et == reflect.Interface ||
					et == reflect.Map) && a != "" {
					tmp[index] = Ofuscate(tmp[index], a)
				} else if (et == reflect.Array || et == reflect.Slice) && next != "" {
					tmp[index] = OfuscateArr(tmp[index], next)
				} else {
					tmp[index] = "XXX"
				}
			}
		} else { // Ofuscate the array
			for i := 0; i < len(tmp); i++ {
				if et := reflect.TypeOf(input).Elem().Kind(); et == reflect.Struct || et == reflect.Interface ||
					et == reflect.Map && a != "" {
					tmp[i] = Ofuscate(tmp[i], a)
				} else if (et == reflect.Array || et == reflect.Slice) && next != "" {
					tmp[i] = OfuscateArr(tmp[i], next)
				} else {
					tmp[i] = "XXX"
				}
			}
		}

		return tmp
	case reflect.Struct:
		return Ofuscate(input, filter)
	default:
		return input
	}
}

func Ofuscate(input any, filter string) (m map[string]interface{}) {
	if k := reflect.TypeOf(input).Kind(); k == reflect.Array || k == reflect.Slice {
		return
	}

	mapstructure.Decode(input, &m)
	DoOfuscate(m, filter)
	return m
}

// TODO: aggregate slicing capability, could be awesome
func DoOfuscate(input map[string]any, filter string) {
	if strings.TrimSpace(filter) == "" {
		return
	}

	actual, after, ok := strings.Cut(filter, ".")
	if !ok { // last one
		propertyName, possibleIndex, isArr := strings.Cut(actual, "[")
		if !isArr || possibleIndex == "]" {
			input[propertyName] = "XXX"
			return
		}

		if index, err := strconv.Atoi(strings.Split(possibleIndex, "]")[0]); err == nil {
			arr, err := processPossibleArrWithIndex(input[propertyName], index)
			if err != nil {
				return
			}
			input[propertyName] = arr.Interface()
		}

		return
	}

	if ofuscateArr(input, actual, after) {
		return
	}

	if attrExists(input, actual) {
		DoOfuscate(input[actual].(map[string]any), after)
	}
}

func ofuscateArr(m map[string]any, actual, after string) bool {
	if !arrRegexp.MatchString(actual) {
		return false
	}

	options := arrRegexp.FindStringSubmatch(actual)

	if index, err := strconv.Atoi(options[2]); err == nil { // index specified
		obj, ok := m[options[1]]
		if !ok {
			return true
		}

		// We are safe, it is an array
		if k := reflect.TypeOf(obj).Kind(); k == reflect.Array || k == reflect.Slice {
			arr := reflect.ValueOf(obj)

			if arr.Len() == 0 || int(index) > arr.Len() || int(index) < 0 {
				return true
			}

			if t := arr.Index(index).Type().Name(); t != "interface" {
				arr = reflect.ValueOf(copySlice(arr))
			}

			var mm map[string]interface{}
			mapstructure.Decode(arr.Index(index).Interface(), &mm)

			arr.Index(index).Set(reflect.ValueOf(mm)) // We set the new map to the actual index

			m[options[1]] = arr.Interface() // We recover the array
			DoOfuscate(mm, after)
		}
	} else if len(options[1]) > 0 { // No index specified
		obj, ok := m[options[1]]
		if !ok {
			return true
		}

		if k := reflect.TypeOf(obj).Kind(); k == reflect.Array || k == reflect.Slice {
			arr := reflect.ValueOf(obj)

			if arr.Len() == 0 {
				return true
			}

			if t := arr.Index(index).Type().Name(); t != "interface" {
				tmp := reflect.MakeSlice(reflect.TypeOf([]any{}), arr.Len(), arr.Len())

				for i := 0; i < tmp.Len(); i++ {
					var mm map[string]interface{}
					mapstructure.Decode(arr.Index(i).Interface(), &mm)

					tmp.Index(i).Set(reflect.ValueOf(mm))
				}

				arr = tmp
			}

			m[options[1]] = arr.Interface()

			for i := 0; i < arr.Len(); i++ {
				DoOfuscate(m[options[1]].([]any)[i].(map[string]interface{}), after)
			}
		}
	}

	return false
}

func parseNumber(input string) (int, bool) {
	n, err := strconv.Atoi(input)
	return n, err == nil
}

func processPossibleArrWithIndex(input any, index int) (reflect.Value, error) {
	if input == nil {
		return reflect.ValueOf(input), errors.New("input is nil")
	}

	arr := reflect.ValueOf(input)
	if k := reflect.TypeOf(input).Kind(); k == reflect.Slice || k == reflect.Array {
		// Limit cases
		if arr.Len() == 0 || index < 0 || index >= arr.Len() {
			return arr, errors.New("index out of bounds")
		}

		// If the arr is not an array of strings or interfaces, then we have to create one
		if t := arr.Index(0).Type().Name(); t != "string" && t != "interface" {
			arr = reflect.ValueOf(copySlice(arr))
		}

		arr.Index(index).Set(reflect.ValueOf("XXX"))

		return arr, nil
	}

	return arr, errors.New("it is not an array")
}

func attrExists(m map[string]any, attr string) bool {
	_, ok := m[attr]
	return ok
}

func copySlice(input reflect.Value) []any {
	if k := input.Type().Kind(); k != reflect.Array && k != reflect.Slice {
		return []any{}
	}

	tmp := make([]any, input.Len(), input.Len())

	for i := 0; i < input.Len(); i++ {
		tmp[i] = input.Index(i).Interface()
	}

	return tmp
}

func join(sep string, input ...string) string {
	var r string
	for _, s := range input {
		if s = strings.TrimSpace(s); s != "" {
			r += s + sep
		}
	}
	return strings.TrimSuffix(r, sep)
}
