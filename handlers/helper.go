package handlers

import (
	"errors"
	"fmt"
	"github.com/aliliin/chitchat/models"
	"html/template"
	"net/http"
)

/** 全局辅助函数 */

/** 通过 Cookie 判断用户是否登录 */
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

/** 解析 HTML 模版（应对应需要传入的多个模版文件的情况，避免重复编写模版代码) */
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

/** 生成响应的 HTML */
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

/** 返回版本号 */
func Version() string {
	return "0.1"
}
