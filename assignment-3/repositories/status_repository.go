package repositories

import (
	"assignment-3/models"
	"encoding/json"
	"log"
	"os"
)

const SOURCE_FILE = "data/status.json"

func GetStatus() (models.Status, error) {
	var status models.Status

	content, err := os.ReadFile(SOURCE_FILE)
	if err != nil {
		log.Println(err)
		return status, err
	}

	err = json.Unmarshal(content, &status)
	if err != nil {
		log.Println(err)
		return status, err
	}

	return status, nil
}

func SetStatus(status models.Status) error {
	content, err := json.Marshal(status)
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile(SOURCE_FILE, content, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
