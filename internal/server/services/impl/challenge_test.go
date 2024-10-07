package impl

import (
	"context"
	"errors"
	"testing"

	"github.com/SladeThe/checked-go/must"
	"github.com/catalinc/hashcash"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
	"github.com/SladeThe/word-of-wisdom/internal/server/services"
)

func TestChallenge_Accept(t *testing.T) {
	ctx := context.Background()
	cfg := ChallengeConfig{ZeroBitCount: 16}

	service, errChallenge := NewChallenge(cfg)
	require.NoError(t, errChallenge)

	type args struct {
		ctx context.Context
		id  entities.ClientID
	}

	type want struct {
		challenge entities.Challenge
		err       error
	}

	test := func(a args, w want, expect func(args, want)) func(t *testing.T) {
		return func(t *testing.T) {
			if expect != nil {
				expect(a, w)
			}

			challenge, err := service.Accept(a.ctx, a.id)

			assert.Truef(t, errors.Is(err, w.err), "invalid error: want = %v, got = %v", w.err, err)
			assert.Equalf(t, w.challenge, challenge, "invalid result: want=%v, got=%v", w.challenge, challenge)
		}
	}

	tests := []struct {
		name   string
		args   args
		want   want
		expect func(args, want)
	}{{
		name: "zero id",
		args: args{
			ctx: ctx,
		},
		want: want{
			err: services.ErrInvalidArguments,
		},
	}, {
		name: "success",
		args: args{
			ctx: ctx,
			id:  entities.ClientID(uuid.New()),
		},
		want: want{
			challenge: entities.Challenge{ZeroBitCount: cfg.ZeroBitCount},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want, tt.expect))
	}
}

func TestChallenge_Solve(t *testing.T) {
	ctx := context.Background()
	cfg := ChallengeConfig{ZeroBitCount: 16}

	service, errChallenge := NewChallenge(cfg)
	require.NoError(t, errChallenge)

	id := entities.ClientID(uuid.New())
	hash := hashcash.New(must.Uint16ToUint(cfg.ZeroBitCount), 8, "")
	header, errMint := hash.Mint(id.String())
	require.NoError(t, errMint)

	type args struct {
		ctx       context.Context
		id        entities.ClientID
		challenge entities.Challenge
		solution  entities.Solution
	}

	type want struct {
		err error
	}

	test := func(a args, w want, expect func(args, want)) func(t *testing.T) {
		return func(t *testing.T) {
			if expect != nil {
				expect(a, w)
			}

			err := service.Solve(a.ctx, a.id, a.challenge, a.solution)

			assert.Truef(t, errors.Is(err, w.err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name   string
		args   args
		want   want
		expect func(args, want)
	}{{
		name: "zero id",
		args: args{
			ctx:       ctx,
			challenge: entities.Challenge{ZeroBitCount: cfg.ZeroBitCount},
			solution:  entities.Solution{Header: header},
		},
		want: want{
			err: services.ErrInvalidArguments,
		},
	}, {
		name: "zero challenge",
		args: args{
			ctx:      ctx,
			id:       id,
			solution: entities.Solution{Header: header},
		},
		want: want{
			err: services.ErrInvalidArguments,
		},
	}, {
		name: "zero solution",
		args: args{
			ctx:       ctx,
			id:        id,
			challenge: entities.Challenge{ZeroBitCount: cfg.ZeroBitCount},
		},
		want: want{
			err: services.ErrInvalidArguments,
		},
	}, {
		name: "invalid solution",
		args: args{
			ctx:       ctx,
			id:        id,
			challenge: entities.Challenge{ZeroBitCount: cfg.ZeroBitCount},
			solution:  entities.Solution{Header: header + "1"},
		},
		want: want{
			err: services.ErrInvalidSolution,
		},
	}, {
		name: "success",
		args: args{
			ctx:       ctx,
			id:        id,
			challenge: entities.Challenge{ZeroBitCount: cfg.ZeroBitCount},
			solution:  entities.Solution{Header: header},
		},
		want: want{},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want, tt.expect))
	}
}
