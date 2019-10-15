package main

import (
	r "restapi-gorillamux/routers"
)

func main() {
	r.Init()
	r.Routers()
}
