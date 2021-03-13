package main

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// We could write `(*p).X`, however, that notation is cumbersome, so
	// the language permits us instead to write just `p.X` without
	// explicit defererence.
	p.X = 1e9
	fmt.Println(v)
}
