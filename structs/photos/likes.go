package structs

type Likes struct {
	Data []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}
