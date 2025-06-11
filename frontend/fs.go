// ElasticView插件前端资源处理包 - 负责嵌入和提供前端静态资源服务
package frontend

import (
	"embed"    // 导入embed包，用于嵌入静态资源到二进制文件中
	"io/fs"    // 导入文件系统接口包
	"net/http" // 导入HTTP包，用于文件服务

	"github.com/gin-gonic/gin" // 导入Gin HTTP框架
)

// 嵌入前端构建产物 - 将Vue.js应用编译后的静态资源嵌入到Go二进制文件中
// 这样可以实现单文件分发，无需单独部署前端资源
//
//go:embed dist/index.html
//go:embed dist/assets/**
var StatisFs embed.FS // ElasticView插件的静态文件系统，包含所有前端资源

// embedFileSystem 结构体 - 扩展标准HTTP文件系统，添加文件存在性检查功能
// 用于ElasticView插件的静态资源服务，支持检查文件是否存在
type embedFileSystem struct {
	http.FileSystem // 嵌入标准HTTP文件系统接口
}

// Exists 方法 - 检查指定路径的文件是否存在
// ElasticView插件路由系统需要此方法来判断是否应该提供静态资源服务
// 参数 prefix: 路径前缀（通常不使用）
// 参数 path: 要检查的文件路径
// 返回值: 文件是否存在的布尔值
func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path) // 尝试打开文件
	if err != nil {        // 如果打开失败
		return false // 文件不存在
	}
	return true // 文件存在
}

// EmbedFolder 函数 - 创建嵌入式文件系统服务
// 将embed.FS转换为可用于HTTP服务的文件系统，支持ElasticView插件的静态资源访问
// 参数 fsEmbed: 嵌入的文件系统
// 参数 targetPath: 目标路径，通常是"dist"目录
// 返回值: 实现了ServeFileSystem接口的文件系统
func EmbedFolder(fsEmbed embed.FS, targetPath string) ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath) // 创建子文件系统，指向目标路径
	if err != nil {                          // 如果创建失败
		panic(err) // 抛出异常，这是启动时的致命错误
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys), // 转换为HTTP文件系统
	}
}

// ServeFileSystem 接口 - 扩展的文件系统接口
// 继承标准HTTP文件系统接口，并添加文件存在性检查方法
// ElasticView插件使用此接口来提供静态资源服务
type ServeFileSystem interface {
	http.FileSystem                         // 标准HTTP文件系统接口
	Exists(prefix string, path string) bool // 文件存在性检查方法
}

// Serve 函数 - 创建静态文件服务中间件
// 为ElasticView插件提供静态资源服务的Gin中间件
// 当请求的文件存在时，直接提供文件服务；否则继续到下一个中间件
// 参数 urlPrefix: URL前缀，用于路径处理
// 参数 fs: 实现了ServeFileSystem接口的文件系统
// 返回值: Gin中间件处理函数
func Serve(urlPrefix string, fs ServeFileSystem) gin.HandlerFunc {
	fileserver := http.FileServer(fs) // 创建HTTP文件服务器
	if urlPrefix != "" {              // 如果有URL前缀
		fileserver = http.StripPrefix(urlPrefix, fileserver) // 去除前缀
	}
	return func(c *gin.Context) { // 返回Gin中间件函数
		if fs.Exists(urlPrefix, c.Request.URL.Path) { // 检查请求的文件是否存在
			fileserver.ServeHTTP(c.Writer, c.Request) // 如果存在，提供文件服务
			c.Abort()                                 // 终止请求处理链，不再执行后续中间件
		}
		// 如果文件不存在，继续执行后续中间件（通常是Vue路由处理）
	}
}
