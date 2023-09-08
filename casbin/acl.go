package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	e, err := casbin.NewEnforcer("casbin/model.conf", "casbin/policy.csv")

	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read"  // 用户对资源执行的操作。

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理err
		log.Panic(err)
	}

	if ok == true {
		// 允许alice读取data1
		log.Println("测试通过")
	} else {
		// 拒绝请求，抛出异常
		log.Println("测试不通过")
	}

	// 您可以使用BatchEnforce()来批量执行一些请求
	// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
	// 例如results[0] 是{"alice", "data1", "read"}的结果
	//results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
}
