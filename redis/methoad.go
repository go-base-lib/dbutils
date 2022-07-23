package redis

import (
	dbutils "github.com/byzk-worker/go-db-utils"
	redisV8 "github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	context dbutils.GetContextFn
	c       *redisV8.Client
}

func (c *Client) GetClient() *redisV8.Client {
	return c.c
}

func (c *Client) Wait(numSlaves int, timeout time.Duration) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}

	result := c.c.Wait(ctx, numSlaves, timeout)
	return result.Result()
}

func (c *Client) Command() (map[string]*redisV8.CommandInfo, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.Command(ctx)
	return result.Result()
}

// ClientGetName returns the name of the connection.
func (c *Client) ClientGetName() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClientGetName(ctx)
	return result.Result()
}

func (c *Client) Echo(message interface{}) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Echo(ctx, message)
	return result.Result()
}

func (c *Client) Ping() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Ping(ctx)
	return result.Result()
}

func (c *Client) Quit() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Quit(ctx)
	return result.Result()
}

func (c *Client) Del(keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.Del(ctx, keys...)
	return result.Result()
}

func (c *Client) Unlink(keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.Unlink(ctx, keys...)
	return result.Result()
}

func (c *Client) Dump(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Dump(ctx, key)
	return result.Result()
}

func (c *Client) Exists(keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.Exists(ctx, keys...)
	return result.Result()
}

func (c *Client) Expire(key string, expiration time.Duration) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.Expire(ctx, key, expiration)
	return result.Result()
}

func (c *Client) ExpireAt(key string, tm time.Time) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.ExpireAt(ctx, key, tm)
	return result.Result()
}

func (c *Client) Keys(pattern string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.Keys(ctx, pattern)
	return result.Result()
}

func (c *Client) Migrate(host, port, key string, db int, timeout time.Duration) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Migrate(ctx, host, port, key, db, timeout)
	return result.Result()
}

func (c *Client) Move(key string, db int) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.Move(ctx, key, db)
	return result.Result()
}

func (c *Client) ObjectRefCount(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ObjectRefCount(ctx, key)
	return result.Result()
}

func (c *Client) ObjectEncoding(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ObjectEncoding(ctx, key)
	return result.Result()
}

func (c *Client) ObjectIdleTime(key string) (time.Duration, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ObjectIdleTime(ctx, key)
	return result.Result()
}

func (c *Client) Persist(key string) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.Persist(ctx, key)
	return result.Result()
}

func (c *Client) PExpire(key string, expiration time.Duration) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.PExpire(ctx, key, expiration)
	return result.Result()
}

func (c *Client) PExpireAt(key string, tm time.Time) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.PExpireAt(ctx, key, tm)
	return result.Result()
}

func (c *Client) PTTL(key string) (time.Duration, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.PTTL(ctx, key)
	return result.Result()
}

func (c *Client) RandomKey() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.RandomKey(ctx)
	return result.Result()
}

func (c *Client) Rename(key, newkey string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Rename(ctx, key, newkey)
	return result.Result()
}

func (c *Client) RenameNX(key, newkey string) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.RenameNX(ctx, key, newkey)
	return result.Result()
}

func (c *Client) Restore(key string, ttl time.Duration, value string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Restore(ctx, key, ttl, value)
	return result.Result()
}

func (c *Client) RestoreReplace(key string, ttl time.Duration, value string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.RestoreReplace(ctx, key, ttl, value)
	return result.Result()
}

func (c *Client) Sort(key string, sort *redisV8.Sort) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.Sort(ctx, key, sort)
	return result.Result()
}

func (c *Client) SortStore(key, store string, sort *redisV8.Sort) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SortStore(ctx, key, store, sort)
	return result.Result()
}

func (c *Client) SortInterfaces(key string, sort *redisV8.Sort) ([]interface{}, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SortInterfaces(ctx, key, sort)
	return result.Result()
}

func (c *Client) Touch(keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.Touch(ctx, keys...)
	return result.Result()
}

func (c *Client) TTL(key string) (time.Duration, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.TTL(ctx, key)
	return result.Result()
}

func (c *Client) Type(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Type(ctx, key)
	return result.Result()
}

func (c *Client) Append(key, value string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.Append(ctx, key, value)
	return result.Result()
}

func (c *Client) Decr(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.Decr(ctx, key)
	return result.Result()
}

func (c *Client) DecrBy(key string, decrement int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.DecrBy(ctx, key, decrement)
	return result.Result()
}

// Get Redis `GET key` command. It returns redis.Nil error when key does not exist.
func (c *Client) Get(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Get(ctx, key)
	return result.Result()
}

func (c *Client) GetRange(key string, start, end int64) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.GetRange(ctx, key, start, end)
	return result.Result()
}

func (c *Client) GetSet(key string, value interface{}) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.GetSet(ctx, key, value)
	return result.Result()
}

