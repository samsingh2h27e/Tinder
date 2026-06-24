package user

import (
	"Tinder/Enums"
	"Tinder/notification"
)

type User struct {
	UserID string
	Up UserProfile
	P Preference
	SwipeHistory map[string]enums.SwipeAction
	ob notification.NotificationOberserver
}

func (U *User) Swipe(userId string,action enums.SwipeAction) {

}

func (U *User) HasLiked(userId string) bool {

}

func (U *User) HasDisliked(userId string) bool {
	return !U.HasLiked(userId)
}

func (u *User) HasInteractedWith(userId string) bool {
	
}

