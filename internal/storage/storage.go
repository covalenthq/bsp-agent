// Package storage contains all function related to local and cloud storage
package storage

import (
	"bytes"
	"context"
	"fmt"
	"os"

	pinapi "github.com/covalenthq/ipfs-pinner"
	pincore "github.com/covalenthq/ipfs-pinner/core"
	"github.com/covalenthq/ipfs-pinner/pinclient"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
)

const (
	uploadTimeout int64 = 50
)

// generateCidFor generates ipfs cid given some content
func generateCidFor(ctx context.Context, pinnode pinapi.PinnerNode, content []byte) (cid.Cid, error) {
	if pinnode == nil {
		return cid.Undef, fmt.Errorf("no pinner node")
	}

	rcid, err := pinnode.UnixfsService().GenerateDag(ctx, bytes.NewReader(content))
	if err != nil {
		return cid.Undef, fmt.Errorf("error generating dag: %w", err)
	}

	return rcid, nil
}

// getPinnerNode get pinner node (web3.storage or pinata supported for now)
func getPinnerNode(pst pincore.PinningService, token string) (pinapi.PinnerNode, error) {
	var pinnode pinapi.PinnerNode
	switch pst {
	case pincore.Pinata, pincore.Web3Storage:
		clientCreateReq := pinclient.NewClientRequest(pst).BearerToken(token)
		cidComputationOnly := (pst == pincore.Pinata)
		nodeCreateReq := pinapi.NewNodeRequest(clientCreateReq).CidVersion(0).CidComputeOnly(cidComputationOnly)
		pinnode = pinapi.NewPinnerNode(*nodeCreateReq)

		return pinnode, nil
	case "":
		return nil, nil
	case pincore.Other:
		fallthrough
	default:
		return nil, fmt.Errorf("unsupported pinning service: %s", pst)
	}
}

func generateCarFile(ctx context.Context, pinnode pinapi.PinnerNode, ccid cid.Cid) (*os.File, error) {
	carf, err := os.CreateTemp(os.TempDir(), "*.car")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	log.Printf("car file location: %s\n", carf.Name())

	err = pinnode.CarExporter().Export(ctx, ccid, carf)
	if err != nil {
		_ = carf.Close()

		return nil, fmt.Errorf("%w", err)
	}

	noffset, err := carf.Seek(0, 0) // reset for Read
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	} else if noffset != 0 {
		return nil, fmt.Errorf("couldn't offset the car file to start")
	}

	return carf, nil
}
