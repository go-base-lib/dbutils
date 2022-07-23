package redis

import (
	redisV8 "github.com/go-redis/redis/v8"
	"time"
)

func GetClient() *redisV8.Client {
	return client.GetClient()
}

func Wait(numSlaves int, timeout time.Duration) (int64, error) {
	return client.Wait(numSlaves, timeout)
}

func Command() (map[string]*redisV8.CommandInfo, error) {
	return client.Command()
}

// ClientGetName returns the name of the connection.
func ClientGetName() (string, error) {
	return client.ClientGetName()
}

func Echo(message interface{}) (string, error) {
	return client.Echo(message)
}

func Ping() (string, error) {
	return client.Ping()
}

func Quit() (string, error) {
	return client.Quit()
}

func Del(keys ...string) (int64, error) {
	return client.Del(keys...)
}

func Unlink(keys ...string) (int64, error) {
	return client.Unlink(keys...)
}

func Dump(key string) (string, error) {
	return client.Dump(key)
}

func Exists(keys ...string) (int64, error) {
	return client.Exists(keys...)
}

func Expire(key string, expiration time.Duration) (bool, error) {
	return client.Expire(key, expiration)
}

func ExpireAt(key string, tm time.Time) (bool, error) {
	return client.ExpireAt(key, tm)
}

func Keys(pattern string) ([]string, error) {
	return client.Keys(pattern)
}

func Migrate(host, port, key string, db int, timeout time.Duration) (string, error) {
	return client.Migrate(host, port, key, db, timeout)
}

func Move(key string, db int) (bool, error) {
	return client.Move(key, db)
}

func ObjectRefCount(key string) (int64, error) {
	return client.ObjectRefCount(key)
}

func ObjectEncoding(key string) (string, error) {
	return client.ObjectEncoding(key)
}

func ObjectIdleTime(key string) (time.Duration, error) {
	return client.ObjectIdleTime(key)
}

func Persist(key string) (bool, error) {
	return client.Persist(key)
}

func PExpire(key string, expiration time.Duration) (bool, error) {
	return client.PExpire(key, expiration)
}

func PExpireAt(key string, tm time.Time) (bool, error) {
	return client.PExpireAt(key, tm)
}

func PTTL(key string) (time.Duration, error) {
	return client.PTTL(key)
}

func RandomKey() (string, error) {
	return client.RandomKey()
}

func Rename(key, newkey string) (string, error) {
	return client.Rename(key, newkey)
}

func RenameNX(key, newkey string) (bool, error) {
	return client.RenameNX(key, newkey)
}

func Restore(key string, ttl time.Duration, value string) (string, error) {
	return client.Restore(key, ttl, value)
}

func RestoreReplace(key string, ttl time.Duration, value string) (string, error) {
	return client.RestoreReplace(key, ttl, value)
}

func Sort(key string, sort *redisV8.Sort) ([]string, error) {
	return client.Sort(key, sort)
}

func SortStore(key, store string, sort *redisV8.Sort) (int64, error) {
	return client.SortStore(key, store, sort)
}

func SortInterfaces(key string, sort *redisV8.Sort) ([]interface{}, error) {
	return client.SortInterfaces(key, sort)
}

func Touch(keys ...string) (int64, error) {
	return client.Touch(keys...)
}

func TTL(key string) (time.Duration, error) {
	return client.TTL(key)
}

func Type(key string) (string, error) {
	return client.Type(key)
}

func Append(key, value string) (int64, error) {
	return client.Append(key, value)
}

func Decr(key string) (int64, error) {
	return client.Decr(key)
}

func DecrBy(key string, decrement int64) (int64, error) {
	return client.DecrBy(key, decrement)
}

// Get Redis `GET key` command. It returns redis.Nil error when key does not exist.
func Get(key string) (string, error) {
	return client.Get(key)
}

func GetRange(key string, start, end int64) (string, error) {
	return client.GetRange(key, start, end)
}

func GetSet(key string, value interface{}) (string, error) {
	return client.GetSet(key, value)
}

// GetEx An expiration of zero removes the TTL associated with the key (i.e. GETEX key persist).
// Requires Redis >= 6.2.0.
func GetEx(key string, expiration time.Duration) (string, error) {
	return client.GetEx(key, expiration)
}