// GetEx An expiration of zero removes the TTL associated with the key (i.e. GETEX key persist).
// Requires Redis >= 6.2.0.
func (c *Client) GetEx(key string, expiration time.Duration) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.GetEx(ctx, key, expiration)
	return result.Result()
}

// GetDel redis-server version >= 6.2.0.
func (c *Client) GetDel(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.GetDel(ctx, key)
	return result.Result()
}

func (c *Client) Incr(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.Incr(ctx, key)
	return result.Result()
}

func (c *Client) IncrBy(key string, value int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.IncrBy(ctx, key, value)
	return result.Result()
}

func (c *Client) IncrByFloat(key string, value float64) (float64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.IncrByFloat(ctx, key, value)
	return result.Result()
}

func (c *Client) MGet(keys ...string) ([]interface{}, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.MGet(ctx, keys...)
	return result.Result()
}

func (c *Client) MSet(values ...interface{}) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.MSet(ctx, values...)
	return result.Result()
}

// MSetNX is like SetNX but accepts multiple values:
//   - MSetNX("key1", "value1", "key2", "value2")
//   - MSetNX([]string{"key1", "value1", "key2", "value2"})
//   - MSetNX(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (c *Client) MSetNX(values ...interface{}) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.MSetNX(ctx, values...)
	return result.Result()
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Set(ctx, key, value, expiration)
	return result.Result()
}

func (c *Client) SetArgs(key string, value interface{}, a redisV8.SetArgs) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.SetArgs(ctx, key, value, a)
	return result.Result()
}

// SetEX Redis `SETEX key expiration value` command.
func (c *Client) SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.SetEX(ctx, key, value, expiration)
	return result.Result()
}

// SetNX Redis `SET key value [expiration] NX` command.
//
// Zero expiration means the key has no expiration time.
// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
// otherwise you will receive an error: (error) ERR syntax error.
func (c *Client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.SetNX(ctx, key, value, expiration)
	return result.Result()
}

// SetXX Redis `SET key value [expiration] XX` command.
//
// Zero expiration means the key has no expiration time.
// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
// otherwise you will receive an error: (error) ERR syntax error.
func (c *Client) SetXX(key string, value interface{}, expiration time.Duration) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.SetXX(ctx, key, value, expiration)
	return result.Result()
}

func (c *Client) SetRange(key string, offset int64, value string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SetRange(ctx, key, offset, value)
	return result.Result()
}

func (c *Client) StrLen(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.StrLen(ctx, key)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) GetBit(key string, offset int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.GetBit(ctx, key, offset)
	return result.Result()
}

func (c *Client) SetBit(key string, offset int64, value int) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SetBit(ctx, key, offset, value)
	return result.Result()
}

func (c *Client) BitCount(key string, bitCount *redisV8.BitCount) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.BitCount(ctx, key, bitCount)
	return result.Result()
}

func (c *Client) BitOpAnd(destKey string, keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.BitOpAnd(ctx, destKey, keys...)
	return result.Result()
}

func (c *Client) BitOpOr(destKey string, keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.BitOpOr(ctx, destKey, keys...)
	return result.Result()
}

func (c *Client) BitOpXor(destKey string, keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.BitOpXor(ctx, destKey, keys...)
	return result.Result()
}

func (c *Client) BitOpNot(destKey string, key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.BitOpNot(ctx, destKey, key)
	return result.Result()
}

func (c *Client) BitPos(key string, bit int64, pos ...int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.BitPos(ctx, key, bit, pos...)
	return result.Result()
}

func (c *Client) BitField(key string, args ...interface{}) ([]int64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.BitField(ctx, key, args...)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, 0, err
	}
	result := c.c.Scan(ctx, cursor, match, count)
	return result.Result()
}

func (c *Client) ScanType(cursor uint64, match string, count int64, keyType string) ([]string, uint64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, 0, err
	}
	result := c.c.ScanType(ctx, cursor, match, count, keyType)
	return result.Result()
}

func (c *Client) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, 0, err
	}
	result := c.c.SScan(ctx, key, cursor, match, count)
	return result.Result()
}

func (c *Client) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, 0, err
	}
	result := c.c.HScan(ctx, key, cursor, match, count)
	return result.Result()
}

func (c *Client) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, 0, err
	}
	result := c.c.ZScan(ctx, key, cursor, match, count)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) HDel(key string, fields ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.HDel(ctx, key, fields...)
	return result.Result()
}

func (c *Client) HExists(key, field string) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.HExists(ctx, key, field)
	return result.Result()
}

func (c *Client) HGet(key, field string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.HGet(ctx, key, field)
	return result.Result()
}

func (c *Client) HGetAll(key string) (map[string]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.HGetAll(ctx, key)
	return result.Result()
}

func (c *Client) HIncrBy(key, field string, incr int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.HIncrBy(ctx, key, field, incr)
	return result.Result()
}

