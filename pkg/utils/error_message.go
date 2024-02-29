package utils

import "github.com/richardktran/KMY-Drug-Store/conf"

func GetMessage(messageCode string) string {
	if message, ok := conf.MessageMap[messageCode]; ok {
		return message
	}

	return messageCode
}
