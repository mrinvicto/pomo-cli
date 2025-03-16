package models

import (
	"encoding/json"
	"log"
)

type Session struct {
	ID        int
	StartTime string
	Title     string
	Tags      []string
	Duration  int
	Status    int
}

func (s Session) GetTagsJSON() []byte {

	tagsJSON, err := json.Marshal(s.Tags)
	if err != nil {
		log.Fatal(err)
	}

	return tagsJSON
}
