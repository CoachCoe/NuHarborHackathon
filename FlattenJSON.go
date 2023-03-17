package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.Open("/home/scoe/polyglot/src/defender.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan() {
		byteValue := scanner.Text()
		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)
		parseMap(result)
	}
}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}

func writeLine(line string) {
	lines := line
	f, _ := os.OpenFile("/home/scoe/polyglot/src/MyActionTypes.txt", os.O_RDWR, 0644)
	defer f.Close()
	dw := bufio.NewWriter(f)

	dw.WriteString(lines + "\n")

	dw.Flush()
}
