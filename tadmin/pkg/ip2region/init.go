package ip2region

import (
	"fmt"
	"os"

	//"tadmin/pkg/logger"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var IpSearcher *xdb.Searcher

func InitIp2region() {
	path, err := os.Getwd()
	dbPath := path + "/etc/ip2region.xdb"
	if err != nil {
		//logger.Error("获取项目路径失败", err)
		return
	}
	// 1、从 dbPath 加载 VectorIndex 缓存，把下述 vIndex 变量全局到内存里面。
	vIndex, err := xdb.LoadVectorIndexFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load vector index from `%s`: %s\n", dbPath, err)
		return
	}

	// 2、用全局的 vIndex 创建带 VectorIndex 缓存的查询对象。
	searcher, err := xdb.NewWithVectorIndex(dbPath, vIndex)
	if err != nil {
		fmt.Printf("failed to create searcher with vector index: %s\n", err)
		return
	}
	IpSearcher = searcher
	// 备注：并发使用，全部 goroutine 共享全局的只读 vIndex 缓存，每个 goroutine 创建一个独立的 searcher 对象
}
