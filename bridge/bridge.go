package bridge

import (
	"fmt"
	"strconv"
	"strings"
)

/**
桥接模式 适用于
1. 一个类存在两个维度的变化, 并且这两个维度都需要扩展
2. 当一个系统不希望使用继承, 或者因为多层次的继承关系会导致类的个数急剧增加
3. 当一个系统需要在组件的抽象化和具体化角色之间增加更多的灵活性时
*/
type Driver interface {
	GetConnection(string)
}

// 可以通过适配器模式代替
type DriverImpl struct {
	MySqlDriverImpl
}

func (d *DriverImpl) GetConnection(url string) {
	splits := strings.Split(url, ":")
	port, _ := strconv.Atoi(splits[1])
	d.MySqlDriverImpl.GetConnection(splits[0], int64(port))
}

type MySqlDriver interface {
	GetConnection(string, int64)
}

type MySqlDriverImpl struct {
}

func (*MySqlDriverImpl) GetConnection(host string, port int64) {
	fmt.Println("connect ", host, port)
}
