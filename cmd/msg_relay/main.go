package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/rattapon001/porter-management-demo/internal/job/msg_relay"
)

func main() {
	err := godotenv.Load("./configs/local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	ctx := context.Background()
	msg_relay.MsgRelay(ctx)

}
