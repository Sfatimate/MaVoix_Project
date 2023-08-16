package main

import (
	"database/sql"
	"fmt"
	//"html/template"
	"log"
	//"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

//var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	server := "LAPTOP-RLPTS42I\\SQLEXPRESS"
	port := 1433
	user := "LAPTOP-RLPTS42I\\SAKSON"
	password := ""
	database := "MaVoix"

	connectionStr := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;",
		server, port, user, password, database)

	db, err := sql.Open("mssql", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to SQL Server!")
}
	/*http.HandleFunc("/home", homeHandler)

	http.HandleFunc("/ascii-art", reponseHandler)

	fmt.Println("http://localhost:9999/home")
	http.Handle("/", http.FileServer(http.Dir("static/")))
	http.ListenAndServe(":9999", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	err := templates.Execute(w, nil) // On exécute ensuite le template de la page home en lui passant url de la page /ascii-art.

	if err != nil {
		http.Error(w, "404 Not Found", http.StatusInternalServerError) // erreur 404
	}

}
func reponseHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 - Not Found", http.StatusNotFound) // Gestion de l'erreur 404 si le chemin de l'URL n'est pas trouvé
		return
	}
	// Exécution du template HTML pour la méthode GET
	if r.Method == "GET" {
	} else if r.Method == "POST" {
	}*/

