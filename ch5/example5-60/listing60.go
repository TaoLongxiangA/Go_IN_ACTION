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
	fmt.Printf("Sending user email to %s<%s>\n", u.Name, u.Email)
}

func (a *Admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.Name, a.Email)
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
	admin.User.notify()
	admin.notify()
}
