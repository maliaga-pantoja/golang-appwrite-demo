package document

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// -------------- Principal struct ----------
type DocumentEntity struct {
	endPoint     string
	CollectionId string
}

// Method to initialize struct, set default values
func (c *DocumentEntity) Init(collectionId string) {
	endpoint := fmt.Sprintf("v1/database/collections/%s/documents", collectionId)
	url := fmt.Sprintf("%s:%s/%s", os.Getenv("BASEPATH"), os.Getenv("PORT"), endpoint)
	c.endPoint = url
}
func (c *DocumentEntity) CreateDocument(input string) int {
	client := &http.Client{}
	// bytes.NewBuffer(data)
	req, reqErr := http.NewRequest("POST", c.endPoint, bytes.NewBuffer([]byte(input)))
	if reqErr != nil {
		panic(reqErr)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Appwrite-Project", os.Getenv("PROJECT_ID"))
	req.Header.Add("X-Appwrite-key", os.Getenv("TOKEN"))

	resp, respErr := client.Do(req)
	if respErr != nil {
		panic(respErr)
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
