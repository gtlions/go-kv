// Package model ...
// Date: 2019/05/02
// Author: gtlions.l@qq.com
package model

// Response request
type Response struct {
	Code  int
	Error string
	Data  interface{}
}

// Cache ,storge KV
type Cache struct {
	Data map[string]string
}

type KV struct {
	Key   string `form:"key" json:"key"`
	Value string `form:"value" json:"value"`
}
