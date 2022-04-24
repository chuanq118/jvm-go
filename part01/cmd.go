package main

import (
	"flag"
	"fmt"
	"os"
)

// 简单的命令行工具

// Cmd 命令行结构体
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

// 如果 Parse（）函数解析失败，它就调用 printUsage（）函数把命令的用法打印到控制台
func printUsage() {
	fmt.Printf("无法识别命令!\n正确用法: %s [-options] class(全限定类名) [args](可选参数). \n", os.Args[0])
}

// 解析命令行
func parseCmd() *Cmd {
	// cmd 定义为 Cmd 结构体的指针
	cmd := &Cmd{}
	// 将 defaultPrintUsage 定义为我们自己的 usage
	flag.Usage = printUsage
	// 定义支持的参数类型
	// flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message.")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit.")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	// 开始解析
	flag.Parse()
	// 返回命令行参数后的其他参数
	args := flag.Args()
	// 除去 options 的其他参数
	if len(args) > 0 {
		// 第一个参数指定为全限定类名
		cmd.class = args[0]
		// 之后的切片定义为类后面的参数列表
		cmd.args = args[1:]
	}
	return cmd
}
