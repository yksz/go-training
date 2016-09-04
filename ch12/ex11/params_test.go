package params

import "fmt"

func Example() {
	var data = struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}{
		Labels:     []string{"golang", "programming"},
		MaxResults: 100,
		Exact:      true,
	}

	url := "http://localhost:12345/search"
	fmt.Println(Pack(url, &data))

	// Output:
	// http://localhost:12345/search?l=golang&l=programming&max=100&x=true
}
