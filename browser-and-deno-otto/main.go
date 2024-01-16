package main

import (
	"io"
	"net/http"

	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	httpGetFunc := func(fcall otto.FunctionCall) otto.Value {
		url := fcall.Argument(0).String()
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		res, _ := vm.ToValue(string(body))
		return res
	}
	vm.Set("httpget", httpGetFunc)
	vm.Run(`
		console.log(httpget('https://httpbin.org/anything'));
	`)
}
