package main

import "flag"
import "fmt"
import "github.com/go-martini/martini"

import "log"
import "net/http"

const defaultPort = 80
const defaultDir = "."

var port = flag.Int("p", defaultPort, "port to listen")
var dir = flag.String("d", defaultDir, "directory to serve")

func main() {
	flag.Parse()
	m := martini.Classic()
	m.Use(martini.Static(*dir))
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("serve http %s on %s\n", *dir, addr)
	log.Fatal(http.ListenAndServe(addr, m))
}
