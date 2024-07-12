package eventConsumer

import (
	"github.com/Sskrill/tgBotTest/events"
	"log"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	proccesor events.Proccesor
	bathSize  int
}

func NewConsumer(fetcher events.Fetcher, proccesor events.Proccesor, bathSize int) Consumer {
	return Consumer{fetcher: fetcher, proccesor: proccesor, bathSize: bathSize}
}
func (c Consumer) Start() error {
	for {
		gotEVents, err := c.fetcher.Fetch(c.bathSize)
		if err != nil {
			log.Printf("Error consumer :%s", err.Error())
			continue
		}
		if len(gotEVents) == 0 {
			time.Sleep(time.Second * 1)
		}
		if err := c.handleEvents(gotEVents); err != nil {
			log.Print(err)
			continue
		}
	}
}
func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event : %s", event.Text)
		if err := c.proccesor.Procces(event); err != nil {
			log.Printf("cant  handle event : %s", err.Error())
			continue
		}
	}
	return nil
}
