package core

// map of id -> tickets
// an growing bank of tickets

type TicketStore struct {
	store map[int]Ticket
}
