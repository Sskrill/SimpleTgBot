package repo

import (
	"crypto/sha1"
	"fmt"
	wrap "github.com/Sskrill/tgBotTest/pkg"
	"io"
)

type Repo interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}
type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	hasher := sha1.New()
	if _, err := io.WriteString(hasher, p.URL); err != nil {
		return "", wrap.Wrap("cant hash", err)
	}
	if _, err := io.WriteString(hasher, p.UserName); err != nil {
		return "", wrap.Wrap("cant hash", err)
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
