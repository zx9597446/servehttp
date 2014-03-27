package main

import "github.com/codegangsta/martini"
import "log"
import "net/http"

const defaultAddr = ":9502"

func main() {
	m := martini.Classic()
	m.Use(martini.Static("."))
	log.Printf("serve http %s\n", defaultAddr)
	log.Fatal(http.ListenAndServe(defaultAddr, m))
}
