package models

import "time"

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

/** 检查会话在数据库中是否有效 */
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where uuid = ?;", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

/** 从数据库删除会话 */
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.Uuid)
	return
}

/** 删除所有session记录 */
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}