func (c *Client) HIncrByFloat(key, field string, incr float64) (float64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.HIncrByFloat(ctx, key, field, incr)
	return result.Result()
}

func (c *Client) HKeys(key string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.HKeys(ctx, key)
	return result.Result()
}

func (c *Client) HLen(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.HLen(ctx, key)
	return result.Result()
}

// HMGet returns the values for the specified fields in the hash stored at key.
// It returns an interface{} to distinguish between empty string and nil value.
func (c *Client) HMGet(key string, fields ...string) ([]interface{}, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.HMGet(ctx, key, fields...)
	return result.Result()
}

// HSet accepts values in following formats:
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
//
// Note that it requires Redis v4 for multiple field/value pairs support.
func (c *Client) HSet(key string, values ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.HSet(ctx, key, values...)
	return result.Result()
}

// HMSet is a deprecated version of HSet left for compatibility with Redis 3.
func (c *Client) HMSet(key string, values ...interface{}) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.HMSet(ctx, key, values...)
	return result.Result()
}

func (c *Client) HSetNX(key, field string, value interface{}) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.HSetNX(ctx, key, field, value)
	return result.Result()
}

func (c *Client) HVals(key string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.HVals(ctx, key)
	return result.Result()
}

// HRandField redis-server version >= 6.2.0.
func (c *Client) HRandField(key string, count int, withValues bool) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.HRandField(ctx, key, count, withValues)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.BLPop(ctx, timeout, keys...)
	return result.Result()
}

func (c *Client) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.BRPop(ctx, timeout, keys...)
	return result.Result()
}

func (c *Client) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.BRPopLPush(ctx, source, destination, timeout)
	return result.Result()
}

func (c *Client) LIndex(key string, index int64) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.LIndex(ctx, key, index)
	return result.Result()
}

func (c *Client) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LInsert(ctx, key, op, pivot, value)
	return result.Result()
}

func (c *Client) LInsertBefore(key string, pivot, value interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LInsertBefore(ctx, key, pivot, value)
	return result.Result()
}

func (c *Client) LInsertAfter(key string, pivot, value interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LInsertAfter(ctx, key, pivot, value)
	return result.Result()
}

func (c *Client) LLen(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LLen(ctx, key)
	return result.Result()
}

func (c *Client) LPop(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.LPop(ctx, key)
	return result.Result()
}

func (c *Client) LPopCount(key string, count int) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.LPopCount(ctx, key, count)
	return result.Result()
}

func (c *Client) LPos(key string, value string, a redisV8.LPosArgs) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LPos(ctx, key, value, a)
	return result.Result()
}

func (c *Client) LPosCount(key string, value string, count int64, a redisV8.LPosArgs) ([]int64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.LPosCount(ctx, key, value, count, a)
	return result.Result()
}

func (c *Client) LPush(key string, values ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LPush(ctx, key, values...)
	return result.Result()
}

func (c *Client) LPushX(key string, values ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LPushX(ctx, key, values...)
	return result.Result()
}

func (c *Client) LRange(key string, start, stop int64) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.LRange(ctx, key, start, stop)
	return result.Result()
}

func (c *Client) LRem(key string, count int64, value interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LRem(ctx, key, count, value)
	return result.Result()
}

func (c *Client) LSet(key string, index int64, value interface{}) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.LSet(ctx, key, index, value)
	return result.Result()
}

func (c *Client) LTrim(key string, start, stop int64) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.LTrim(ctx, key, start, stop)
	return result.Result()
}

func (c *Client) RPop(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.RPop(ctx, key)
	return result.Result()
}

func (c *Client) RPopCount(key string, count int) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.RPopCount(ctx, key, count)
	return result.Result()
}

func (c *Client) RPopLPush(source, destination string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.RPopLPush(ctx, source, destination)
	return result.Result()
}

func (c *Client) RPush(key string, values ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.RPush(ctx, key, values...)
	return result.Result()
}

func (c *Client) RPushX(key string, values ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.RPushX(ctx, key, values...)
	return result.Result()
}

func (c *Client) LMove(source, destination, srcpos, destpos string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.LMove(ctx, source, destination, srcpos, destpos)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) SAdd(key string, members ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SAdd(ctx, key, members...)
	return result.Result()
}

func (c *Client) SCard(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SCard(ctx, key)
	return result.Result()
}

func (c *Client) SDiff(keys ...string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SDiff(ctx, keys...)
	return result.Result()
}

func (c *Client) SDiffStore(destination string, keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SDiffStore(ctx, destination, keys...)
	return result.Result()
}

func (c *Client) SInter(keys ...string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SInter(ctx, keys...)
	return result.Result()
}

func (c *Client) SInterStore(destination string, keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SInterStore(ctx, destination, keys...)
	return result.Result()
}

