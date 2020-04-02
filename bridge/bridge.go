package bridge

import (
	"fmt"
	"strconv"
	"strings"
)

type Driver interface {
	GetConnection(string)
}

// 相当于一个接口适配
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
