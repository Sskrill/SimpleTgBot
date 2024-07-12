package telegram2

import (
	"github.com/Sskrill/tgBotTest/clients/telegram"
	wrap "github.com/Sskrill/tgBotTest/pkg"
	"github.com/Sskrill/tgBotTest/repo"
	"log"
	"net/url"
	"strings"
)

const (
	RndCmd   = "/rnd"
	StartCmd = "/start"
	HelpCmd  = "/help"
)

func (p *Proccesor) doCmd(text string, chatId int, username string) error {
	text = strings.TrimSpace(text)
	log.Printf("got new command '%s' from '%s", text, username)

	if isAddCmd(text) {
		return p.savePage(chatId, text, username)
	}
	switch text {
	case RndCmd:
		return p.SendRandom(chatId, username)
	case StartCmd:
		return p.SendHello(chatId)
	case HelpCmd:
		return p.SendHelp(chatId)
	default:
		return p.tg.SendMessage(chatId, msgUnknownCommand)

	}
}
func (p *Proccesor) savePage(chatId int, pageUrl, username string) error {
	page := &repo.Page{URL: pageUrl, UserName: username}
	//	send := NewMessageSender(chatId, p.tg)  -Альтернативный вариант отправки сообщения
	isExists, err := p.repo.IsExists(page)
	if err != nil {
		return wrap.Wrap("cant do command :save page", err)
	}
	if isExists {
		return p.tg.SendMessage(chatId, msgAlreadyExists)
	}
	if err = p.repo.Save(page); err != nil {
		return err
	}
	if err = p.tg.SendMessage(chatId, msgSaved); err != nil {
		return err
	}
	return nil
}
func (p *Proccesor) SendRandom(chatId int, username string) error {
	page, err := p.repo.PickRandom(username)
	if err != nil {
		return wrap.Wrap("cant do command send random page ", err)
	}
	if err = p.tg.SendMessage(chatId, page.URL); err != nil {
		return err
	}
	return p.repo.Remove(page)
}
func NewMessageSender(chatId int, tg *telegram.Client) func(string) error {
	return func(s string) error {
		return tg.SendMessage(chatId, s)
	}
}
func isAddCmd(text string) bool {
	return isUrl(text)
}
func isUrl(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
func (p *Proccesor) SendHelp(chatId int) error {
	return p.tg.SendMessage(chatId, msgHelp)
}
func (p *Proccesor) SendHello(chatId int) error {
	return p.tg.SendMessage(chatId, msgHello)
}
