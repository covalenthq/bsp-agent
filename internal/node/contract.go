// Package node contains structs/interfaces around the bsp-agent node, which aggregates all the services and orchestrates
// processing of redis block specimen messages.
package node

import "context"

// AgentNode defines the interface to interact with the bsp-agent node.
// The lifecycle of an AgentNode goes from start -> stop (processing) -> close.
// It is assumed that all the setup work has been done at Start()
type AgentNode interface {
	NodeChainType() ChainType
	Start(ctx context.Context)
	StopProcessing()
	Close()
}
