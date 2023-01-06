package csvs

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVHandle struct {
	records [][]string
}

func NewCSVHandle() *CSVHandle {
	obj := CSVHandle{}
	return &obj
}

func (obj *CSVHandle) OpenCSV(strFile string) error {
	fileCSV, err := os.Open(strFile)
	if nil != err {
		fmt.Println(err)
		return err
	}
	defer fileCSV.Close()

	csvReader := csv.NewReader(fileCSV)

	obj.records, err = csvReader.ReadAll()
	if nil != err {
		fmt.Println(err)
		return err
	}

	fmt.Printf("> DBG\t%v is opened: %v rows %v columns\n\n", strFile, len(obj.records), len(obj.records[0]))
	return nil
}

func (obj *CSVHandle) RowCount() int {
	return len(obj.records)
}

func (obj *CSVHandle) CoulmnCount() int {
	return len(obj.records[0])
}

func (obj *CSVHandle) GetField(x, y int) string {
	return obj.records[x][y]
}
