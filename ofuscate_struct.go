package ofuscatestruct

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

var arrRegexp = regexp.MustCompile("([a-zA-Z0-9]+)\\[([0-9]*)\\]")

func Ofuscate(input any, filter string) map[string]any {
	m, err := anyToMap(input)
	if err != nil {
		panic(err)
	}
	DoOfuscate(m, filter)
	return m
}

func DoOfuscate(input map[string]any, filter string) {
	actual, after, ok := strings.Cut(filter, ".")
	if !ok { // last one
		propertyName := strings.Split(actual, "[")[0]
		if _, ok := input[propertyName]; ok {
			input[propertyName] = "XXX"
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

func anyToMap(input any) (map[string]any, error) {
	output := make(map[string]any)
	buf, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func attrExists(m map[string]any, attr string) bool {
	_, ok := m[attr]
	return ok
}

func ofuscateArr(m map[string]any, actual, after string) bool {
	if !arrRegexp.MatchString(actual) {
		return false
	}

	options := arrRegexp.FindStringSubmatch(actual)

	if index, err := strconv.Atoi(options[2]); err == nil { // index specified
		arr := m[options[1]].([]any)
		if len(arr) <= int(index) {
			return true
		}
		DoOfuscate(arr[index].(map[string]interface{}), after)
	} else if len(options[1]) > 0 { // No index specified
		arr := m[options[1]].([]any)
		for i := range arr {
			DoOfuscate(arr[i].(map[string]interface{}), after)
		}
	}

	return false
}
