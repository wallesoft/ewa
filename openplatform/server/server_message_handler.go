package server

type TicketHandler struct {
	responseText string
}

func (t *TicketHandler) ServeMessage(Message) bool {
	//缓存
}
