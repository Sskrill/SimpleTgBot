package telegram

import "github.com/Sskrill/tgBotTest/clients/telegram"

type Proccesor struct {
	tg     *telegram.Client
	offset int
	// storage
}
