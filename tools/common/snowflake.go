package common

import (
	"github.com/bwmarrin/snowflake"
)

func SnowflakeNew(nodeId int64) *snowflake.Node {
	//Node number must be between 0 and 1023
	node, _ := snowflake.NewNode(nodeId)
	return node
}

func SnowflakeIdInt64(node *snowflake.Node) int64 {
	return node.Generate().Int64()
}

func SnowflakeIdString(node *snowflake.Node) string {
	return node.Generate().String()
}
