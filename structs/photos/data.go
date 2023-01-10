package structs

type Data struct {
	Data []struct {
		ID          string `json:"id"`
		CreatedTime string `json:"created_time"`
		From        From   `json:"from"`
		Height      int    `json:"height"`
		Icon        string `json:"icon"`
		Images      Images `json:"images"`
		Link        string `json:"link"`
		Name        string `json:"name"`
		Picture     string `json:"picture"`
		Place       Place  `json:"place"`
		Source      string `json:"source"`
		UpdatedTime string `json:"updated_time"`
		Width       int    `json:"width"`
		Tags        Tags   `json:"tags"`
		Likes       Likes  `json:"likes"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}
