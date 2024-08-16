package router

import (
	"github.com/1Lindo/metrics_collector/internal/server/controller"
	"log"
	"net/http"
)

func InitServer(c controller.Controller) {
	err := http.ListenAndServe("localhost:8080", newRouter(c))
	if err != nil {
		log.Printf("Init server error:%v", err)
	}
}

func newRouter(c controller.Controller) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", c.UpdateMetrics)
	mux.HandleFunc("/metrics", c.GetMetrics)

	return mux
}
