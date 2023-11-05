package models

type Gender uint8

const (
	MALE Gender = iota + 1
	FEMALE
)

type Role uint8

const (
	USER Role = iota + 1
	ADMIN
)

type MessageType uint8

const (
	LEAVING_MESSAGES MessageType = iota + 1
	COMPLAINT
	INQUIRY
	POST_SALE
	WANT_TO_BUY
)