func (c *Client) SIsMember(key string, member interface{}) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.SIsMember(ctx, key, member)
	return result.Result()
}

// SMIsMember Redis `SMISMEMBER key member [member ...]` command.
func (c *Client) SMIsMember(key string, members ...interface{}) ([]bool, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SMIsMember(ctx, key, members...)
	return result.Result()
}

// SMembers Redis `SMEMBERS key` command output as a slice.
func (c *Client) SMembers(key string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SMembers(ctx, key)
	return result.Result()
}

// SMembersMap Redis `SMEMBERS key` command output as a map.
func (c *Client) SMembersMap(key string) (map[string]struct{}, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SMembersMap(ctx, key)
	return result.Result()
}

func (c *Client) SMove(source, destination string, member interface{}) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.SMove(ctx, source, destination, member)
	return result.Result()
}

// SPop Redis `SPOP key` command.
func (c *Client) SPop(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.SPop(ctx, key)
	return result.Result()
}

// SPopN Redis `SPOP key count` command.
func (c *Client) SPopN(key string, count int64) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SPopN(ctx, key, count)
	return result.Result()
}

// SRandMember Redis `SRANDMEMBER key` command.
func (c *Client) SRandMember(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.SRandMember(ctx, key)
	return result.Result()
}

// SRandMemberN Redis `SRANDMEMBER key count` command.
func (c *Client) SRandMemberN(key string, count int64) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SRandMemberN(ctx, key, count)
	return result.Result()
}

func (c *Client) SRem(key string, members ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SRem(ctx, key, members...)
	return result.Result()
}

func (c *Client) SUnion(keys ...string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SUnion(ctx, keys...)
	return result.Result()
}

func (c *Client) SUnionStore(destination string, keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.SUnionStore(ctx, destination, keys...)
	return result.Result()
}

// XAdd a.Limit has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
func (c *Client) XAdd(a *redisV8.XAddArgs) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.XAdd(ctx, a)
	return result.Result()
}

func (c *Client) XDel(stream string, ids ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XDel(ctx, stream, ids...)
	return result.Result()
}

func (c *Client) XLen(stream string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XLen(ctx, stream)
	return result.Result()
}

func (c *Client) XRange(stream, start, stop string) ([]redisV8.XMessage, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XRange(ctx, stream, start, stop)
	return result.Result()
}

func (c *Client) XRangeN(stream, start, stop string, count int64) ([]redisV8.XMessage, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XRangeN(ctx, stream, start, stop, count)
	return result.Result()
}

func (c *Client) XRevRange(stream, start, stop string) ([]redisV8.XMessage, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XRevRange(ctx, stream, start, stop)
	return result.Result()
}

func (c *Client) XRevRangeN(stream, start, stop string, count int64) ([]redisV8.XMessage, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XRevRangeN(ctx, stream, start, stop, count)
	return result.Result()
}

func (c *Client) XRead(a *redisV8.XReadArgs) ([]redisV8.XStream, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XRead(ctx, a)
	return result.Result()
}

func (c *Client) XReadStreams(streams ...string) ([]redisV8.XStream, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XReadStreams(ctx, streams...)
	return result.Result()
}

func (c *Client) XGroupCreate(stream, group, start string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.XGroupCreate(ctx, stream, group, start)
	return result.Result()
}

func (c *Client) XGroupCreateMkStream(stream, group, start string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.XGroupCreateMkStream(ctx, stream, group, start)
	return result.Result()
}

func (c *Client) XGroupSetID(stream, group, start string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.XGroupSetID(ctx, stream, group, start)
	return result.Result()
}

func (c *Client) XGroupDestroy(stream, group string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XGroupDestroy(ctx, stream, group)
	return result.Result()
}

func (c *Client) XGroupCreateConsumer(stream, group, consumer string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XGroupCreateConsumer(ctx, group, stream, consumer)
	return result.Result()
}

func (c *Client) XGroupDelConsumer(stream, group, consumer string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XGroupDelConsumer(ctx, group, stream, consumer)
	return result.Result()
}

func (c *Client) XReadGroup(a *redisV8.XReadGroupArgs) ([]redisV8.XStream, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XReadGroup(ctx, a)
	return result.Result()
}

func (c *Client) XAck(stream, group string, ids ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XAck(ctx, stream, group, ids...)
	return result.Result()
}

func (c *Client) XPending(stream, group string) (*redisV8.XPending, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XPending(ctx, stream, group)
	return result.Result()
}

func (c *Client) XPendingExt(a *redisV8.XPendingExtArgs) ([]redisV8.XPendingExt, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XPendingExt(ctx, a)
	return result.Result()
}

func (c *Client) XAutoClaim(a *redisV8.XAutoClaimArgs) (messages []redisV8.XMessage, start string, err error) {
	ctx, err := c.context()
	if err != nil {
		return nil, "", err
	}
	result := c.c.XAutoClaim(ctx, a)
	return result.Result()
}

