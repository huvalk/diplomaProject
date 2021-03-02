package channel

type Channel interface {
	Run()
	SendNotification(n *Notification) (send bool, err error)
	HandleMessage(rawMessage []byte)
	Forward([]byte)
	Join(user *ConnectedUser)
	Leave(user *ConnectedUser)
}
