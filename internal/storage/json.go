package storage

import (
	"encoding/json"
	"os"
	"github.com/dakotaodev/brag/internal/models"
)

func SaveBrag(filepath string, entry models.Brag) error {
	var brags []models.Brag 
	
	// read in file 
	data, err := os.ReadFile(filepath)
	if err == nil {
		// this means that the file exists
		json.Unmarshal(data,&brags)
	}
	brags=append(brags, entry)

	// convert back to json
	updatedData, err := json.MarshalIndent(brags, "","  ")
	if err != nil {
		// hit an error, return int
		return err
	}

	return os.WriteFile(filepath, updatedData, 0644)
}