func (c *Client) XAutoClaimJustID(a *redisV8.XAutoClaimArgs) (ids []string, start string, err error) {
	ctx, err := c.context()
	if err != nil {
		return nil, "", err
	}
	result := c.c.XAutoClaimJustID(ctx, a)
	return result.Result()
}

func (c *Client) XClaim(a *redisV8.XClaimArgs) ([]redisV8.XMessage, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XClaim(ctx, a)
	return result.Result()
}

func (c *Client) XClaimJustID(a *redisV8.XClaimArgs) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XClaimJustID(ctx, a)
	return result.Result()
}

// XTrimMaxLen No `~` rules are used, `limit` cannot be used.
// cmd: XTRIM key MAXLEN maxLen
func (c *Client) XTrimMaxLen(key string, maxLen int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XTrimMaxLen(ctx, key, maxLen)
	return result.Result()
}

// XTrimMaxLenApprox LIMIT has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
// cmd: XTRIM key MAXLEN ~ maxLen LIMIT limit
func (c *Client) XTrimMaxLenApprox(key string, maxLen, limit int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XTrimMaxLenApprox(ctx, key, maxLen, limit)
	return result.Result()
}

// XTrimMinID No `~` rules are used, `limit` cannot be used.
// cmd: XTRIM key MINID minID
func (c *Client) XTrimMinID(key string, minID string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XTrimMinID(ctx, key, minID)
	return result.Result()
}

// XTrimMinIDApprox LIMIT has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
// cmd: XTRIM key MINID ~ minID LIMIT limit
func (c *Client) XTrimMinIDApprox(key string, minID string, limit int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.XTrimMinIDApprox(ctx, key, minID, limit)
	return result.Result()
}

func (c *Client) XInfoConsumers(key string, group string) ([]redisV8.XInfoConsumer, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XInfoConsumers(ctx, key, group)
	return result.Result()
}

func (c *Client) XInfoGroups(key string) ([]redisV8.XInfoGroup, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XInfoGroups(ctx, key)
	return result.Result()
}

func (c *Client) XInfoStream(key string) (*redisV8.XInfoStream, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XInfoStream(ctx, key)
	return result.Result()
}

// XInfoStreamFull XINFO STREAM FULL [COUNT count]
// redis-server >= 6.0.
func (c *Client) XInfoStreamFull(key string, count int) (*redisV8.XInfoStreamFull, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.XInfoStreamFull(ctx, key, count)
	return result.Result()
}

//------------------------------------------------------------------------------

// BZPopMax Redis `BZPOPMAX key [key ...] timeout` command.
func (c *Client) BZPopMax(timeout time.Duration, keys ...string) (*redisV8.ZWithKey, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.BZPopMax(ctx, timeout, keys...)
	return result.Result()
}

// BZPopMin Redis `BZPOPMIN key [key ...] timeout` command.
func (c *Client) BZPopMin(timeout time.Duration, keys ...string) (*redisV8.ZWithKey, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.BZPopMin(ctx, timeout, keys...)
	return result.Result()
}

func (c *Client) ZAddArgs(key string, args redisV8.ZAddArgs) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZAddArgs(ctx, key, args)
	return result.Result()
}

func (c *Client) ZAddArgsIncr(key string, args redisV8.ZAddArgs) (float64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZAddArgsIncr(ctx, key, args)
	return result.Result()
}

// ZAdd Redis `ZADD key score member [score member ...]` command.
func (c *Client) ZAdd(key string, members ...*redisV8.Z) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZAdd(ctx, key, members...)
	return result.Result()
}

// ZAddNX Redis `ZADD key NX score member [score member ...]` command.
func (c *Client) ZAddNX(key string, members ...*redisV8.Z) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZAddNX(ctx, key, members...)
	return result.Result()
}

// ZAddXX Redis `ZADD key XX score member [score member ...]` command.
func (c *Client) ZAddXX(key string, members ...*redisV8.Z) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZAddXX(ctx, key, members...)
	return result.Result()
}

func (c *Client) ZCard(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZCard(ctx, key)
	return result.Result()
}

func (c *Client) ZCount(key, min, max string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZCount(ctx, key, min, max)
	return result.Result()
}

func (c *Client) ZLexCount(key, min, max string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZLexCount(ctx, key, min, max)
	return result.Result()
}

func (c *Client) ZIncrBy(key string, increment float64, member string) (float64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZIncrBy(ctx, key, increment, member)
	return result.Result()
}

func (c *Client) ZInterStore(destination string, store *redisV8.ZStore) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZInterStore(ctx, destination, store)
	return result.Result()
}

func (c *Client) ZInter(store *redisV8.ZStore) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZInter(ctx, store)
	return result.Result()
}

