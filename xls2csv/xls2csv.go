package xls2csv

/*
#cgo CFLAGS: -I/usr/local/libxls/include
#cgo LDFLAGS: -L/usr/local/libxls/lib -lxlsreader
#include <stdio.h>
#include <stdlib.h>
#include <libxls/xls.h>
#include "xls2csv.h"
*/
import "C"

import (
	"encoding/csv"
	"fmt"
	"strings"
	"unsafe"
)

func XLS2CSV(xlsFile string, sheetId int) (records [][]string, err error) {
	var buf *C.char = nil
	var r *csv.Reader = nil

	f := C.CString(xlsFile)
	// C string should be free after use.
	defer C.free(unsafe.Pointer(f))

	id := C.int(sheetId)

	// xls2csv() will return a buffer(char *) contains CSV string.
	// The buffer should be free in C.
	buf = C.xls2csv(f, id)
	if buf == nil {
		err = fmt.Errorf("xls2csv() error.")
		goto end
	}

	// Free memory block after use.
	defer C.free(unsafe.Pointer(buf))

	r = csv.NewReader(strings.NewReader(C.GoString(buf)))
	records, err = r.ReadAll()

end:
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
