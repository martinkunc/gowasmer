// +build !js

package exec

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/martinkunc/wagather/wasm"
)

func NewWebAssembly() (*WebAssembly, func(), error) {
	n := time.Now()
	b, dealloc, err := wasm.BridgeFromFile("test", "wasmrun.wasm", nil)
	if err != nil {
		return nil, nil, err
	}
	// 12s
	log.Println(time.Now().Sub(n))

	w := &WebAssembly{}

	init := make(chan error)
	ctx, cancF := context.WithCancel(context.Background())
	//defer cancF()
	_ = cancF
	go b.Run(ctx, init)
	err = <-init
	if err != nil {
		return nil, nil, err
	}

	w.bridge = b
	return w, dealloc, nil
}

type WebAssembly struct {
	bridge *wasm.Bridge
}

func (w *WebAssembly) createString(s string) (interface{}, error) {
	return w.bridge.CallFunc("createString", []interface{}{s})
}

// readString reads js.Val value from WA runtime and converts it to string
func (w *WebAssembly) readString(jsVal interface{}) (string, error) {
	nb, err := wasm.Bytes(jsVal)
	if err != nil {
		return "", fmt.Errorf("readString Bytes error: %w", err)
	}
	sb := strings.Builder{}
	for i := 0; i < len(nb); i++ {
		sb.WriteByte(nb[i])
	}
	return sb.String(), nil
}

func (w *WebAssembly) Run() (string, error) {

	kubeVal, err := w.createString(string("Hello world from Golang"))
	if err != nil {
		return "", err
	}
	resVal, err := w.bridge.CallFunc("run", []interface{}{kubeVal})
	if err != nil {
		return "", fmt.Errorf("Run callFunc error: %w", err)
	}
	result, err := w.readString(resVal)
	if err != nil {
		return "", fmt.Errorf("Run readString error: %w", err)
	}

	return result, nil
}
