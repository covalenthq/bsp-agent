#!/bin/bash
if [ "$BLOCKCHAIN" == "elrond" ]
then
    timeout 120s ./bsp-agent --redis-url=redis://username:@redis:6379/0?topic=replication#replicate \
    --avro-codec-path=./codec/block-elrond.avsc \
    --binary-file-path=./bin/block-elrond/ \
    --replica-bucket=covalenthq-geth-block-specimen \
    --segment-length=10 \
    --proof-chain-address=0xEa2ff902dbeEECcc828757B881b343F9316752e5 \
    --consumer-timeout=60 \
    --websocket-urls="34.66.210.112:20000 34.66.210.112:20001 34.66.210.112:20002 34.66.210.112:20003" 
else
    ./bsp-agent --redis-url=redis://username:@redis:6379/0?topic=replication#replicate  \
    --avro-codec-path=./codec/block-ethereum.avsc \
    --binary-file-path=./bin/block-ethereum/ \
    --replica-bucket=covalenthq-geth-block-specimen \
    --segment-length=10 \
    --proof-chain-address=0xEa2ff902dbeEECcc828757B881b343F9316752e5 \
    --consumer-timeout=60
fi
