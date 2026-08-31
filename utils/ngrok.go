package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.ngrok.com/ngrok/v2"
)

type ngrokContext struct {
	Url string
}

var NgroxCtx *ngrokContext

func StartNgrok(HOST string, PORT string) error {
	ctx := context.Background()

	n := os.Getenv("NGROK_AUTHTOKEN")
	if n == "" {
		log.Fatalln("NGROK_AUTHTOKEN not found. Make sure to include it in the .env file and provide valid NGROK auth token key.")
	}

	agent, err := ngrok.NewAgent(ngrok.WithAuthtoken(os.Getenv("NGROK_AUTHTOKEN")))
	if err != nil {
		return err
	}

	ln, err := agent.Forward(ctx,
		ngrok.WithUpstream(HOST+PORT),
	)

	if err != nil {
		return err
	}

	NgroxCtx = &ngrokContext{
		Url: ln.URL().String(),
	}

	fmt.Println("Endpoint online: forwarding from", ln.URL(), "to", HOST+PORT)
	return nil
}
