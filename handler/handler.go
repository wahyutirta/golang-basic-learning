package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Is Happening -load file", http.StatusInternalServerError)
		return
	}

	//slice of struct
	data := []entity.Product{
		{Id: 1, Name: "Mobilio", Price: 280000000, Stock: 2},
		{Id: 2, Name: "Xpander", Price: 380000000, Stock: 11},
		{Id: 3, Name: "Jazz", Price: 295000000, Stock: 9},
	}

	// data := map[string]interface{}{
	// 	"title":   "Learning Golang",
	// 	"content": "I'M Learning Golang Web Site Development",
	// }

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Is Happening, -tmpl exe", http.StatusInternalServerError)
		return
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world, golang web"))

}
func NameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello mrs...., golang web"))

}
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"content": idNum,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Is Happening, -Loading html", http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Is Happening, -tmpl exe", http.StatusInternalServerError)
	}

}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini Adalah GET"))

	case "POST":
		w.Write([]byte("Ini Adalah POST"))
	default:
		http.Error(w, "Wrong method error, calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening on loading files, calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening on running files, calm", http.StatusInternalServerError)
			return
		}
		return

	}
	http.Error(w, "Error is happening", http.StatusBadRequest)

}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening on processing form, calm", http.StatusInternalServerError)
			return
		}
		name := r.Form.Get("name")
		message := r.Form.Get("message")

		w.Write([]byte(name))
		w.Write([]byte(message))

		return

	}
	http.Error(w, "Error HTTP Method, Calm", http.StatusInternalServerError)
}
