package main

import (
	"fmt"
	"net/http"
)

type employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// //default mux
	// mux := http.NewServeMux()
	// //mux handle func
	// mux.HandleFunc("/v1/teachers", teacherHandler)
	// //mux handle type
	// sHandler := studentHandler{}
	// mux.Handle("/v1/students", sHandler)
	// //create server
	// s := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// }
	// s.ListenAndServe()

	createEmployeeHandler := http.HandlerFunc(createEmployee)
	http.Handle("/employee", createEmployeeHandler)
	http.ListenAndServe(":8080", nil)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()
	fmt.Println("request.Form::")
	for key, value := range r.Form {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}
	fmt.Println("\nrequest.PostForm::")
	for key, value := range r.PostForm {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}

	fmt.Printf("\nName field in Form:%s\n", r.Form["name"])
	fmt.Printf("\nName field in PostForm:%s\n", r.PostForm["name"])
	fmt.Printf("\nHobbies field in FormValue:%s\n", r.FormValue("hobbies"))

	w.WriteHeader(200)
	for key, value := range r.Form {
		w.Write([]byte(key)])
		w.Write([]byte(value)])
	}
}

// func teacherHandler(res http.ResponseWriter, req *http.Request) {
// 	data := []byte("v1 of teachers called")
// 	res.WriteHeader(200)
// 	res.Write(data)
// }

// type studentHandler struct{}

// func (h studentHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	data := []byte("v1 of students called")
// 	res.WriteHeader(200)
// 	res.Write(data)
// }
