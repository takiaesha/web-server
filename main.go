package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	_ "github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strconv"
	"time"
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
	_, err := fmt.Fprintf(response, "welcome to the home page of central Library Collections\n")
	if err != nil {
		return
	}
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

	_, err = fmt.Fprintf(w, "Archived Information of : \n")
	if err != nil {
		return
	}
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
	var list []bookinfo //it's a book info type array for creating  book information of multiples
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		//fmt.Println("jugsfkjugvkgrk")
		return
	}

	booklist = append(booklist, list...)

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

func generateJWTKey(SigningKey []byte) (string, error) {
	//var samplekey = []byte("dynamic")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(3 * time.Minute).Unix()
	claims["authorized"] = "true"
	claims["user"] = "Esha"

	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		fmt.Println("Are you really a Human?!!")
		return "", nil
	}
	return tokenString, nil
}

func verifyAuth(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var sampleJWT = []byte("dynamic")

		bearer := r.Header.Get("Authorization")
		if len(bearer) > 0 {
			bearer = bearer[7:]

			token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
				_, alg := token.Method.(*jwt.SigningMethodHMAC)

				if !alg {
					return nil, fmt.Errorf("invalid Error")
				}
				return sampleJWT, nil
			})

			if err != nil {
				_, err2 := fmt.Fprintf(w, err.Error())
				if err2 != nil {
					return
				}
			}

			if token.Valid {
				endpoint(w, r)
			} else {
				_, err2 := fmt.Fprintf(w, "! You are not authorized")
				if err2 != nil {
					return
				}
			}
		}
	})

}

func main() {
	fmt.Println("library collections ----")
	fmt.Println("Staring to show server status, ")

	//for authorization
	var samplekey = []byte("dynamic")
	fmt.Println("generated key is: ")
	fmt.Println(generateJWTKey(samplekey))

	ch := chi.NewRouter()

	ch.HandleFunc("/", verifyAuth(homepage))
	ch.Get("/library", verifyAuth(p1))
	ch.Get("/library/{id}", verifyAuth(specificBookID))

	ch.Post("/library", verifyAuth(createList))
	ch.Put("/library/{id}", verifyAuth(bookUpdate))
	ch.Delete("/library/{id}", verifyAuth(deleteBook))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(8080), ch))

}
