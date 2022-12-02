package model

type User struct {
	Id        int
	Username  string
	Password  string
	Telephone string
}
type MessageBoard struct {
	NickName string
	Message  string
	Time     string
}
type Comment struct {
	MessageId string
	NickName  string
	Comment   string
	Time      string
}
