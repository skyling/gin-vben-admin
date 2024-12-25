package models

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"testing"
)

func TestA(t *testing.T) {
	snowflake.Epoch = 1701736313123
	snowflake.NodeBits = 3
	snowflake.StepBits = 7
	node, _ = snowflake.NewNode(1)
	for i := 0; i < 100; i++ {
		fmt.Println(node.Generate())
	}

}