// GetDel redis-server version >= 6.2.0.
func GetDel(key string) (string, error) {
	return client.GetDel(key)
}

func Incr(key string) (int64, error) {
	return client.Incr(key)
}

func IncrBy(key string, value int64) (int64, error) {
	return client.IncrBy(key, value)
}

func IncrByFloat(key string, value float64) (float64, error) {
	return client.IncrByFloat(key, value)
}

func MGet(keys ...string) ([]interface{}, error) {
	return client.MGet(keys...)
}

func MSet(values ...interface{}) (string, error) {
	return client.MSet(values...)
}

// MSetNX is like SetNX but accepts multiple values:
//   - MSetNX("key1", "value1", "key2", "value2")
//   - MSetNX([]string{"key1", "value1", "key2", "value2"})
//   - MSetNX(map[string]interface{}{"key1": "value1", "key2": "value2"})
func MSetNX(values ...interface{}) (bool, error) {
	return client.MSetNX(values...)
}

func Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return client.Set(key, value, expiration)
}

func SetArgs(key string, value interface{}, a redisV8.SetArgs) (string, error) {
	return client.SetArgs(key, value, a)
}

// SetEX Redis `SETEX key expiration value` command.
func SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	return client.SetEX(key, value, expiration)
}

// SetNX Redis `SET key value [expiration] NX` command.
//
// Zero expiration means the key has no expiration time.
// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
// otherwise you will receive an error: (error) ERR syntax error.
func SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return client.SetNX(key, value, expiration)
}

// SetXX Redis `SET key value [expiration] XX` command.
//
// Zero expiration means the key has no expiration time.
// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
// otherwise you will receive an error: (error) ERR syntax error.
func SetXX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return client.SetXX(key, value, expiration)
}

func SetRange(key string, offset int64, value string) (int64, error) {
	return client.SetRange(key, offset, value)
}

func StrLen(key string) (int64, error) {
	return client.StrLen(key)
}

//------------------------------------------------------------------------------

func GetBit(key string, offset int64) (int64, error) {
	return client.GetBit(key, offset)
}

func SetBit(key string, offset int64, value int) (int64, error) {
	return client.SetBit(key, offset, value)
}

func BitCount(key string, bitCount *redisV8.BitCount) (int64, error) {
	return client.BitCount(key, bitCount)
}

func BitOpAnd(destKey string, keys ...string) (int64, error) {
	return client.BitOpAnd(destKey, keys...)
}

func BitOpOr(destKey string, keys ...string) (int64, error) {
	return client.BitOpOr(destKey, keys...)
}

func BitOpXor(destKey string, keys ...string) (int64, error) {
	return client.BitOpXor(destKey, keys...)
}

func BitOpNot(destKey string, key string) (int64, error) {
	return client.BitOpNot(destKey, key)
}

func BitPos(key string, bit int64, pos ...int64) (int64, error) {
	return client.BitPos(key, bit, pos...)
}

func BitField(key string, args ...interface{}) ([]int64, error) {
	return client.BitField(key, args...)
}

//------------------------------------------------------------------------------

func Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.Scan(cursor, match, count)
}

func ScanType(cursor uint64, match string, count int64, keyType string) ([]string, uint64, error) {
	return client.ScanType(cursor, match, count, keyType)
}

func SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.SScan(key, cursor, match, count)
}

func HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.HScan(key, cursor, match, count)
}

func ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return client.ZScan(key, cursor, match, count)
}

//------------------------------------------------------------------------------

func HDel(key string, fields ...string) (int64, error) {
	return client.HDel(key, fields...)
}

func HExists(key, field string) (bool, error) {
	return client.HExists(key, field)
}

func HGet(key, field string) (string, error) {
	return client.HGet(key, field)
}

func HGetAll(key string) (map[string]string, error) {
	return client.HGetAll(key)
}

func HIncrBy(key, field string, incr int64) (int64, error) {
	return client.HIncrBy(key, field, incr)
}

func HIncrByFloat(key, field string, incr float64) (float64, error) {
	return client.HIncrByFloat(key, field, incr)
}

func HKeys(key string) ([]string, error) {
	return client.HKeys(key)
}

