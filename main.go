package main 

import "fmt"
import "net/http"

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

 func main(){
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
} 
