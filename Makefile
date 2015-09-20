all: tests

tests:
	@ go test -v ./conn
	@ go test -v ./models

deps:
	go get github.com/blevesearch/bleve/... # bleve package
	go get github.com/mattn/go-sqlite3      # sqlite3 package
	go get github.com/jinzhu/gorm           # orm package
