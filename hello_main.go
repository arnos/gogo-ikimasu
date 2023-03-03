package main

func main() {

	// load the configurations
	loadConfiguration("helloworld.toml")

	// define core entity
	w := world{}

	// call logic on the core entity
	HelloWorld(w)
}
