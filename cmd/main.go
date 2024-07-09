package main

import (
	"flag"
	"github.com/Sskrill/tgBotTest/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.NewClient(tgBotHost, mustToken())
	fether := fetcher.NewFetcher(tgClient)

}
func mustToken() string {
	token := flag.String("t", "", "bot token")
	if *token == "" {
		log.Fatal("token is empty")
	}
	return *token
}
