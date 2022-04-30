package models

type Pair struct{
	User1 User `json:"user1"`
	User2 User `json:"user2"`
}

type Response struct{
	FromId uint	`json:"who_likes"`
	ToId   uint	`json:"who_is_liked"`
}