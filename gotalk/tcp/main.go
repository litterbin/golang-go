package main

type GreetIn struct {
	Name string `json:"name"`
}

type GreetOut struct {
	Greeting string `json:"greeting"`
}

func main() {
	port := "1234"
	server(port)
	client(port)
}
