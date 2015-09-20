package conn

import "github.com/blevesearch/bleve"

var bleveIdx bleve.Index

// Bleve connect or create the index persistence
func Bleve(indexPath string) (bleve.Index, error) {

	// with bleveIdx isn't set...
	if bleveIdx == nil {
		var err error
		// try to open de persistence file...
		bleveIdx, err = bleve.Open(indexPath)
		// if doesn't exists or something goes wrong...
		if err != nil {
			// create a new mapping file and create a new index
			mapping := bleve.NewIndexMapping()
			bleveIdx, err = bleve.New(indexPath, mapping)
			if err != nil {
				return nil, err
			}
		}
	}

	// return de index
	return bleveIdx, nil
}
