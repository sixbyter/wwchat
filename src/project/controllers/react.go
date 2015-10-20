package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sixbyte/framework"
)

type ReactController struct {
	comments string
}

var React = &ReactController{
	`[{"author": "### Pete Hunt", "text": "This is one comment"},{"author": "### Jordan Walke", "text": "This is *another* comment"}]`,
}

/**
 * 你好首页
 * @param  {[type]} react *ReactController) Nihao(rw http.ResponseWriter, req *http.Request [description]
 * @return {[type]}       [description]
 */
func (react *ReactController) Nihao(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("HI", "写在body前面呀")
	rw.Header().Set("Set-Cookie", "a=2&b=3")
	b := framework.ViewByte("/react/nihao.html")
	rw.Write(b)
}

/**
 * 获取所有评论
 * @param  {[type]} react *ReactController) GetComment(rw http.ResponseWriter, req *http.Request [description]
 * @return {[type]}       [description]
 */
func (react *ReactController) GetComment(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.Write([]byte(react.comments))
}

/**
 * 提交评论
 * @param  {[type]} react *ReactController) PostComment(rw http.ResponseWriter, req *http.Request [description]
 * @return {[type]}       [description]
 */
func (react *ReactController) PostComment(rw http.ResponseWriter, req *http.Request) {

	newComment := make(map[string]string)
	req.ParseForm()
	newComment["author"] = req.Form.Get("author")
	newComment["text"] = req.Form.Get("text")

	var f []map[string]string
	json.Unmarshal([]byte(react.comments), &f)
	f = append(f, newComment)
	b, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err.Error())
	}
	react.comments = string(b)

	fmt.Println(react.comments)
	rw.Write(b)
}

func (react *ReactController) AddonsAnimation(rw http.ResponseWriter, req *http.Request) {
	b := framework.ViewByte("/react/animation.html")
	rw.Write(b)
}

func (react *ReactController) AmazeUi(rw http.ResponseWriter, req *http.Request) {
	b := framework.ViewByte("/react/amaze.html")
	rw.Write(b)
}
