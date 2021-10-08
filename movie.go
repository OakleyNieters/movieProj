package entities

import "github.com/google/uuid"

type Movies struct {
	ID          string   `json:"ID"`
	Title       string   `json:"Title"`
	Genre       []string `json:"Genre"`
	Description string   `json:"Description"`
	Director    string   `json:"Director"`
	Actors      []string `json:"Actors"`
	Rating      float32  `json:"Rating"`
}

func (m *Movies) SetId() {
	m.ID = uuid.New().String()
}
