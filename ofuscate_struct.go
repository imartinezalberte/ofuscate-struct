package ofuscatestruct

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
)

var arrRegexp = regexp.MustCompile("([a-zA-Z0-9]+)\\[([0-9]*)\\]")

func Ofuscate(input any, filter string) (m map[string]any) {
	mapstructure.Decode(input, &m)
	DoOfuscate(m, filter)
	return m
}

// TODO: aggregate slicing capability, could be awesome
func DoOfuscate(input map[string]any, filter string) {
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

			m[options[1]] = arr.Interface()           // We recover the array
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
