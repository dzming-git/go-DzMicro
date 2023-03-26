package message_handler

import (
	"fmt"
)

func MessageHandler(message string, sourceID []string) {
	//TODO 测试代码，还未实现
	fmt.Println(message)
	for _, s := range sourceID {
		fmt.Println(s)
	}
}
