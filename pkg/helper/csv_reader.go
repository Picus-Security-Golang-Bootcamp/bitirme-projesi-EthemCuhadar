package helper

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
)

// ReadCSV takes csv file and returns category body.
func ReadCSV(csvPartFile multipart.File) (*models.Category, error) {
	csvLines, readErr := csv.NewReader(csvPartFile).ReadAll()
	if readErr != nil {
		return nil, readErr
	}

	var categoryBody models.Category
	categoryBody.Name = csvLines[1][0]

	return &categoryBody, nil
}
