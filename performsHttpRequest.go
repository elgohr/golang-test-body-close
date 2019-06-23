package closing

import "net/http"

type MyClient struct {
	Client http.Client
}

func (c *MyClient) Closing() (err error) {
	res, err := c.Client.Get("http://localhost")
	defer res.Body.Close()
	return
}

func (c *MyClient) NotClosing() (err error) {
	_, err = c.Client.Get("http://localhost")
	return
}
