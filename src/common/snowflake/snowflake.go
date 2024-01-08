package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(machineID int64) (err error) {
	node, err = snowflake.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
