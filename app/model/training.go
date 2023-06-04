package model

type Training struct {
	Id          int     `json:"id"`
	TrainerId   int     `json:"trainer_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Severity    int     `json:"severity"`
	TypeId      int     `json:"type_id"`
	Type        string  `json:"type"`
	Blocked     bool    `json:"blocked"`
	Score       float32 `json:"score"`
}
