package notification

type Channel interface {
	Run()
	SendNotification(n *Notification) error
	HandleMessage(rawMessage []byte)
	Forward([]byte)
	Join(user *ChannelUser)
	Leave(user *ChannelUser)
}
