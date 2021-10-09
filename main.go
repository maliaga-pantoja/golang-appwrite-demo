package main

import (
	Collection "github.com/maliaga-pantoja/golang-appwrite-demo/src/database/collection"
	Document "github.com/maliaga-pantoja/golang-appwrite-demo/src/database/document"
	Util "github.com/maliaga-pantoja/golang-appwrite-demo/src/util"
)

var collection Collection.CollectionEntity
var document Document.DocumentEntity

func init() {
	// -- initializen vars
	collection.Init()
}
func main() {

	// ---------
	collectionInput := createCollectionInput("demo7Collection")
	createCollectionResult := collection.CreateCollection(collectionInput)
	Util.Present{}.CreateCollection(createCollectionResult)
	// ----------
	collections := collection.GetCollections()
	Util.Present{}.ListCollections(collections)
	// ----------
	collectionLen := len(collections.Collections) - 1
	document.Init(collections.Collections[collectionLen].Id)
	createDocumentResult := document.CreateDocument(`{"data": {"key2": "value4"} }`)
	Util.Present{}.CreateDocument(createDocumentResult)
}

func createCollectionInput(collectionName string) Collection.CollectionCreateInput {
	inputRule := Collection.CollectionCreateInputRule{
		Label:    "label",
		Key:      "key2",
		Type:     "text",
		Default:  "no set",
		Required: false,
		Array:    false,
	}
	var ruleGroup []Collection.CollectionCreateInputRule
	ruleGroup = append(ruleGroup, inputRule)

	collectionInput := Collection.CollectionCreateInput{
		Name:  collectionName,
		Read:  []string{"*"},
		Write: []string{"*"},
		Rules: ruleGroup,
	}
	return collectionInput
}
