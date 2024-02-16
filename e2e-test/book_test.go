package e2e_test

import (
	"bytes"
	"example.com/mod/authentication"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
)

var _ = Describe("Book", func() {
	var (
		//server http.Server
		client    *http.Client
		samplekey = []byte("dynamic")
		url       string
	)
	Context("api-server test", func() {
		BeforeEach(func() {
			url = fmt.Sprintf("http://localhost:%d", tunnel.Local)
		})
		//GET
		Context("give the list of book-library", func() {
			It("returns the list of all books through GET", func() {
				request, _ := http.NewRequest("GET", url+"/library", bytes.NewBuffer([]byte(nil)))
				request.Header.Set("Content-Type", "application/json")
				token, err := authentication.GenerateJWTKey(samplekey)
				Expect(err).NotTo(HaveOccurred())
				bearer := "Bearer " + token

				request.Header.Set("Authorization", bearer)
				request.Header.Add("Accept", "application/json")

				client = &http.Client{}
				response, err := client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				By("library book list(GET)")
				Expect(response.StatusCode).To(Equal(http.StatusOK))

				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						return
					}
				}(response.Body)
			})
		})

		//get-specific id
		Context("Give a specific book info", func() {
			It("return a specific book info through GET", func() {

				request, _ := http.NewRequest("GET", url+"/library/104", bytes.NewBuffer([]byte(nil)))
				request.Header.Set("Content-Type", "application/json")

				client = &http.Client{}
				response, err := client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				By("Specific Book from Library (GET)")
				Expect(response.StatusCode).To(Equal(http.StatusOK))

				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						return
					}
				}(response.Body)
			})
		})

		//post book
		Context("Book Information add", func() {
			It("While creating a new book information", func() {
				requestBody := `{"id": 110, "name": "Farm House", "author": "harvey", "language": "English"}`
				request, _ := http.NewRequest("POST", url+"/library", bytes.NewBuffer([]byte(requestBody)))
				request.Header.Set("Content-Type", "application/json")

				client = &http.Client{}
				response, err := client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				By("library book add (POST)")
				Expect(response.StatusCode).To(Equal(http.StatusOK))

				if err != nil {
					By(err.Error())
				}

				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						return
					}
				}(response.Body)
			})
		})

		//DELETE bookID-106
		Context("Delete a specific book info", func() {
			It("delete a book details", func() {

				request, _ := http.NewRequest("DELETE", url+"/library/106", bytes.NewBuffer([]byte(nil)))
				request.Header.Set("Content-Type", "application/json")

				client = &http.Client{}
				response, err := client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				By("library book delete(DELETE)")
				Expect(response.StatusCode).To(Equal(http.StatusOK))

				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						return
					}
				}(response.Body)

			})
		})
	})

})
