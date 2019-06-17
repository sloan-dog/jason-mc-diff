package main

import (
	helpers "jasonMcDiff/testing-helpers"
)

func main() {
	obj := helpers.GenerateRandomObj(10, 20)
	helpers.DumpObjToFileAsJson(obj, "bigObj")
}
