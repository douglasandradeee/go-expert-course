package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

// Possibilidade de webHook para resolver meu problema da unifique
func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
		t = template.Must(t.ParseFiles(templates...))

		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 20},
			{"Python", 30},
		})

		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)

}
