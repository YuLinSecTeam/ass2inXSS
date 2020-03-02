package core

import (
	"strings"
)

//=====flags添加其他请求头
func (i *Headers) String() string {
	return ""
}
func (i *Headers) Set(value string) error {
	*i = append(*i, value)
	return nil
}

//=====flags添加请求cookie

func (c *Cookies) String() string {
	return ""
}
func (c *Cookies) Set(value string) error {
	*c = strings.Split(value,";")
	return nil
}
