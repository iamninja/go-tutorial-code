package main

import (
	"fmt"
	"html/template"
	"os"
)

type Article struct {
	Name		string
	AuthorName	string
	Draft		bool
}

func (a Article) Byline() string {
	return fmt.Sprintf("Written by %s", a.AuthorName)
}

func Multiply(a, b float64) float64 {
	return a * b
}

type Product struct {
	Price		float64
	Quantity	float64
}

func main() {
	goArticle := Article{
		Name:		"The GO html/template package",
		AuthorName:	"Mal Curtis",
	}
	// article := map[string]string {
	// 	"Name": 		"The Go html/template package",
	// 	"AuthorName":	"Mal Curtis",
	// }


	//
	tmpl, err := template.New("Foo").Parse("'{{.Name}}' {{.Byline}} {{if .Draft}}(Draft){{else}}(Published){{end}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil { panic(err) }

	//
	tmpl, err = template.New("Foo").Parse(`
	{{range .}}
		<p>{{.Name}} by {{.AuthorName}}</p>
	{{else}}
		<p>No articles yet</p>
	{{end}}
	`)
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, []Article{})
	if err != nil { panic(err) }

	//
	tmpl, err = template.New("Foo").Parse(`
	{{define "ArticleResource"}}
		<p>{{.Name}} by {{.AuthorName}}</p>
	{{end}}

	{{define "ArticleLoop"}}
		{{range .}}
			{{template "ArticleResource" .}}
		{{else}}
			<p>No published articles yet</p>
		{{end}}
	{{end}}

	{{template "ArticleLoop" .}}
	`)
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, []Article{})
	if err != nil { panic(err) }

	//
	tmpl, _ = template.New("Foo").Parse(
		"Price: ${{printf \"%.2f\" .}}\n",
	)
	tmpl.Execute(os.Stdout, 12.3)

	//
	tmpl, _ = template.New("Foo").Parse(
		"Price: ${{. | printf \"%.2f\"}}\n",
	)
	tmpl.Execute(os.Stdout, 12.4)

	//
	tmpl = template.New("Foo")
	tmpl.Funcs(template.FuncMap{ "multiply": Multiply })

	tmpl, err = tmpl.Parse(
		"Price: ${{ multiply .Price .Quantity | printf \"%.2f\"}}\n",
	)
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, Product{
		Price:		12.3,
		Quantity:	2,
	})
	if err != nil { panic(err) }

	//
	tmpl = template.New("Foo")
	tmpl.Funcs(template.FuncMap{ "multiply": Multiply })

	tmpl, err = tmpl.Parse(`
	{{$total := multiply .Price .Quantity}}
	Price: ${{ printf "%.2f" $total }}
	`)
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, Product{
		Price:		12.3,
		Quantity:	2,
	})
	if err != nil { panic(err) }
}