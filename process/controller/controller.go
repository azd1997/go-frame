package controller

import (
	"net/http"
	"time"
)

type ServerTimeReq struct {}

type ServerTimeResp struct {
	ServerTime int64 `json:"server_time"`
}

func GetServerTime() (code int, msg string, data ServerTimeResp) {
	data = ServerTimeResp{ServerTime:time.Now().Unix()}
	return http.StatusOK, "", data
}
