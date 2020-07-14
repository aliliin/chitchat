package handlers

import (
	"github.com/aliliin/chitchat/models"
	"net/http"
)

/** 论坛首页路由处理器方法 */
func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "auth.navbar", "index")
		}

	}
}

/** 错误路由处理方法 */
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}
