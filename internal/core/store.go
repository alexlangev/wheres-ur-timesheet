package core

// map of id -> tickets
// an growing bank of tickets

type ticketStore struct {
	store map[int]Ticket
}
