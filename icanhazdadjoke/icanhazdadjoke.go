package icanhazdadjoke

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

var (
	defaultTimeout   = 10 * time.Second
	defaultBaseURL   = "https://icanhazdadjoke.com/"
	defaultUserAgent = "Brandless Take Home Test Library (https://github.com/chriskaukis/brandless)"
)

type Joke struct {
	ID     string
	Joke   string
	Status int
}

type Client struct {
	*http.Client
	BaseURL   *url.URL
	UserAgent string
}

func New() *Client {
	// We have legitimate maybe reason to ignore this error since we aren't
	// allowing to set a URL via our API yet.
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		Client:    &http.Client{Timeout: defaultTimeout},
		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
	}
}

func (c *Client) Random() (*Joke, error) {
	req, err := c.request()
	if err != nil {
		return nil, err
	}

	var joke Joke
	_, err = c.do(req, &joke)
	if err != nil {
		return nil, err
	}
	return &joke, nil

}

func (c *Client) request() (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)
	return res, err
}
