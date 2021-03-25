package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/wzshiming/openapiui/v2/redoc"
	"github.com/wzshiming/openapiui/v2/swaggereditor"
	"github.com/wzshiming/openapiui/v2/swaggerui"
)

var port int

func init() {
	flag.IntVar(&port, "p", 9000, "port")
}

func main() {

	p := fmt.Sprintf(":%d", port)
	http.Handle("/swaggerui/", http.FileServer(http.FS(swaggerui.FS)))
	http.Handle("/swaggereditor/", http.FileServer(http.FS(swaggereditor.FS)))
	http.Handle("/redoc/", http.FileServer(http.FS(redoc.FS)))
	fmt.Printf("Open http://127.0.0.1:%d/swaggerui/ or http://127.0.0.1:%d/swaggereditor/ or http://127.0.0.1:%d/redoc/ with your browser.\n", port, port, port)
	err := http.ListenAndServe(p, nil)
	if err != nil {
		log.Println(err)
	}

	return
}
