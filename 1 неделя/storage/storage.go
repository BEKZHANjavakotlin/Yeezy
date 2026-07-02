package storage

import (
	"encoding/json"
	"os"
	"week1/models"
)

func LoadEntries() ([]models.TimeEntry, error) {
	var entries []models.TimeEntry
	data, err := os.ReadFile("traker.json")
	if os.IsNotExist(err) {
		return nil, nil
	}
	json.Unmarshal(data, &entries)
	return entries, nil
}
func SaveEntries(entries []models.TimeEntry) error {
	data, err := json.Marshal(entries)
	if err != nil {
		return err
	}
	return os.WriteFile("traker.json", data, 0644)
}
