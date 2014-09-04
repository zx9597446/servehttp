package main

import "github.com/codegangsta/martini"
import "io/ioutil"
import "log"
import "net/http"

const defaultAddr = ":80"

func viascan() string {
	b, err := ioutil.ReadFile("viascan.html")
	if err != nil {
		log.Panicln(err, b)
	}
	return string(b)
}

func main() {
	m := martini.Classic()
	m.Use(martini.Static("."))
	m.Get("/viascan", viascan)
	log.Printf("serve http %s\n", defaultAddr)
	log.Fatal(http.ListenAndServe(defaultAddr, m))
}
