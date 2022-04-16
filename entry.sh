#!/bin/bash
if [ "$BLOCKCHAIN" == "elrond" ]
then
    timeout 120s ./bsp-agent --redis-url=redis://username:@redis:6379/0?topic=replication-1#replicate \
    --avro-codec-path=./codec/block-elrond.avsc \
    --binary-file-path=./bin/block-elrond/ \
    --proof-chain-address=0xea2ff902dbeeeccc828757b881b343f9316752e5\
    --consumer-timeout=15 \
    --websocket-urls="34.66.210.112:20000 34.66.210.112:20001 34.66.210.112:20002 34.66.210.112:20003"

else
    ./bsp-agent --redis-url=redis://username:@redis:6379/0?topic=replication-1#replicate  \
    --avro-codec-path=./codec/block-ethereum.avsc \
    --binary-file-path=./bin/block-ethereum/ \
    --block-divisor=3 \
    --proof-chain-address=0xea2ff902dbeeeccc828757b881b343f9316752e5 \
    --consumer-timeout=15
fi
