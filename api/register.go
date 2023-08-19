package api

import (
	"ms-go-blog/common"
	"ms-go-blog/logger"
	"ms-go-blog/service"
	"net/http"
)

func (*Api) Register(w http.ResponseWriter, r *http.Request) {
	//接收用户名和密码 返回 对应的json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	err := service.Register(userName, passwd)
	if err != nil {
		logger.Error("注册失败", err)
		common.Error(w, err)
	}
	data := 0
	common.Success(w, data)
}
