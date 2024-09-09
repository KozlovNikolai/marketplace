package httpserver

import (
	"context"
	"testing"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
	"github.com/stretchr/testify/assert"
)

func Test_getUserFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    domain.User
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				ctx: context.WithValue(
					context.Background(),
					"user",
					domain.User{}),
			},
			want:    domain.User{},
			wantErr: false,
		},
		{
			name: "invalid key",
			args: args{
				ctx: context.WithValue(
					context.Background(),
					"product",
					domain.User{}),
			},
			want:    domain.User{},
			wantErr: true,
		},
		{
			name: "invalid struct",
			args: args{
				ctx: context.WithValue(
					context.Background(),
					"user",
					domain.Provider{}),
			},
			want:    domain.User{},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			user, err := getUserFromContext(tt.args.ctx)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, user)
				assert.NoError(t, err)
			}
		})
	}
}
