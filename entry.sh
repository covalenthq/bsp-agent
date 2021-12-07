#!/bin/bash

sleep 80

./main --redis-url=redis://username:@redis:6379/0?topic=replication#replicate  \
 --codec-path=./codec/block-replica.avsc \
 --binary-file-path=./bin/block-replica/ \
 --gcp-svc-account=/Users/user/.config/gcloud/bsp.json \
 --replica-bucket=covalenthq-geth-block-specimen \
 --segment-length=10 \
 --eth-client=http://ganache-cli:8545  \
 --proof-chain-address=0xEa2ff902dbeEECcc828757B881b343F9316752e5 \
 --consumer-timeout=80
