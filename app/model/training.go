package model

type Training struct {
	Id          uint   `json:"id"`
	TrainerId   int    `json:"trainer_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Severity    int    `json:"severity"`
	TypeId      int    `json:"type_id"`
	Type        string `json:"type"`
	Blocked     bool   `json:"blocked"`
	Score       string `json:"score"`
}
