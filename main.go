package main

import (
	"fmt"
	"net/http"

	"github.com/guregu/kami"
	"golang.org/x/net/context"
)

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", kami.Param(ctx, "name"))
}

func main() {
	kami.Get("/hello/:name", hello)
	kami.Serve()
}
