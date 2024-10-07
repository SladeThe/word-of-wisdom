package server

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	"github.com/SladeThe/checked-go/must"
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vnumber"
	"github.com/google/uuid"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
	"github.com/SladeThe/word-of-wisdom/internal/common/network"
	"github.com/SladeThe/word-of-wisdom/internal/common/network/raw"
	"github.com/SladeThe/word-of-wisdom/internal/server/services"
)

type Config struct {
	Port          uint16 `default:"9999"`
	ListenerCount uint16 `default:"8"`
}

func (cfg Config) Validate() error {
	return yav.Join(
		yav.Chain(
			"Port", cfg.Port,
			vnumber.MinUint16(1),
		),
		yav.Chain(
			"ListenerCount", cfg.ListenerCount,
			vnumber.BetweenUint16(1, 1000),
		),
	)
}

type Server struct {
	ctx      context.Context
	cancel   context.CancelFunc
	listener *net.TCPListener
	wg       sync.WaitGroup
}

func Start(ctx context.Context, cfg Config) (*Server, error) {
	if errValidate := cfg.Validate(); errValidate != nil {
		return nil, errValidate
	}

	listener, errListen := net.ListenTCP("tcp", &net.TCPAddr{Port: must.Uint16ToInt(cfg.Port)})
	if errListen != nil {
		return nil, errListen
	}

	ctx, cancel := context.WithCancel(ctx)
	server := &Server{ctx: ctx, cancel: cancel, listener: listener}
	server.wg.Add(must.Uint16ToInt(cfg.ListenerCount))

	for i := uint16(0); i < cfg.ListenerCount; i++ {
		go func() {
			defer server.wg.Done()

			for {
				conn, errAccept := listener.AcceptTCP()
				if errors.Is(errAccept, net.ErrClosed) {
					return
				}
				if errAccept != nil {
					log.Print("[ERROR] failed accepting TPC connection: ", errAccept)
					continue
				}

				server.wg.Add(1)
				go server.process(conn)
			}
		}()
	}

	go func() {
		<-server.ctx.Done()
		_ = server.listener.Close()
	}()

	return server, nil
}

func (s *Server) Shutdown() {
	s.cancel()
	s.wg.Wait()
}

func (s *Server) process(conn net.Conn) {
	defer func() {
		_ = conn.Close()
		s.wg.Done()
	}()

	var server network.Server = raw.FromConnection(conn)
	clientID := entities.ClientID(uuid.New())

	if errWriteID := server.WriteClientID(clientID); errWriteID != nil {
		log.Print("[ERROR] failed writing client ID: ", errWriteID)
		return
	}

	ss := services.Must(s.ctx)

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		challenge, errAccept := ss.Challenge.Accept(s.ctx, clientID)
		if errAccept != nil {
			log.Print("[ERROR] failed accepting challenge: ", errAccept)
			return
		}

		log.Print("[INFO] challenge client: ", clientID.String())
		if errWriteChallenge := server.WriteChallenge(challenge); errWriteChallenge != nil {
			log.Print("[ERROR] failed writing challenge: ", errWriteChallenge)
			return
		}

		solution, errReadSolution := server.ReadSolution()
		if errReadSolution != nil {
			log.Print("[ERROR] failed reading solution: ", errReadSolution)
			return
		}

		log.Print("[INFO] check solution: ", solution.Header)
		if errSolve := ss.Challenge.Solve(s.ctx, clientID, challenge, solution); errSolve != nil {
			log.Print("[ERROR] failed solving challenge: ", errSolve)
			return
		}

		word, errRandom := ss.WordOfWisdom.OneRandom(s.ctx)
		if errRandom != nil {
			log.Print("[ERROR] failed getting random word: ", errRandom)
			return
		}

		log.Print("[INFO] send word to: ", clientID.String())
		if errWriteWord := server.WriteWordOfWisdom(word); errWriteWord != nil {
			log.Print("[ERROR] failed writing word: ", errWriteWord)
			return
		}
	}
}
