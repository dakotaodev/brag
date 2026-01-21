package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dakotaodev/brag/internal/models"
	"github.com/dakotaodev/brag/internal/storage"
)

func main() {
	testBrag := models.Brag{
		ID:        "1",
		Timestamp: time.Now(),
		Category:  "Test",
		Content:   "Hello",
	}
	err := storage.SaveBrag("brags.json", testBrag)
	if err != nil {
		log.Fatalf("Failed to save brag: %v", err)
	}
	fmt.Println("Well done!")
}
