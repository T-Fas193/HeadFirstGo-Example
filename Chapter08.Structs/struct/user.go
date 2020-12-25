package _struct

type User struct {
	username string

	mobile string

	Email string

	Dept Dept
}

type EmbedUser struct {
	username string
	mobile   string
	Email    string
	Dept
}
