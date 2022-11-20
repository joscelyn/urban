package server

import (
	"testing"

	"context"

	"github.com/joscelyn/urban"
	"github.com/stretchr/testify/assert"
)

func TestWhois(t *testing.T) {
	ctx := context.Background()
	req := &urban.WhoisRequest{Username: "fra"}

	res, err := cli.Whois(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestEmptyWhois(t *testing.T) {
	ctx := context.Background()
	req := &urban.WhoisRequest{}

	res, err := cli.Whois(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}
