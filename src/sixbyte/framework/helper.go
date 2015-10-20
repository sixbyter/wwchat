package framework

import (
	"fmt"
	"html/template"
	"io/ioutil"
)

func SetConfig(k, v string) error {
	App.Configer.Configs[k] = v
	return nil
}

func GetConfig(k string) string {
	return App.Configer.Configs[k]
}

func ViewByte(path string) []byte {
	b, err := ioutil.ReadFile(App.Viewer.ViewDir + path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return b
}

func ViewTemplate(path string) *template.Template {
	t, err := template.ParseFiles(App.Viewer.ViewDir + path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return t
}
