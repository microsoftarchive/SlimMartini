package slim

import (
	"fmt"
	"github.com/SlyMarbo/spdy"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"runtime"
	"time"
)

func productionize() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	martini.Env = martini.Prod
}

func redirector(w http.ResponseWriter, req *http.Request) {
	url := "https://" + req.Host + req.RequestURI
	code := http.StatusMovedPermanently
	http.Redirect(w, req, url, code)
}

func RedirectServer(port int) {
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(redirector),
	}
	server.ListenAndServe()
	log.Println("redirecting http traffic to https")
}

func HttpServer(handler *Handler, port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Println("http://0.0.0.0" + addr)
	server := &http.Server{
		Addr:        addr,
		Handler:     handler,
		ReadTimeout: 10 * time.Second,
	}
	server.SetKeepAlivesEnabled(false)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func SpdyServer(handler *Handler, port int, cert string, key string) {
	productionize()
	addr := fmt.Sprintf(":%d", port)
	log.Println("https://0.0.0.0" + addr)
	err := spdy.ListenAndServeTLS(addr, cert, key, handler)
	if err != nil {
		log.Fatal(err)
	}
}
