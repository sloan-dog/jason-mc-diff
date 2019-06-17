package parse

import (
	"encoding/json"
	"fmt"
	"jasonMcDiff/types"
	// "os"
)

func ParseString(source string) (types.UnknownStringMapStructure, error) {
	bytes, err := marshalJsonFromString(source)
	if err != nil {
		return nil, err
	}
	structure, err := parseJsonToNative(bytes)
	if err != nil {
		return nil, err
	}
	return structure, nil
}

func marshalJsonFromString(source string) ([]byte, error) {
	rawIn := json.RawMessage(source)
	jsonBytes, err := rawIn.MarshalJSON()
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonBytes))
	return jsonBytes, nil
}

func parseJsonToNative(source []byte) (types.UnknownStringMapStructure, error) {
	// we have no idea
	result := types.UnknownStringMapStructure{}
	err := json.Unmarshal(source, &result)
	fmt.Println(result, "poops")
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("fuck", result["foo"])
	return result, nil
}
