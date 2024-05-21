package routes

import (
	"html/template"
	"log"
	"net/http"
)

func Enrollment(w http.ResponseWriter, r *http.Request) {

	// r.ParseForm()
	// if r.Method == "POST" {
	// 	fmt.Println("Form obtained")
	// }
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	//debug failure to laod templates

	err := tpl.ExecuteTemplate(w, "enrollstudent.html", nil)

	if err != nil {
		log.Fatal(err)
	}
}
