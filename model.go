package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Table struct {
	Name          string
	Params        string
	ExampleNumber int
	TableColumns  []TableColumn
	MetaList      []string
}

func (t Table) CreateQueryHeader() string {
	return fmt.Sprintf("CREATE TABLE `%s` (", t.Name)
}

func (t Table) CreateQueryFooter() string {
	return fmt.Sprintf(") %s;", t.Params)
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
	StartDate   string
	EndDate     string
	StringRange []string
}

func (t ExampleTableColumn) GetValue(i int) string {
	var s string
	switch t.Type {
	case intType:
		s = fmt.Sprintf("%d", t.Text)
	case intIncType:
		s = fmt.Sprintf("%d", i)
	case intRangeType:
		s = fmt.Sprintf("%d", random(t.Min, t.Max))
	case stringType:
		s = fmt.Sprintf("'%s'", t.Text)
	case stringIncType:
		s = fmt.Sprintf("'%d-%s'", i, t.Text)
	case stringRangeType:
		idx := rand.Intn(len(t.StringRange))
		s = fmt.Sprintf("'%s'", t.StringRange[idx])
	case dateRangeType:
		layout := "2006-01-02 15:04:05"
		startDate, _ := time.Parse(layout, t.StartDate)
		endDate, _ := time.Parse(layout, t.EndDate)
		unixtime := randomForDate(startDate.Unix(), endDate.Unix())
		s = fmt.Sprintf("'%s'", time.Unix(unixtime, 0).Format(layout))
	default:
		fmt.Println("Dose not exists type:", t.Type)
	}
	return s
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func randomForDate(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}
