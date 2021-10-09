package main

import (
	Database "github.com/maliaga-pantoja/golang-appwrite-demo/src/database"
	Util "github.com/maliaga-pantoja/golang-appwrite-demo/src/util"
)

func main() {
	v := Database.CollectionEntity{}
	v.Init()
	// ---------
	input := createCollectionInput("demo2Collection")
	createResult := v.CreateCollection(input)
	Util.Present{}.CreateCollection(createResult)
	// ----------
	collections := v.GetCollections()
	Util.Present{}.ListCollections(collections)
	// ----------
}

func createCollectionInput(collectionName string) Database.CollectionCreateInput {
	inputRule := Database.CollectionCreateInputRule{
		Label:    "label",
		Key:      "key",
		Type:     "text",
		Default:  "no set",
		Required: false,
		Array:    false,
	}
	var ruleGroup []Database.CollectionCreateInputRule
	ruleGroup = append(ruleGroup, inputRule)

	collectionInput := Database.CollectionCreateInput{
		Name:  collectionName,
		Read:  []string{"*"},
		Write: []string{"*"},
		Rules: ruleGroup,
	}
	return collectionInput
}
