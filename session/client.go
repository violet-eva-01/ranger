package session

type Client struct {
	host     string
	port     int
	path     string
	proxy    string
	userName string
	passWord string
	headers  map[string]string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetHost(host string) *Client {
	c.host = host
	return c
}

func (c *Client) SetPort(port int) *Client {
	c.port = port
	return c
}

func (c *Client) SetPath(path string) *Client {
	c.path = path
	return c
}

func (c *Client) SetProxy(proxy string) *Client {
	c.proxy = proxy
	return c
}

func (c *Client) SetUserName(userName string) *Client {
	c.userName = userName
	return c
}

func (c *Client) SetPassWord(passWord string) *Client {
	c.passWord = passWord
	return c
}

func (c *Client) SetHeaders(headers map[string]string) *Client {
	for k, v := range headers {
		c.headers[k] = v
	}
	return c
}

func (c *Client) GetSession() *Session {
	return &Session{
		host:     c.host,
		port:     c.port,
		path:     c.path,
		userName: c.userName,
		passWord: c.passWord,
		headers:  c.headers,
	}
}
