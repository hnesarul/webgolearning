package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tmpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufecturer string
	Model        string
	Doors        int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tmpl = template.Must(template.New("tpl.html").Funcs(fm).ParseFiles("tpl.html"))
}

func main() {

	m := sage{
		Name:  "Muhammed",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	j := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	b := sage{
		Name:  "Buddha",
		Motto: "The Belief of no beliefs",
	}

	f := car{
		Manufecturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufecturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{m, g, j, b}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tmpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
