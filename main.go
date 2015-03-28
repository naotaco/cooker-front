package main

import (
	"fmt"
	"net/http"
	"strconv"

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

func setTarget(c web.C, w http.ResponseWriter, r *http.Request) {
	target := r.PostFormValue("target")
	fmt.Println(target)

	t, err := strconv.ParseFloat(target, 32)
	if err != nil || t < 0 || t > 100 {
		// Invalid format or out of correct range
		http.Redirect(w, r, "/", 400)
		return
	}

	fmt.Println("Set " + target)
	setTargetTemp(target)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)	// workaround to suppress transition page.
}

func forceOff(c web.C, w http.ResponseWriter, r *http.Request) {
	setTargetTemp("0.0")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func forceOn(c web.C, w http.ResponseWriter, r *http.Request) {
	setTargetTemp("100.0")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func main() {
	goji.Get("/", root)
	goji.Post("/set", setTarget)
	goji.Post("/force_on", forceOn)
	goji.Post("/force_off", forceOff)
	goji.Serve()
}
