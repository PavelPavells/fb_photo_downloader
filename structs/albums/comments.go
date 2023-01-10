package structs

type Comments struct {
	Data []struct {
		CanRemove   bool   `json:"can_remove"`
		CreatedTime string `json:"created_time"`
		From        From   `json:"from"`
		ID          string `json:"id"`
		LikeCount   int    `json:"like_count"`
		Message     string `json:"message"`
		UserLikes   bool   `json:"user_likes"`
	} `json:"data"`
	Paging Paging `json:"paging"`
}
