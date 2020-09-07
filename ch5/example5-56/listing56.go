package main

import "fmt"

type notifier interface {
	notify()
}

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Email string
}

func (u *User) notify() {
	fmt.Printf("Sending user wmail to %s<%s>\n", u.Name, u.Email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	admin := Admin{
		User: User{
			Name:  "tao",
			Email: "123@qq.com",
		},
		Email: "456@qq.com",
	}
	sendNotification(&admin)
}
