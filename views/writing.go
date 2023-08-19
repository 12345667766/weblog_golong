package views

import (
	"ms-go-blog/common"
	"ms-go-blog/dao"
	"ms-go-blog/logger"
	"ms-go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}

func (*HTMLApi) RemoveArtical(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI
	pIDSlice := strings.Split(url, "=")
	pIdStr := pIDSlice[len(pIDSlice)-1]
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		logger.Error(err)
	}
	err = dao.DeletePost(pId)
	if err != nil {
		logger.Error("数据库操作删除文章ID %d 失败", pId)
	}
	HTML.Index(w, r)
}
