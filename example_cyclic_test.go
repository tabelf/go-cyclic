package main

import (
	"testing"

	"github.com/elza2/go-cyclic/tool"
)

func TestCyclic(t *testing.T) {
	tool.CheckCycleDepend("/Users/yuanyou/go/src/daji")
}
