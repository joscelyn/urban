package server

import (
	"errors"
	"fmt"
	"time"

	"context"

	"github.com/joscelyn/urban"
	"github.com/peterhellberg/hn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Whois fetch user information of the given hacker news user
func (s UrbanServer) Whois(ctx context.Context, r *urban.WhoisRequest) (*urban.WhoisResponse, error) {
	if r.Username == "" {
		return nil, errors.New("username must not be empty")
	}

	user, err := hn.DefaultClient.User(r.Username)
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to fetch hn user", err))
	}

	return &urban.WhoisResponse{
		User:   r.Username,
		Karma:  int64(user.Karma),
		About:  user.About,
		Joined: timestamppb.New(time.Unix(int64(user.Created), 0)),
	}, nil
}
