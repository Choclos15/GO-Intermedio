package main

import (
	hellomod "github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Mario")
}
