package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brodyd795/lenslocked/controllers"
	"github.com/brodyd795/lenslocked/templates"
	"github.com/brodyd795/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := views.ParseFS(templates.FS, filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "contact.gohtml")
}

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// 	executeTemplate(w, "faq.gohtml")
// }

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "layout-page.gohtml", "home-page.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "layout-page.gohtml", "contact-page.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
