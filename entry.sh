#!/bin/bash
if [ "$BLOCKCHAIN" == "elrond" ]
then
    timeout 120s ./bsp-agent --redis-url=redis://username:@redis:6379/0?topic=replication-1#replicate \
    --avro-codec-path=./codec/block-elrond.avsc \
    --binary-file-path=./bin/block-elrond/ \
    --segment-length=1 \
    --proof-chain-address=0x6fbe1051956b1b2416c49e5e5076588fd4f072a1\
    --consumer-timeout=15 \
    --websocket-urls="34.66.210.112:20000 34.66.210.112:20001 34.66.210.112:20002 34.66.210.112:20003" 
else
    ./bsp-agent --redis-url=redis://username:@redis:6379/0?topic=replication-1#replicate  \
    --avro-codec-path=./codec/block-ethereum.avsc \
    --binary-file-path=./bin/block-ethereum/ \
    --segment-length=1 \
    --proof-chain-address=0x6fbe1051956b1b2416c49e5e5076588fd4f072a1 \
    --consumer-timeout=15
fi
