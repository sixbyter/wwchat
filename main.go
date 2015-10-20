package main

import (
	_ "project/conf"
	_ "project/routes"
	"sixbyte/framework"
)

func main() {
	framework.Run()
}
