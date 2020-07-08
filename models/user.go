package models

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

/* 为现有用户创建一个新会话 */
func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid,email,user_id,created_at) values(?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Email, user.Id, time.Now())

	stmtout, err := Db.Prepare("select id,uuid,email,user_id,created_at from sessions where uuid = ?")

	if err != nil {
		return
	}
	defer stmtout.Close()
	/* 使用QueryRow返回一行并将返回的ID扫描到Session结构中 */
	err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

/**获取现有用户的会话*/
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id,uuid,email,user_id,create_at FROM session WHERE user_id = ?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

/**创建一个新用户，将用户信息保存到数据库中*/
func (user *User) Create() (err error) {
	statement := "insert into users(uuid,name,email,password,created_at) value(?,?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())

	stmtout, err := Db.Prepare("select id,uuid,created_at from users where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

/** 删除一个用户 */
func (user *User) Delete() (err error) {
	statement := "delete from users where id = ?"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()
	_, err = stmtin.Exec(user.Id)
	return
}

/** 从数据库中删除所有用户 */
func (user *User) DeleteUserAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

/** 更新数据库中的用户信息 */
func (user *User) Update() (err error) {
	statement := "update users set name = ?,email = ? where id = ?"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()
	_, err = stmtin.Exec(user.Name, user.Email, user.Id)
	return
}

/** 获取数据库中的所有用户并返回 */
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id,uuid,name,email,password,created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

/** 根据电子邮件获得一个用户 */
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id,uuid,name,email,password,created_at FROM users where email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

/** 根据UUID获得一个用户 */
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id,uuid,name,email,password,created_at FROM users where uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

/** 创建一个新主题 */
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "insert into threads (uuid,topic,user_id,created_at) values(?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()
	uuid := createUUID()
	stmtin.Exec(uuid, topic, user.Id, time.Now())
	stmtout, err := Db.Prepare("select id,uuid,topic,user_id,created_at from threads where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

/** 在主题中创建新帖子 */
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, body, user.Id, conv.Id, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, body, user_id, thread_id, created_at from posts where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// 使用 QueryRow 返回一行并将返回的ID扫描到 Session 结构中
	err = stmtout.QueryRow(uuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}
