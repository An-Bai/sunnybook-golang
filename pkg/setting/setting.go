package setting

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

var (
	Cfg          *ini.File     // ini 文件对象
	RunMode      string        // 运行环境
	HTTPPort     int           // http服务端口
	ReadTimeout  time.Duration // 读取文件超时时间
	WriteTimeout time.Duration // 写文件超时时间
	PageSize     int           // 分页默认设置的大小
	JwtSecret    string        // JWT秘钥
	LogFile      string        // 日志输出位置
)

// 初始化配置
func init() {
	var err error
	// 获取配置文件对象
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
		os.Exit(1)
	}
	// 调用配置初始化方法
	LoadBase()
	LoadServer()
	LoadApp()
}

/**
 * 获取ini文件方法示例说明：
 * Section() 分区设置；
 * Key() 指定要获取值的key；
 * MustString()：设置默认值候补
 */

// LoadBase 加载基础配置
func LoadBase() {
	// 设置运行输出日志级别（在gin配置参数时使用）
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	// 日志输出配置，将日志输出到指定文件内
	LogFile = Cfg.Section("").Key("LOG_FILE").MustString("error.log")
	f, err := os.OpenFile(LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	log.SetOutput(f)
	log.Println("服务器启动中...")
}

// LoadApp 项目应用配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("@@sunny@@book$go")
	// 设置默认分页值（不存在默认值为10）
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

// LoadServer 服务相关配置
func LoadServer() {
	// Cfg.GetSection() 获取整个区
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	// 设置访问端口（不存在默认值为8000）
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