func HLen(key string) (int64, error) {
	return client.HLen(key)
}

// HMGet returns the values for the specified fields in the hash stored at key.
// It returns an interface{} to distinguish between empty string and nil value.
func HMGet(key string, fields ...string) ([]interface{}, error) {
	return client.HMGet(key, fields...)
}

// HSet accepts values in following formats:
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
//
// Note that it requires Redis v4 for multiple field/value pairs support.
func HSet(key string, values ...interface{}) (int64, error) {
	return client.HSet(key, values...)
}

// HMSet is a deprecated version of HSet left for compatibility with Redis 3.
func HMSet(key string, values ...interface{}) (bool, error) {
	return client.HMSet(key, values...)
}

func HSetNX(key, field string, value interface{}) (bool, error) {
	return client.HSetNX(key, field, value)
}

func HVals(key string) ([]string, error) {
	return client.HVals(key)
}

// HRandField redis-server version >= 6.2.0.
func HRandField(key string, count int, withValues bool) ([]string, error) {
	return client.HRandField(key, count, withValues)
}

//------------------------------------------------------------------------------

func BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return client.BLPop(timeout, keys...)
}

func BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	return client.BRPop(timeout, keys...)
}

func BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	return client.BRPopLPush(source, destination, timeout)
}

func LIndex(key string, index int64) (string, error) {
	return client.LIndex(key, index)
}

func LInsert(key, op string, pivot, value interface{}) (int64, error) {
	return client.LInsert(key, op, pivot, value)
}

func LInsertBefore(key string, pivot, value interface{}) (int64, error) {
	return client.LInsertBefore(key, pivot, value)
}

func LInsertAfter(key string, pivot, value interface{}) (int64, error) {
	return client.LInsertAfter(key, pivot, value)
}

func LLen(key string) (int64, error) {
	return client.LLen(key)
}

func LPop(key string) (string, error) {
	return client.LPop(key)
}

func LPopCount(key string, count int) ([]string, error) {
	return client.LPopCount(key, count)
}

func LPos(key string, value string, a redisV8.LPosArgs) (int64, error) {
	return client.LPos(key, value, a)
}

func LPosCount(key string, value string, count int64, a redisV8.LPosArgs) ([]int64, error) {
	return client.LPosCount(key, value, count, a)
}

func LPush(key string, values ...interface{}) (int64, error) {
	return client.LPush(key, values...)
}

func LPushX(key string, values ...interface{}) (int64, error) {
	return client.LPushX(key, values...)
}

func LRange(key string, start, stop int64) ([]string, error) {
	return client.LRange(key, start, stop)
}

func LRem(key string, count int64, value interface{}) (int64, error) {
	return client.LRem(key, count, value)
}

func LSet(key string, index int64, value interface{}) (string, error) {
	return client.LSet(key, index, value)
}

func LTrim(key string, start, stop int64) (string, error) {
	return client.LTrim(key, start, stop)
}

func RPop(key string) (string, error) {
	return client.RPop(key)
}

func RPopCount(key string, count int) ([]string, error) {
	return client.RPopCount(key, count)
}

func RPopLPush(source, destination string) (string, error) {
	return client.RPopLPush(source, destination)
}

func RPush(key string, values ...interface{}) (int64, error) {
	return client.RPush(key, values...)
}

func RPushX(key string, values ...interface{}) (int64, error) {
	return client.RPushX(key, values...)
}

func LMove(source, destination, srcpos, destpos string) (string, error) {
	return client.LMove(source, destination, srcpos, destpos)
}

//------------------------------------------------------------------------------

func SAdd(key string, members ...interface{}) (int64, error) {
	return client.SAdd(key, members...)
}

func SCard(key string) (int64, error) {
	return client.SCard(key)
}

func SDiff(keys ...string) ([]string, error) {
	return client.SDiff(keys...)
}

func SDiffStore(destination string, keys ...string) (int64, error) {
	return client.SDiffStore(destination, keys...)
}

func SInter(keys ...string) ([]string, error) {
	return client.SInter(keys...)
}

func SInterStore(destination string, keys ...string) (int64, error) {
	return client.SInterStore(destination, keys...)
}

