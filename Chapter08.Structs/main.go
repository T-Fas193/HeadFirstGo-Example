package main

import "fmt"
import "./struct"

func main() {

	var user _struct.User

	var department _struct.Dept

	user.Email = "33"
	user.Dept.DeptName = "financial"
	department.DeptName = "develop"

	fmt.Printf("%#v\n", user)
	fmt.Printf("%#v\n", department)

	var user2 _struct.EmbedUser
	user2.DeptName = "success"
	changeDept(&user2)
	fmt.Printf("%#v\n", user2)

	fmt.Printf("%#v\n", addNewUser())
}

func changeDept(user *_struct.EmbedUser) {
	user.DeptId = 1123
}

func addNewUser() *_struct.EmbedUser {
	user := _struct.EmbedUser{Email: "123@123.com"}
	return &user
}
