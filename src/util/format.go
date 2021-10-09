package util

import (
	"fmt"

	Database "github.com/maliaga-pantoja/golang-appwrite-demo/src/database"
)

type Present struct{}

func (p Present) ListCollections(collections Database.CollectionResponse) {
	fmt.Println("List of available collections:")
	for _, v := range collections.Collections {
		fmt.Printf("collection Name: %s \n", v.Name)
	}
}

func (p Present) CreateCollection(status int) {
	fmt.Println("Creating a new collection")
	fmt.Printf("Server response with status %d \n", status)
}