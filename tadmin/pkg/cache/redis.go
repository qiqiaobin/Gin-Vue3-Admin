package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"tadmin/conf"
	"tadmin/pkg/logger"
)

const (
	CacheTime                 = time.Minute * 30       //缓存时间（可应用大部分业务缓存）
	CacheKeySysUserMenu       = "sys_user_menu_"       //菜单
	CacheKeySysUserMenuIds    = "sys_user_menuids_"    //菜单ID
	CacheKeySysUserPermission = "sys_user_permission_" //用户权限标识
)

type RedisTool struct{}

var Redis = new(RedisTool)
var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Config.Redis.Host, conf.Config.Redis.Port),
		Password: conf.Config.Redis.Password, // no password set
		DB:       conf.Config.Redis.DB,       // use default DB
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Error("连接Redis出现错误：", zap.Error(err))
	} else {
		logger.Info("连接Redis成功:", zap.String("pong", pong))
		rdb = client
	}
}

// Set 设置缓存
func (*RedisTool) Set(key string, value interface{}, expiration time.Duration) {
	//序列化
	data, _ := json.Marshal(value)
	rdb.Set(ctx, key, string(data), expiration)
}

// Get 获取缓存
func (*RedisTool) Get(key string, value interface{}) {
	str, _ := rdb.Get(ctx, key).Result()
	//反序列化
	json.Unmarshal([]byte(str), &value)

}

// Del 删除缓存
func (*RedisTool) Del(key string) {
	rdb.Del(ctx, key)
}

// DelByPattern 模糊删除缓存
func (*RedisTool) DelByPattern(pattern string) {
	keys, _ := rdb.Keys(ctx, pattern+"*").Result()
	for i := 0; i < len(keys); i++ {
		rdb.Del(ctx, keys[i])
	}
}

// GetAllKeys 获取所有key
func (*RedisTool) GetAllKeys() (keys []string) {
	iter := rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		keys = append(keys, key)
	}
	return keys
}

// GetValue 获取缓存
func (*RedisTool) GetValue(key string) string {
	str, _ := rdb.Get(ctx, key).Result()
	return str
}
