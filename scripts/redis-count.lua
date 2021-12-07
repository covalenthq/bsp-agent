local T = redis.call('XRANGE', KEYS[1], ARGV[1], ARGV[2])
local count = 0
for _ in pairs(T) do count = count + 1 end
return count

-- call with params <stream-key> , "<first-stream-id>" "<last-stream-id>" -- get to these with XINFO STREAM <stream-key>
-- redis-cli --eval redis-count.lua replication , 3  "1637349831400-35"
-- (integer) 6288
