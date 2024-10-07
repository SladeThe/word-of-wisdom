package client

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/SladeThe/checked-go/must"
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vnumber"
	"github.com/SladeThe/yav/vstring"
	"github.com/catalinc/hashcash"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
	"github.com/SladeThe/word-of-wisdom/internal/common/network"
	"github.com/SladeThe/word-of-wisdom/internal/common/network/raw"
)

type Config struct {
	Host string `default:"server"`
	Port uint16 `default:"9999"`
}

func (cfg Config) Validate() error {
	return yav.Join(
		yav.Chain(
			"Host", cfg.Host,
			vstring.Between(2, 200),
			yav.Or2[string](vstring.Hostname, vstring.HostnameRFC1123),
		),
		yav.Chain(
			"Port", cfg.Port,
			vnumber.MinUint16(1),
		),
	)
}

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func Start(ctx context.Context, cfg Config) (*Client, error) {
	if errValidate := cfg.Validate(); errValidate != nil {
		return nil, errValidate
	}

	conn, errDial := net.Dial("tcp4", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if errDial != nil {
		return nil, errDial
	}

	ctx, cancel := context.WithCancel(ctx)
	client := &Client{ctx: ctx, cancel: cancel}

	go client.process(conn)

	return client, nil
}

func (c *Client) Shutdown() {
	c.cancel()
}

func (c *Client) process(conn net.Conn) {
	defer func() { _ = conn.Close() }()

	var client network.Client = raw.FromConnection(conn)

	clientID, errReadID := client.ReadClientID()
	if errReadID != nil {
		log.Fatal("[ERROR] failed reading client ID: ", errReadID)
		return
	}

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		challenge, errReadChallenge := client.ReadChallenge()
		if errReadChallenge != nil {
			log.Fatal("[ERROR] failed reading challenge: ", errReadChallenge)
			return
		}

		hash := hashcash.New(must.Uint16ToUint(challenge.ZeroBitCount), 8, "")

		header, errMint := hash.Mint(clientID.String())
		if errMint != nil {
			log.Fatal("[ERROR] failed solving challenge: ", errMint)
			return
		}

		if errWriteSolution := client.WriteSolution(entities.Solution{Header: header}); errWriteSolution != nil {
			log.Fatal("[ERROR] failed writing solution: ", errWriteSolution)
			return
		}

		word, errReadWord := client.ReadWordOfWisdom()
		if errReadWord != nil {
			log.Fatal("[ERROR] failed reading word: ", errReadWord)
			return
		}

		log.Print("[INFO] remember carefully: ", word.Text)
	}
}
