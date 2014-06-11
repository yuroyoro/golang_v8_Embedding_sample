package main

// #cgo LDFLAGS: -L. -lv8wrapper -lv8  -lstdc++
// #include <stdlib.h>
// #include "v8wrapper.h"
import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func RunV8(script string, result interface{}) error {

	// convert Go string to nul terminated C-string
	cstr := C.CString(script)
	defer C.free(unsafe.Pointer(cstr))

	// run script and convert returned C-string to Go string
	rcstr := C.runv8(cstr)
	defer C.free(unsafe.Pointer(rcstr))

	jsonstr := C.GoString(rcstr)

	fmt.Printf("Runv8 json -> %s\n", jsonstr)
	// unmarshal result
	err := json.Unmarshal([]byte(jsonstr), result)
	if err != nil {
		return err
	}

	fmt.Printf("Runv8 Result -> %T: %#+v\n", result, result)
	return nil
}
