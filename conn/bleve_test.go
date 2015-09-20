package conn

import (
	"os"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
	testIdx   = "test.index"
)

func TestCreateIndex(t *testing.T) {
	should := "Should be able to create an index if don't exists"
	idx, err := Bleve(testIdx)
	if err != nil || idx == nil {
		t.Error(should, ballotX, err)
	} else {
		t.Log(should, checkMark)
	}

	idxDestroy()
}

func TestOpenIndex(t *testing.T) {
	should := "Should be able to open an index if don't exists"

	Bleve("test.index")
	idx, err := Bleve(testIdx)
	if err != nil || idx == nil {
		t.Error(should, ballotX, err)
	} else {
		t.Log(should, checkMark)
	}

	idxDestroy()
}

func idxDestroy() {
	os.RemoveAll(testIdx)
}
