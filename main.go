package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Taufiq Sap"}
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Taufiq Sap"}
		var tmpl = template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	var address = "localhost:9000"
	fmt.Printf("Starting server at http://%s/\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}


/*func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"Title" : "Halo saya sedang belajar golang web",
			"name" : "TAUFIQ SAP TAMPAN",
		}

		err = tmpl.Execute(w, data)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	var address = "localhost:9000"
	fmt.Printf("Starting server at http://%s/\n", address)
	err:= http.ListenAndServe(address, nil)
	if( err != nil){
		fmt.Println("Error starting server:", err)
	}
} */


 /* func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {

	var message = "Hello World"
	w.Write([]byte(message))
}

func main(){

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)


	var address = "localhost:9000"
	fmt.Printf("starting server at %s\n", address)
	err:= http.ListenAndServe(address, nil)
	if err != nil{
		fmt.Println("Error starting server:", err)
	}
} */

 /*func main(){
	handlerIndex := func(w http.ResponseWriter, r *http.Request){
		var message = "Hello"
		w.Write([]byte(message))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	http.HandleFunc("/data",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello again"))
	})

	http.Handle("/static/",http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	var address = "localhost:9000"
	fmt.Printf("starting server at %s\n", address)
	err:= http.ListenAndServe(address, nil)
	if err != nil{
		fmt.Println("Error starting server:", err)
	}
} */