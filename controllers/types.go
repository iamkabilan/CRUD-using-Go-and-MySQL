package controllers

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRedis struct {
	User
	Timestamp int64 `json:"timestamp"`
}

type Response struct {
	Message string `json:"message"`
}
