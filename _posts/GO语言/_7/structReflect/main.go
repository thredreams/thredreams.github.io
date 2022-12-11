package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string "用户名称"
	UserId   int    `json:"user_id" bson:"b_user_id"`
}

func main() {
	u := User{}
	rt := reflect.TypeOf(u)

	// 遍历结构体所有成员
	for i := 0; i < rt.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := rt.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}

	//分解tag 里的值
	field, ok := rt.FieldByName("UserId")
	if ok {
		fmt.Printf(" tag json value: %v\n", field.Tag.Get("json"))
		fmt.Printf(" tag bson value: %v\n", field.Tag.Get("bson"))
	}

}
