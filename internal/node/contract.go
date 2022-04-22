package node

import "context"

type AgentNode interface {
	NodeChainType() ChainType
	Start(ctx context.Context)
	StopProcessing()
	Close()
}
