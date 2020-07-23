package models

import (
	"fmt"
	"log"
	"time"
	uuid2 "github.com/tim/chitchat/common/uuid"
)

type User struct {
	Id int
	Uuid string
	Name string
	Email string
	Passwd string
	CreatedAt time.Time

}


func (user *User) CreateSession()(session Session,err error){
	statment := "insert into sessions (uuid, email, user_id, created_at) values (?, ?, ?, ?) "
	stmtin,err := Db.Prepare(statment)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmtin.Close()

	uuid := uuid2.CreateUuid()
	stmtin.Exec(uuid,user.Email,user.Id,time.Now())

	stmtout, err := Db.Prepare(" Select id,uuid,email,user_id,created_at from sessions where uuid = ?")
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// create user
func (user *User) Create() (err error){
	statment := "Insert into user (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)"
	stmtin , err := Db.Prepare(statment)
	if err != nil {
		fmt.Println("inser error",err.Error())
		return
	}
	defer stmtin.Close()
	uuid := uuid2.CreateUuid()

	stmtin.Exec(uuid,user.Name,user.Email,user.Passwd,time.Now())

	// check the user create

	stmtout ,err := Db.Prepare("Select id,uuid,created_at from user where uuid = ?")
	defer stmtout.Close()

	if err != nil {
		return
	}
	err = stmtout.QueryRow(uuid).Scan(&user.Id,&user.Uuid,&user.CreatedAt)

	return
}

func UserFindByEmail(email string) (user User,err error){
	statment := "Select * from user where email = ?"

	user = User{}
	err = Db.QueryRow(statment,email).Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.Passwd,&user.CreatedAt)
	return
}

func (user *User)CreateThread(topic string) (thread Threads,err error){
	statment := "Insert into threads (uuid,topic,user_id,created_at) values(?,?,?,?)"
	stmtin,err := Db.Prepare(statment)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmtin.Close()
	uuid := uuid2.CreateUuid()
	stmtin.Exec(uuid,topic,user.Id,time.Now())

	stmtout,err := Db.Prepare("Select * from threads where uuid = ? ")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&thread.Id,&thread.Uuid,&thread.Topic,&thread.UserId,&thread.CreatedAt)
	return
}

func (user *User) CreatePosts(thread Threads,body string)(post Posts,err error){
	statment := "Insert into posts (uuid,body,user_id,thread_id,created_at) values(?,?,?,?,?)"
	stmtin,err := Db.Prepare(statment)
	if err != nil {
		return
	}
	defer stmtin.Close()
	uuid := uuid2.CreateUuid()
	stmtin.Exec(uuid,body,user.Id,thread.Id,time.Now())

	err = Db.QueryRow("select * from posts where uuid = ?",uuid).Scan(
		&post.Id,&post.Uuid,&post.Body,&post.UserId,&post.ThreadId,&post.CreatedAt)
	return
}

func UserByUuid(uuid string)(user User) {
	user = User{}
	Db.QueryRow("select * from user where uuid = ?",uuid).Scan(
		&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Passwd, &user.CreatedAt)
	return
}
