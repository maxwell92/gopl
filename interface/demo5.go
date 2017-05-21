package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	fmt.Printf("User: Notifying email to %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

func (u User) Send() error {
	fmt.Printf("User: Sending email to %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

type Notifier interface {
	Notify() error
	Send() error
}

func Notification(notify Notifier) error {
	return notify.Notify()
}

func Sending(notify Notifier) error {
	return notify.Send()
}

/*
func Notification1(notify *Notifier) error {
	return *notify.Notify()
}

func Sending1(notify *Notifier) error {
	return *notify.Send()
}
*/

func main() {
	/*--------------------- general struct method test -------------------*/
	/*
		user1 := User{
			Name:  "Mushroom",
			Email: "mushroom@mworks.com",
		}
		user1.Notify()
		user1.Send()

		user2 := &User{
			Name:  "Magic",
			Email: "magic@mworks.com",
		}
		user2.Notify()
		user2.Send()
	*/

	/*---------------------- interface test ---------------------*/
	/*
		user := User{
			Name:  "Mushroom",
			Email: "mushroom@mworks.com",
		}
	*/

	/* cannot use user (type User) as type Notifier in argument to Notification:
		User does not implement Notifier (Notify method has pointer receiver)

	Notification(user)
	*/

	/*
		Notification(&user)
	*/

	/* cannot use user (type User) as type Notifier in argument to Notification:
		User does not implement Notifier (Notify method has pointer receiver)

	Sending(user)
	*/

	/*
		Sending(&user)
	*/

	/* ------------------- interface test2 ----------------------*/
	/*
		user1 := User{
			Name:  "mushroom",
			Email: "mushroom@mworks.com",
		}

		Notification(user1)
		Sending(user1)
	*/

	user2 := &User{
		Name:  "magic",
		Email: "magic@mworks.com",
	}

	Notification(user2)
	Sending(user2)
}
