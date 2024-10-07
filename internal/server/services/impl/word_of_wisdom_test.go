package impl

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
	"github.com/SladeThe/word-of-wisdom/internal/server/repositories"
	mockRepositories "github.com/SladeThe/word-of-wisdom/internal/server/repositories/mock"
	"github.com/SladeThe/word-of-wisdom/internal/server/services"
)

func TestWordOfWisdom_OneRandom(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := context.Background()
	ctx = repositories.Set(ctx, repositories.Repositories{WordOfWisdom: mockRepositories.NewMockWordOfWisdom(ctrl)})

	service := NewWordOfWisdom()

	word := "Word of Wisdom"

	type args struct {
		ctx context.Context
	}

	type want struct {
		word entities.WordOfWisdom
		err  error
	}

	test := func(a args, w want, expect func(args, want)) func(t *testing.T) {
		return func(t *testing.T) {
			if expect != nil {
				expect(a, w)
			}

			word, err := service.OneRandom(a.ctx)

			assert.Truef(t, errors.Is(err, w.err), "invalid error: want = %v, got = %v", w.err, err)
			assert.Equalf(t, w.word, word, "invalid result: want=%v, got=%v", w.word, word)
		}
	}

	tests := []struct {
		name   string
		args   args
		want   want
		expect func(args, want)
	}{{
		name: "not found",
		args: args{
			ctx: ctx,
		},
		want: want{
			err: services.ErrNotFound,
		},
		expect: func(a args, w want) {
			repositories.Must(a.ctx).WordOfWisdom.(*mockRepositories.MockWordOfWisdom).EXPECT().
				OneRandom().
				Return(entities.WordOfWisdom{}, repositories.ErrNotFound).
				Times(1)
		},
	}, {
		name: "internal repository error",
		args: args{
			ctx: ctx,
		},
		want: want{
			err: services.ErrInternal,
		},
		expect: func(a args, w want) {
			repositories.Must(a.ctx).WordOfWisdom.(*mockRepositories.MockWordOfWisdom).EXPECT().
				OneRandom().
				Return(entities.WordOfWisdom{}, repositories.ErrInternal).
				Times(1)
		},
	}, {
		name: "success",
		args: args{
			ctx: ctx,
		},
		want: want{
			word: entities.WordOfWisdom{Text: word},
		},
		expect: func(a args, w want) {
			repositories.Must(a.ctx).WordOfWisdom.(*mockRepositories.MockWordOfWisdom).EXPECT().
				OneRandom().
				Return(entities.WordOfWisdom{Text: word}, nil).
				Times(1)
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want, tt.expect))
	}
}
