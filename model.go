package main

import (
	"fmt"
	"math/rand"
)

type Table struct {
	Name          string
	Meta          string
	ExampleNumber int
	TableColumns  []TableColumn
}

func (t Table) CreateQueryHeader() string {
	return fmt.Sprintf("CREATE TABLE `%s` (", t.Name)
}

func (t Table) CreateQueryFooter() string {
	return fmt.Sprintf(") %s;", t.Meta)
}

func (t Table) DropQery() string {
	return fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", t.Name)
}

func (t Table) InsertQery(colLine string, query string) string {
	return fmt.Sprintf("INSERT INTO `%s` (%s) VALUES\n%s;\n\n", t.Name, colLine, query)
}

type TableColumn struct {
	Name    string
	Type    string
	Example ExampleTableColumn
}

func (t TableColumn) CreateQueryColumn(withcomma bool) string {
	format := "  `%s` %s"
	if withcomma {
		format = "  `%s` %s,"
	}
	return fmt.Sprintf(format, t.Name, t.Type)
}

type ExampleTableColumn struct {
	Type        string
	Text        string
	Min         int
	Max         int
	StringRange []string
}

func (t ExampleTableColumn) GetExample(i int) string {
	var s string
	switch t.Type {
	case intType:
		s = fmt.Sprintf("%d", t.Text)
	case intIncType:
		s = fmt.Sprintf("%d", i)
	case intRangeType:
		rand := rand.Intn(t.Max-t.Min) + t.Min
		s = fmt.Sprintf("%d", rand)
	case stringType:
		s = fmt.Sprintf("'%s'", t.Text)
	case stringIncType:
		s = fmt.Sprintf("'%d-%s'", i, t.Text)
	case stringRangeType:
		idx := rand.Intn(len(t.StringRange))
		s = fmt.Sprintf("'%s'", t.StringRange[idx])
	default:
		fmt.Println("Dose not exists type:", t.Type)
	}
	return s
}
