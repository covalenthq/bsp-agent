local streamLen = redis.call('XLEN', KEYS[1])
if streamLen >= tonumber(ARGV[1]) then 
    return redis.call('XTRIM', KEYS[1], 'MAXLEN', streamLen - tonumber(ARGV[1]))
else
    return redis.error_reply(KEYS[1]..' has less than '..ARGV[1]..' items')
end

-- call with params <stream-key> , <number-of-elements-to-trim-from-start>[int] --get to these with XINFO STREAM <stream-key>
-- redis-cli --eval redis-trim.lua replication , 10
-- (integer) 10