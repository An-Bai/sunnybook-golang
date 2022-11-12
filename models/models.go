package models

// 数据库连接配置
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"sunnybook-golang/pkg/setting"
	"time"
)

var db *gorm.DB

// gorm.Model 源码定义
//type Model struct {
//	ID        uint `gorm:"primarykey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt DeletedAt `gorm:"index"`
//}

// Model 自定义 gorm.Model
type Model struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"-"`
}

func init() {
	// 定义要从配置中获取的值
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	// 使用变量拼接
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 正式环境去掉sql日志
		//Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,        // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	//表示最大的连接数，这个我们不设置默认就是不限制，可以无限创建连接，问题就在数据库本身有瓶颈，无限创建，会损耗性能。
	//所以我们要根据我们自己的数据库瓶颈情况来进行相关的设置。当出现连接数超出了我们设定的数量时候，后面的用户等待超时时间之前，
	//有连接释放就会自动获得操作的权限，否则返回连接超时。（每个公司的使用情况不同，所以根据情况自己设定，个人建议不要采用默认无限制创建连接）
	sqlDB.SetMaxOpenConns(100)

	// 设置最大的可空闲连接数，该函数的作用就是保持等待连接操作状态的连接数，这个主要就是避免操作过程中频繁的获取连接，释放连接。
	//默认情况下会保持的连接数量为2.就是说会有两个连接一直保持，不释放，等待需要使用的用户使用。
	sqlDB.SetMaxIdleConns(10)

}

// CloseDB 关闭数据库连接
func CloseDB() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
