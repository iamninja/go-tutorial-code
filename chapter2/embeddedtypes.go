package main

import "fmt"

type User struct {
	Name string
}
func (u User) IsAdmin() bool {
	return false
}
func (u User) DisplayName() string {
	return u.Name
}

type Admin struct {
	User
}
func (a Admin) IsAdmin() bool {
	return true
}
func (a Admin) DisplayName() string {
	return "[Admin] " + a.User.DisplayName()
}

func main() {
	u := User{"Normal User"}
	fmt.Println(u.Name)
	fmt.Println(u.DisplayName())
	fmt.Println(u.IsAdmin())

	a := Admin{User{"Admin User"}}
	fmt.Println(a.Name)
	fmt.Println(a.User.Name)
	fmt.Println(a.User.DisplayName())
	fmt.Println(a.DisplayName())
	fmt.Println(a.IsAdmin())
}