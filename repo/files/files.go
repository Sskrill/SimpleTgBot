package files

import (
	wrap "github.com/Sskrill/tgBotTest/pkg"
	"github.com/Sskrill/tgBotTest/repo"
	"os"
	"path/filepath"
)

type Repo struct {
	basePath string
}

const (
	defaultPerm = 0774
)

func NewRepo(basePath string) Repo {
	return Repo{basePath: basePath}
}
func (r Repo) Save(p *repo.Page) (err error) {
	fPath := filepath.Join(r.basePath, p.UserName)
	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return wrap.Wrap("cant save", err)
	}
	fName, err := fileName(p)
	if err != nil {
		return wrap.Wrap("cant save", err)
	}
	fPath = filepath.Join(fPath, fName)
	file, err := os.Create(fPath)
	if err != nil {
		return wrap.Wrap("cant save", err)
	}
	defer func() { _ = file.Close() }()
}
func fileName(p *repo.Page) (string, error) {
	return p.Hash()
}
