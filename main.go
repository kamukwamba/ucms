package main

import (
	"fmt"
	"log"
	"net/http"
	"ucmps/dbcode"
	"ucmps/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))

	dbread := dbcode.SqlRead()

	defer dbread.DB.Close()
	sqlStmt := `
	create table if not exists admin(id text not null , admin_name text, admin_email text, admin_password text);
	`
	_, err := dbread.DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	defer dbread.DB.Close()

	fmt.Println("Server running")
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/aboutus", routes.AboutUs)
	http.HandleFunc("/programs", routes.Programs)
	http.HandleFunc("/login", routes.LoginPage)
	http.HandleFunc("/enroll", routes.Enrollment)
	http.HandleFunc("/adminlogin", routes.AdminLogin)
	http.HandleFunc("/admindashboard", routes.AdminDashboard)

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.ListenAndServe(":3000", nil)

}
