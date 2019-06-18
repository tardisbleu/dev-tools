package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"github.com/tardisbleu/dev-tools/backend/pdfgenerator"
)

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	pdfGenerator := &pdfgenerator.PdfGenerator{}

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "dev-tools",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(pdfGenerator)
	app.Run()
}
