package routes

import (
	"github.com/aliliin/chitchat/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/** 声明 WebRoute 切片存放所有的 web 路由 */
type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
}

/** 返回一个 mux.Router 类型指针，从而可以当作处理器使用 */
func NewRouter() *mux.Router {

	/** 创建  mux.Router 路由示例 */
	router := mux.NewRouter().StrictSlash(true)

	// 遍历 web.go 中定义的所有 rouoter
	for _, route := range webRoutes {
		// 将每个 web 路由应用到路由器
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
