package database

// -----------Inputs for create collections -----------
type CollectionCreateInputRule struct {
	Label    string      `json:"label"`
	Key      string      `json:"key"`
	Type     string      `json:"type"` // https://appwrite.io/docs/rules#types
	Default  interface{} `json:"default"`
	Required bool        `json:"required"`
	Array    bool        `json:"array"`
}
type CollectionCreateInput struct {
	Name  string                      `json:"name"`
	Read  []string                    `json:"read"`  // https://appwrite.io/docs/permissions
	Write []string                    `json:"write"` // https://appwrite.io/docs/permissions
	Rules []CollectionCreateInputRule `json:"rules"` // https://appwrite.io/docs/rules
}

// ------------- end ----------
// -----------Inputs for list collections -----------
type Collection struct {
	Id        string `json:"$id"`
	Name      string `json:"name"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}
type CollectionResponse struct {
	Sum         int          `json:"sum"`
	Collections []Collection `json:"collections"`
}

// ------------- end ----------
