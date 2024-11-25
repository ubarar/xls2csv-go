package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ubarar/xls2csv/xls2csv"
)

func main() {
    var err error
    home, err := os.UserHomeDir()
    f := filepath.Join(home, "xls2csv-go/xls2csv/files/my.xls")
    sheetId := 0
    records := [][]string{}

    // Call XLS2CSV() to convert XLS and get all records.
    if records, err = xls2csv.XLS2CSV(f, sheetId); err != nil {
        fmt.Printf("err: %v\n", err)
        return
    }

    fmt.Printf("records: %v\n", records)
}
