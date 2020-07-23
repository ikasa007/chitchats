package models

import (
	"log"
	"time"
)

type Threads struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreatedAt time.Time
}

func Thread() (threads []Threads,err error) {
	statment := "Select id,uuid,topic,user_id,created_at from threads order by created_at desc "

	rows,err := Db.Query(statment)
	if err != nil {
		log.Fatal(err)
		return
	}

	for rows.Next() {
		conv := Threads{}
		if err = rows.Scan(&conv.Id,&conv.Uuid,&conv.Topic,&conv.UserId,&conv.CreatedAt);err != nil {
			return
		}
		threads = append(threads,conv)
	}

	rows.Close()
	return
}

func ThreadByUuid(uuid string)(conv Threads,err error){
	conv = Threads{}
	statment := " select * from threads where uuid = ? "
	err = Db.QueryRow(statment,uuid).Scan(&conv.Id,
		&conv.Uuid,&conv.Topic,&conv.UserId,&conv.CreatedAt)
	return

}

func (thread *Threads) User() (user User) {
	user = User{}
	statment := "SELECT id, uuid, name, email, created_at FROM user WHERE id = ? "
	Db.QueryRow(statment,thread.UserId).Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)
	return
}

func (thread *Threads) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (thread *Threads) NumReplies() (count int){
	statment := "Select count(*) from posts where thread_id = ?"
	rows,err := Db.Query(statment,thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil{
			return
		}
	}
	rows.Close()
	return
}


func (thread Threads) Posts()(posts []Posts,err error) {
	rows,err := Db.Query("select id, uuid, body, user_id, thread_id, created_at from posts where thread_id = ?",thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Posts{}
		if err = rows.Scan(&post.Id,&post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt);err != nil {
			return
		}
		posts = append(posts,post)
	}
	return

}
