package telegram

import (
	"encoding/json"
	wrap "github.com/Sskrill/tgBotTest/pkg"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   *http.Client
}

const (
	getUpdates  = "getUpdates"
	sendMessage = "sendMessage"
)

func NewClient(host, token string) *Client {
	return &Client{host: host, basePath: newBasePath(token), client: &http.Client{}}
}
func newBasePath(token string) string {
	return "bot" + token
}
func (c *Client) Updates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))
	data, err := c.doReq(getUpdates, q)
	if err != nil {
		return nil, err
	}
	var res UpdateResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Result, nil
}
func (c *Client) SendMessage(chatId int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", text)
	_, err := c.doReq(sendMessage, q)
	if err != nil {
		return wrap.Wrap("cant send msg", err)
	}
	return nil
}
func (c *Client) doReq(method string, query url.Values) ([]byte, error) {
	const errMsg = "cant do req"
	u := url.URL{Scheme: "https", Host: c.host, Path: path.Join(c.basePath, method)}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, wrap.Wrap(errMsg, err)
	}
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, wrap.Wrap(errMsg, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrap.Wrap(errMsg, err)
	}
	return body, nil
}
