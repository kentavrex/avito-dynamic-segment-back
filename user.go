package app

type User struct {
	Id   int    `json:"-"`
	Name string `json:"name" binding:"required"`
}