func SIsMember(key string, member interface{}) (bool, error) {
	return client.SIsMember(key, member)
}

// SMIsMember Redis `SMISMEMBER key member [member ...]` command.
func SMIsMember(key string, members ...interface{}) ([]bool, error) {
	return client.SMIsMember(key, members...)
}

// SMembers Redis `SMEMBERS key` command output as a slice.
func SMembers(key string) ([]string, error) {
	return client.SMembers(key)
}

// SMembersMap Redis `SMEMBERS key` command output as a map.
func SMembersMap(key string) (map[string]struct{}, error) {
	return client.SMembersMap(key)
}

func SMove(source, destination string, member interface{}) (bool, error) {
	return client.SMove(source, destination, member)
}

// SPop Redis `SPOP key` command.
func SPop(key string) (string, error) {
	return client.SPop(key)
}

// SPopN Redis `SPOP key count` command.
func SPopN(key string, count int64) ([]string, error) {
	return client.SPopN(key, count)
}

// SRandMember Redis `SRANDMEMBER key` command.
func SRandMember(key string) (string, error) {
	return client.SRandMember(key)
}

// SRandMemberN Redis `SRANDMEMBER key count` command.
func SRandMemberN(key string, count int64) ([]string, error) {
	return client.SRandMemberN(key, count)
}

func SRem(key string, members ...interface{}) (int64, error) {
	return client.SRem(key, members...)
}

func SUnion(keys ...string) ([]string, error) {
	return client.SUnion(keys...)
}

func SUnionStore(destination string, keys ...string) (int64, error) {
	return client.SUnionStore(destination, keys...)
}

// XAdd a.Limit has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
func XAdd(a *redisV8.XAddArgs) (string, error) {
	return client.XAdd(a)
}

func XDel(stream string, ids ...string) (int64, error) {
	return client.XDel(stream, ids...)
}

func XLen(stream string) (int64, error) {
	return client.XLen(stream)
}

func XRange(stream, start, stop string) ([]redisV8.XMessage, error) {
	return client.XRange(stream, start, stop)
}

func XRangeN(stream, start, stop string, count int64) ([]redisV8.XMessage, error) {
	return client.XRangeN(stream, start, stop, count)
}

func XRevRange(stream, start, stop string) ([]redisV8.XMessage, error) {
	return client.XRevRange(stream, start, stop)
}

func XRevRangeN(stream, start, stop string, count int64) ([]redisV8.XMessage, error) {
	return client.XRevRangeN(stream, start, stop, count)
}

func XRead(a *redisV8.XReadArgs) ([]redisV8.XStream, error) {
	return client.XRead(a)
}

func XReadStreams(streams ...string) ([]redisV8.XStream, error) {
	return client.XReadStreams(streams...)
}

func XGroupCreate(stream, group, start string) (string, error) {
	return client.XGroupCreate(stream, group, start)
}

func XGroupCreateMkStream(stream, group, start string) (string, error) {
	return client.XGroupCreateMkStream(stream, group, start)
}

func XGroupSetID(stream, group, start string) (string, error) {
	return client.XGroupSetID(stream, group, start)
}

func XGroupDestroy(stream, group string) (int64, error) {
	return client.XGroupDestroy(stream, group)
}

func XGroupCreateConsumer(stream, group, consumer string) (int64, error) {
	return client.XGroupCreateConsumer(stream, group, consumer)
}

func XGroupDelConsumer(stream, group, consumer string) (int64, error) {
	return client.XGroupDelConsumer(stream, group, consumer)
}

func XReadGroup(a *redisV8.XReadGroupArgs) ([]redisV8.XStream, error) {
	return client.XReadGroup(a)
}

func XAck(stream, group string, ids ...string) (int64, error) {
	return client.XAck(stream, group, ids...)
}

func XPending(stream, group string) (*redisV8.XPending, error) {
	return client.XPending(stream, group)
}

func XPendingExt(a *redisV8.XPendingExtArgs) ([]redisV8.XPendingExt, error) {
	return client.XPendingExt(a)
}

func XAutoClaim(a *redisV8.XAutoClaimArgs) (messages []redisV8.XMessage, start string, err error) {
	return client.XAutoClaim(a)
}

