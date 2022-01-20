#!/bin/bash
if [ "$BLOCKCHAIN" == "elrond" ]
then
    timeout 120s ./main --redis-url=redis://username:@redis:6379/0?topic=replication#replicate \
   	--codec-path=./codec/block-elrond.avsc \
    --binary-file-path=./bin/block-elrond/ \
    --gcp-svc-account=/Users/user/.config/gcloud/bsp.json \
    --replica-bucket=covalenthq-geth-block-specimen \
    --segment-length=10 \
 	--eth-client=http://ganache-cli:8545  \
 	--proof-chain-address=0xEa2ff902dbeEECcc828757B881b343F9316752e5 \
	--consumer-timeout=15 \
	--websocket-urls="34.66.210.112:20000 34.66.210.112:20001 34.66.210.112:20002 34.66.210.112:20003" 
else
    ./main --redis-url=redis://username:@redis:6379/0?topic=replication#replicate  \
    --codec-path=./codec/block-ethereum.avsc \
    --binary-file-path=./bin/block-ethereum/ \
    --gcp-svc-account=/Users/user/.config/gcloud/bsp.json \
    --replica-bucket=covalenthq-geth-block-specimen \
    --segment-length=10 \
    --eth-client=http://ganache-cli:8545  \
    --proof-chain-address=0xEa2ff902dbeEECcc828757B881b343F9316752e5 \
    --consumer-timeout=15
fi