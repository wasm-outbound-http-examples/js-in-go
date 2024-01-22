package main

import (
	"io"
	"net/http"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func main() {
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)

	httpGetFunc := func(fcall goja.FunctionCall) goja.Value {
		url := fcall.Argument(0).String()
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return vm.ToValue(string(body))
	}
	vm.Set("httpget", httpGetFunc)
	vm.RunString(`
		console.log(httpget('https://httpbin.org/anything'));
	`)
}
