package models

import "time"

type Posts struct {
	Id int
	Uuid string
	Body string
	UserId int
	ThreadId int
	CreatedAt time.Time
}

func (post *Posts) CreatedAtDate() string{
	return post.CreatedAt.Format("2006-01-02 15:04:05")
}

func (post *Posts) User()(user User){
	user = User{}
	Db.QueryRow("select id,uuid,name,email,created_at from user where id = ?",post.UserId).Scan(
		&user.Id,&user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
