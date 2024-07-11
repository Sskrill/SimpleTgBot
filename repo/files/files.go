package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	wrap "github.com/Sskrill/tgBotTest/pkg"
	"github.com/Sskrill/tgBotTest/repo"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Repo struct {
	basePath string
}

const defaultPerm = 0774

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
	if gob.NewEncoder(file).Encode(p); err != nil {

	}
	return nil
}
func (r Repo) IsExists(p *repo.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, wrap.Wrap("cant find file ", err)
	}
	path := filepath.Join(r.basePath, p.UserName, fileName)
	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("cant check if file %s exists", path)
		return false, wrap.Wrap(msg, err)
	}
	return true, nil
}
func (r Repo) Remove(p *repo.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return wrap.Wrap("cant remove file ", err)
	}
	path := filepath.Join(r.basePath, p.UserName, fileName)
	msg := fmt.Sprintf("cant remove file %s", path)
	if err := os.Remove(path); err != nil {
		return wrap.Wrap(msg, err)
	}
	return nil
}
func (r Repo) PickRandom(userName string) (page *repo.Page, err error) {
	path := filepath.Join(r.basePath, page.UserName)
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, wrap.Wrap("cant pick randaom page ", err)
	}
	if len(files) == 0 {
		return nil, wrap.Wrap("cant pick randaom page ", errors.New("not saved files"))
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))
	file := files[n]
	return r.decodePage(filepath.Join(path, file.Name()))
}
func fileName(p *repo.Page) (string, error) {
	return p.Hash()
}
func (r Repo) decodePage(filePath string) (*repo.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, wrap.Wrap("cant decode page ", err)
	}
	defer func() { _ = f.Close() }()
	var p repo.Page
	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, wrap.Wrap("cant decode page ", err)
	}
	return &p, nil
}
