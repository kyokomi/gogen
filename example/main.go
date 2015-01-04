package main

// Hoge hoge
//go:generate genexample
type Hoge struct {
	Name    string
	Num     int
	Message string
}

func main() {

	h := Hoge{
		Name:    "hoge",
		Num:     1,
		Message: "test",
	}
	h.Sample()
}
