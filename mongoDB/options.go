package mongoDB

type Options struct {
	Connection string `json:"connection"`
	Name       string `json:"name"`
	Collection string `json:"collection"`
}
