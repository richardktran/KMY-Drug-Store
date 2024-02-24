package model

type ItemStatus string

const (
	ItemStatusTodo    ItemStatus = "TODO"
	ItemStatusDoing   ItemStatus = "DOING"
	ItemStatusDone    ItemStatus = "DONE"
	ItemStatusDeleted ItemStatus = "DELETED"
)
