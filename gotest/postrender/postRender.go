// practice on templets in Go
package postrender

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
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

// /*v-2.1*/
// var (
// 	//go:embed "templates/*"
// 	postTemplates embed.FS
// )

// func (pr *postRender) Render(w io.Writer, p Post) error {

// 	if err := pr.templ.Execute(w, p); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // doing parser and render in separate to see effect via benchmark test
type postRender struct {
	templ    *template.Template
	mdParser *parser.Parser //for last version
}

// func NewPostParser() (*postRender, error) {
// 	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &postRender{templ: templ}, nil

// }

// func (r *postRender) RenderIndex(w io.Writer, posts []Post) error {
// 	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`

// 	templ, err := template.New("index").Funcs(template.FuncMap{
// 		"sanitiseTitle": func(title string) string {
// 			return strings.ToLower(strings.Replace(title, " ", "-", -1)) //-1 means no limit on the #of replacementt
// 		},
// 	}).Parse(indexTemplate)
// 	if err != nil {
// 		return err
// 	}

// 	if err := templ.Execute(w, posts); err != nil {
// 		return err
// 	}

// 	return nil
// }
/*last-version*/

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// SanitisedTitle returns the title of the post with spaces replaced by dashes for pleasant URLs
func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

// NewPostRenderer creates a new PostRenderer
func NewPostRenderer() (*postRender, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &postRender{templ: templ, mdParser: parser}, nil
}

// Render renders post into HTML
func (r *postRender) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

// RenderIndex creates an HTML index page given a collection of posts
func (r *postRender) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

// Rather than our rendering code being coupled to the domain object, Post
type postViewModel struct {
	Post     //embedded field, sub struct
	HTMLBody template.HTML
}

func newPostVM(p Post, r *postRender) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