func XAutoClaimJustID(a *redisV8.XAutoClaimArgs) (ids []string, start string, err error) {
	return client.XAutoClaimJustID(a)
}

func XClaim(a *redisV8.XClaimArgs) ([]redisV8.XMessage, error) {
	return client.XClaim(a)
}

func XClaimJustID(a *redisV8.XClaimArgs) ([]string, error) {
	return client.XClaimJustID(a)
}

// XTrimMaxLen No `~` rules are used, `limit` cannot be used.
// cmd: XTRIM key MAXLEN maxLen
func XTrimMaxLen(key string, maxLen int64) (int64, error) {
	return client.XTrimMaxLen(key, maxLen)
}

// XTrimMaxLenApprox LIMIT has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
// cmd: XTRIM key MAXLEN ~ maxLen LIMIT limit
func XTrimMaxLenApprox(key string, maxLen, limit int64) (int64, error) {
	return client.XTrimMaxLenApprox(key, maxLen, limit)
}

// XTrimMinID No `~` rules are used, `limit` cannot be used.
// cmd: XTRIM key MINID minID
func XTrimMinID(key string, minID string) (int64, error) {
	return client.XTrimMinID(key, minID)
}

// XTrimMinIDApprox LIMIT has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
// cmd: XTRIM key MINID ~ minID LIMIT limit
func XTrimMinIDApprox(key string, minID string, limit int64) (int64, error) {
	return client.XTrimMinIDApprox(key, minID, limit)
}

func XInfoConsumers(key string, group string) ([]redisV8.XInfoConsumer, error) {
	return client.XInfoConsumers(key, group)
}

func XInfoGroups(key string) ([]redisV8.XInfoGroup, error) {
	return client.XInfoGroups(key)
}

func XInfoStream(key string) (*redisV8.XInfoStream, error) {
	return client.XInfoStream(key)
}

// XInfoStreamFull XINFO STREAM FULL [COUNT count]
// redis-server >= 6.0.
func XInfoStreamFull(key string, count int) (*redisV8.XInfoStreamFull, error) {
	return client.XInfoStreamFull(key, count)
}

//------------------------------------------------------------------------------

// BZPopMax Redis `BZPOPMAX key [key ...] timeout` command.
func BZPopMax(timeout time.Duration, keys ...string) (*redisV8.ZWithKey, error) {
	return client.BZPopMax(timeout, keys...)
}

// BZPopMin Redis `BZPOPMIN key [key ...] timeout` command.
func BZPopMin(timeout time.Duration, keys ...string) (*redisV8.ZWithKey, error) {
	return client.BZPopMin(timeout, keys...)
}

func ZAddArgs(key string, args redisV8.ZAddArgs) (int64, error) {
	return client.ZAddArgs(key, args)
}

func ZAddArgsIncr(key string, args redisV8.ZAddArgs) (float64, error) {
	return client.ZAddArgsIncr(key, args)
}

// ZAdd Redis `ZADD key score member [score member ...]` command.
func ZAdd(key string, members ...*redisV8.Z) (int64, error) {
	return client.ZAdd(key, members...)
}

// ZAddNX Redis `ZADD key NX score member [score member ...]` command.
func ZAddNX(key string, members ...*redisV8.Z) (int64, error) {
	return client.ZAddNX(key, members...)
}

// ZAddXX Redis `ZADD key XX score member [score member ...]` command.
func ZAddXX(key string, members ...*redisV8.Z) (int64, error) {
	return client.ZAddXX(key, members...)
}

func ZCard(key string) (int64, error) {
	return client.ZCard(key)
}

func ZCount(key, min, max string) (int64, error) {
	return client.ZCount(key, min, max)
}

func ZLexCount(key, min, max string) (int64, error) {
	return client.ZLexCount(key, min, max)
}

func ZIncrBy(key string, increment float64, member string) (float64, error) {
	return client.ZIncrBy(key, increment, member)
}

func ZInterStore(destination string, store *redisV8.ZStore) (int64, error) {
	return client.ZInterStore(destination, store)
}

func ZInter(store *redisV8.ZStore) ([]string, error) {
	return client.ZInter(store)
}

