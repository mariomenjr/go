package main

import (
	"fmt"
	"net/http"

	"github.com/mariomenjr/handlr"
)

func main() {
	h := handlr.New()

	h.Handler("/:id", linkHandler)

	h.Start(1993)
}

func linkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Method+" "+r.URL.Path)
}
