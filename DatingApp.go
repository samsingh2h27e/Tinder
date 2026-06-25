package main

import (
	chatroom "Tinder/ChatRoom"
	enums "Tinder/Enums"
	matcher "Tinder/Matcher"
	user "Tinder/user"
)

type Tinder struct {
	Users []user.User
	Rooms []chatroom.ChatRoom
	Matcher matcher.Matcher
}

func (t *Tinder) SetMatcher (mt enums.MatcherType) {
	factory := matcher.MatcherFactory{}
	matcher := factory.CreateMatcher(mt)
	t.Matcher = matcher
}

func (t *Tinder) CreateUser (userId string) {
	user := user.User{}
	user.UserID = userId
	t.Users = append(t.Users , user)
}

func (t *Tinder) Swipe(userId string,targetId string,sa enums.SwipeAction) {
	for _,val := range t.Users {
		if val.UserID == userId {
			val.Swipe(targetId,sa)
			return
		}
	}
}