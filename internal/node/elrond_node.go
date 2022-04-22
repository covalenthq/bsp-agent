package node

import (
	"context"

	"github.com/covalenthq/bsp-agent/internal/config"
)

type elrondAgentNode struct {
	agentNode
}

func newElrondAgentNode(aconfig *config.AgentConfig) *elrondAgentNode {
	// TODO
	return &elrondAgentNode{}
}

func (node *elrondAgentNode) NodeChainType() chainType {
	return Elrond
}

func (node *elrondAgentNode) Start(ctx context.Context) {

}
func (node *elrondAgentNode) StopProcessing() {

}

func (node *elrondAgentNode) Close() {

}
