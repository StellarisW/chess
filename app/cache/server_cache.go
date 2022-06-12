/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-03
* Time: 15:23
 */

package cache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"main/app/internal/model"
	"strconv"
)

const (
	serversHashKey       = "acc:hash:servers" // 全部的服务器
	serversHashCacheTime = 2 * 60 * 60        // key过期时间
	serversHashTimeout   = 3 * 60             // 超时时间
)

func getServersHashKey() (key string) {
	key = fmt.Sprintf("%s", serversHashKey)

	return
}

// 设置服务器信息
func SetServerInfo(server *model.Server, currentTime uint64) (err error) {
	key := getServersHashKey()

	value := fmt.Sprintf("%d", currentTime)

	number, err := g.Redis().Do(context.Background(), "hSet", key, server.String(), value)
	if err != nil {
		fmt.Println("SetServerInfo", key, number, err)

		return
	}

	if number.Int() != 1 {

		return
	}

	g.Redis().Do(context.Background(), "Expire", key, serversHashCacheTime)

	return
}

// 下线服务器信息
func DelServerInfo(server *model.Server) (err error) {
	key := getServersHashKey()
	number, err := g.Redis().Do(context.Background(), "hDel", key, server.String())
	if err != nil {
		fmt.Println("DelServerInfo", key, number, err)

		return
	}

	if number.Int() != 1 {

		return
	}

	g.Redis().Do(context.Background(), "Expire", key, serversHashCacheTime)

	return
}

func GetServerAll(currentTime uint64) (servers []*model.Server, err error) {

	servers = make([]*model.Server, 0)
	key := getServersHashKey()

	val, err := g.Redis().Do(context.Background(), "hGetAll", key)

	fmt.Println("GetServerAll", key, string(val.Bytes()))

	serverMap, err := g.Redis().Do(context.Background(), "hGetAll", key)
	if err != nil {
		fmt.Println("SetServerInfo", key, err)

		return
	}

	for key, value := range serverMap.MapStrStr() {
		valueUint64, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			fmt.Println("GetServerAll", key, err)

			return nil, err
		}

		// 超时
		if valueUint64+serversHashTimeout <= currentTime {
			continue
		}

		server, err := model.StringToServer(key)
		if err != nil {
			fmt.Println("GetServerAll", key, err)

			return nil, err
		}

		servers = append(servers, server)
	}

	return
}
