package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	TB_CONFIG_FILE = "conf/tables.json"
)

type Column struct {
	Name  string
	Ctype string
	Size  int
}

type Table struct {
	Name       string
	Columnlist []Column
}

type Tables struct {
	Tablelist []Table
}

func (tbs *Tables) GetTableList() []Table {
	return tbs.Tablelist
}

func (tbs *Tables) GetTable(tbname string) *Table {
	for _, tb := range tbs.Tablelist {
		if tb.Name == tbname {
			return &tb
		}
	}
	return nil
}

var Schema = new(Tables)

func init() {
	tbfile, err := os.Open(TB_CONFIG_FILE)

	if err != nil {
		if len(TB_CONFIG_FILE) > 1 {
			fmt.Printf("Error: Could not read tables file: %s\n %s\n", TB_CONFIG_FILE, err)
		}
	}

	tbdecoder := json.NewDecoder(tbfile)
	err = tbdecoder.Decode(&Schema)
	if err != nil {
		fmt.Printf("Error decoding: %s \n %s \n", tbfile, err)
	}
}
