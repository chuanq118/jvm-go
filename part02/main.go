package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 1.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

// 启动 jvm
func startJVM(cmd *Cmd) {
	fmt.Printf("启动 JVM classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}

/**
测试
PS D:\go\workspace\bin> .\part01.exe -version
version 1.0
PS D:\go\workspace\bin> .\part01.exe -help
无法识别命令!
正确用法: D:\go\workspace\bin\part01.exe [-options] class(全限定类名) [args](可选参数).
PS D:\go\workspace\bin> .\part01.exe HelloWorld arg1 arg2
classpath: class:HelloWorld args:[arg1 arg2]
*/
