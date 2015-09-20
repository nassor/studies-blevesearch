package models

import (
	"time"

	"github.com/blevesearch/bleve"
)

// Event is an event! wow! ;D
type Event struct {
	ID          int
	Name        string
	Description string
	Local       string
	Website     string
	Start       time.Time
	End         time.Time
}

// Index is used to add the event in the bleve index.
func (e *Event) Index(index bleve.Index) error {
	err := index.Index(string(e.ID), e)
	return err
}
