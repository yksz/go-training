package display

func Example_arrayKeyMap() {
	m := make(map[[2]string]string)
	key := [2]string{"item1", "item2"}
	m[key] = "value"
	Display("m", m)

	// Output:
	// Display m (map[[2]string]string):
	// m[{"item1", "item2"}] = "value"
}

func Example_structKeyMap() {
	type object struct {
		id   int
		name string
	}
	m := make(map[object]string)
	key := object{1, "obj"}
	m[key] = "value"
	Display("m", m)

	// Output:
	// Display m (map[display.object]string):
	// m[{id:1, name:"obj"}] = "value"
}