func (c *Client) ZInterWithScores(store *redisV8.ZStore) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZInterWithScores(ctx, store)
	return result.Result()
}

func (c *Client) ZMScore(key string, members ...string) ([]float64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZMScore(ctx, key, members...)
	return result.Result()
}

func (c *Client) ZPopMax(key string, count ...int64) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZPopMax(ctx, key, count...)
	return result.Result()
}

func (c *Client) ZPopMin(key string, count ...int64) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZPopMin(ctx, key, count...)
	return result.Result()
}

func (c *Client) ZRangeArgs(z redisV8.ZRangeArgs) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRangeArgs(ctx, z)
	return result.Result()
}

func (c *Client) ZRangeArgsWithScores(z redisV8.ZRangeArgs) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRangeArgsWithScores(ctx, z)
	return result.Result()
}

func (c *Client) ZRange(key string, start, stop int64) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRange(ctx, key, start, stop)
	return result.Result()
}

func (c *Client) ZRangeWithScores(key string, start, stop int64) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRangeWithScores(ctx, key, start, stop)
	return result.Result()
}

func (c *Client) ZRangeByScore(key string, opt *redisV8.ZRangeBy) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRangeByScore(ctx, key, opt)
	return result.Result()
}

func (c *Client) ZRangeByLex(key string, opt *redisV8.ZRangeBy) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRangeByLex(ctx, key, opt)
	return result.Result()
}

func (c *Client) ZRangeByScoreWithScores(key string, opt *redisV8.ZRangeBy) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRangeByScoreWithScores(ctx, key, opt)
	return result.Result()
}

func (c *Client) ZRangeStore(dst string, z redisV8.ZRangeArgs) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZRangeStore(ctx, dst, z)
	return result.Result()
}

func (c *Client) ZRank(key, member string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZRank(ctx, key, member)
	return result.Result()
}

func (c *Client) ZRem(key string, members ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZRem(ctx, key, members...)
	return result.Result()
}

func (c *Client) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZRemRangeByRank(ctx, key, start, stop)
	return result.Result()
}

func (c *Client) ZRemRangeByScore(key, min, max string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZRemRangeByScore(ctx, key, min, max)
	return result.Result()
}

func (c *Client) ZRemRangeByLex(key, min, max string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZRemRangeByLex(ctx, key, min, max)
	return result.Result()
}

func (c *Client) ZRevRange(key string, start, stop int64) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRevRange(ctx, key, start, stop)
	return result.Result()
}

func (c *Client) ZRevRangeWithScores(key string, start, stop int64) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRevRangeWithScores(ctx, key, start, stop)
	return result.Result()
}

func (c *Client) ZRevRangeByScore(key string, opt *redisV8.ZRangeBy) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRevRangeByScore(ctx, key, opt)
	return result.Result()
}

func (c *Client) ZRevRangeByLex(key string, opt *redisV8.ZRangeBy) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRevRangeByLex(ctx, key, opt)
	return result.Result()
}

func (c *Client) ZRevRangeByScoreWithScores(key string, opt *redisV8.ZRangeBy) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRevRangeByScoreWithScores(ctx, key, opt)
	return result.Result()
}

func (c *Client) ZRevRank(key, member string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZRevRank(ctx, key, member)
	return result.Result()
}

func (c *Client) ZScore(key, member string) (float64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZScore(ctx, key, member)
	return result.Result()
}

func (c *Client) ZUnion(store redisV8.ZStore) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZUnion(ctx, store)
	return result.Result()
}

func (c *Client) ZUnionWithScores(store redisV8.ZStore) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZUnionWithScores(ctx, store)
	return result.Result()
}

func (c *Client) ZUnionStore(dest string, store *redisV8.ZStore) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZUnionStore(ctx, dest, store)
	return result.Result()
}

// ZRandMember redis-server version >= 6.2.0.
func (c *Client) ZRandMember(key string, count int, withScores bool) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZRandMember(ctx, key, count, withScores)
	return result.Result()
}

// ZDiff redis-server version >= 6.2.0.
func (c *Client) ZDiff(keys ...string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZDiff(ctx, keys...)
	return result.Result()
}

// ZDiffWithScores redis-server version >= 6.2.0.
func (c *Client) ZDiffWithScores(keys ...string) ([]redisV8.Z, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ZDiffWithScores(ctx, keys...)
	return result.Result()
}

// ZDiffStore redis-server version >=6.2.0.
func (c *Client) ZDiffStore(destination string, keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ZDiffStore(ctx, destination, keys...)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) PFAdd(key string, els ...interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.PFAdd(ctx, key, els)
	return result.Result()
}

func (c *Client) PFCount(keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.PFCount(ctx, keys...)
	return result.Result()
}

func (c *Client) PFMerge(dest string, keys ...string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.PFMerge(ctx, dest, keys...)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) BgRewriteAOF() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.BgRewriteAOF(ctx)
	return result.Result()
}

