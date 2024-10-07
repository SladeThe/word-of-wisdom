package raw

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync/atomic"
	"time"

	"github.com/google/uuid"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
	"github.com/SladeThe/word-of-wisdom/internal/common/network"
)

const (
	maxRemoteBytesLength = 1 << 20
)

var (
	byteOrder = binary.LittleEndian

	ErrRemoteDataExceedsLimit = errors.New("remote data exceeds limit")
)

type Remote struct {
	conn net.Conn

	reader io.Reader
	writer io.Writer

	err atomic.Pointer[error]
}

var _ network.Client = (*Remote)(nil)
var _ network.Server = (*Remote)(nil)

func FromConnection(conn net.Conn) *Remote {
	return &Remote{
		conn:   conn,
		reader: conn,
		writer: conn,
	}
}

func (r *Remote) Close() error {
	return r.conn.Close()
}

func (r *Remote) ReadClientID() (entities.ClientID, error) {
	r.setDeadline(time.Second)

	id, err := uuid.FromBytes(r.readBytes())
	if err != nil {
		return entities.ClientID{}, err
	}

	return handleErrors(r, entities.ClientID(id))
}

func (r *Remote) WriteClientID(id entities.ClientID) error {
	r.setDeadline(time.Second)
	r.writeBytes(id[:])
	return r.flush()
}

func (r *Remote) ReadChallenge() (entities.Challenge, error) {
	r.setDeadline(time.Second)
	challenge := entities.Challenge{ZeroBitCount: r.readUint16()}
	return handleErrors(r, challenge)
}

func (r *Remote) WriteChallenge(challenge entities.Challenge) error {
	r.setDeadline(time.Second)
	r.writeUint16(challenge.ZeroBitCount)
	return r.flush()
}

func (r *Remote) ReadSolution() (entities.Solution, error) {
	r.setDeadline(time.Minute)
	solution := entities.Solution{Header: r.readString()}
	return handleErrors(r, solution)
}

func (r *Remote) WriteSolution(solution entities.Solution) error {
	r.setDeadline(time.Second)
	r.writeString(solution.Header)
	return r.flush()
}

func (r *Remote) ReadWordOfWisdom() (entities.WordOfWisdom, error) {
	r.setDeadline(time.Second)
	word := entities.WordOfWisdom{Text: r.readString()}
	return handleErrors(r, word)
}

func (r *Remote) WriteWordOfWisdom(word entities.WordOfWisdom) error {
	r.setDeadline(time.Second)
	r.writeString(word.Text)
	return r.flush()
}
