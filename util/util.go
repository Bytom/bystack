package util

import (
	"context"

	"github.com/bytom/bystack/api"
	"github.com/bytom/bystack/blockchain/rpc"
	"github.com/bytom/bystack/env"
	jww "github.com/spf13/jwalterweatherman"
)

const (
	// Success indicates the rpc calling is successful.
	Success = iota
	// ErrLocalExe indicates error occurs before the rpc calling.
	ErrLocalExe
	// ErrConnect indicates error occurs connecting to the bystackd, e.g.,
	// bystackd can't parse the received arguments.
	ErrConnect
	// ErrLocalParse indicates error occurs locally when parsing the response.
	ErrLocalParse
	// ErrRemote indicates error occurs in bystackd.
	ErrRemote
)

var (
	coreURL = env.String("BYTOM_URL", "http://127.0.0.1:9888")
)

// Wraper rpc's client
func MustRPCClient() *rpc.Client {
	env.Parse()
	return &rpc.Client{BaseURL: *coreURL}
}

// Wrapper rpc call api.
func ClientCall(path string, req ...interface{}) (interface{}, int) {

	var response = &api.Response{}
	var request interface{}

	if req != nil {
		request = req[0]
	}

	client := MustRPCClient()
	client.Call(context.Background(), path, request, response)

	switch response.Status {
	case api.FAIL:
		jww.ERROR.Println(response.Msg)
		return nil, ErrRemote
	case "":
		jww.ERROR.Println("Unable to connect to the bystackd")
		return nil, ErrConnect
	}

	return response.Data, Success
}
