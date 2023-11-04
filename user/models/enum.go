package models


type Gender uint8

const (
	MALE Gender = iota
	FEMALE
)

type Role uint8

const (
	USER Role = iota
	ADMIN
)

type MessageType uint8

const (
	LEAVING_MESSAGES MessageType = iota
	COMPLAINT
	INQUIRY
	POST_SALE
	WANT_TO_BUY
)