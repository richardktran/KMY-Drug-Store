package utils

import "github.com/richardktran/MyBlogBE/conf"

func GetMessage(messageCode string) string {
	if message, ok := conf.MessageMap[messageCode]; ok {
		return message
	}

	return messageCode
}
