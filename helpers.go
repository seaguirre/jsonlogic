package jsonlogic

import (
	"errors"
	"reflect"
	"strconv"
)

func is(obj interface{}, kind reflect.Kind) bool {
	return obj != nil && reflect.TypeOf(obj).Kind() == kind
}

func isBool(obj interface{}) bool {
	return is(obj, reflect.Bool)
}

func isString(obj interface{}) bool {
	return is(obj, reflect.String)
}

func isNumber(obj interface{}) bool {
	return is(obj, reflect.Float64)
}

func isPrimitive(obj interface{}) bool {
	return isBool(obj) || isString(obj) || isNumber(obj)
}

func isMap(obj interface{}) bool {
	return is(obj, reflect.Map)
}

func isSlice(obj interface{}) bool {
	return is(obj, reflect.Slice)
}

func isTrue(obj interface{}) bool {
	if isBool(obj) {
		return obj.(bool)
	}

	if isNumber(obj) {
		n := toNumber(obj)
		return n != 0
	}

	if isString(obj) || isSlice(obj) || isMap(obj) {
		length := reflect.ValueOf(obj).Len()
		return length > 0
	}

	return false
}

func toBool(value interface{}) bool {
	if isString(value) {
		w, _ := strconv.ParseBool(value.(string))
		return w
	}

	return value.(bool)
}

func safeToBool(value interface{}) (parsed bool, err error) {
	if isBool(value) {
		return value.(bool), err
	}

	if isString(value) {
		w, err := strconv.ParseBool(value.(string))
		return w, err
	}

	err = errors.New("value format is not correct")
	return false, err
}

func toNumber(value interface{}) float64 {
	if isString(value) {
		w, _ := strconv.ParseFloat(value.(string), 64)

		return w
	}

	return value.(float64)
}

func safeToNumber(value interface{}) (parsed float64, err error) {
	if isNumber(value) {
		return value.(float64), err
	}

	if isString(value) {
		w, err := strconv.ParseFloat(value.(string), 64)
		return w, err
	}

	err = errors.New("value format is not correct")
	return 0, err
}

func toString(value interface{}) string {
	if isNumber(value) {
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	}

	if value == nil {
		return ""
	}

	return value.(string)
}

func safeToString(value interface{}) (parsed string, err error) {
	if isNumber(value) {
		return strconv.FormatFloat(value.(float64), 'f', -1, 64), err
	}

	if value == nil {
		err = errors.New("value is nil")
		return "", err
	} else {
		return value.(string), err
	}
}