package flags

// 用于解析命令行参数
import (
	"flag"
	"os"
)

// Options 定义命令行可选参数结构体
type Options struct {
	File    string
	DB      bool
	Version bool
}

var FlagOptions = new(Options)

func Parse() {
	// 第一个参数解析命令行参数后接收的参数, 第二个参数命令行参数的名称, 第三个参数当没有传递该命令行参数时默认值, 第四个参数命令行参数说明
	// 命令行参数不传递会赋默认值
	flag.StringVar(&FlagOptions.File, "f", "settings.yml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.BoolVar(&FlagOptions.Version, "v", false, "版本")
	flag.Parse()
}

// Run 数据库迁移
func Run() {
	if FlagOptions.DB {
		MigrateDB()
		os.Exit(0)
	}
}
