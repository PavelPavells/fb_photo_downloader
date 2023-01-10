package structs

type Cover struct {
	CoverID string `json:"cover_id"`
	OffsetX int    `json:"offset_x"`
	OffsetY int    `json:"offset_y"`
	Source  string `json:"source"`
	ID      string `json:"id"`
}
