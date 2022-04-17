package helper

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/internal/entity/models"
)

func ReadCSV(csvPartFile multipart.File) (*models.Category, error) {
	csvLines, readErr := csv.NewReader(csvPartFile).ReadAll()
	if readErr != nil {
		return nil, readErr
	}

	var categoryBody models.Category
	categoryBody.ID = csvLines[1][0]
	categoryBody.Name = csvLines[1][1]

	return &categoryBody, nil
}
