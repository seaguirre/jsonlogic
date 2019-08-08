package jsonlogic

import (
	"reflect"
)

func less(a, b interface{}) bool {
	if isNumber(a) {
		return toNumber(b) > toNumber(a)
	}

	return toString(b) > toString(a)
}

func hardEquals(a, b interface{}) bool {
	ra := reflect.ValueOf(a).Kind()
	rb := reflect.ValueOf(b).Kind()

	if ra != rb {
		return false
	}

	return equals(a, b)
}

func equals(a, b interface{}) bool {

	if isBool(a) {
		return toBool(a) == toBool(b)
	}

	if isNumber(a) {
		return toNumber(a) == toNumber(b)
	}

	return toString(a) == toString(b)
}

func safeEquals(a, b interface{}) bool {

	if pb, err := safeToBool(b); err == nil {
		if pa, err := safeToBool(a); err == nil {
			return pa == pb
		} else {
			return false
		}
	}

	if pb, err := safeToNumber(b); err == nil {
		if pa, err := safeToNumber(a); err == nil {
			return pa == pb
		} else {
			return false
		}
	}

	if pb, err := safeToString(b); err == nil {
		if pa, err := safeToString(a); err == nil {
			return pa == pb
		} else {
			return false
		}
	}

	if a == nil && b == nil {
		return true
	} else {
		return false
	}

}