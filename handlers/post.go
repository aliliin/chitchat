package handlers

import (
	"fmt"
	"github.com/aliliin/chitchat/models"
	"net/http"
)

/** 指定群组下创建新主题 POST /thread/post */
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger("Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger("Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)

		if err != nil {
			error_message(writer, request, "cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger("Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
