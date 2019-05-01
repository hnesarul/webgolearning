package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tmpl *template.Template

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func init() {
	tmpl = template.Must(template.New("tpl.html").Funcs(fm).ParseFiles("tpl.html"))
}

func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "tpl.html", time.Now())

	if err != nil {
		log.Fatalln(err)
	}
}
