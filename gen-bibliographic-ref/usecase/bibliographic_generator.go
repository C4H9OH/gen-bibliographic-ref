package usecase

import (
	"fmt"
	"gen-bibliographic-ref/model"
)

func CreateBibliographicRef(bibliographicComponents model.BibliographicComponents) string {
	return fmt.Sprintf(
		"%s %s - %d-е изд. - %s: %s, %d. - %d с.",
		bibliographicComponents.Author,
		bibliographicComponents.BookName,
		bibliographicComponents.PublicationNumber,
		bibliographicComponents.PublicationCity,
		bibliographicComponents.PublicationName,
		bibliographicComponents.PublicationYear,
		bibliographicComponents.PagesNumber,
	)
}
