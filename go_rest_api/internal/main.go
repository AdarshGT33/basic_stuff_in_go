package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID        string   `json:"id,omitempty"`
	Title     string   `json:"title,omitempty"`
	Author    string   `json:"author,omitempty"`
	Publishor *Company `json:"publishor,omitempty"`
}

//`json...` is a tag(metadata) which JSON liberary uses to map Go struct field(ID, Title, etc) and JSON field

type Company struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

var books []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	//getting params out of request
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	/* func here is the handler function which is being matched by
	// 	DefaultServeMux(idk what the hell this is), this matching is done
	// 	by HandleFunc, moreover when a request is made from client side,
	// 	URL is used and when its server, URI is used (in ref to URL.Path)
	// 	This is not an api, all we have done here is started a server.
	// 	*/
	// 	fmt.Fprintln(w, "Hello", html.EscapeString(r.URL.Path))
	// 	fmt.Print()
	// })

	// log.Println("Listening on localhost:8000")

	// log.Fatal(http.ListenAndServe(":8000", nil))

	// router := mux.NewRouter()

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "hello world")
	// })

	// http.ListenAndServe(":8000", router)

	// router.HandleFunc("/users", getUsersHandlers).Methods(http.MethodGet)

	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Title: "Title 1", Author: "Author 1", Publishor: &Company{Name: "Comapany 1", Address: "Address 1"}})
	books = append(books, Book{ID: "2", Title: "Title 2", Author: "Author 2", Publishor: &Company{Name: "Comapany 2", Address: "Address 2"}})
	books = append(books, Book{ID: "3", Title: "Title 3", Author: "Author 3", Publishor: &Company{Name: "Comapany 3", Address: "Address 3"}})

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

// func getUsersHandlers(w http.ResponseWriter, r *http.Request) {
// 	// getting users from database
// 	users := getUsersFromDB()

// 	//converting list of users to JSON
// 	userJson, err := json.Marshal(users)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	//setting the content type of the response to JSON
// 	w.Header().Set("Content-Type", "application/json")

// 	//writing users to client in json format
// 	w.Write(userJson)
// }
