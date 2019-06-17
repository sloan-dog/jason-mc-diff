package diff

import (
	"fmt"
	"jasonMcDiff/types"

	sets "github.com/deckarep/golang-set"
)

func DiffStructures(left types.UnknownStringMapStructure, right types.UnknownStringMapStructure) (bool, error) {
	return mapsAreSame(left, right)
}

func getKeysFromLeftAndRight(left map[string]interface{}, right map[string]interface{}) sets.Set {
	keys := sets.NewSet()

	for key := range left {
		keys.Add(key)
	}
	for key := range right {
		keys.Add(key)
	}
	return keys
}

func mapsAreSame(left interface{}, right interface{}) (bool, error) {
	leftMap, leftOk := left.(map[string]interface{})
	rightMap, rightOk := right.(map[string]interface{})
	keys := getKeysFromLeftAndRight(leftMap, rightMap)
	if !leftOk || !rightOk {
		return false, nil
	}
	same := true
	var err error
	for _, keyI := range keys.ToSlice() {
		key := keyI.(string)
		valLeft := leftMap[key]
		valRight := rightMap[key]
		// fmt.Printf("key: %v\nval type: %v\nval: %v\n", key, reflect.TypeOf(val), val)
		switch valLeft.(type) {
		case string:
			same, err = diffString(valLeft, valRight)
		case int:
			same, err = diffInteger(valLeft, valRight)
		case map[string]interface{}:
			// recursion
			same, err = mapsAreSame(valLeft, valRight)
		case []interface{}:
			same, err = diffArray(valLeft, valRight)
		case float64:
			same, err = diffFloat64(valLeft, valRight)
		case float32:
			same, err = diffFloat32(valLeft, valRight)
		case bool:
			same, err = diffBool(valLeft, valRight)
		default:
			return false, fmt.Errorf("invalid type")
		}
		if same != true && err == nil {
			return false, nil
		} else if same != true {
			return false, err
		}
	}
	return true, nil
}

func diffArray(left interface{}, right interface{}) (bool, error) {
	leftVal, leftOk := left.([]interface{})
	rightVal, rightOk := right.([]interface{})
	if !leftOk || !rightOk {
		return false, nil
	}
	if len(leftVal) != len(rightVal) {
		return false, nil
	}
	same := true
	var err error
	for i := range leftVal {
		leftEl := leftVal[i]
		rightEl := rightVal[i]
		switch leftEl.(type) {
		case string:
			same, err = diffString(leftEl, rightEl)
		case int:
			same, err = diffInteger(leftEl, rightEl)
		case map[string]interface{}:
			same, err = mapsAreSame(leftEl, rightEl)
		case []interface{}:
			same, err = diffArray(leftEl, rightEl)
		case float64:
			same, err = diffFloat64(leftEl, rightEl)
		case float32:
			same, err = diffFloat32(leftEl, rightEl)
		case bool:
			same, err = diffBool(leftEl, rightEl)
		default:
			return false, fmt.Errorf("invalid type")
		}
		if same != true && err == nil {
			return false, nil
		} else if same != true {
			return false, err
		}
	}
	// TODO - actually diff this you lazy man
	return true, nil
}

func diffInteger(left interface{}, right interface{}) (bool, error) {
	leftVal, leftOk := left.(int)
	rightVal, rightOk := right.(int)
	if !leftOk || !rightOk {
		return false, fmt.Errorf("invalid int")
	}
	return leftVal == rightVal, nil
}

func diffFloat64(left interface{}, right interface{}) (bool, error) {
	leftVal, leftOk := left.(float64)
	rightVal, rightOk := right.(float64)
	if !leftOk || !rightOk {
		return false, fmt.Errorf("invalid float64")
	}
	return leftVal == rightVal, nil
}

func diffFloat32(left interface{}, right interface{}) (bool, error) {
	leftVal, leftOk := left.(float32)
	rightVal, rightOk := right.(float32)
	if !leftOk || !rightOk {
		return false, fmt.Errorf("invalid float32")
	}
	return leftVal == rightVal, nil
}

func diffString(left interface{}, right interface{}) (bool, error) {
	leftVal, leftOk := left.(string)
	rightVal, rightOk := right.(string)
	if !leftOk || !rightOk {
		return false, fmt.Errorf("invalid string")
	}
	return leftVal == rightVal, nil
}

func diffBool(left interface{}, right interface{}) (bool, error) {
	leftVal, leftOk := left.(bool)
	rightVal, rightOk := right.(bool)
	if !leftOk || !rightOk {
		return false, fmt.Errorf("invalid bool")
	}
	return leftVal == rightVal, nil
}
