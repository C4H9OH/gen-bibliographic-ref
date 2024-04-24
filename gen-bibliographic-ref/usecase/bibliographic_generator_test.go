package usecase

import (
	"gen-bibliographic-ref/model"
	"testing"
)

func TestCreateBibliographicRef(t *testing.T) {
	testData := model.BibliographicComponents{
		Author:            "Бердюгин А.В.",
		BookName:          "Человеко-машинное взаимодействие. Методические указания.",
		PublicationNumber: 3,
		PublicationCity:   "М",
		PublicationName:   "МГТУ СТАНКИН",
		PublicationYear:   2020,
		PagesNumber:       109,
	}

	expectedOutput := "Бердюгин А.В. Человеко-машинное взаимодействие. Методические указания. - 3-е изд. - М: МГТУ СТАНКИН, 2020. - 109 с."

	res := CreateBibliographicRef(testData)

	if res != expectedOutput {
		t.Errorf(
			"ошибка! ожидалось: %s получено: %s",
			expectedOutput,
			res,
		)
	}
}
