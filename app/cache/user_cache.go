package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"main/app/internal/model"
)

const (
	userOnlinePrefix    = "acc:user:online:" // 用户在线状态
	userOnlineCacheTime = 24 * 60 * 60
)

/*********************  查询用户是否在线  ************************/
func getUserOnlineKey(userKey string) (key string) {
	key = fmt.Sprintf("%s%s", userOnlinePrefix, userKey)

	return
}

func GetUserOnlineInfo(userKey string) (userOnline *model.UserOnline, err error) {
	key := getUserOnlineKey(userKey)
	data, err := g.Redis().Do(context.Background(), "GET", key)
	if err != nil {
		fmt.Println("GetUserOnlineInfo", userKey, err)
		return
	}

	userOnline = &model.UserOnline{}
	err = json.Unmarshal(data.Bytes(), userOnline)
	if err != nil {
		fmt.Println("获取用户在线数据 json Unmarshal", userKey, err)

		return
	}

	fmt.Println("获取用户在线数据", userKey, "time", userOnline.LoginTime, userOnline.HeartbeatTime, "AccIp", userOnline.AccIp, userOnline.IsLogoff)

	return
}

// 设置用户在线数据
func SetUserOnlineInfo(userKey string, userOnline *model.UserOnline) (err error) {

	key := getUserOnlineKey(userKey)

	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		fmt.Println("设置用户在线数据 json Marshal", key, err)

		return
	}

	_, err = g.Redis().Do(context.Background(), "setEx", key, userOnlineCacheTime, string(valueByte))
	if err != nil {
		fmt.Println("设置用户在线数据 ", key, err)
		return
	}

	return
}
