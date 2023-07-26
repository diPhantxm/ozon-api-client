package notifications

import (
	"net/http"
)

type NotificationServer struct{}

func NewNotificationServer() *NotificationServer {
	return &NotificationServer{}
}

func (ns *NotificationServer) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ns.handler)
	return http.ListenAndServe("", mux)
}

func (ns *NotificationServer) handler(rw http.ResponseWriter, req *http.Request) {

}
