package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	expected := `rune	count
' '	1
','	1
'H'	1
'e'	1
'l'	2
'o'	1
'世'	1
'界'	1

len	count
1	7
2	0
3	2
4	0
`
	in = strings.NewReader("Hello, 世界")
	out = new(bytes.Buffer)
	charcount()
	actual := out.(*bytes.Buffer).String()
	if expected != actual {
		t.Errorf("\nexpected:\n%s\nactual:\n%s", expected, actual)
	}
}
