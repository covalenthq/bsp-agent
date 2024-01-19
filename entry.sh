#!/bin/bash
if [ "$BLOCKCHAIN" == "elrond" ]
then
    timeout 120s ./bsp-agent --redis-url=redis://username:@redis:6379/0?topic=replication-2#replicate \
    --avro-codec-path=./codec/block-elrond.avsc \
    --binary-file-path=./bin/block-elrond/ \
    --proof-chain-address=0xea2ff902dbeeeccc828757b881b343f9316752e5\
    --consumer-timeout=15 \
    --websocket-urls="34.69.250.147:20000 34.69.250.147:20001 34.69.250.147:20002 34.69.250.147:20003"

else
    ./bsp-agent --redis-url=redis://username:@redis:6379/0?topic=replication-2#replicate  \
    --avro-codec-path=./codec/block-ethereum.avsc \
    --binary-file-path=./bin/block-ethereum/ \
    --block-divisor=3 \
    --proof-chain-address=0xea2ff902dbeeeccc828757b881b343f9316752e5 \
    --consumer-timeout=15 \
    --ipfs-pinner-server="http://ipfs-pinner:3001/"
fi
