package main

import (
	"auth/router"
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed static
var static embed.FS

//go:embed templates
var tmpl embed.FS

func main() {
	r := router.StartRouter()
	r.StaticFS("/static", http.FS(static))
	t, _ := template.ParseFS(tmpl, "templates/*.html")
	r.SetHTMLTemplate(t)

	if err := r.Run(":8083"); err != nil {
		fmt.Println("运行时出错", err)
	}
}