func (c *Client) BgSave() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.BgSave(ctx)
	return result.Result()
}

func (c *Client) ClientKill(ipPort string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClientKill(ctx, ipPort)
	return result.Result()
}

// ClientKillByFilter is new style syntax, while the ClientKill is old
//
//   CLIENT KILL <option> [value] ... <option> [value]
func (c *Client) ClientKillByFilter(keys ...string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ClientKillByFilter(ctx, keys...)
	return result.Result()
}

func (c *Client) ClientList() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClientList(ctx)
	return result.Result()
}

func (c *Client) ClientPause(dur time.Duration) (bool, error) {
	ctx, err := c.context()
	if err != nil {
		return false, err
	}
	result := c.c.ClientPause(ctx, dur)
	return result.Result()
}

func (c *Client) ClientID() (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ClientID(ctx)
	return result.Result()
}

func (c *Client) ClientUnblock(id int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ClientUnblock(ctx, id)
	return result.Result()
}

func (c *Client) ClientUnblockWithError(id int64) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ClientUnblockWithError(ctx, id)
	return result.Result()
}

func (c *Client) ConfigGet(parameter string) ([]interface{}, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ConfigGet(ctx, parameter)
	return result.Result()
}

func (c *Client) ConfigResetStat() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ConfigResetStat(ctx)
	return result.Result()
}

func (c *Client) ConfigSet(parameter, value string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ConfigSet(ctx, parameter, value)
	return result.Result()
}

func (c *Client) ConfigRewrite() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ConfigRewrite(ctx)
	return result.Result()
}

func (c *Client) DBSize() (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.DBSize(ctx)
	return result.Result()
}

func (c *Client) FlushAll() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.FlushAll(ctx)
	return result.Result()
}

func (c *Client) FlushAllAsync() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.FlushAllAsync(ctx)
	return result.Result()
}

func (c *Client) FlushDB() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.FlushDB(ctx)
	return result.Result()
}

func (c *Client) FlushDBAsync() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.FlushDBAsync(ctx)
	return result.Result()
}

func (c *Client) Info(section ...string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Info(ctx, section...)
	return result.Result()
}

func (c *Client) LastSave() (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.LastSave(ctx)
	return result.Result()
}

func (c *Client) Save() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Save(ctx)
	return result.Result()
}

func (c *Client) Shutdown() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.Shutdown(ctx)
	return result.Result()
}

func (c *Client) ShutdownSave() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ShutdownSave(ctx)
	return result.Result()
}

func (c *Client) ShutdownNoSave() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ShutdownNoSave(ctx)
	return result.Result()
}

func (c *Client) SlaveOf(host, port string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.SlaveOf(ctx, host, port)
	return result.Result()
}

func (c *Client) SlowLogGet(num int64) ([]redisV8.SlowLog, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.SlowLogGet(ctx, num)
	return result.Result()
}

func (c *Client) Time() (time.Time, error) {
	ctx, err := c.context()
	if err != nil {
		return time.Time{}, err
	}
	result := c.c.Time(ctx)
	return result.Result()
}

func (c *Client) DebugObject(key string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.DebugObject(ctx, key)
	return result.Result()
}

func (c *Client) ReadOnly() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ReadOnly(ctx)
	return result.Result()
}

func (c *Client) ReadWrite() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ReadWrite(ctx)
	return result.Result()
}

func (c *Client) MemoryUsage(key string, samples ...int) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.MemoryUsage(ctx, key, samples...)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.Eval(ctx, script, keys, args...)
	return result.Result()
}

func (c *Client) EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.EvalSha(ctx, sha1, keys, args...)
	return result.Result()
}

func (c *Client) ScriptExists(hashes ...string) ([]bool, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ScriptExists(ctx, hashes...)
	return result.Result()
}

func (c *Client) ScriptFlush() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ScriptFlush(ctx)
	return result.Result()
}

func (c *Client) ScriptKill() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ScriptKill(ctx)
	return result.Result()
}

func (c *Client) ScriptLoad(script string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ScriptLoad(ctx, script)
	return result.Result()
}

//------------------------------------------------------------------------------

// Publish posts the message to the channel.
func (c *Client) Publish(channel string, message interface{}) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.Publish(ctx, channel, message)
	return result.Result()
}

func (c *Client) PubSubChannels(pattern string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.PubSubChannels(ctx, pattern)
	return result.Result()
}

func (c *Client) PubSubNumSub(channels ...string) (map[string]int64, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.PubSubNumSub(ctx, channels...)
	return result.Result()
}

func (c *Client) PubSubNumPat() (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.PubSubNumPat(ctx)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) ClusterSlots() ([]redisV8.ClusterSlot, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ClusterSlots(ctx)
	return result.Result()
}

