package main

import (
    "fmt"
    "os"
    "github.com/urfave/cli"
)

func main() {
    var m_port int
    app := cli.NewApp()
    app.Name = "greet"           // 指定程序名称
    app.Usage = "say a greeting" //  程序功能描述
    app.Flags = []cli.Flag{
        cli.IntFlag{
            Name:        "port, p",        // 配置名称
            Value:       8000,             // 缺省配置值
            Usage:       "listening port", // 配置描述
            Destination: &m_port,          // 保存配置值
        },
    }
    app.Commands = []cli.Command{
        {
            Name:     "add",         //命令名称
            Aliases:  []string{"a"}, // 命令的别名列表
            Usage:    "calc 1+1",    // 命令描述
            Category: "arithmetic",  // 命令所属的类别
            Action: func(c *cli.Context) error { // 函数入口
                fmt.Println("1 + 1 = ", 1+1)
                return nil
            },
        },
        {
            Name:     "sub",
            Aliases:  []string{"s"},
            Usage:    "calc 5-3",
            Category: "arithmetic",
            Action: func(c *cli.Context) error {
                fmt.Println("5 - 3 = ", 5-3)
                return nil
            },
        },
        {
            Name:     "db",
            Usage:    "database operations",
            Category: "database",
            Subcommands: []cli.Command{ // 子命令列表
                {
                    Name:  "insert",
                    Usage: "insert data",
                    Action: func(c *cli.Context) error {
                        fmt.Println("insert subcommand")
                        return nil
                    },
                },
                {
                    Name:  "delete",
                    Usage: "delete data",
                    Action: func(c *cli.Context) error {
                        fmt.Println("delete subcommand")
                        return nil
                    },
                },
            },
        },
    }

    app.Action = func(c *cli.Context) error {
        println("Greetings")
        fmt.Println(c.Int("port"))
        fmt.Println(m_port)
        return nil
    }

    app.Run(os.Args)
}
