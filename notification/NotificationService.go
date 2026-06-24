package notification

type NotificationService struct {
	Observers map[string]NotificationOberserver
}

func (n *NotificationService) Add(UserId string, ob NotificationOberserver) {
	n.Observers[UserId] = ob
} 

func (n *NotificationService) Remove(userId string) {
	delete(n.Observers,userId)
}

func (n *NotificationService) Notify(userId string,msg string) {
	observer := n.Observers[userId]
	observer.Update(msg)
}

func (n *NotificationService) NotifyAll (msg string) {
	for _,val := range n.Observers {
		val.Update(msg)
	}
}