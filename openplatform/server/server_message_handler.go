package server

type TicketHandler struct {
	responseText string
}

function (t *TicketHandler) ServeMessage(Message) bool {
	//缓存
}