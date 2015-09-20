package models

import (
	"os"
	"testing"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nassor/studies-blevesearch/conn"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
	testIdx   = "test.bleve"
	dbFile    = "test.sqlite3.db"
)

func TestIndexing(t *testing.T) {
	_, eventList := dbCreate()
	idx := idxCreate()

	err := eventList[0].Index(idx)
	if err != nil {
		t.Error("Wasn't possible create the index", err, ballotX)
	} else {
		t.Log("Should create an event index", checkMark)
	}

	idxDestroy()
	dbDestroy()
}

func TestFindByAnything(t *testing.T) {
	db, eventList := dbCreate()
	idx := idxCreate()
	indexEvents(idx, eventList)

	// We are looking to an Event with some string which match with dotGo
	query := bleve.NewMatchQuery("dotGo")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := idx.Search(searchRequest)
	if err != nil {
		t.Error("Something wrong happen with the search", err, ballotX)
	} else {
		t.Log("Should search the query", checkMark)
	}

	if searchResult.Total != 1 {
		t.Error("Only 1 result are expected, got ", searchResult.Total, ballotX)
	} else {
		t.Log("Should return only one result", checkMark)
	}

	event := &Event{}
	db.First(&event, &searchResult.Hits[0].ID)
	if event.Name != "dotGo 2015" {
		t.Error("Expected \"dotGo 2015\", Receive: ", event.Name)
	} else {
		t.Log("Should return an event with the name equal a", event.Name, checkMark)
	}

	idxDestroy()
	dbDestroy()
}

// indexEvents add the eventList to the index
func indexEvents(idx bleve.Index, eventList []Event) {
	for _, event := range eventList {
		event.Index(idx)
	}
}

// fill the database with some data
func fillDatabase(db *gorm.DB) []Event {
	eventList := []Event{
		{1, "dotGo 2015", "The European Go conference", "Paris", "http://www.dotgo.eu/", time.Date(2015, 11, 19, 9, 0, 0, 0, time.UTC), time.Date(2015, 11, 19, 18, 30, 0, 0, time.UTC)},

		{2, "GopherCon INDIA 2016", "The Go Conference in India", "Bengaluru", "http://www.gophercon.in/", time.Date(2016, 2, 19, 0, 0, 0, 0, time.UTC), time.Date(2016, 2, 20, 23, 59, 0, 0, time.UTC)},

		{3, "GopherCon 2016", "GopherCon, It is the largest event in the world dedicated solely to the Go programming language. It's attended by the best and the brightest of the Go team and community.", "Denver", "http://gophercon.com/", time.Date(2016, 7, 11, 0, 0, 0, 0, time.UTC), time.Date(2016, 7, 13, 23, 59, 0, 0, time.UTC)},
	}

	// inserting the events
	for _, event := range eventList {
		db.Create(&event)
	}

	return eventList
}

func idxCreate() bleve.Index {
	idx, _ := conn.Bleve(testIdx)
	return idx
}

// create a SQLite3 database file, create an events table and fill with some data.
func dbCreate() (gorm.DB, []Event) {
	db, _ := gorm.Open("sqlite3", dbFile)
	db.DropTableIfExists(&Event{})
	db.CreateTable(&Event{})
	eventList := fillDatabase(&db)
	return db, eventList
}

func idxDestroy() {
	os.RemoveAll(testIdx)
}

func dbDestroy() {
	os.RemoveAll(dbFile)
}