func ZInterWithScores(store *redisV8.ZStore) ([]redisV8.Z, error) {
	return client.ZInterWithScores(store)
}

func ZMScore(key string, members ...string) ([]float64, error) {
	return client.ZMScore(key, members...)
}

func ZPopMax(key string, count ...int64) ([]redisV8.Z, error) {
	return client.ZPopMax(key, count...)
}

func ZPopMin(key string, count ...int64) ([]redisV8.Z, error) {
	return client.ZPopMin(key, count...)
}

func ZRangeArgs(z redisV8.ZRangeArgs) ([]string, error) {
	return client.ZRangeArgs(z)
}

func ZRangeArgsWithScores(z redisV8.ZRangeArgs) ([]redisV8.Z, error) {
	return client.ZRangeArgsWithScores(z)
}

func ZRange(key string, start, stop int64) ([]string, error) {
	return client.ZRange(key, start, stop)
}

func ZRangeWithScores(key string, start, stop int64) ([]redisV8.Z, error) {
	return client.ZRangeWithScores(key, start, stop)
}

func ZRangeByScore(key string, opt *redisV8.ZRangeBy) ([]string, error) {
	return client.ZRangeByScore(key, opt)
}

func ZRangeByLex(key string, opt *redisV8.ZRangeBy) ([]string, error) {
	return client.ZRangeByLex(key, opt)
}

func ZRangeByScoreWithScores(key string, opt *redisV8.ZRangeBy) ([]redisV8.Z, error) {
	return client.ZRangeByScoreWithScores(key, opt)
}

func ZRangeStore(dst string, z redisV8.ZRangeArgs) (int64, error) {
	return client.ZRangeStore(dst, z)
}

func ZRank(key, member string) (int64, error) {
	return client.ZRank(key, member)
}

func ZRem(key string, members ...interface{}) (int64, error) {
	return client.ZRem(key, members...)
}

func ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	return client.ZRemRangeByRank(key, start, stop)
}

func ZRemRangeByScore(key, min, max string) (int64, error) {
	return client.ZRemRangeByScore(key, min, max)
}

func ZRemRangeByLex(key, min, max string) (int64, error) {
	return client.ZRemRangeByLex(key, min, max)
}

func ZRevRange(key string, start, stop int64) ([]string, error) {
	return client.ZRevRange(key, start, stop)
}

func ZRevRangeWithScores(key string, start, stop int64) ([]redisV8.Z, error) {
	return client.ZRevRangeWithScores(key, start, stop)
}

func ZRevRangeByScore(key string, opt *redisV8.ZRangeBy) ([]string, error) {
	return client.ZRevRangeByScore(key, opt)
}

func ZRevRangeByLex(key string, opt *redisV8.ZRangeBy) ([]string, error) {
	return client.ZRevRangeByLex(key, opt)
}

func ZRevRangeByScoreWithScores(key string, opt *redisV8.ZRangeBy) ([]redisV8.Z, error) {
	return client.ZRevRangeByScoreWithScores(key, opt)
}

func ZRevRank(key, member string) (int64, error) {
	return client.ZRevRank(key, member)
}

func ZScore(key, member string) (float64, error) {
	return client.ZScore(key, member)
}

func ZUnion(store redisV8.ZStore) ([]string, error) {
	return client.ZUnion(store)
}

func ZUnionWithScores(store redisV8.ZStore) ([]redisV8.Z, error) {
	return client.ZUnionWithScores(store)
}

func ZUnionStore(dest string, store *redisV8.ZStore) (int64, error) {
	return client.ZUnionStore(dest, store)
}

// ZRandMember redis-server version >= 6.2.0.
func ZRandMember(key string, count int, withScores bool) ([]string, error) {
	return client.ZRandMember(key, count, withScores)
}

// ZDiff redis-server version >= 6.2.0.
func ZDiff(keys ...string) ([]string, error) {
	return client.ZDiff(keys...)
}

// ZDiffWithScores redis-server version >= 6.2.0.
func ZDiffWithScores(keys ...string) ([]redisV8.Z, error) {
	return client.ZDiffWithScores(keys...)
}

// ZDiffStore redis-server version >=6.2.0.
func ZDiffStore(destination string, keys ...string) (int64, error) {
	return client.ZDiffStore(destination, keys...)
}

