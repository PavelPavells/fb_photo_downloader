package main

import structs "structs/user"

type FBUser struct {
	ID                string           `json:"id"`
	About             string           `json:"about"`
	Awards            string           `json:"awards"`
	CanPost           bool             `json:"can_post"`
	Category          string           `json:"category"`
	Checkins          int              `json:"checkins"`
	Cover             structs.Cover    `json:"cover"`
	HasAddedApp       bool             `json:"has_added_app"`
	IsComunityPage    bool             `json:"is_community_page"`
	IsPublished       bool             `json:"is_published"`
	Likes             int              `json:"likes"`
	Link              string           `json:"link"`
	Location          structs.Location `json:"location"`
	Name              string           `json:"name"`
	Parking           structs.Parking  `json:"parking"`
	PersonalInfo      string           `json:"personal_info"`
	PersonalInterests string           `json:"personal_interests"`
	TalkingAboutCount int              `json:"talking_about_count"`
	UserName          string           `json:"username"`
	WereHereCount     int              `json:"were_here_count"`
}
