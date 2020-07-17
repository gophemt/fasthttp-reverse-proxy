package main

import (
	"log"

	proxy "github.com/gophemt/fasthttp-reverse-proxy"
	"github.com/valyala/fasthttp"
)

var (
	weights = map[string]proxy.Weight{
		"localhost:9090": 20,
		"localhost:9091": 30,
		"localhost:9092": 50,
	}

	proxyServer = proxy.NewReverseProxy("", proxy.WithBalancer(weights))
)

// ProxyHandler ... fasthttp.RequestHandler func
func ProxyHandler(ctx *fasthttp.RequestCtx) {
	// all proxy to localhost
	proxyServer.ServeHTTP(ctx)
}

func main() {
	if err := fasthttp.ListenAndServe(":8081", ProxyHandler); err != nil {
		log.Fatal(err)
	}
}