//------------------------------------------------------------------------------

func PFAdd(key string, els ...interface{}) (int64, error) {
	return client.PFAdd(key, els)
}

func PFCount(keys ...string) (int64, error) {
	return client.PFCount(keys...)
}

func PFMerge(dest string, keys ...string) (string, error) {
	return client.PFMerge(dest, keys...)
}

//------------------------------------------------------------------------------

func BgRewriteAOF() (string, error) {
	return client.BgRewriteAOF()
}

func BgSave() (string, error) {
	return client.BgSave()
}

func ClientKill(ipPort string) (string, error) {
	return client.ClientKill(ipPort)
}

// ClientKillByFilter is new style syntax, while the ClientKill is old
//
//   CLIENT KILL <option> [value] ... <option> [value]
func ClientKillByFilter(keys ...string) (int64, error) {
	return client.ClientKillByFilter(keys...)
}

func ClientList() (string, error) {
	return client.ClientList()
}

func ClientPause(dur time.Duration) (bool, error) {
	return client.ClientPause(dur)
}

func ClientID() (int64, error) {
	return client.ClientID()
}

func ClientUnblock(id int64) (int64, error) {
	return client.ClientUnblock(id)
}

func ClientUnblockWithError(id int64) (int64, error) {
	return client.ClientUnblockWithError(id)
}

func ConfigGet(parameter string) ([]interface{}, error) {
	return client.ConfigGet(parameter)
}

func ConfigResetStat() (string, error) {
	return client.ConfigResetStat()
}

func ConfigSet(parameter, value string) (string, error) {
	return client.ConfigSet(parameter, value)
}

func ConfigRewrite() (string, error) {
	return client.ConfigRewrite()
}

func DBSize() (int64, error) {
	return client.DBSize()
}

func FlushAll() (string, error) {
	return client.FlushAll()
}

func FlushAllAsync() (string, error) {
	return client.FlushAllAsync()
}

func FlushDB() (string, error) {
	return client.FlushDB()
}

func FlushDBAsync() (string, error) {
	return client.FlushDBAsync()
}

func Info(section ...string) (string, error) {
	return client.Info(section...)
}

func LastSave() (int64, error) {
	return client.LastSave()
}

func Save() (string, error) {
	return client.Save()
}

func Shutdown() (string, error) {
	return client.Shutdown()
}

func ShutdownSave() (string, error) {
	return client.ShutdownSave()
}

func ShutdownNoSave() (string, error) {
	return client.ShutdownNoSave()
}

func SlaveOf(host, port string) (string, error) {
	return client.SlaveOf(host, port)
}

func SlowLogGet(num int64) ([]redisV8.SlowLog, error) {
	return client.SlowLogGet(num)
}

func Time() (time.Time, error) {
	return client.Time()
}

func DebugObject(key string) (string, error) {
	return client.DebugObject(key)
}

func ReadOnly() (string, error) {
	return client.ReadOnly()
}

func ReadWrite() (string, error) {
	return client.ReadWrite()
}

func MemoryUsage(key string, samples ...int) (int64, error) {
	return client.MemoryUsage(key, samples...)
}

//------------------------------------------------------------------------------

func Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	return client.Eval(script, keys, args...)
}

func EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return client.EvalSha(sha1, keys, args...)
}

func ScriptExists(hashes ...string) ([]bool, error) {
	return client.ScriptExists(hashes...)
}

func ScriptFlush() (string, error) {
	return client.ScriptFlush()
}

func ScriptKill() (string, error) {
	return client.ScriptKill()
}

func ScriptLoad(script string) (string, error) {
	return client.ScriptLoad(script)
}

//------------------------------------------------------------------------------

// Publish posts the message to the channel.
func Publish(channel string, message interface{}) (int64, error) {
	return client.Publish(channel, message)
}

func PubSubChannels(pattern string) ([]string, error) {
	return client.PubSubChannels(pattern)
}

func PubSubNumSub(channels ...string) (map[string]int64, error) {
	return client.PubSubNumSub(channels...)
}

func PubSubNumPat() (int64, error) {
	return client.PubSubNumPat()
}

