./main --redis-url=redis://username:@redis:6379/0?topic=replication#replicate  \
 --codec-path=./codec/block-replica.avsc \
 --binary-file-path=./bin/block-replica/ \
 --gcp-svc-account=/Users/user/.config/gcloud/bsp.json \
 --replica-bucket=covalenthq-geth-block-specimen \
 --segment-length=10 \
 --eth-client=http://127.0.0.1:7545  \
 --proof-chain-address=0xca59d70517cc581E2277EdCA8587A0dd2BeC5eb9 \
 --consumer-timeout=80