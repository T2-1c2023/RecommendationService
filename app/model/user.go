package model

import "encoding/json"

type Interest struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
}

type User struct {
	Id            int        `json:"id"`
	Fullname      string     `json:"fullname"`
	Mail          string     `json:"mail"`
	PhoneNumber   string     `json:"phone_number"`
	IsAthlete     bool       `json:"is_athlete"`
	IsTrainer     bool       `json:"is_trainer"`
	IsVerified    bool       `json:"is_verified"`
	IsAdmin       bool       `json:"is_admin"`
	PhotoId       string     `json:"photo_id"`
	ExpoPushToken string     `json:"expo_push_token"`
	Latitude      string     `json:"latitude"`
	Longitude     string     `json:"longitude"`
	Interests     []Interest `json:"interests"`
}

func NewUserFromUserInfo(userInfo string) (User, error) {
	var user User
	err := json.Unmarshal([]byte(userInfo), &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
