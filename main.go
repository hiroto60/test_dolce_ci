package main

import (
	"net/http"

	"github.com/hiroto60/test_dolce_ci/gen/greet/greetconnect"
	"github.com/hiroto60/test_dolce_ci/gen/hoge/hogeconnect"
	"github.com/hiroto60/test_dolce_ci/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	greeter := handler.NewGreetHandler()
	greetPath, greetHandler := greetconnect.NewGreeterHandler(greeter)

	hoge := handler.NewHogeHandler()
	hogePath, hogeHandler := hogeconnect.NewHogeHandler(hoge)

	mux := http.NewServeMux()
	mux.Handle(greetPath, greetHandler)
	mux.Handle(hogePath, hogeHandler)

	http.ListenAndServe(
		"localhost:8084",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
