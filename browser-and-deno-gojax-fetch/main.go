// Based on https://github.com/olebedev/gojax/blob/master/fetch/README.md#usage
package main

import (
	"sync"

	goproxy "gopkg.in/elazarl/goproxy.v1"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/olebedev/gojax/fetch"
)

func main() {
	loop := eventloop.NewEventLoop()
	loop.Start()
	defer loop.Stop()

	fetch.Enable(loop, goproxy.NewProxyHttpServer())

	var wg sync.WaitGroup
	wg.Add(1)
	loop.RunOnLoop(func(vm *goja.Runtime) {
		vm.Set("exitloop", func(call goja.FunctionCall) goja.Value {
			wg.Done()
			return nil
		})

		vm.RunString(`
			fetch('https://httpbin.org/anything')
			.then( function(res) {
				return res.text();
			}).then( function(txt) {
				console.log(txt);
				exitloop();
			});
		`)
	})
	wg.Wait()
}
