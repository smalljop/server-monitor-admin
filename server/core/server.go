package core

import (
	"net/http"
	"server-monitor-admin/initialize"
	"time"
)

/**
启动服务
*/
func RunServer() {

	routers := initialize.Routers()

	s := &http.Server{
		Addr:           ":8086",
		Handler:        routers,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
