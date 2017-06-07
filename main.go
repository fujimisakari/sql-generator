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

	if len(args) <= 1 {
		fmt.Println("Usage: ", path.Base(args[0]), " yaml-path command")
		OutputHelp()
		return
	}

	// Select command
	command := args[1]
	var f func(Table)
	switch command {
	case "create":
		f = OutputCreateTableQuery
	case "drop":
		f = OutputDropTableQuery
	case "example":
		f = OutputExampletQuery
	case "help":
		OutputHelp()
		return
	case "sample":
		OutputSampleTableSchema()
		return
	default:
		fmt.Println("Dose not exists command:", command)
		OutputHelp()
		return
	}

	if len(args) < 2 {
		fmt.Println("Dose not exists yaml-path")
		return
	}

	// Load yaml data
	yamlPath := args[2]
	if _, err := os.Stat(yamlPath); err != nil {
		fmt.Println("YamlFile not found:", yamlPath)
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
	tblParams := schema["params"].(string)
	exNumber := schema["ex-number"].(int)

	// Create TableColumn List
	columns := schema["columns"].([]interface{})
	colCount := len(columns)
	tblColumns := make([]TableColumn, colCount)
	for colIdx, colData := range columns {
		c, _ := colData.(map[interface{}]interface{})

		exText := ""
		exMin := 0
		exMax := 0
		exStartDate := ""
		exEndDate := ""
		exStrRange := make([]string, 0)

		switch c["ex-type"] {
		case intType:
			exText = c["ex-text"].(string)
		case intIncType:
		case intRangeType:
			rangeSchema := c["ex-range"].([]interface{})
			exMin = rangeSchema[0].(map[interface{}]interface{})["min"].(int)
			exMax = rangeSchema[1].(map[interface{}]interface{})["max"].(int)
		case stringType, stringIncType:
			exText = c["ex-text"].(string)
		case stringRangeType:
			rangeSchema := c["ex-range"].([]interface{})
			strList := make([]string, len(rangeSchema))
			for idx, rData := range rangeSchema {
				r, _ := rData.(map[interface{}]interface{})
				strList[idx] = r["value"].(string)
			}
			exStrRange = strList
		case dateRangeType:
			rangeSchema := c["ex-range"].([]interface{})
			exStartDate = rangeSchema[0].(map[interface{}]interface{})["start"].(string)
			exEndDate = rangeSchema[1].(map[interface{}]interface{})["end"].(string)
		}

		exTblCol := ExampleTableColumn{
			c["ex-type"].(string),
			exText,
			exMin,
			exMax,
			exStartDate,
			exEndDate,
			exStrRange,
		}

		tblCol := TableColumn{
			c["name"].(string),
			c["type"].(string),
			exTblCol,
		}
		tblColumns[colIdx] = tblCol
	}

	// Create Meta List
	tblMetaList := make([]string, 0)
	if metaList, ok := schema["meta-list"]; ok {
		metaList := metaList.([]interface{})
		metaListCount := len(metaList)
		tblMetaList = make([]string, metaListCount)
		for metaIdx, metaData := range metaList {
			m, _ := metaData.(map[interface{}]interface{})
			tblMetaList[metaIdx] = m["value"].(string)
		}
	}

	context := Table{
		tblName,
		tblParams,
		exNumber,
		tblColumns,
		tblMetaList,
	}
	return context
}
