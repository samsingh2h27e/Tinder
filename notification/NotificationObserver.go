package notification

type NotificationOberserver interface {
	Update(msg string) 
}