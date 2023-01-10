package structs

type Tags struct {
	Data []struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		CreatedTime string  `json:"created_time"`
		X           float64 `json:"x"`
		Y           float64 `json:"y"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}
