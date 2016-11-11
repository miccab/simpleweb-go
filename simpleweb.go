package main

import (
	"net/http"
	"github.com/miccab/simpleweb-go/nonblockingjava"
)

func main() {
	http.HandleFunc("/productGo", nonblockingjava.Handler())
	http.ListenAndServe(":8080", nil)
}
