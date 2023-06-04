package model

type User struct {
	Id            int     `json:"id"`
	Fullname      string  `json:"fullname"`
	Mail          string  `json:"mail"`
	PhoneNumber   string  `json:"phone_number"`
	IsAthlete     bool    `json:"is_athlete"`
	IsTrainer     bool    `json:"is_trainer"`
	IsVerified    bool    `json:"is_verified"`
	PhotoId       string  `json:"photo_id"`
	ExpoPushToken string  `json:"expo_push_token"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}
