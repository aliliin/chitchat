package models

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	. "github.com/aliliin/chitchat/config"
	_ "github.com/go-sql-driver/mysql"
)

// Db 变量代表数据库连接池
var Db *sql.DB

/*通过 init 方法在 Web 应用启动时自动初始化数据库连接*/
func init() {
	var err error
	driver := ViperConfig.Db.Driver
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true", ViperConfig.Db.User, ViperConfig.Db.Password,
		ViperConfig.Db.Address, ViperConfig.Db.Database)

	Db, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	////设置以下内容的四个最高有效位（第12至15位）
	//将time_hi_and_version字段更改为4位版本号。
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
