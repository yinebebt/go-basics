// practice on templets in Go
package postrender

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

// if you're continuing from the read files chapter, you shouldn't redefine this
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

/*v-1*/
// func Render(w io.Writer, p Post) error {

// 	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
// 	return err
// }
/*v-1.1*/
// func Render(w io.Writer, p Post) error {
// 	_, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, p.Description)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = fmt.Fprint(w, "Tags: <ul>")
// 	if err != nil {
// 		return err
// 	}

// 	for _, tag := range p.Tags {
// 		_, err = fmt.Fprintf(w, "<li>%s</li>", tag)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	_, err = fmt.Fprint(w, "</ul>")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
/*v-2*/

// const (
// 	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
// )

// func Render(w io.Writer, p Post) error {
// 	templ, err := template.New("blog").Parse(postTemplate)
// 	if err != nil {
// 		return err
// 	}

// 	if err := templ.Execute(w, p); err != nil {
// 		return err
// 	}

// 	return nil
// }

/*v-2.1*/
var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func (pr *postRender) Render(w io.Writer, p Post) error {

	if err := pr.templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}

// doing parser and render in separate to see effect via benchmark test
type postRender struct {
	templ *template.Template
}

func NewPostParser() (*postRender, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &postRender{templ: templ}, nil

}

func (r *postRender) RenderIndex(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`

	templ, err := template.New("index").Funcs(template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.Replace(title, " ", "-", -1)) //-1 means no limit on the #of replacementt
		},
	}).Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, posts); err != nil {
		return err
	}

	return nil
}
