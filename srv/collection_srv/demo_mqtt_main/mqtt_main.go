package main

import (
	"collection_srv/internal/logic"
	"common/initialize"
)

func main() {
	initialize.InitMongoDB()
	_, err := logic.DataCollection()
	if err != nil {
		return
	}
}
