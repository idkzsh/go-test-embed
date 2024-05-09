package handler

// package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed all:src
var staticFiles embed.FS
var templates = template.Must(template.ParseFS(staticFiles, "src/templates/*.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	// switch r.URL.Path {
	// case "/":
	// 	data := r.URL.Path
	// 	err := templates.ExecuteTemplate(w, "index.html", data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// // case "/src/style.css":
	// // 	file, err := staticFiles.Open("src/style.css")
	// // 	if err != nil {
	// // 		http.Error(w, err.Error(), http.StatusNotFound)
	// // 		return
	// // 	}
	// // 	defer file.Close()

	// // 	w.Header().Set("Content-Type", "text/css")
	// // 	_, err = io.Copy(w, file)
	// // 	if err != nil {
	// // 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// // 		return
	// // 	}
	// default:
	// 	http.NotFound(w, r)
	// }
	// data := r.URL.Path
	templates.ExecuteTemplate(w, "index.html", nil)
}

func Main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Handler)
	router.Handle("GET /src/", http.StripPrefix("/src/", http.FileServer(http.FS(staticFiles))))

	log.Println("LISTENING AT PORT:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
