package telegram

import (
	"errors"
	"github.com/Sskrill/tgBotTest/clients/telegram"
	"github.com/Sskrill/tgBotTest/events"
	wrap "github.com/Sskrill/tgBotTest/pkg"
	"github.com/Sskrill/tgBotTest/repo"
)

type Proccesor struct {
	tg     *telegram.Client
	offset int
	repo   repo.Repo
}
type Meta struct {
	ChatId   int
	Usernaem string
}

var ErrUnkwownMetaType = errors.New("unkwown meta type")
var ErrUnknouwnEventType = errors.New("unknown event type")

func NewProccesor(client *telegram.Client, repo repo.Repo) *Proccesor {
	return &Proccesor{tg: client, repo: repo}
}
func (p *Proccesor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, wrap.Wrap("cant fetch", err)
	}
	res := make([]events.Event, 0, len(updates))
	if len(updates) == 0 {
		return nil, nil
	}
	for _, u := range updates {
		res = append(res, event(u))
	}
	p.offset = updates[len(updates)-1].Id + 1
	return res, nil
}
func (p Proccesor) Procces(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.proccesMessage(event)
	default:
		return wrap.Wrap("cant procces message", ErrUnknouwnEventType)
	}
}
func (p Proccesor) proccesMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return wrap.Wrap("cant procces message", err)
	}
	if err = p.doCmd(event.Text, meta.ChatId, meta.Usernaem); err != nil {
		return wrap.Wrap("cant process message", err)
	}
	return nil
}
func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, wrap.Wrap("cant get meta", ErrUnkwownMetaType)
	}
	return res, nil
}
func event(u telegram.Update) events.Event {
	uType := fetchType(u)
	res := events.Event{Type: uType, Text: fecthText(u)}
	if uType == events.Message {
		res.Meta = Meta{ChatId: u.Message.Chat.Id, Usernaem: u.Message.From.Username}
	}
	return res
}

func fetchType(u telegram.Update) events.Type {
	if u.Message == nil {
		return events.Unknouwn
	}
	return events.Message
}
func fecthText(u telegram.Update) string {
	if u.Message == nil {
		return ""
	}
	return u.Message.Text
}
