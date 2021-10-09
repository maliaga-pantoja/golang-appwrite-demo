package database
import (
	"net/http"
)
// -----------------------
type Collection struct {
	Id string `json:"$id"`
	Name string `json:"name"`
	CreatedAt int `json:"createdAt"`
	UpdatedAt int `json:"updatedAt"`
}
type CollectionResponse struct {
	Sum int `json:"sum"`
	Collections  Collection[] `json:"collections"`
}
// ------------------------
type CollectionEntity struct{
	EndPoint string
}
func (c *CollectionEntity) Init() {
	endpoint := "v1/database/collections"
	c.EndPoint = fmt.Sprintf("%s:%s/%s", os.GetEnv("BASEPATH"), os.GetEnv("PORT"), endpoint)
}
func(c *CollectionEntity) GetCollections () (response CollectionResponse) {
	// build request
	client := &http.Client{}
	req,reqErr := http.NewRequest("GET", c.EndPoint)
	if reqErr != nil {
		panic(reqErr.Error())
	}
	req.Header.Add("Content-Type": "application/json")
	req.Header.Add("X-Appwrite-Project", os.GetEnv("PROJECT_ID"))
	req.Header.Add("X-Appwrite-key", os.GetEnv("TOKEN"))
	// execute request
	resp, respErr := client.Do(req)
	if respErr != nil {
		panic(resErr.Error())
	}
	defer resp.Body.Close()
	// parse response
	var collectionResponse CollectionResponse{}
	parseErr := json.NewDecoder(resp.Body).Decode(&collectionResponse)
	if parseErr != nil {
		panic(parseErr.Error())
	}
	return collectionResponse
}
