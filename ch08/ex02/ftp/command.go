package ftp

type command int

const (
	Unknown command = iota
	User
	Quit
	Cwd
	Port
	Type
	Retr
	Stor
	List
	Nlst
)

var commands = map[string]command{
	"USER": User,
	"QUIT": Quit,
	"CWD":  Cwd,
	"PORT": Port,
	"TYPE": Type,
	"RETR": Retr,
	"STOR": Stor,
	"LIST": List,
	"NLST": Nlst,
}
