package ftp

type dataType int

const (
	Ascii dataType = iota
	Image
)

var dataTypes = map[string]dataType{
	"A": Ascii,
	"I": Image,
}
