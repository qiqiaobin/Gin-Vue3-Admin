package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 配置文件映射对象
var Config *Server

func InitConfig() {
	//设置日志打印格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//configFilePath := flag.String("e", "./etc/config.yaml", "配置文件路径，默认为可执行文件目录")
	//flag.Parse()
	// 获取启动参数中，配置文件的绝对路径
	//path, _ := filepath.Abs(*configFilePath)

	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		//panic(err)
		log.Fatal("获取项目路径失败", err)
		return
	}
	//cfg := viper.New()
	// 设置读取的文件路径

	// 读取配置文件信息
	cfg := &Server{}
	//viper.SetConfigFile(path)
	// 设置读取的文件路径
	viper.AddConfigPath(path + "/etc")
	// 设置读取的文件名
	viper.SetConfigName("config")
	// 设置文件的类型
	viper.SetConfigType("yaml")
	// 尝试进行配置读取
	err = viper.ReadInConfig()
	// 读取配置文件内容
	//data, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		log.Fatalf("读取配置文件%s失败: %v\n", path, err)
	}
	// 监听配置文件
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件有变更:", e.Name)
		if err = viper.Unmarshal(cfg); err != nil {
			log.Fatalln(err)
		}
	})
	// 解析配置赋值,将配置文件绑定到 config 上
	if err = viper.Unmarshal(cfg); err != nil {
		log.Fatalln(err)
	}
	// 校验配置文件内容信息
	//cfg.Valid()
	// 尝试使用系统环境变量替换配置信息
	cfg.ReplaceOsEnv()
	Config = cfg

	log.Printf("成功解析配置文件。")
}

/*
// 配置文件内容校验
func (c *Conf) Valid() {
	assert.IsTrue(c.Jwt != nil, "配置文件的[jwt]信息不能为空")
	c.Jwt.Valid()
	if c.Aes != nil {
		c.Aes.Valid()
	}
}
*/
// 替换系统环境变量，如果环境变量中存在该值，则优先使用环境变量设定的值
func (c *Server) ReplaceOsEnv() {
	dbHost := os.Getenv("MAYFLY_DB_HOST")
	if dbHost != "" {
		c.Database.Host = dbHost
	}

	dbName := os.Getenv("MAYFLY_DB_NAME")
	if dbName != "" {
		c.Database.DbName = dbName
	}

	dbUser := os.Getenv("MAYFLY_DB_USER")
	if dbUser != "" {
		c.Database.User = dbUser
	}

	dbPwd := os.Getenv("MAYFLY_DB_PASS")
	if dbPwd != "" {
		c.Database.Password = dbPwd
	}

	//aesKey := os.Getenv("MAYFLY_AES_KEY")
	//if aesKey != "" {
	//	c.Aes.Key = aesKey
	//}

	//jwtKey := os.Getenv("MAYFLY_JWT_KEY")
	//if jwtKey != "" {
	//	c.Jwt.Key = jwtKey
	//}
}
