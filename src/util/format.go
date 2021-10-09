package util

import (
	"fmt"

	Collection "github.com/maliaga-pantoja/golang-appwrite-demo/src/database/collection"
)

type Present struct{}

func (p Present) ListCollections(collections Collection.CollectionResponse) {
	fmt.Println("List of available collections:")
	for _, v := range collections.Collections {
		fmt.Printf("collection Name: %s \n", v.Name)
	}
}

func (p Present) CreateCollection(status int) {
	fmt.Println("Creating a new collection")
	fmt.Printf("Server response with status %d \n", status)
}

// ------------------------------
func (p Present) CreateDocument(status int) {
	fmt.Println("Creating a new document")
	fmt.Printf("Server response with status %d \n", status)
}
