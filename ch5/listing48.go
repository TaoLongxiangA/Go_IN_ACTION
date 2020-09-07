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
	Name  string
	Email string
}

func (u *User) notify() {
	fmt.Printf("Sending user wmail to %s<%s>\n", u.Name, u.Email)
}

func (a *Admin) notify() {
	fmt.Printf("Sending user wmail to %s<%s>\n", a.Name, a.Email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	user := User{
		Name:  "tao",
		Email: "123@qq.com",
	}

	admin := Admin{
		Name:  "zhang",
		Email: "456@qq.com",
	}

	sendNotification(&user)
	sendNotification(&admin)

	//user.notify()
	//admin.notify()
}