//------------------------------------------------------------------------------

func ClusterSlots() ([]redisV8.ClusterSlot, error) {
	return client.ClusterSlots()
}

func ClusterNodes() (string, error) {
	return client.ClusterNodes()
}

func ClusterMeet(host, port string) (string, error) {
	return client.ClusterMeet(host, port)
}

func ClusterForget(nodeID string) (string, error) {
	return client.ClusterForget(nodeID)
}

func ClusterReplicate(nodeID string) (string, error) {
	return client.ClusterReplicate(nodeID)
}

func ClusterResetSoft() (string, error) {
	return client.ClusterResetSoft()
}

func ClusterResetHard() (string, error) {
	return client.ClusterResetHard()
}

func ClusterInfo() (string, error) {
	return client.ClusterInfo()
}

func ClusterKeySlot(key string) (int64, error) {
	return client.ClusterKeySlot(key)
}

func ClusterGetKeysInSlot(slot int, count int) ([]string, error) {
	return client.ClusterGetKeysInSlot(slot, count)
}

func ClusterCountFailureReports(nodeID string) (int64, error) {
	return client.ClusterCountFailureReports(nodeID)
}

func ClusterCountKeysInSlot(slot int) (int64, error) {
	return client.ClusterCountKeysInSlot(slot)
}

func ClusterDelSlots(slots ...int) (string, error) {
	return client.ClusterDelSlots(slots...)
}

func ClusterDelSlotsRange(min, max int) (string, error) {
	return client.ClusterDelSlotsRange(min, max)
}

func ClusterSaveConfig() (string, error) {
	return client.ClusterSaveConfig()
}

func ClusterSlaves(nodeID string) ([]string, error) {
	return client.ClusterSlaves(nodeID)
}

func ClusterFailover() (string, error) {
	return client.ClusterFailover()
}

func ClusterAddSlots(slots ...int) (string, error) {
	return client.ClusterAddSlots(slots...)
}

func ClusterAddSlotsRange(min, max int) (string, error) {
	return client.ClusterAddSlotsRange(min, max)
}

//------------------------------------------------------------------------------

func GeoAdd(key string, geoLocation ...*redisV8.GeoLocation) (int64, error) {
	return client.GeoAdd(key, geoLocation...)
}

// GeoRadius is a read-only GEORADIUS_RO command.
func GeoRadius(key string, longitude, latitude float64, query *redisV8.GeoRadiusQuery) ([]redisV8.GeoLocation, error) {
	return client.GeoRadius(key, longitude, latitude, query)
}

// GeoRadiusStore is a writing GEORADIUS command.
func GeoRadiusStore(key string, longitude, latitude float64, query *redisV8.GeoRadiusQuery) (int64, error) {
	return client.GeoRadiusStore(key, longitude, latitude, query)
}

// GeoRadiusByMember is a read-only GEORADIUSBYMEMBER_RO command.
func GeoRadiusByMember(key, member string, query *redisV8.GeoRadiusQuery) ([]redisV8.GeoLocation, error) {
	return client.GeoRadiusByMember(key, member, query)
}

// GeoRadiusByMemberStore is a writing GEORADIUSBYMEMBER command.
func GeoRadiusByMemberStore(key, member string, query *redisV8.GeoRadiusQuery) (int64, error) {
	return client.GeoRadiusByMemberStore(key, member, query)
}

func GeoSearch(key string, q *redisV8.GeoSearchQuery) ([]string, error) {
	return client.GeoSearch(key, q)
}

func GeoSearchLocation(key string, q *redisV8.GeoSearchLocationQuery) ([]redisV8.GeoLocation, error) {
	return client.GeoSearchLocation(key, q)
}

func GeoSearchStore(key, store string, q *redisV8.GeoSearchStoreQuery) (int64, error) {
	return client.GeoSearchStore(key, store, q)
}

func GeoDist(key string, member1, member2, unit string) (float64, error) {
	return client.GeoDist(key, member1, member2, unit)
}

func GeoHash(key string, members ...string) ([]string, error) {
	return client.GeoHash(key, members...)
}

func GeoPos(key string, members ...string) ([]*redisV8.GeoPos, error) {
	return client.GeoPos(key, members...)
}
