package user

import (
	"Tinder/Enums"
	userprofile "Tinder/UserProfile"
	"Tinder/notification"
	"fmt"
)

type User struct {
	UserID string
	Up userprofile.UserProfile
	P Preference
	SwipeHistory map[string]enums.SwipeAction
	ob notification.NotificationOberserver
}

func (U *User) Swipe(userId string,action enums.SwipeAction) {
	U.SwipeHistory[userId] = action
	fmt.Println("User: "+ U.UserID + "has swiped " + action.String() + "for " + userId)
}

func (U *User) HasLiked(userId string) bool {
	if U.HasInteractedWith(userId) {
		if U.SwipeHistory[userId].String() == "left" {
			return false
		}else {
			return true
		}
	}
	return false
}

func (U *User) HasDisliked(userId string) bool {
	return !U.HasLiked(userId)
}

func (u *User) HasInteractedWith(userId string) bool {
	if _,ok := u.SwipeHistory[userId]; ok {
		return true
	}
	return false
}

