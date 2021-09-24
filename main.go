package main

const serverAddress = "localhost:8888"

func main() {
	SetupRouter().Run(serverAddress)
}
