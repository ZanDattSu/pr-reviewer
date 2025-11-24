package model

type Team struct {
	TeamName string       `json:"team_name"`
	Members  []TeamMember `json:"members"`
}

type TeamMember struct {
	UserUUID string `json:"user_uuid"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}
