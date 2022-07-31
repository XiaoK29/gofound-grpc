package searcher

import (
	"fmt"
	"gofound-grpc/internal/searcher/words"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"unsafe"
)

type Container struct {
	Dir       string             // 文件夹
	engines   map[string]*Engine // 引擎
	Tokenizer *words.Tokenizer   // 分词器
	Shard     int32              // 分片
}

func (c *Container) Init(e *Engine) error {

	c.engines = make(map[string]*Engine)

	//读取当前路径下的所有目录，就是数据库名称
	dirs, err := ioutil.ReadDir(c.Dir)
	if err != nil {
		if os.IsNotExist(err) {
			//创建
			err := os.MkdirAll(c.Dir, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	//初始化数据库
	for _, dir := range dirs {
		if dir.IsDir() {
			c.engines[dir.Name()] = e.GetDataBase(dir.Name(), c)
			log.Println("db:", dir.Name())
		}
	}

	return nil
}

// GetDataBases 获取数据库列表
func (c *Container) GetDataBases() map[string]*Engine {
	for _, engine := range c.engines {
		size := unsafe.Sizeof(&engine)
		fmt.Printf("%s:%d\n", engine.DatabaseName, size)
	}
	return c.engines
}

func (c *Container) GetDataBaseNumber() int {
	return len(c.engines)
}

func (c *Container) GetIndexCount() int64 {
	var count int64
	for _, engine := range c.engines {
		count += engine.GetIndexCount()
	}
	return count
}

func (c *Container) GetDocumentCount() int64 {
	var count int64
	for _, engine := range c.engines {
		count += engine.GetDocumentCount()
	}
	return count
}

// DropDataBase 删除数据库
func (c *Container) DropDataBase(name string) error {
	err := c.engines[name].Drop()
	if err != nil {
		return err
	}

	delete(c.engines, name)
	//释放资源
	runtime.GC()

	return nil
}
