package main

import (
	"fmt"
	"math/rand"
	"time"
)

func OutputHelp() {
	helpMsg := `Usage:

        sqlgen command [yaml-path]


The commands are:

        create      output create table query
        example     output example insert query
        drop        output drop table query
        schema      output sample table schema yaml file`

	fmt.Println(helpMsg)
}

func OutputCreateTableQuery(t Table) {
	totalColCount := len(t.TableColumns) + len(t.MetaList)

	fmt.Println(t.CreateQueryHeader())

	offset := 1
	for _, v := range t.TableColumns {
		withcomma := true
		if offset == totalColCount {
			withcomma = false
		}
		fmt.Println(v.CreateQueryColumn(withcomma))
		offset++
	}
	for _, v := range t.MetaList {
		if offset == totalColCount {
			fmt.Println(fmt.Sprintf("  %s", v))
		} else {
			fmt.Println(fmt.Sprintf("  %s,", v))
		}
		offset++
	}

	fmt.Println(t.CreateQueryFooter())
}

func OutputDropTableQuery(t Table) {
	fmt.Println(t.DropQery())
}

func OutputExampletQuery(t Table) {
	rand.Seed(time.Now().UnixNano())
	colCount := len(t.TableColumns) - 1

	colLine := ""
	for i, v := range t.TableColumns {
		colLine += v.Name
		if i != colCount {
			colLine += ", "
		}
	}

	start := 1
	end := t.ExampleNumber
	if exampleLimit < t.ExampleNumber {
		end = exampleLimit
	}
	for {
		if start >= t.ExampleNumber {
			break
		}

		query := ""
		for i := start; i <= end; i++ {
			columnLine := "("
			for idx, v := range t.TableColumns {
				columnLine += v.Example.GetValue(i)
				if idx != colCount {
					columnLine += ", "
				}
			}

			if i == end {
				query += fmt.Sprintf("%s)", columnLine)
			} else {
				query += fmt.Sprintf("%s),\n", columnLine)
			}
		}

		fmt.Println(t.InsertQery(colLine, query))
		start += exampleLimit
		end += exampleLimit
		if t.ExampleNumber < end {
			end = t.ExampleNumber
		}
	}
}

func OutputSampleTableSchema() {
	fmt.Println(sampleSchema)
}
