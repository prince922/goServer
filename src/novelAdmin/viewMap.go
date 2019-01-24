package main

import (
	"gopkg.in/mgo.v2/bson"
)

// mongo id 转换成 string 类型
func Id2string(id bson.ObjectId) string {
	return id.Hex()
}

//func GetImageUrl(path string)string{
//	return
//}
