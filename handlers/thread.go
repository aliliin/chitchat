package handlers

import (
	"fmt"
	"github.com/aliliin/chitchat/models"
	"net/http"
)

/** 创建群组页面 GET /threads/new */
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "new.thread")
	}
}

/** 执行群组创建逻辑 POST /thread/create */
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			fmt.Println(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			fmt.Println(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

/** 通过 ID 渲染指定群组页面 GET /thread/read */
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		fmt.Println("Cannot read thread")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}