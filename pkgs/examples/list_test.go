package examples

import (
	"gitee.com/aurora-engine/aurora/pkgs/list"
	"testing"
)

func TestLinkList_Add(t *testing.T) {
	l := list.ArrayList[int]{1, 2, 3}
	t.Log(l)
	l.Delete(0)
	t.Log(l)
}