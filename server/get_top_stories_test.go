package server

import (
	"testing"

	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetTopStories(t *testing.T) {
	ctx := context.Background()
	req := &emptypb.Empty{}

	res, err := cli.GetTopStories(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
