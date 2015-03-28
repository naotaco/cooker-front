package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/flosch/pongo2"
)

type Page struct {
	Title              string
	CurrentTemperature float64
	TargetTemperature  float64
	IsHeating          string
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func root(c web.C, w http.ResponseWriter, r *http.Request) {

	var page Page
	page.Title = "My Cooker"

	page.TargetTemperature = getTargetTemp()
	page.CurrentTemperature = getCurrentTemp()
	page.IsHeating = isHeating()

	tpl, err := pongo2.DefaultSet.FromFile("template.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteWriter(pongo2.Context{"page": page}, w)
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Get("/", root)
	goji.Serve()

}