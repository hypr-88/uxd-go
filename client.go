package uxd

import (
	"github.com/gagliardetto/solana-go/rpc"
	"go.uber.org/zap"
)

type Client struct {
	Env             Env
	RPC            *rpc.Client
	WebSocketURL   string
	Log            *zap.Logger
}

func NewClient(env ENV, rpcURL string , wsURL string) *Client {
	return &Client{
		Env: env,
		RPC: rpc.New(rpcURL),
		WebSocket: wsURL,
		Log: zap.NewNop(),
	}
}