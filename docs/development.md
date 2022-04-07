# Development 

## Running agent without geth

`data/redis` has a redis configuration file (which points to rdb file present in the directory). Starting the redis server with this configuration will load the rdb messages into the redis stream and the agent can run off that (rather than running bsp-geth).

```bash
>> cd data/redis
>> redis-server redis.conf
```