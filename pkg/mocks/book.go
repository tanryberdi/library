package mocks

import "library/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
	{
		Id: 2,
		Title: "Python",
		Author: "Orazz",
		Desc: "Good tutorial",
	},
}
