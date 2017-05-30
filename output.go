package main

import (
	"fmt"
	"math/rand"
	"time"
)

func OutputCreateTableQuery(t Table) {
	colCount := len(t.TableColumns) - 1

	fmt.Println(t.CreateQueryHeader())
	for i, v := range t.TableColumns {
		withcomma := true
		if i == colCount {
			withcomma = false
		}
		fmt.Println(v.CreateQueryColumn(withcomma))
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
