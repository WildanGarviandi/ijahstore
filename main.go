package main

func main() {
	router := SetRouter()
	router.Run(":8081")
}
