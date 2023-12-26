package apiCall

import (
	"encoding/json"
	"example.com/mod/authentication"
	"example.com/mod/data"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

// for homepage
func homepage(response http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(response, "welcome to the home page of central Library Collections\n")
	if err != nil {
		return
	}
}

// give all book info
func p1(w http.ResponseWriter, r *http.Request) {
	for _, ll := range data.Booklist {
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

	_, err = fmt.Fprintf(w, "Archived Information of : \n")
	if err != nil {
		return
	}
	flag := false

	for _, t := range data.Booklist {
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
		_, err := fmt.Fprintf(w, "Book information showed Successfully")
		if err != nil {
			return
		}
	} else {
		_, err := fmt.Fprintf(w, "BookID doesn't found in this archive")
		if err != nil {
			return
		}
	}
}

// add book to the book list
func createList(w http.ResponseWriter, r *http.Request) {

	var list []data.Bookinfo //it's a book info type array for creating  book information of multiples
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		//fmt.Println("jugsfkjugvkgrk")
		return
	}

	data.Booklist = append(data.Booklist, list...) //list... because for creating array, automatically handle array

	_, err = fmt.Fprintf(w, "New Book is added to the Archive")
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated) //201
}

func bookUpdate(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return
	}

	fmt.Println("Updated Info about Book, BookID : ", bookID)

	for i, ll := range data.Booklist {
		if ll.Id == bookID {
			var list data.Bookinfo
			err := json.NewDecoder(r.Body).Decode(&list) //single info will be decoded, not a bunch of
			if err != nil {
				return
			}
			data.Booklist[i] = list
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

	for i, ll := range data.Booklist {
		if ll.Id == bookID {
			data.Booklist = append(data.Booklist[:i], data.Booklist[i+1:]...)
			_, err2 := fmt.Fprintf(w, "Deleted the Book information from the Archive")
			if err2 != nil {
				return
			}
			return
		}
	}
	_, err = fmt.Fprintf(w, "Deleted Book from archive ")
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ServerStartPoint(port int) {
	fmt.Println("library collections ----")
	fmt.Println("Staring to show server status, ")

	fmt.Printf("Server Start point is at %v : ", port)

	//for authorization
	var samplekey = []byte("dynamic")
	fmt.Println("generated key is: ")
	fmt.Println(authentication.GenerateJWTKey(samplekey))

	ch := chi.NewRouter()

	ch.HandleFunc("/", authentication.VerifyAuth(homepage))
	ch.Get("/library", authentication.VerifyAuth(p1))
	ch.Get("/library/{id}", authentication.VerifyAuth(specificBookID))

	ch.Post("/library", authentication.VerifyAuth(createList))
	ch.Put("/library/{id}", authentication.VerifyAuth(bookUpdate))
	ch.Delete("/library/{id}", authentication.VerifyAuth(deleteBook))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(8080), ch))

}
