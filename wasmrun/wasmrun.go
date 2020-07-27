// +
// build js,wasm

package main

import (
	"log"
	"syscall/js"
)

func createString(this js.Value, args []js.Value) interface{} {
	if args == nil {
		logf("createString input is nil")
	}

	b := args[0].String()
	logf("createString: b = %s\n len(b) %d", b, len(b))

	v := js.Global().Get("Uint8Array").New(len(b)) // no
	js.CopyBytesToJS(v, []byte(b))
	return v
}

// readStringVal reads js.Val from js runtime to Go
func readStringVal(s js.Value) string {
	st := s.String()
	l := s.Length()
	logf("readStringVal: b = %s length %d\n", st, l)
	// read bytes
	buf := make([]byte, l, l)
	js.CopyBytesToGo(buf, s)
	return string(buf)
}

// writeStringVal writes js.Val from Go to js runtime
func writeStringVal(b string) interface{} {
	r := []byte(b)
	logf("writeStringVal with bytes: r = %s\n len(r) %d", r, len(r))

	// write bytes
	v := js.Global().Get("Uint8Array").New(len(r))
	js.CopyBytesToJS(v, r)
	return v
}

func run(this js.Value, args []js.Value) interface{} {
	logf("inside WASM run")
	kcVal := args[0]
	logf("kc : b = %s\n", kcVal.String())

	kc := readStringVal(kcVal)

	res := string(kc + " and from Wasm in Wasmer")

	logf("Returning: %s", res)
	return writeStringVal(res)
}

func logf(format string, v ...interface{}) {
	if LOGENABLED {
		log.Printf(format, v...)
	}
}

var LOGENABLED = false

func main() {
	ch := make(chan bool)
	logf("inside wasm main")
	// register functions
	js.Global().Set("createString", js.FuncOf(createString))
	js.Global().Set("run", js.FuncOf(run))

	<-ch
}
