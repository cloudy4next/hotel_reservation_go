package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

// RenderTemple from html
func Rendetemplate(w http.ResponseWriter, tmpl string) {

	tempplateCache, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	template, ok := tempplateCache[tmpl]

	if !ok {
		log.Fatal(err)
	}

	byteBuffer := new(bytes.Buffer)

	_ = template.Execute(byteBuffer, nil)

	_, err = byteBuffer.WriteTo(w)

	if err != nil {
		fmt.Println("error while writing on browser!!")
	}
}

// create a ttemplate cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cachedTemp := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return cachedTemp, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cachedTemp, err
		}

		matches, err := filepath.Glob("./templates/*s.layout.html")
		if err != nil {
			return cachedTemp, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*s.layout.html")
			if err != nil {
				return cachedTemp, err
			}
		}
		cachedTemp[name] = templateSet

	}
	return cachedTemp, nil

}
