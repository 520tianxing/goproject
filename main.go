package main

import (
	"goproject/src/go_code/ginproject/router"
)

var r *router.Router

func init() {
	r = router.NewRouter()
}

func main() {
	r.Start()
	r.Run("127.0.0.1:8081")
	// 监听并在 0.0.0.0:8080 上启动服务
}
