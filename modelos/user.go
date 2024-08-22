package modelos

import "time"

type User struct {
	Id       	int
	Nombre   	string
	CreatedAt 	time.Time
	Status 		bool
}

func (user *User) AddUser(id int,  name string, createdAt time.Time, status bool) {
	user.Id = id
	user.Nombre = name
	user.CreatedAt = createdAt
	user.Status = status
}