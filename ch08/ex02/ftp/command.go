package ftp

type command int

const (
	Unknown command = iota
	User
	Quit
	Port
	Type
	Retr
	Stor
)

var commands = map[string]command{
	"USER": User,
	"QUIT": Quit,
	"PORT": Port,
	"TYPE": Type,
	"RETR": Retr,
	"STOR": Stor,
}
