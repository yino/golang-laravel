package redis
//
//import (
//	"context"
//	"gin-api/config"
//	"github.com/go-redis/redis/v8"
//	"time"
//)
//
//var (
//	ctx     = Ctx()
//	connect = Connection()
//	prefix  = config.RedisConf()["RedisPrefix"].(string) // 前缀
//)
//
//func Del(keys ...string) *redis.IntCmd {
//	return connect.Del(ctx, prefix+keys)
//}
//
//// ---------------------------SET--------------------------------------
//
//// Redis `SET key value [expiration]` command.
//// Use expiration for `SETEX`-like behavior.
////
//// Zero expiration means the key has no expiration time.
//// KeepTTL(-1) expiration is a Redis KEEPTTL option to keep existing TTL.
//func Set(key string, value interface{}, expiration time.Duration) error {
//	return connect.Set(ctx, prefix+key, value, expiration).Err()
//}
//
//// Redis `SETEX key expiration value` command.
//func SetEX(key string, value interface{}, expiration time.Duration) error {
//	return connect.SetEX(ctx, prefix+key, value, expiration).Err()
//}
//
//// Redis `SetNX key expiration value` command.
//func SetNX(key string, value interface{}, expiration time.Duration) error {
//	return connect.SetNX(ctx, prefix+key, value, expiration).Err()
//}
//
//// Redis `SET key value [expiration] XX` command.
////
//// Zero expiration means the key has no expiration time.
//// KeepTTL(-1) expiration is a Redis KEEPTTL option to keep existing TTL.
//func SetXX(key string, value interface{}, expiration time.Duration) error {
//	return connect.SetXX(ctx, prefix+key, value, expiration).Err()
//}
//
//func SetRange(key string, offset int64, value string) *redis.IntCmd {
//	return connect.SetRange(ctx, prefix+key, offset, value)
//}
//
//func Incr(key string) *redis.IntCmd {
//	return connect.Incr(ctx, prefix+key)
//}
//func IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
//	return connect.IncrBy(ctx, prefix+key, value)
//}
//
//func IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
//	return connect.IncrByFloat(ctx, prefix+key, value)
//}
//
//// -----------------------------SET END----------------------------------
//// -------------------------------GET------------------------------------
//func Get(key string) *redis.StringCmd {
//	return connect.Get(ctx, prefix+key)
//}
//
//func GetRange(key string, start, end int64) *redis.StringCmd {
//	return connect.GetRange(ctx, prefix+key, start, end)
//}
//
//func GetSet(key string, value interface{}) *redis.StringCmd {
//	return connect.GetSet(ctx, prefix+key, value)
//}
//
//func StrLen(key string) *redis.IntCmd {
//	return connect.StrLen(ctx, prefix+key)
//}
//
//// -----------------------------GET END----------------------------------
//// -----------------------------HASH SET---------------------------------
//
//func HIncrBy(key, field string, incr int64) *redis.IntCmd {
//	return connect.HIncrBy(ctx, prefix+key, field, incr)
//}
//
//func HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
//	return connect.HIncrByFloat(ctx, prefix+key, field, incr)
//}
//
//func HMSet( key string, values ...interface{}) *redis.BoolCmd{
//	return connect.HMSet(ctx, prefix+key, values)
//}
//
//// -----------------------------HASH SET END-----------------------------
//// -----------------------------HASH GET---------------------------------
//func HKeys(key string) *redis.StringSliceCmd {
//	return connect.HKeys(ctx, prefix+key)
//}
//
//func HLen(key string) *redis.IntCmd {
//	return connect.HLen(ctx, key)
//}
//
//func HDel(key string, fields ...string) *redis.IntCmd {
//	return connect.HDel(ctx, key, fields)
//}
//
//func HExists(key, field string) *redis.BoolCmd {
//	return connect.HExists(ctx, prefix+key, field)
//}
//
//func HGet(key, field string) *redis.StringCmd {
//	return connect.HGet(ctx, prefix+key, field)
//}
//
//func HGetAll(key string) *redis.StringStringMapCmd {
//	return connect.HGetAll(ctx, prefix+key)
//}
//
//func HMGet(key string, fields ...string) *redis.SliceCmd{
//	return connect.HMGet(ctx, prefix+key, fields)
//}
//// ---------------------------HASH GET END-----------------------------
