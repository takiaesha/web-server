package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type bookinfo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Language string `json:"language"`
}

type library []bookinfo

var booklist = library{

	bookinfo{
		Id:       101,
		Name:     "Opekkha",
		Author:   "Humayun Ahmed",
		Language: "Bengali",
	},

	bookinfo{
		Id:       102,
		Name:     "The Diary of Anna Frank",
		Author:   "Anna Frank",
		Language: "Jewish",
	},
	bookinfo{
		Id:       103,
		Name:     "Keep Going",
		Author:   "Austin Kleon",
		Language: "English",
	},
	bookinfo{
		Id:       104,
		Name:     "Emma",
		Author:   "Jane Austin",
		Language: "English",
	},
	bookinfo{
		Id:       105,
		Name:     "Kaizen",
		Author:   "Sarah Harvey",
		Language: "English",
	},
	bookinfo{
		Id:       106,
		Name:     "The old man and the sea",
		Author:   "Ernest Hemingway",
		Language: "English",
	},
}

// for homepage
func homepage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "welcome to the home page of central Library Collections\n")
}

// give all book info
func p1(w http.ResponseWriter, r *http.Request) { // ~all sales
	for _, ll := range booklist {
		err := json.NewEncoder(w).Encode(ll)
		if err != nil {
			return
		}
	}
}

// for specific book info
func specificBookID(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(chi.URLParam(r, "id"))
	// as used integer type id, and URLParam takes string type, so t.Id == bookID creates conflict between integer and string type
	//so converts the string bookID to integer
	if err != nil {
		return
	}

	fmt.Fprintf(w, "Archived Information of : \n")
	flag := false

	for _, t := range booklist {
		if t.Id == bookID {
			err := json.NewEncoder(w).Encode(t)
			if err != nil {
				return
			}
			flag = true
		}
	}

	if flag {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Book information showed Successfully")
	} else {
		fmt.Fprintf(w, "BookID doesn't found in this archive")
	}
}

// add book to the booklist
func createlist(w http.ResponseWriter, r *http.Request) {
	var list []bookinfo //it's a bookinfo type array for creating  book information of multiples
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		//fmt.Println("jugsfkjugvkgrk")
		return
	}

	booklist = append(booklist, list...)

	fmt.Fprintf(w, "New Book is added to the Archive")
	w.WriteHeader(http.StatusCreated) //201
}

func bookUpdate(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return
	}

	fmt.Println("Updated Info about Book, BookID : ", bookID)

	for i, ll := range booklist {
		if ll.Id == bookID {
			var list bookinfo
			err := json.NewDecoder(r.Body).Decode(&list) //single info will be decoded, not a bunch of
			if err != nil {
				return
			}
			booklist[i] = list
			err = json.NewEncoder(w).Encode(list)
			if err != nil {
				return
			}
			fmt.Println("Book Information is Successfully Updated ")
			return
		}
	}
	_, err = fmt.Fprintf(w, " BookID is not found in the archive ")
	if err != nil {
		return
	}
	//w.WriteHeader(http.StatusOK)   //200
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		return
	}

	for i, ll := range booklist {
		if ll.Id == bookID {
			booklist = append(booklist[:i], booklist[i+1:]...)
			fmt.Fprintf(w, "Deleted the Book information from the Archive")
			return
		}
	}
	fmt.Fprintf(w, "Deleted BookID ")
	w.WriteHeader(http.StatusCreated)
}
func main() {
	fmt.Println("library collections ----")

	ch := chi.NewRouter()

	fmt.Println("Staring to show server status, ")

	//printing homepage
	ch.HandleFunc("/", homepage)
	//ch.Use(middleware.Logger())

	ch.Get("/library", p1)
	ch.Get("/library/{id}", specificBookID)
	ch.Post("/library", createlist)

	ch.Put("/library/{id}", bookUpdate)
	ch.Delete("/library/{id}", deleteBook)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(8080), ch))

}
