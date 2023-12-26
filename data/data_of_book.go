package data

type Bookinfo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Language string `json:"language"`
}

type library []Bookinfo

var Booklist = library{

	Bookinfo{
		Id:       101,
		Name:     "Opekkha",
		Author:   "Humayun Ahmed",
		Language: "Bengali",
	},

	Bookinfo{
		Id:       102,
		Name:     "The Diary of Anna Frank",
		Author:   "Anna Frank",
		Language: "Jewish",
	},
	Bookinfo{
		Id:       103,
		Name:     "Keep Going",
		Author:   "Austin Kleon",
		Language: "English",
	},
	Bookinfo{
		Id:       104,
		Name:     "Emma",
		Author:   "Jane Austin",
		Language: "English",
	},
	Bookinfo{
		Id:       105,
		Name:     "Kaizen",
		Author:   "Sarah Harvey",
		Language: "English",
	},
	Bookinfo{
		Id:       106,
		Name:     "The old man and the sea",
		Author:   "Ernest Hemingway",
		Language: "English",
	},
}
