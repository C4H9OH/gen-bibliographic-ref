package model

type BibliographicComponents struct {
	Author            string `json:"author"`
	BookName          string `json:"book_name"`
	PublicationNumber int    `json:"publication_number"`
	PublicationCity   string `json:"publication_city"`
	PublicationName   string `json:"publication_name"`
	PublicationYear   int    `json:"publication_year"`
	PagesNumber       int    `json:"pages_number"`
}
