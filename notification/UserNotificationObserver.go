package notification

import "fmt"

type UserNotificationObserver struct {
	UserId string
}

func (u *UserNotificationObserver) Update(msg string) {
	fmt.Println("User : " + u.UserId + "got message as : " + msg)
}