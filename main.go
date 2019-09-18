package main

import (
	"github.com/erikdubbelboer/fasthttp"
)

func main() {
	fasthttp.ListenAndServe(":8080", Router().InitRouter().HandleRequest)
}
