package testingHelpers

import (
	"jasonMcDiff/types"
	"encoding/json"
	"os"
	"fmt"
	randData "github.com/Pallinder/go-randomdata"
)

func GenerateRandomObj(maxWidth int, maxDepth int) *types.UnknownStringMapStructure {
	d := map[string]interface{}{}
	for i := 0; i < maxWidth; i++ {
		key := randData.Letters(randData.Number(1, 20))
		val := generateRandomDataType(maxWidth - 1, maxDepth - 1)
		d[key] = val
	}
	return &d
}

func generateRandomArr(maxWidth int, maxDepth int) []interface{} {
	d := make([]interface{}, maxWidth)
	flag := randData.Number(0, 5)
	for i := 0; i < maxWidth; i++ {
		switch flag {
		case 0:
			d[i] = randData.Number(1,10e6)
		case 1:
			d[i] = randData.Boolean()
		case 2:
			d[i] = GenerateRandomObj(maxWidth - 1, maxDepth - 1)
		case 3:
			d[i] = generateRandomArr(maxWidth - 1, maxDepth - 1)
		default:
			d[i] = randData.Letters(randData.Number(0, 15))
		}
	}
	return d
}

func generateRandomDataType(maxWidth, maxDepth int) interface{} {
	if maxWidth < 1 || maxDepth < 1 {
		return "finish"
	}
	flag := randData.Number(0, 5)
	switch flag {
	case 0:
		return randData.Number(1,10e6)
	case 1:
		return randData.Boolean()
	case 2:
		return GenerateRandomObj(maxWidth - 1, maxDepth - 1)
	case 3:
		return generateRandomArr(maxWidth - 1, maxDepth - 1)
	default:
		return randData.Letters(randData.Number(0, 15))
	}
}

func DumpObjToFileAsJson(data *types.UnknownStringMapStructure, name string) {
	jsonData, err := json.MarshalIndent(data, "  ", "  ")

	jsonFile, err := os.Create(fmt.Sprintf("./%v.json", name))
	defer jsonFile.Close()
	if err != nil {
			panic(err)
	}

	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}