func (c *Client) ClusterNodes() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterNodes(ctx)
	return result.Result()
}

func (c *Client) ClusterMeet(host, port string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterMeet(ctx, host, port)
	return result.Result()
}

func (c *Client) ClusterForget(nodeID string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterForget(ctx, nodeID)
	return result.Result()
}

func (c *Client) ClusterReplicate(nodeID string) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterReplicate(ctx, nodeID)
	return result.Result()
}

func (c *Client) ClusterResetSoft() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterResetSoft(ctx)
	return result.Result()
}

func (c *Client) ClusterResetHard() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterResetHard(ctx)
	return result.Result()
}

func (c *Client) ClusterInfo() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterInfo(ctx)
	return result.Result()
}

func (c *Client) ClusterKeySlot(key string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ClusterKeySlot(ctx, key)
	return result.Result()
}

func (c *Client) ClusterGetKeysInSlot(slot int, count int) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ClusterGetKeysInSlot(ctx, slot, count)
	return result.Result()
}

func (c *Client) ClusterCountFailureReports(nodeID string) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ClusterCountFailureReports(ctx, nodeID)
	return result.Result()
}

func (c *Client) ClusterCountKeysInSlot(slot int) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.ClusterCountKeysInSlot(ctx, slot)
	return result.Result()
}

func (c *Client) ClusterDelSlots(slots ...int) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterDelSlots(ctx, slots...)
	return result.Result()
}

func (c *Client) ClusterDelSlotsRange(min, max int) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterDelSlotsRange(ctx, min, max)
	return result.Result()
}

func (c *Client) ClusterSaveConfig() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterSaveConfig(ctx)
	return result.Result()
}

func (c *Client) ClusterSlaves(nodeID string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.ClusterSlaves(ctx, nodeID)
	return result.Result()
}

func (c *Client) ClusterFailover() (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterFailover(ctx)
	return result.Result()
}

func (c *Client) ClusterAddSlots(slots ...int) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterAddSlots(ctx, slots...)
	return result.Result()
}

func (c *Client) ClusterAddSlotsRange(min, max int) (string, error) {
	ctx, err := c.context()
	if err != nil {
		return "", err
	}
	result := c.c.ClusterAddSlotsRange(ctx, min, max)
	return result.Result()
}

//------------------------------------------------------------------------------

func (c *Client) GeoAdd(key string, geoLocation ...*redisV8.GeoLocation) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.GeoAdd(ctx, key, geoLocation...)
	return result.Result()
}

// GeoRadius is a read-only GEORADIUS_RO command.
func (c *Client) GeoRadius(key string, longitude, latitude float64, query *redisV8.GeoRadiusQuery) ([]redisV8.GeoLocation, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.GeoRadius(ctx, key, longitude, latitude, query)
	return result.Result()
}

// GeoRadiusStore is a writing GEORADIUS command.
func (c *Client) GeoRadiusStore(key string, longitude, latitude float64, query *redisV8.GeoRadiusQuery) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.GeoRadiusStore(ctx, key, longitude, latitude, query)
	return result.Result()
}

// GeoRadiusByMember is a read-only GEORADIUSBYMEMBER_RO command.
func (c *Client) GeoRadiusByMember(key, member string, query *redisV8.GeoRadiusQuery) ([]redisV8.GeoLocation, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.GeoRadiusByMember(ctx, key, member, query)
	return result.Result()
}

// GeoRadiusByMemberStore is a writing GEORADIUSBYMEMBER command.
func (c *Client) GeoRadiusByMemberStore(key, member string, query *redisV8.GeoRadiusQuery) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.GeoRadiusByMemberStore(ctx, key, member, query)
	return result.Result()
}

func (c *Client) GeoSearch(key string, q *redisV8.GeoSearchQuery) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.GeoSearch(ctx, key, q)
	return result.Result()
}

func (c *Client) GeoSearchLocation(key string, q *redisV8.GeoSearchLocationQuery) ([]redisV8.GeoLocation, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.GeoSearchLocation(ctx, key, q)
	return result.Result()
}

func (c *Client) GeoSearchStore(key, store string, q *redisV8.GeoSearchStoreQuery) (int64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.GeoSearchStore(ctx, key, store, q)
	return result.Result()
}

func (c *Client) GeoDist(key string, member1, member2, unit string) (float64, error) {
	ctx, err := c.context()
	if err != nil {
		return 0, err
	}
	result := c.c.GeoDist(ctx, key, member1, member2, unit)
	return result.Result()
}

func (c *Client) GeoHash(key string, members ...string) ([]string, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.GeoHash(ctx, key, members...)
	return result.Result()
}

func (c *Client) GeoPos(key string, members ...string) ([]*redisV8.GeoPos, error) {
	ctx, err := c.context()
	if err != nil {
		return nil, err
	}
	result := c.c.GeoPos(ctx, key, members...)
	return result.Result()
}
