package main

import (
	"flag"
	"github.com/Sskrill/tgBotTest/clients/telegram"
	"github.com/Sskrill/tgBotTest/consumer/eventConsumer"
	telegram2 "github.com/Sskrill/tgBotTest/events/telegram"
	"github.com/Sskrill/tgBotTest/repo/files"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	repoPath  = "repo"
	batchSize = 100
)

func main() {
	tgClient := telegram.NewClient(tgBotHost, mustToken())
	eventsProccesor := telegram2.NewProccesor(tgClient, files.NewRepo(repoPath))
	consumer := eventConsumer.NewConsumer(eventsProccesor, eventsProccesor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}
}
func mustToken() string {
	token := flag.String("t", "", "bot token")
	if *token == "" {
		log.Fatal("token is empty")
	}
	return *token
}
