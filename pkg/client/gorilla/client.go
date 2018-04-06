package gorilla

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/xakep666/wurl/pkg/client"
	"github.com/xakep666/wurl/pkg/config"
)

var (
	_ client.Client      = &Client{}
	_ client.Constructor = NewClient
)

type Client struct {
	conn           *websocket.Conn
	connWriteMutex sync.Mutex

	opts *config.Options
	log  *logrus.Entry
}

func setupDialer(opts *config.Options) *websocket.Dialer {
	return &websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: opts.AllowInsecureSSL,
		},
	}
}

func (c *Client) prepareConnection() {
	c.conn.SetPingHandler(func(appData string) error {
		if c.opts.RespondPings {
			c.log.Debugf("ping received from %s, payload %s", c.conn.RemoteAddr(), appData)
			return c.conn.WriteMessage(websocket.PingMessage, nil)
		}
		return nil
	})
	c.conn.SetPongHandler(func(appData string) error {
		c.log.Debugf("pong received from %s, payload %s", c.conn.RemoteAddr(), appData)
		return nil
	})
}

func (c *Client) periodicPinger() {
	ticker := time.NewTicker(c.opts.PingPeriod)
	defer ticker.Stop()
	for {
		if err := c.Ping(nil); err != nil {
			c.log.WithError(err).Error("websocket ping failed")
			return
		}

		<-ticker.C
	}
}

func NewClient(url string, opts *config.Options) (client.Client, error) {
	dialer := setupDialer(opts)
	conn, resp, err := dialer.Dial(url, opts.AdditionalHeaders)
	switch err {
	case nil:
		if printErr := client.WriteHandshakeResponse(resp, opts); printErr != nil {
			return nil, printErr
		}
	case websocket.ErrBadHandshake:
		if printErr := client.WriteHandshakeResponse(resp, opts); printErr != nil {
			return nil, printErr
		}
		return nil, fmt.Errorf("bad handshake: status %s", resp.Status)
	default:
		return nil, err
	}

	log := logrus.StandardLogger().WithField("client", "gorilla")

	ret := &Client{conn: conn, opts: opts, log: log}

	ret.prepareConnection()

	if opts.PingPeriod > 0 {
		go ret.periodicPinger()
	}

	return ret, nil
}

func (c *Client) Close() error {
	c.log.Debugf("closing websocket client")
	return c.conn.Close()
}
