package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pingproto "github.com/enginoid/monorepo-base/services/ping/proto"
	cli "github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "ping",
		Usage: "ping the server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "address",
				Value: "localhost:50051",
				Usage: "grpc address to dial",
			},
		},
		Action: func(c *cli.Context) error {
			address := c.String("address")
			log.Printf("dialing %s...", address)

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("failed to connect: %v", err)
			}
			defer conn.Close()

			client := pingproto.NewPingClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			message := c.Args().First()
			log.Printf("ping message: %#v", message)

			response, err := client.Ping(ctx, &pingproto.PingRequest{Message: message})
			if err != nil {
				return fmt.Errorf("could not ping: %v", err)
			}
			log.Printf("ping reply: %#v", response.GetMessage())

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
