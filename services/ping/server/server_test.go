package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/enginoid/monorepo-base/services/ping/proto"
)

func TestPing(t *testing.T) {
	server := NewServer()
	response, err := server.Ping(context.TODO(), &pb.PingRequest{
		Message: "hello",
	})

	assert.NoError(t, err)
	assert.Equal(t, &pb.PingReply{
		Message: "hello",
	}, response)
}
