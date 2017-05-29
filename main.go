package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 2 {
		fmt.Println("Usage: ", path.Base(args[0]), " yaml-path command")
		return
	}

	// Load yaml data
	target := args[1]
	// yamlPath := fmt.Sprintf("./%s.yaml", target)
	yamlPath := target
	if _, err := os.Stat(yamlPath); err != nil {
		fmt.Println("YamlFile not found:", yamlPath)
		return
	}

	// Select command
	command := args[2]
	var f func(Table)
	switch command {
	case "create":
		f = OutputCreateTableQuery
	case "drop":
		f = OutputDropTableQuery
	case "example":
		f = OutputExampletQuery
	default:
		fmt.Println("Dose not exists command:", command)
		return
	}

	// Execute command
	context := makeTableContext(yamlPath)
	f(context)
}

func getYamlData(yamlPath string) map[interface{}]interface{} {
	buf, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	return m
}

func makeTableContext(yamlPath string) Table {
	yamlMap := getYamlData(yamlPath)
	schema := yamlMap["table-schema"].(map[interface{}]interface{})
	name, _ := schema["name"].(string)
	tblName := strings.ToLower(name)
	tblMeta := schema["meta"].(string)
	exampleNumber := schema["example-number"].(int)
	columns := schema["columns"].([]interface{})
	colCount := len(columns)

	// Create TableColumn list
	tblColumns := make([]TableColumn, colCount)
	for colIdx, colData := range columns {
		c, _ := colData.(map[interface{}]interface{})

		exText := ""
		exMin := 0
		exMax := 0
		exStrRange := make([]string, 0)

		switch c["example-type"] {
		case intType:
			exText = c["example-text"].(string)
		case intIncType:
		case intRangeType:
			rangeSchema := c["example-range"].([]interface{})
			minTag := rangeSchema[0].(map[interface{}]interface{})
			maxTag := rangeSchema[1].(map[interface{}]interface{})
			exMin = minTag["min"].(int)
			exMax = maxTag["max"].(int)
		case stringType, stringIncType:
			exText = c["example-text"].(string)
		case stringRangeType:
			rangeSchema := c["example-range"].([]interface{})
			strList := make([]string, len(rangeSchema))
			for idx, rData := range rangeSchema {
				r, _ := rData.(map[interface{}]interface{})
				strList[idx] = r["value"].(string)
			}
			exStrRange = strList
		}

		exTblCol := ExampleTableColumn{
			c["example-type"].(string),
			exText,
			exMin,
			exMax,
			exStrRange,
		}

		tblCol := TableColumn{
			c["name"].(string),
			c["type"].(string),
			exTblCol,
		}
		tblColumns[colIdx] = tblCol
	}

	context := Table{
		tblName,
		tblMeta,
		exampleNumber,
		tblColumns,
	}
	return context
}
