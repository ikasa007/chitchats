package main

import "github.com/nicksnyder/go-i18n/v2/i18n"

var message = []i18n.Message{
	i18n.Message{
		ID: "thread_not_found",
		Description: "Thread not exists in db",
		Other: "Cannot read thread",
	},
	i18n.Message{
		ID: "thread_create_failed",
		Description: "Thread saved into db failed",
		Other: "Cannot saved thread",
	},
}
