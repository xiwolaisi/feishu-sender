package config

import (
	"fmt"
	"os"
	"time"

	"feishu-sender/corp"
	"github.com/toolkits/pkg/logger"
)

// InitLogger init logger toolkits
func InitLogger() {
	c := Get().Logger

	lb, err := logger.NewFileBackend(c.Dir)
	if err != nil {
		fmt.Println("cannot init logger:", err)
		os.Exit(1)
	}

	lb.SetRotateByHour(true)
	lb.SetKeepHours(c.KeepHours)

	logger.SetLogging(c.Level, lb)
}

func Test(webhook string, args []string) {

	if len(args) == 0 {
		fmt.Println("im not given")
		os.Exit(1)
	}
	for i := 0; i < len(args); i++ {
		err := corp.Send(webhook, corp.Message{
			Msgtype: "text",
			Text:    corp.Content{Content: fmt.Sprintf("鲸算告警中心test message from n9e at %v", time.Now())},
		})

		if err != nil {
			fmt.Printf("send to %s fail: %v\n", args[i], err)
		} else {
			fmt.Printf("send to %s succ\n", args[i])
		}
	}
}
