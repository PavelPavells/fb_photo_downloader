package structs

type Data struct {
	CanUpload   bool     `json:"can_upload"`
	Comments    Comments `json:"comments"`
	Count       int      `json:"count"`
	CoverPhoto  string   `json:"cover_photo"`
	CreatedTime string   `json:"created_time"`
	From        From     `json:"from"`
	ID          string   `json:"id"`
	Link        string   `json:"link"`
	Name        string   `json:"name"`
	Privacy     string   `json:"privacy"`
	Type        string   `json:"type"`
	UpdatedTime string   `json:"updated_time"`
}
