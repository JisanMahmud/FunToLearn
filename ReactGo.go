package main

import (
	"github.com/GoRockypt/FunToLearn/GoSupport"
	"github.com/GoRockypt/FunToLearn/gorrilaRouter"
)

func main() {
	godb.DbOpen()
	gorrilaRouter.StartServer()
}
