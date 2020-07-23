package models

import "time"

type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}

func (session *Session) Check() (valid bool, err error){
	statment := "select id,uuid,email,user_id,created_at from sessions where uuid = ?"

	err = Db.QueryRow(statment,session.Uuid).Scan(
		&session.Id,&session.Uuid,&session.Email,&session.UserId,&session.CreatedAt)

	if err != nil {
		valid = false
		return
	}

	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteByUuid() (err error){
	statment := "Delete from sessions where uuid = ?"
	stmtin,err := Db.Prepare(statment)
	if err != nil {
		return
	}
	defer stmtin.Close()
	_,err = stmtin.Exec(session.Uuid)
	return
}
//
func (session *Session) SessionUser() (user  User,err error){
	statment := "Select id,uuid,name,email,created_at from user where id = ? "
	err = Db.QueryRow(statment,session.UserId).Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)
	return
}

func (session *Session) SessionTruncate() (err error ){
	statment := "Delete from sessions"
	stmt, err := Db.Prepare(statment)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return

}
