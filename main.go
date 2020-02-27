//+build !test

package main

func main() {
	router := NewCommandRouter()
	router.Run()
}
