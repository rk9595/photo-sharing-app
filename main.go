package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rk9595/photoSharingApp/controllers"
	"github.com/rk9595/photoSharingApp/templates"
	"github.com/rk9595/photoSharingApp/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
