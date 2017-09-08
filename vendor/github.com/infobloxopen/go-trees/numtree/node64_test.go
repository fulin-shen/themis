package numtree

import (
	"fmt"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

func TestInsert64(t *testing.T) {
	var r *Node64

	r = r.Insert(0, 64, "test")
	assertTree64(r, TestTree64WithSingleNodeInserted,
		"64-tree with single node inserted", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 18, "bottom")
	r = r.Insert(0xAAAAAAAA00000000, 9, "top")
	assertTree64(r, TestTree64WithTopAfterBottomToLeftNodesInserted,
		"64-tree with top after bottom to left nodes inserted", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 18, "bottom")
	r = r.Insert(0xAAAAAAAA00000000, 10, "top")
	assertTree64(r, TestTree64WithTopAfterBottomToRightNodesInserted,
		"64-tree with top after bottom to right nodes inserted", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 18, "bottom")
	r = r.Insert(0xABAAAAAA00000000, 10, "top")
	assertTree64(r, TestTree64WithTopAfterBottomAndAdditionalNotLeafNodesInserted,
		"64-tree with top after bottom and addtional not leaf nodes inserted", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 18, "bottom")
	oldR := r.Insert(0xABAAAAAA00000000, 10, "top")
	newR := oldR.Insert(0xABAAAAAA00000000, 7, "root")
	assertTree64(oldR, TestTree64WithOldTopReplacingTopAfterBottomNodesInserted,
		"64-tree with old top replacing top after bottom nodes inserted", t)
	assertTree64(newR, TestTree64WithNewTopReplacingTopAfterBottomNodesInserted,
		"64-tree with new top replacing top after bottom nodes inserted", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 9, "top")
	r = r.Insert(0xAAAAAAAA00000000, 18, "bottom")
	assertTree64(r, TestTree64WithTopBeforeBottomToLeftNodesInserted,
		"64-tree with top before bottom to left nodes inserted", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 10, "top")
	r = r.Insert(0xAAAAAAAA00000000, 18, "bottom")
	assertTree64(r, TestTree64WithTopBeforeBottomToRightNodesInserted,
		"64-tree with top before bottom to right nodes inserted", t)

	r = nil
	r = r.Insert(0xAAAAAAA00000000, 7, "L1")
	r = r.Insert(0xABAAAAA00000000, 9, "L2")
	r = r.Insert(0xAAAAAAA00000000, 18, "L3")
	r = r.Insert(0xAABAAAA00000000, 19, "L4")
	assertTree64(r, TestTree64WithTopBeforeBottomSeveralLevelNodesInserted,
		"64-tree with top before bottom several level nodes inserted", t)

	r = nil
	r = r.Insert(0, -10, nil)
	assertTree64(r, TestTree64WithNegativeNumberOfBits,
		"64-tree with negative number of significant bits", t)

	r = nil
	r = r.Insert(0, 65, nil)
	assertTree64(r, TestTree64WithTooBigNumberOfBits,
		"64-tree with too big number of significant bits", t)

	r = nil
	for i := uint64(0); i < 256; i++ {
		r = r.Insert(i, 64, fmt.Sprintf("%02x", i))
	}
	assertTree64(r, TestTree64BigTreeInsertions,
		"64-tree big tree", t)

	r = nil
	for i := uint64(0); i < 256; i++ {
		r = r.Insert(inv64[i]<<56, 64, fmt.Sprintf("%02x", inv64[i]))
	}
	assertTree64(r, TestTree64BigTreeInvertedInsertions,
		"64-tree big tree with inverted keys", t)
}

func TestInplaceInsert64(t *testing.T) {
	var r *Node64

	r = r.InplaceInsert(0, 64, "test")
	assertTree64(r, TestTree64WithSingleNodeInserted,
		"64-tree with single node inplace inserted", t)

	r = nil
	r = r.InplaceInsert(0xAAAAAAAA00000000, 18, "bottom")
	r = r.InplaceInsert(0xAAAAAAAA00000000, 9, "top")
	assertTree64(r, TestTree64WithTopAfterBottomToLeftNodesInserted,
		"64-tree with top after bottom to left nodes inplace inserted", t)

	r = nil
	r = r.InplaceInsert(0xAAAAAAAA00000000, 18, "bottom")
	r = r.InplaceInsert(0xAAAAAAAA00000000, 10, "top")
	assertTree64(r, TestTree64WithTopAfterBottomToRightNodesInserted,
		"64-tree with top after bottom to right nodes inplace inserted", t)

	r = nil
	r = r.InplaceInsert(0xAAAAAAAA00000000, 18, "bottom")
	r = r.InplaceInsert(0xABAAAAAA00000000, 10, "top")
	assertTree64(r, TestTree64WithTopAfterBottomAndAdditionalNotLeafNodesInserted,
		"64-tree with top after bottom and addtional not leaf nodes inplace inserted", t)

	r = nil
	r = r.InplaceInsert(0xAAAAAAAA00000000, 18, "bottom")
	r = r.InplaceInsert(0xABAAAAAA00000000, 10, "top")
	assertTree64(r, TestTree64WithOldTopReplacingTopAfterBottomNodesInserted,
		"64-tree with old top replacing top after bottom nodes inplace inserted", t)
	r = r.InplaceInsert(0xABAAAAAA00000000, 7, "root")
	assertTree64(r, TestTree64WithNewTopReplacingTopAfterBottomNodesInserted,
		"64-tree with new top replacing top after bottom nodes inplace inserted", t)

	r = nil
	r = r.InplaceInsert(0xAAAAAAAA00000000, 9, "top")
	r = r.InplaceInsert(0xAAAAAAAA00000000, 18, "bottom")
	assertTree64(r, TestTree64WithTopBeforeBottomToLeftNodesInserted,
		"64-tree with top before bottom to left nodes inplace inserted", t)

	r = nil
	r = r.InplaceInsert(0xAAAAAAAA00000000, 10, "top")
	r = r.InplaceInsert(0xAAAAAAAA00000000, 18, "bottom")
	assertTree64(r, TestTree64WithTopBeforeBottomToRightNodesInserted,
		"64-tree with top before bottom to right nodes inplace inserted", t)

	r = nil
	r = r.InplaceInsert(0xAAAAAAA00000000, 7, "L1")
	r = r.InplaceInsert(0xABAAAAA00000000, 9, "L2")
	r = r.InplaceInsert(0xAAAAAAA00000000, 18, "L3")
	r = r.InplaceInsert(0xAABAAAA00000000, 19, "L4")
	assertTree64(r, TestTree64WithTopBeforeBottomSeveralLevelNodesInserted,
		"64-tree with top before bottom several level nodes inplace inserted", t)

	r = nil
	r = r.InplaceInsert(0, -10, nil)
	assertTree64(r, TestTree64WithNegativeNumberOfBits,
		"64-tree with negative number of significant bits (inplace)", t)

	r = nil
	r = r.InplaceInsert(0, 65, nil)
	assertTree64(r, TestTree64WithTooBigNumberOfBits,
		"64-tree with too big number of significant bits (inplace)", t)

	r = nil
	for i := uint64(0); i < 256; i++ {
		r = r.InplaceInsert(i, 64, fmt.Sprintf("%02x", i))
	}
	assertTree64(r, TestTree64BigTreeInsertions,
		"64-tree big tree (inplace)", t)

	r = nil
	for i := uint64(0); i < 256; i++ {
		r = r.InplaceInsert(inv64[i]<<56, 64, fmt.Sprintf("%02x", inv64[i]))
	}
	assertTree64(r, TestTree64BigTreeInvertedInsertions,
		"64-tree big tree with inverted keys (inplace)", t)
}

func TestEnumerate64(t *testing.T) {
	var r *Node64

	ch := r.Enumerate()
	assertSequence64(ch, t, "64-tree empty tree")

	r = r.Insert(0xAAAAAAAA00000000, 7, "L1")
	r = r.Insert(0xA8AAAAAA00000000, 9, "L2.1")
	r = r.Insert(0xABAAAAAA00000000, 9, "L2.2")
	r = r.Insert(0xAAAAAAAA00000000, 18, "L3")
	r = r.Insert(0xAAABAAAA00000000, 24, "L5")
	r = r.Insert(0xAABAAAAA00000000, 19, "L4")
	ch = r.Enumerate()
	assertSequence64(ch, t, "64-tree for enumeration",
		"0xa8aaaaaa00000000/9: \"L2.1\"",
		"0xaaaaaaaa00000000/7: \"L1\"",
		"0xaaaaaaaa00000000/18: \"L3\"",
		"0xaaabaaaa00000000/24: \"L5\"",
		"0xaabaaaaa00000000/19: \"L4\"",
		"0xabaaaaaa00000000/9: \"L2.2\"")
}

func TestMatch64(t *testing.T) {
	var r *Node64

	v, ok := r.Match(0, 0)
	assertTreeMatch(v, ok, nil,
		"64-bit empty tree", t)

	r = r.Insert(0xAAAAAAAA00000000, 7, "L1")
	r = r.Insert(0xA8AAAAAA00000000, 9, "L2.1")
	r = r.Insert(0xABAAAAAA00000000, 9, "L2.2")
	r = r.Insert(0xAAAAAAAA00000000, 18, "L3")
	r = r.Insert(0xAABAAAAA00000000, 19, "L4")

	v, ok = r.Match(0, -1)
	assertTreeMatch(v, ok, nil,
		"64-tree match with negative significant bits", t)

	v, ok = r.Match(0xAAAAAAAA00000000, 67)
	assertTreeMatch(v, ok, wrapStr("L3"),
		"64-tree match with overflow significant bits number", t)

	v, ok = r.Match(0xAAAAAAAA00000000, 5)
	assertTreeMatch(v, ok, nil,
		"64-tree match with small significant bits number", t)

	v, ok = r.Match(0xA8AAAAAA00000000, 9)
	assertTreeMatch(v, ok, wrapStr("L2.1"),
		"64-tree match with exact match to a node", t)

	v, ok = r.Match(0xA9AAAAAA00000000, 9)
	assertTreeMatch(v, ok, nil,
		"64-tree match with exact not match to a node", t)

	v, ok = r.Match(0xAABAAACA00000000, 64)
	assertTreeMatch(v, ok, wrapStr("L4"),
		"64-tree match with contains match to child node", t)

	v, ok = r.Match(0xABAAAAAA00000000, 9)
	assertTreeMatch(v, ok, wrapStr("L2.2"),
		"64-tree match with exact match to child node", t)

	v, ok = r.Match(0xA80AAAAA00000000, 11)
	assertTreeMatch(v, ok, nil,
		"64-tree match with contains match to non-leaf node", t)
}

func TestExactMatch64(t *testing.T) {
	var r *Node64

	v, ok := r.ExactMatch(0, 0)
	assertTreeMatch(v, ok, nil,
		"64-bit empty tree", t)

	r = r.Insert(0xAAAAAAAA00000000, 7, "L1")
	r = r.Insert(0xA8AAAAAA00000000, 9, "L2.1")
	r = r.Insert(0xABAAAAAA00000000, 9, "L2.2")
	r = r.Insert(0xAAAAAAAA00000000, 18, "L3")
	r = r.Insert(0xAABAAAAA00000000, 19, "L4")

	v, ok = r.ExactMatch(0, -1)
	assertTreeMatch(v, ok, nil,
		"64-tree exact match with negative significant bits", t)

	v, ok = r.ExactMatch(0xAAAAAAAA00000000, 65)
	assertTreeMatch(v, ok, nil,
		"64-tree exact match with overflow significant bits number", t)

	v, ok = r.ExactMatch(0xAAAAAAAA00000000, 5)
	assertTreeMatch(v, ok, nil,
		"64-tree exact match with small significant bits number", t)

	v, ok = r.ExactMatch(0xA8AAAAAA00000000, 9)
	assertTreeMatch(v, ok, wrapStr("L2.1"),
		"64-tree exact match with exact match to a node", t)

	v, ok = r.ExactMatch(0xA9AAAAAA00000000, 9)
	assertTreeMatch(v, ok, nil,
		"64-tree exact match with exact not match to a node", t)

	v, ok = r.ExactMatch(0xAABAAACA00000000, 64)
	assertTreeMatch(v, ok, nil,
		"64-tree exact match with contains not match to child node", t)

	v, ok = r.ExactMatch(0xABAAAAAA00000000, 9)
	assertTreeMatch(v, ok, wrapStr("L2.2"),
		"64-tree match with exact match to child node", t)

	v, ok = r.ExactMatch(0xA80AAAAA00000000, 11)
	assertTreeMatch(v, ok, nil,
		"64-tree match with contains match to non-leaf node", t)
}

func TestDelete64(t *testing.T) {
	var (
		r  *Node64
		ok bool
	)

	r, ok = r.Delete(0, 64)
	assertTree64Delete(r, ok, "", "64-bit empty tree", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 18, "test")
	r, ok = r.Delete(0xAAAAAAAA00000000, 9)
	assertTree64Delete(r, ok, TestTree64EmptyTree,
		"64-tree with contained node", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 18, "test")
	r, ok = r.Delete(0xBBBBBBBB00000000, 9)
	assertTree64Delete(r, ok, "", "64-tree with not contained node", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 9, "test")
	r = r.Insert(0xAAAAAAAA00000000, 18, "test")
	r, ok = r.Delete(0xBBBBBBBB00000000, 10)
	assertTree64Delete(r, ok, "", "64-tree with not containing node", t)

	r, ok = r.Delete(0xAAEAAAAA00000000, 10)
	assertTree64Delete(r, ok, "", "64-tree with empty branch", t)

	r, ok = r.Delete(0xAAABBBBB00000000, 16)
	assertTree64Delete(r, ok, "", "64-tree with not contained branch", t)

	r, ok = r.Delete(0xAAAAAAAA00000000, 16)
	assertTree64Delete(r, ok, TestTree64WithDeletedChildNode,
		"64-tree with deleted child node", t)

	r, ok = r.Delete(0, -10)
	assertTree64Delete(r, ok, TestTree64EmptyTree,
		"64-tree with deleted all nodes by negative number of significant bits", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 9, "test")
	r = r.Insert(0xAAAAAAAA00000000, 64, "test")
	r, ok = r.Delete(0xAAAAAAAA00000000, 65)
	assertTree64Delete(r, ok, TestTree64WithDeletedChildNode,
		"64-tree with deleted child node by too big number of significant bits", t)

	r = nil
	r = r.Insert(0xAAAAAAAA00000000, 7, "L1")
	r = r.Insert(0xA8AAAAAA00000000, 9, "L2.1")
	r = r.Insert(0xABAAAAAA00000000, 9, "L2.2")
	r = r.Insert(0xAAAAAAAA00000000, 18, "L3")
	r = r.Insert(0xAAABAAAA00000000, 24, "L5")
	r = r.Insert(0xAABAAAAA00000000, 19, "L4")

	r, ok = r.Delete(0xAABAAAAA00000000, 19)
	assertTree64Delete(r, ok, TestTree64WithDeletedChildAndNonLeafNodes,
		"64-tree with deleted child and non-leaf node", t)

	r, ok = r.Delete(0xAAAAAAAA00000000, 18)
	assertTree64Delete(r, ok, TestTree64WithDeletedTwoChildrenAndNonLeafNodes,
		"64-tree with deleted two children and non-leaf nodes", t)
}

func TestClz64(t *testing.T) {
	assertClz64(0x0000000000000000, 64, t)
	assertClz64(0x0000000000000001, 63, t)
	assertClz64(0x0000000000000002, 62, t)
	assertClz64(0x0000000000000003, 62, t)
	assertClz64(0x0000000000000004, 61, t)
	assertClz64(0x0000000000000005, 61, t)
	assertClz64(0x0000000000000006, 61, t)
	assertClz64(0x0000000000000007, 61, t)
	assertClz64(0x0000000000000008, 60, t)

	assertClz64(0x0000000000000010, 59, t)
	assertClz64(0x0000000000000020, 58, t)
	assertClz64(0x0000000000000040, 57, t)
	assertClz64(0x0000000000000080, 56, t)

	assertClz64(0x0000000000000100, 55, t)
	assertClz64(0x0000000000000200, 54, t)
	assertClz64(0x0000000000000400, 53, t)
	assertClz64(0x0000000000000800, 52, t)

	assertClz64(0x0000000000001000, 51, t)
	assertClz64(0x0000000000002000, 50, t)
	assertClz64(0x0000000000004000, 49, t)
	assertClz64(0x0000000000008000, 48, t)

	assertClz64(0x0000000000010000, 47, t)
	assertClz64(0x0000000000020000, 46, t)
	assertClz64(0x0000000000040000, 45, t)
	assertClz64(0x0000000000080000, 44, t)

	assertClz64(0x0000000000100000, 43, t)
	assertClz64(0x0000000000200000, 42, t)
	assertClz64(0x0000000000400000, 41, t)
	assertClz64(0x0000000000800000, 40, t)

	assertClz64(0x0000000001000000, 39, t)
	assertClz64(0x0000000002000000, 38, t)
	assertClz64(0x0000000004000000, 37, t)
	assertClz64(0x0000000008000000, 36, t)

	assertClz64(0x0000000010000000, 35, t)
	assertClz64(0x0000000020000000, 34, t)
	assertClz64(0x0000000040000000, 33, t)
	assertClz64(0x0000000080000000, 32, t)

	assertClz64(0x0000000100000000, 31, t)
	assertClz64(0x0000000200000000, 30, t)
	assertClz64(0x0000000400000000, 29, t)
	assertClz64(0x0000000800000000, 28, t)

	assertClz64(0x0000001000000000, 27, t)
	assertClz64(0x0000002000000000, 26, t)
	assertClz64(0x0000004000000000, 25, t)
	assertClz64(0x0000008000000000, 24, t)

	assertClz64(0x0000010000000000, 23, t)
	assertClz64(0x0000020000000000, 22, t)
	assertClz64(0x0000040000000000, 21, t)
	assertClz64(0x0000080000000000, 20, t)

	assertClz64(0x0000100000000000, 19, t)
	assertClz64(0x0000200000000000, 18, t)
	assertClz64(0x0000400000000000, 17, t)
	assertClz64(0x0000800000000000, 16, t)

	assertClz64(0x0001000000000000, 15, t)
	assertClz64(0x0002000000000000, 14, t)
	assertClz64(0x0004000000000000, 13, t)
	assertClz64(0x0008000000000000, 12, t)

	assertClz64(0x0010000000000000, 11, t)
	assertClz64(0x0020000000000000, 10, t)
	assertClz64(0x0040000000000000, 9, t)
	assertClz64(0x0080000000000000, 8, t)

	assertClz64(0x0100000000000000, 7, t)
	assertClz64(0x0200000000000000, 6, t)
	assertClz64(0x0400000000000000, 5, t)
	assertClz64(0x0800000000000000, 4, t)

	assertClz64(0x1000000000000000, 3, t)
	assertClz64(0x2000000000000000, 2, t)
	assertClz64(0x4000000000000000, 1, t)
	assertClz64(0x8000000000000000, 0, t)
}

func assertTree64(r *Node64, e, desc string, t *testing.T) {
	assertStringLists(difflib.SplitLines(r.Dot()), difflib.SplitLines(e), desc, t)
}

func assertSequence64(ch chan *Node64, t *testing.T, desc string, e ...string) {
	items := []string{}
	for n := range ch {
		if n == nil {
			items = append(items, fmt.Sprintf("%#v\n", n))
			continue
		}

		s, ok := n.Value.(string)
		if ok {
			s = fmt.Sprintf("%q", s)
		} else {
			s = fmt.Sprintf("%#v (non-string type %T)", n.Value, n.Value)
		}

		items = append(items, fmt.Sprintf("0x%016x/%d: %s\n", n.Key, n.Bits, s))
	}

	eItems := make([]string, len(e))
	for i, item := range e {
		eItems[i] = item + "\n"
	}

	assertStringLists(items, eItems, desc, t)
}

func assertTree64Delete(r *Node64, ok bool, e string, desc string, t *testing.T) {
	if len(e) > 0 {
		if !ok {
			t.Errorf("Expected something to be deleted from %s but it isn't and got old root:\n%s\n", desc, r.Dot())
			return
		}

		assertTree64(r, e, desc, t)
	} else if ok {
		t.Errorf("Expected nothing to be deleted from %s but it is and got new root:\n%s\n", desc, r.Dot())
	}
}

func assertClz64(x uint64, c uint8, t *testing.T) {
	r := clz64(x)
	if r != c {
		t.Errorf("Expected %d as result of clz64(0x%016x) but got %d", c, x, r)
	}
}

const (
	TestTree64WithSingleNodeInserted = `digraph d {
N0 [label="k: 0000000000000000, b: 64, v: \"\"test\"\""]
}
`

	TestTree64WithTopAfterBottomToLeftNodesInserted = `digraph d {
N0 [label="k: aaaaaaaa00000000, b: 9, v: \"\"top\"\""]
N0 -> { N1 N2 }
N1 [label="k: aaaaaaaa00000000, b: 18, v: \"\"bottom\"\""]
N2 [label="nil"]
}
`

	TestTree64WithTopAfterBottomToRightNodesInserted = `digraph d {
N0 [label="k: aaaaaaaa00000000, b: 10, v: \"\"top\"\""]
N0 -> { N1 N2 }
N1 [label="nil"]
N2 [label="k: aaaaaaaa00000000, b: 18, v: \"\"bottom\"\""]
}
`

	TestTree64WithTopAfterBottomAndAdditionalNotLeafNodesInserted = `digraph d {
N0 [label="k: aa00000000000000, b: 7"]
N0 -> { N1 N2 }
N1 [label="k: aaaaaaaa00000000, b: 18, v: \"\"bottom\"\""]
N2 [label="k: abaaaaaa00000000, b: 10, v: \"\"top\"\""]
}
`

	TestTree64WithOldTopReplacingTopAfterBottomNodesInserted = `digraph d {
N0 [label="k: aa00000000000000, b: 7"]
N0 -> { N1 N2 }
N1 [label="k: aaaaaaaa00000000, b: 18, v: \"\"bottom\"\""]
N2 [label="k: abaaaaaa00000000, b: 10, v: \"\"top\"\""]
}
`

	TestTree64WithNewTopReplacingTopAfterBottomNodesInserted = `digraph d {
N0 [label="k: abaaaaaa00000000, b: 7, v: \"\"root\"\""]
N0 -> { N1 N2 }
N1 [label="k: aaaaaaaa00000000, b: 18, v: \"\"bottom\"\""]
N2 [label="k: abaaaaaa00000000, b: 10, v: \"\"top\"\""]
}
`

	TestTree64WithTopBeforeBottomToLeftNodesInserted = `digraph d {
N0 [label="k: aaaaaaaa00000000, b: 9, v: \"\"top\"\""]
N0 -> { N1 N2 }
N1 [label="k: aaaaaaaa00000000, b: 18, v: \"\"bottom\"\""]
N2 [label="nil"]
}
`

	TestTree64WithTopBeforeBottomToRightNodesInserted = `digraph d {
N0 [label="k: aaaaaaaa00000000, b: 10, v: \"\"top\"\""]
N0 -> { N1 N2 }
N1 [label="nil"]
N2 [label="k: aaaaaaaa00000000, b: 18, v: \"\"bottom\"\""]
}
`

	TestTree64WithTopBeforeBottomSeveralLevelNodesInserted = `digraph d {
N0 [label="k: 0aaaaaaa00000000, b: 7, v: \"\"L1\"\""]
N0 -> { N1 N2 }
N1 [label="k: 0abaaaaa00000000, b: 9, v: \"\"L2\"\""]
N1 -> { N3 N4 }
N2 [label="nil"]
N3 [label="k: 0aaa000000000000, b: 15"]
N3 -> { N5 N6 }
N4 [label="nil"]
N5 [label="k: 0aaaaaaa00000000, b: 18, v: \"\"L3\"\""]
N6 [label="k: 0aabaaaa00000000, b: 19, v: \"\"L4\"\""]
}
`

	TestTree64WithNegativeNumberOfBits = `digraph d {
N0 [label="k: 0000000000000000, b: 0, v: \"<nil>\""]
}
`

	TestTree64WithTooBigNumberOfBits = `digraph d {
N0 [label="k: 0000000000000000, b: 64, v: \"<nil>\""]
}
`

	TestTree64BigTreeInsertions = `digraph d {
N0 [label="k: 0000000000000000, b: 56"]
N0 -> { N1 N2 }
N1 [label="k: 0000000000000000, b: 57"]
N1 -> { N3 N4 }
N2 [label="k: 0000000000000080, b: 57"]
N2 -> { N5 N6 }
N3 [label="k: 0000000000000000, b: 58"]
N3 -> { N7 N8 }
N4 [label="k: 0000000000000040, b: 58"]
N4 -> { N9 N10 }
N5 [label="k: 0000000000000080, b: 58"]
N5 -> { N11 N12 }
N6 [label="k: 00000000000000c0, b: 58"]
N6 -> { N13 N14 }
N7 [label="k: 0000000000000000, b: 59"]
N7 -> { N15 N16 }
N8 [label="k: 0000000000000020, b: 59"]
N8 -> { N17 N18 }
N9 [label="k: 0000000000000040, b: 59"]
N9 -> { N19 N20 }
N10 [label="k: 0000000000000060, b: 59"]
N10 -> { N21 N22 }
N11 [label="k: 0000000000000080, b: 59"]
N11 -> { N23 N24 }
N12 [label="k: 00000000000000a0, b: 59"]
N12 -> { N25 N26 }
N13 [label="k: 00000000000000c0, b: 59"]
N13 -> { N27 N28 }
N14 [label="k: 00000000000000e0, b: 59"]
N14 -> { N29 N30 }
N15 [label="k: 0000000000000000, b: 60"]
N15 -> { N31 N32 }
N16 [label="k: 0000000000000010, b: 60"]
N16 -> { N33 N34 }
N17 [label="k: 0000000000000020, b: 60"]
N17 -> { N35 N36 }
N18 [label="k: 0000000000000030, b: 60"]
N18 -> { N37 N38 }
N19 [label="k: 0000000000000040, b: 60"]
N19 -> { N39 N40 }
N20 [label="k: 0000000000000050, b: 60"]
N20 -> { N41 N42 }
N21 [label="k: 0000000000000060, b: 60"]
N21 -> { N43 N44 }
N22 [label="k: 0000000000000070, b: 60"]
N22 -> { N45 N46 }
N23 [label="k: 0000000000000080, b: 60"]
N23 -> { N47 N48 }
N24 [label="k: 0000000000000090, b: 60"]
N24 -> { N49 N50 }
N25 [label="k: 00000000000000a0, b: 60"]
N25 -> { N51 N52 }
N26 [label="k: 00000000000000b0, b: 60"]
N26 -> { N53 N54 }
N27 [label="k: 00000000000000c0, b: 60"]
N27 -> { N55 N56 }
N28 [label="k: 00000000000000d0, b: 60"]
N28 -> { N57 N58 }
N29 [label="k: 00000000000000e0, b: 60"]
N29 -> { N59 N60 }
N30 [label="k: 00000000000000f0, b: 60"]
N30 -> { N61 N62 }
N31 [label="k: 0000000000000000, b: 61"]
N31 -> { N63 N64 }
N32 [label="k: 0000000000000008, b: 61"]
N32 -> { N65 N66 }
N33 [label="k: 0000000000000010, b: 61"]
N33 -> { N67 N68 }
N34 [label="k: 0000000000000018, b: 61"]
N34 -> { N69 N70 }
N35 [label="k: 0000000000000020, b: 61"]
N35 -> { N71 N72 }
N36 [label="k: 0000000000000028, b: 61"]
N36 -> { N73 N74 }
N37 [label="k: 0000000000000030, b: 61"]
N37 -> { N75 N76 }
N38 [label="k: 0000000000000038, b: 61"]
N38 -> { N77 N78 }
N39 [label="k: 0000000000000040, b: 61"]
N39 -> { N79 N80 }
N40 [label="k: 0000000000000048, b: 61"]
N40 -> { N81 N82 }
N41 [label="k: 0000000000000050, b: 61"]
N41 -> { N83 N84 }
N42 [label="k: 0000000000000058, b: 61"]
N42 -> { N85 N86 }
N43 [label="k: 0000000000000060, b: 61"]
N43 -> { N87 N88 }
N44 [label="k: 0000000000000068, b: 61"]
N44 -> { N89 N90 }
N45 [label="k: 0000000000000070, b: 61"]
N45 -> { N91 N92 }
N46 [label="k: 0000000000000078, b: 61"]
N46 -> { N93 N94 }
N47 [label="k: 0000000000000080, b: 61"]
N47 -> { N95 N96 }
N48 [label="k: 0000000000000088, b: 61"]
N48 -> { N97 N98 }
N49 [label="k: 0000000000000090, b: 61"]
N49 -> { N99 N100 }
N50 [label="k: 0000000000000098, b: 61"]
N50 -> { N101 N102 }
N51 [label="k: 00000000000000a0, b: 61"]
N51 -> { N103 N104 }
N52 [label="k: 00000000000000a8, b: 61"]
N52 -> { N105 N106 }
N53 [label="k: 00000000000000b0, b: 61"]
N53 -> { N107 N108 }
N54 [label="k: 00000000000000b8, b: 61"]
N54 -> { N109 N110 }
N55 [label="k: 00000000000000c0, b: 61"]
N55 -> { N111 N112 }
N56 [label="k: 00000000000000c8, b: 61"]
N56 -> { N113 N114 }
N57 [label="k: 00000000000000d0, b: 61"]
N57 -> { N115 N116 }
N58 [label="k: 00000000000000d8, b: 61"]
N58 -> { N117 N118 }
N59 [label="k: 00000000000000e0, b: 61"]
N59 -> { N119 N120 }
N60 [label="k: 00000000000000e8, b: 61"]
N60 -> { N121 N122 }
N61 [label="k: 00000000000000f0, b: 61"]
N61 -> { N123 N124 }
N62 [label="k: 00000000000000f8, b: 61"]
N62 -> { N125 N126 }
N63 [label="k: 0000000000000000, b: 62"]
N63 -> { N127 N128 }
N64 [label="k: 0000000000000004, b: 62"]
N64 -> { N129 N130 }
N65 [label="k: 0000000000000008, b: 62"]
N65 -> { N131 N132 }
N66 [label="k: 000000000000000c, b: 62"]
N66 -> { N133 N134 }
N67 [label="k: 0000000000000010, b: 62"]
N67 -> { N135 N136 }
N68 [label="k: 0000000000000014, b: 62"]
N68 -> { N137 N138 }
N69 [label="k: 0000000000000018, b: 62"]
N69 -> { N139 N140 }
N70 [label="k: 000000000000001c, b: 62"]
N70 -> { N141 N142 }
N71 [label="k: 0000000000000020, b: 62"]
N71 -> { N143 N144 }
N72 [label="k: 0000000000000024, b: 62"]
N72 -> { N145 N146 }
N73 [label="k: 0000000000000028, b: 62"]
N73 -> { N147 N148 }
N74 [label="k: 000000000000002c, b: 62"]
N74 -> { N149 N150 }
N75 [label="k: 0000000000000030, b: 62"]
N75 -> { N151 N152 }
N76 [label="k: 0000000000000034, b: 62"]
N76 -> { N153 N154 }
N77 [label="k: 0000000000000038, b: 62"]
N77 -> { N155 N156 }
N78 [label="k: 000000000000003c, b: 62"]
N78 -> { N157 N158 }
N79 [label="k: 0000000000000040, b: 62"]
N79 -> { N159 N160 }
N80 [label="k: 0000000000000044, b: 62"]
N80 -> { N161 N162 }
N81 [label="k: 0000000000000048, b: 62"]
N81 -> { N163 N164 }
N82 [label="k: 000000000000004c, b: 62"]
N82 -> { N165 N166 }
N83 [label="k: 0000000000000050, b: 62"]
N83 -> { N167 N168 }
N84 [label="k: 0000000000000054, b: 62"]
N84 -> { N169 N170 }
N85 [label="k: 0000000000000058, b: 62"]
N85 -> { N171 N172 }
N86 [label="k: 000000000000005c, b: 62"]
N86 -> { N173 N174 }
N87 [label="k: 0000000000000060, b: 62"]
N87 -> { N175 N176 }
N88 [label="k: 0000000000000064, b: 62"]
N88 -> { N177 N178 }
N89 [label="k: 0000000000000068, b: 62"]
N89 -> { N179 N180 }
N90 [label="k: 000000000000006c, b: 62"]
N90 -> { N181 N182 }
N91 [label="k: 0000000000000070, b: 62"]
N91 -> { N183 N184 }
N92 [label="k: 0000000000000074, b: 62"]
N92 -> { N185 N186 }
N93 [label="k: 0000000000000078, b: 62"]
N93 -> { N187 N188 }
N94 [label="k: 000000000000007c, b: 62"]
N94 -> { N189 N190 }
N95 [label="k: 0000000000000080, b: 62"]
N95 -> { N191 N192 }
N96 [label="k: 0000000000000084, b: 62"]
N96 -> { N193 N194 }
N97 [label="k: 0000000000000088, b: 62"]
N97 -> { N195 N196 }
N98 [label="k: 000000000000008c, b: 62"]
N98 -> { N197 N198 }
N99 [label="k: 0000000000000090, b: 62"]
N99 -> { N199 N200 }
N100 [label="k: 0000000000000094, b: 62"]
N100 -> { N201 N202 }
N101 [label="k: 0000000000000098, b: 62"]
N101 -> { N203 N204 }
N102 [label="k: 000000000000009c, b: 62"]
N102 -> { N205 N206 }
N103 [label="k: 00000000000000a0, b: 62"]
N103 -> { N207 N208 }
N104 [label="k: 00000000000000a4, b: 62"]
N104 -> { N209 N210 }
N105 [label="k: 00000000000000a8, b: 62"]
N105 -> { N211 N212 }
N106 [label="k: 00000000000000ac, b: 62"]
N106 -> { N213 N214 }
N107 [label="k: 00000000000000b0, b: 62"]
N107 -> { N215 N216 }
N108 [label="k: 00000000000000b4, b: 62"]
N108 -> { N217 N218 }
N109 [label="k: 00000000000000b8, b: 62"]
N109 -> { N219 N220 }
N110 [label="k: 00000000000000bc, b: 62"]
N110 -> { N221 N222 }
N111 [label="k: 00000000000000c0, b: 62"]
N111 -> { N223 N224 }
N112 [label="k: 00000000000000c4, b: 62"]
N112 -> { N225 N226 }
N113 [label="k: 00000000000000c8, b: 62"]
N113 -> { N227 N228 }
N114 [label="k: 00000000000000cc, b: 62"]
N114 -> { N229 N230 }
N115 [label="k: 00000000000000d0, b: 62"]
N115 -> { N231 N232 }
N116 [label="k: 00000000000000d4, b: 62"]
N116 -> { N233 N234 }
N117 [label="k: 00000000000000d8, b: 62"]
N117 -> { N235 N236 }
N118 [label="k: 00000000000000dc, b: 62"]
N118 -> { N237 N238 }
N119 [label="k: 00000000000000e0, b: 62"]
N119 -> { N239 N240 }
N120 [label="k: 00000000000000e4, b: 62"]
N120 -> { N241 N242 }
N121 [label="k: 00000000000000e8, b: 62"]
N121 -> { N243 N244 }
N122 [label="k: 00000000000000ec, b: 62"]
N122 -> { N245 N246 }
N123 [label="k: 00000000000000f0, b: 62"]
N123 -> { N247 N248 }
N124 [label="k: 00000000000000f4, b: 62"]
N124 -> { N249 N250 }
N125 [label="k: 00000000000000f8, b: 62"]
N125 -> { N251 N252 }
N126 [label="k: 00000000000000fc, b: 62"]
N126 -> { N253 N254 }
N127 [label="k: 0000000000000000, b: 63"]
N127 -> { N255 N256 }
N128 [label="k: 0000000000000002, b: 63"]
N128 -> { N257 N258 }
N129 [label="k: 0000000000000004, b: 63"]
N129 -> { N259 N260 }
N130 [label="k: 0000000000000006, b: 63"]
N130 -> { N261 N262 }
N131 [label="k: 0000000000000008, b: 63"]
N131 -> { N263 N264 }
N132 [label="k: 000000000000000a, b: 63"]
N132 -> { N265 N266 }
N133 [label="k: 000000000000000c, b: 63"]
N133 -> { N267 N268 }
N134 [label="k: 000000000000000e, b: 63"]
N134 -> { N269 N270 }
N135 [label="k: 0000000000000010, b: 63"]
N135 -> { N271 N272 }
N136 [label="k: 0000000000000012, b: 63"]
N136 -> { N273 N274 }
N137 [label="k: 0000000000000014, b: 63"]
N137 -> { N275 N276 }
N138 [label="k: 0000000000000016, b: 63"]
N138 -> { N277 N278 }
N139 [label="k: 0000000000000018, b: 63"]
N139 -> { N279 N280 }
N140 [label="k: 000000000000001a, b: 63"]
N140 -> { N281 N282 }
N141 [label="k: 000000000000001c, b: 63"]
N141 -> { N283 N284 }
N142 [label="k: 000000000000001e, b: 63"]
N142 -> { N285 N286 }
N143 [label="k: 0000000000000020, b: 63"]
N143 -> { N287 N288 }
N144 [label="k: 0000000000000022, b: 63"]
N144 -> { N289 N290 }
N145 [label="k: 0000000000000024, b: 63"]
N145 -> { N291 N292 }
N146 [label="k: 0000000000000026, b: 63"]
N146 -> { N293 N294 }
N147 [label="k: 0000000000000028, b: 63"]
N147 -> { N295 N296 }
N148 [label="k: 000000000000002a, b: 63"]
N148 -> { N297 N298 }
N149 [label="k: 000000000000002c, b: 63"]
N149 -> { N299 N300 }
N150 [label="k: 000000000000002e, b: 63"]
N150 -> { N301 N302 }
N151 [label="k: 0000000000000030, b: 63"]
N151 -> { N303 N304 }
N152 [label="k: 0000000000000032, b: 63"]
N152 -> { N305 N306 }
N153 [label="k: 0000000000000034, b: 63"]
N153 -> { N307 N308 }
N154 [label="k: 0000000000000036, b: 63"]
N154 -> { N309 N310 }
N155 [label="k: 0000000000000038, b: 63"]
N155 -> { N311 N312 }
N156 [label="k: 000000000000003a, b: 63"]
N156 -> { N313 N314 }
N157 [label="k: 000000000000003c, b: 63"]
N157 -> { N315 N316 }
N158 [label="k: 000000000000003e, b: 63"]
N158 -> { N317 N318 }
N159 [label="k: 0000000000000040, b: 63"]
N159 -> { N319 N320 }
N160 [label="k: 0000000000000042, b: 63"]
N160 -> { N321 N322 }
N161 [label="k: 0000000000000044, b: 63"]
N161 -> { N323 N324 }
N162 [label="k: 0000000000000046, b: 63"]
N162 -> { N325 N326 }
N163 [label="k: 0000000000000048, b: 63"]
N163 -> { N327 N328 }
N164 [label="k: 000000000000004a, b: 63"]
N164 -> { N329 N330 }
N165 [label="k: 000000000000004c, b: 63"]
N165 -> { N331 N332 }
N166 [label="k: 000000000000004e, b: 63"]
N166 -> { N333 N334 }
N167 [label="k: 0000000000000050, b: 63"]
N167 -> { N335 N336 }
N168 [label="k: 0000000000000052, b: 63"]
N168 -> { N337 N338 }
N169 [label="k: 0000000000000054, b: 63"]
N169 -> { N339 N340 }
N170 [label="k: 0000000000000056, b: 63"]
N170 -> { N341 N342 }
N171 [label="k: 0000000000000058, b: 63"]
N171 -> { N343 N344 }
N172 [label="k: 000000000000005a, b: 63"]
N172 -> { N345 N346 }
N173 [label="k: 000000000000005c, b: 63"]
N173 -> { N347 N348 }
N174 [label="k: 000000000000005e, b: 63"]
N174 -> { N349 N350 }
N175 [label="k: 0000000000000060, b: 63"]
N175 -> { N351 N352 }
N176 [label="k: 0000000000000062, b: 63"]
N176 -> { N353 N354 }
N177 [label="k: 0000000000000064, b: 63"]
N177 -> { N355 N356 }
N178 [label="k: 0000000000000066, b: 63"]
N178 -> { N357 N358 }
N179 [label="k: 0000000000000068, b: 63"]
N179 -> { N359 N360 }
N180 [label="k: 000000000000006a, b: 63"]
N180 -> { N361 N362 }
N181 [label="k: 000000000000006c, b: 63"]
N181 -> { N363 N364 }
N182 [label="k: 000000000000006e, b: 63"]
N182 -> { N365 N366 }
N183 [label="k: 0000000000000070, b: 63"]
N183 -> { N367 N368 }
N184 [label="k: 0000000000000072, b: 63"]
N184 -> { N369 N370 }
N185 [label="k: 0000000000000074, b: 63"]
N185 -> { N371 N372 }
N186 [label="k: 0000000000000076, b: 63"]
N186 -> { N373 N374 }
N187 [label="k: 0000000000000078, b: 63"]
N187 -> { N375 N376 }
N188 [label="k: 000000000000007a, b: 63"]
N188 -> { N377 N378 }
N189 [label="k: 000000000000007c, b: 63"]
N189 -> { N379 N380 }
N190 [label="k: 000000000000007e, b: 63"]
N190 -> { N381 N382 }
N191 [label="k: 0000000000000080, b: 63"]
N191 -> { N383 N384 }
N192 [label="k: 0000000000000082, b: 63"]
N192 -> { N385 N386 }
N193 [label="k: 0000000000000084, b: 63"]
N193 -> { N387 N388 }
N194 [label="k: 0000000000000086, b: 63"]
N194 -> { N389 N390 }
N195 [label="k: 0000000000000088, b: 63"]
N195 -> { N391 N392 }
N196 [label="k: 000000000000008a, b: 63"]
N196 -> { N393 N394 }
N197 [label="k: 000000000000008c, b: 63"]
N197 -> { N395 N396 }
N198 [label="k: 000000000000008e, b: 63"]
N198 -> { N397 N398 }
N199 [label="k: 0000000000000090, b: 63"]
N199 -> { N399 N400 }
N200 [label="k: 0000000000000092, b: 63"]
N200 -> { N401 N402 }
N201 [label="k: 0000000000000094, b: 63"]
N201 -> { N403 N404 }
N202 [label="k: 0000000000000096, b: 63"]
N202 -> { N405 N406 }
N203 [label="k: 0000000000000098, b: 63"]
N203 -> { N407 N408 }
N204 [label="k: 000000000000009a, b: 63"]
N204 -> { N409 N410 }
N205 [label="k: 000000000000009c, b: 63"]
N205 -> { N411 N412 }
N206 [label="k: 000000000000009e, b: 63"]
N206 -> { N413 N414 }
N207 [label="k: 00000000000000a0, b: 63"]
N207 -> { N415 N416 }
N208 [label="k: 00000000000000a2, b: 63"]
N208 -> { N417 N418 }
N209 [label="k: 00000000000000a4, b: 63"]
N209 -> { N419 N420 }
N210 [label="k: 00000000000000a6, b: 63"]
N210 -> { N421 N422 }
N211 [label="k: 00000000000000a8, b: 63"]
N211 -> { N423 N424 }
N212 [label="k: 00000000000000aa, b: 63"]
N212 -> { N425 N426 }
N213 [label="k: 00000000000000ac, b: 63"]
N213 -> { N427 N428 }
N214 [label="k: 00000000000000ae, b: 63"]
N214 -> { N429 N430 }
N215 [label="k: 00000000000000b0, b: 63"]
N215 -> { N431 N432 }
N216 [label="k: 00000000000000b2, b: 63"]
N216 -> { N433 N434 }
N217 [label="k: 00000000000000b4, b: 63"]
N217 -> { N435 N436 }
N218 [label="k: 00000000000000b6, b: 63"]
N218 -> { N437 N438 }
N219 [label="k: 00000000000000b8, b: 63"]
N219 -> { N439 N440 }
N220 [label="k: 00000000000000ba, b: 63"]
N220 -> { N441 N442 }
N221 [label="k: 00000000000000bc, b: 63"]
N221 -> { N443 N444 }
N222 [label="k: 00000000000000be, b: 63"]
N222 -> { N445 N446 }
N223 [label="k: 00000000000000c0, b: 63"]
N223 -> { N447 N448 }
N224 [label="k: 00000000000000c2, b: 63"]
N224 -> { N449 N450 }
N225 [label="k: 00000000000000c4, b: 63"]
N225 -> { N451 N452 }
N226 [label="k: 00000000000000c6, b: 63"]
N226 -> { N453 N454 }
N227 [label="k: 00000000000000c8, b: 63"]
N227 -> { N455 N456 }
N228 [label="k: 00000000000000ca, b: 63"]
N228 -> { N457 N458 }
N229 [label="k: 00000000000000cc, b: 63"]
N229 -> { N459 N460 }
N230 [label="k: 00000000000000ce, b: 63"]
N230 -> { N461 N462 }
N231 [label="k: 00000000000000d0, b: 63"]
N231 -> { N463 N464 }
N232 [label="k: 00000000000000d2, b: 63"]
N232 -> { N465 N466 }
N233 [label="k: 00000000000000d4, b: 63"]
N233 -> { N467 N468 }
N234 [label="k: 00000000000000d6, b: 63"]
N234 -> { N469 N470 }
N235 [label="k: 00000000000000d8, b: 63"]
N235 -> { N471 N472 }
N236 [label="k: 00000000000000da, b: 63"]
N236 -> { N473 N474 }
N237 [label="k: 00000000000000dc, b: 63"]
N237 -> { N475 N476 }
N238 [label="k: 00000000000000de, b: 63"]
N238 -> { N477 N478 }
N239 [label="k: 00000000000000e0, b: 63"]
N239 -> { N479 N480 }
N240 [label="k: 00000000000000e2, b: 63"]
N240 -> { N481 N482 }
N241 [label="k: 00000000000000e4, b: 63"]
N241 -> { N483 N484 }
N242 [label="k: 00000000000000e6, b: 63"]
N242 -> { N485 N486 }
N243 [label="k: 00000000000000e8, b: 63"]
N243 -> { N487 N488 }
N244 [label="k: 00000000000000ea, b: 63"]
N244 -> { N489 N490 }
N245 [label="k: 00000000000000ec, b: 63"]
N245 -> { N491 N492 }
N246 [label="k: 00000000000000ee, b: 63"]
N246 -> { N493 N494 }
N247 [label="k: 00000000000000f0, b: 63"]
N247 -> { N495 N496 }
N248 [label="k: 00000000000000f2, b: 63"]
N248 -> { N497 N498 }
N249 [label="k: 00000000000000f4, b: 63"]
N249 -> { N499 N500 }
N250 [label="k: 00000000000000f6, b: 63"]
N250 -> { N501 N502 }
N251 [label="k: 00000000000000f8, b: 63"]
N251 -> { N503 N504 }
N252 [label="k: 00000000000000fa, b: 63"]
N252 -> { N505 N506 }
N253 [label="k: 00000000000000fc, b: 63"]
N253 -> { N507 N508 }
N254 [label="k: 00000000000000fe, b: 63"]
N254 -> { N509 N510 }
N255 [label="k: 0000000000000000, b: 64, v: \"\"00\"\""]
N256 [label="k: 0000000000000001, b: 64, v: \"\"01\"\""]
N257 [label="k: 0000000000000002, b: 64, v: \"\"02\"\""]
N258 [label="k: 0000000000000003, b: 64, v: \"\"03\"\""]
N259 [label="k: 0000000000000004, b: 64, v: \"\"04\"\""]
N260 [label="k: 0000000000000005, b: 64, v: \"\"05\"\""]
N261 [label="k: 0000000000000006, b: 64, v: \"\"06\"\""]
N262 [label="k: 0000000000000007, b: 64, v: \"\"07\"\""]
N263 [label="k: 0000000000000008, b: 64, v: \"\"08\"\""]
N264 [label="k: 0000000000000009, b: 64, v: \"\"09\"\""]
N265 [label="k: 000000000000000a, b: 64, v: \"\"0a\"\""]
N266 [label="k: 000000000000000b, b: 64, v: \"\"0b\"\""]
N267 [label="k: 000000000000000c, b: 64, v: \"\"0c\"\""]
N268 [label="k: 000000000000000d, b: 64, v: \"\"0d\"\""]
N269 [label="k: 000000000000000e, b: 64, v: \"\"0e\"\""]
N270 [label="k: 000000000000000f, b: 64, v: \"\"0f\"\""]
N271 [label="k: 0000000000000010, b: 64, v: \"\"10\"\""]
N272 [label="k: 0000000000000011, b: 64, v: \"\"11\"\""]
N273 [label="k: 0000000000000012, b: 64, v: \"\"12\"\""]
N274 [label="k: 0000000000000013, b: 64, v: \"\"13\"\""]
N275 [label="k: 0000000000000014, b: 64, v: \"\"14\"\""]
N276 [label="k: 0000000000000015, b: 64, v: \"\"15\"\""]
N277 [label="k: 0000000000000016, b: 64, v: \"\"16\"\""]
N278 [label="k: 0000000000000017, b: 64, v: \"\"17\"\""]
N279 [label="k: 0000000000000018, b: 64, v: \"\"18\"\""]
N280 [label="k: 0000000000000019, b: 64, v: \"\"19\"\""]
N281 [label="k: 000000000000001a, b: 64, v: \"\"1a\"\""]
N282 [label="k: 000000000000001b, b: 64, v: \"\"1b\"\""]
N283 [label="k: 000000000000001c, b: 64, v: \"\"1c\"\""]
N284 [label="k: 000000000000001d, b: 64, v: \"\"1d\"\""]
N285 [label="k: 000000000000001e, b: 64, v: \"\"1e\"\""]
N286 [label="k: 000000000000001f, b: 64, v: \"\"1f\"\""]
N287 [label="k: 0000000000000020, b: 64, v: \"\"20\"\""]
N288 [label="k: 0000000000000021, b: 64, v: \"\"21\"\""]
N289 [label="k: 0000000000000022, b: 64, v: \"\"22\"\""]
N290 [label="k: 0000000000000023, b: 64, v: \"\"23\"\""]
N291 [label="k: 0000000000000024, b: 64, v: \"\"24\"\""]
N292 [label="k: 0000000000000025, b: 64, v: \"\"25\"\""]
N293 [label="k: 0000000000000026, b: 64, v: \"\"26\"\""]
N294 [label="k: 0000000000000027, b: 64, v: \"\"27\"\""]
N295 [label="k: 0000000000000028, b: 64, v: \"\"28\"\""]
N296 [label="k: 0000000000000029, b: 64, v: \"\"29\"\""]
N297 [label="k: 000000000000002a, b: 64, v: \"\"2a\"\""]
N298 [label="k: 000000000000002b, b: 64, v: \"\"2b\"\""]
N299 [label="k: 000000000000002c, b: 64, v: \"\"2c\"\""]
N300 [label="k: 000000000000002d, b: 64, v: \"\"2d\"\""]
N301 [label="k: 000000000000002e, b: 64, v: \"\"2e\"\""]
N302 [label="k: 000000000000002f, b: 64, v: \"\"2f\"\""]
N303 [label="k: 0000000000000030, b: 64, v: \"\"30\"\""]
N304 [label="k: 0000000000000031, b: 64, v: \"\"31\"\""]
N305 [label="k: 0000000000000032, b: 64, v: \"\"32\"\""]
N306 [label="k: 0000000000000033, b: 64, v: \"\"33\"\""]
N307 [label="k: 0000000000000034, b: 64, v: \"\"34\"\""]
N308 [label="k: 0000000000000035, b: 64, v: \"\"35\"\""]
N309 [label="k: 0000000000000036, b: 64, v: \"\"36\"\""]
N310 [label="k: 0000000000000037, b: 64, v: \"\"37\"\""]
N311 [label="k: 0000000000000038, b: 64, v: \"\"38\"\""]
N312 [label="k: 0000000000000039, b: 64, v: \"\"39\"\""]
N313 [label="k: 000000000000003a, b: 64, v: \"\"3a\"\""]
N314 [label="k: 000000000000003b, b: 64, v: \"\"3b\"\""]
N315 [label="k: 000000000000003c, b: 64, v: \"\"3c\"\""]
N316 [label="k: 000000000000003d, b: 64, v: \"\"3d\"\""]
N317 [label="k: 000000000000003e, b: 64, v: \"\"3e\"\""]
N318 [label="k: 000000000000003f, b: 64, v: \"\"3f\"\""]
N319 [label="k: 0000000000000040, b: 64, v: \"\"40\"\""]
N320 [label="k: 0000000000000041, b: 64, v: \"\"41\"\""]
N321 [label="k: 0000000000000042, b: 64, v: \"\"42\"\""]
N322 [label="k: 0000000000000043, b: 64, v: \"\"43\"\""]
N323 [label="k: 0000000000000044, b: 64, v: \"\"44\"\""]
N324 [label="k: 0000000000000045, b: 64, v: \"\"45\"\""]
N325 [label="k: 0000000000000046, b: 64, v: \"\"46\"\""]
N326 [label="k: 0000000000000047, b: 64, v: \"\"47\"\""]
N327 [label="k: 0000000000000048, b: 64, v: \"\"48\"\""]
N328 [label="k: 0000000000000049, b: 64, v: \"\"49\"\""]
N329 [label="k: 000000000000004a, b: 64, v: \"\"4a\"\""]
N330 [label="k: 000000000000004b, b: 64, v: \"\"4b\"\""]
N331 [label="k: 000000000000004c, b: 64, v: \"\"4c\"\""]
N332 [label="k: 000000000000004d, b: 64, v: \"\"4d\"\""]
N333 [label="k: 000000000000004e, b: 64, v: \"\"4e\"\""]
N334 [label="k: 000000000000004f, b: 64, v: \"\"4f\"\""]
N335 [label="k: 0000000000000050, b: 64, v: \"\"50\"\""]
N336 [label="k: 0000000000000051, b: 64, v: \"\"51\"\""]
N337 [label="k: 0000000000000052, b: 64, v: \"\"52\"\""]
N338 [label="k: 0000000000000053, b: 64, v: \"\"53\"\""]
N339 [label="k: 0000000000000054, b: 64, v: \"\"54\"\""]
N340 [label="k: 0000000000000055, b: 64, v: \"\"55\"\""]
N341 [label="k: 0000000000000056, b: 64, v: \"\"56\"\""]
N342 [label="k: 0000000000000057, b: 64, v: \"\"57\"\""]
N343 [label="k: 0000000000000058, b: 64, v: \"\"58\"\""]
N344 [label="k: 0000000000000059, b: 64, v: \"\"59\"\""]
N345 [label="k: 000000000000005a, b: 64, v: \"\"5a\"\""]
N346 [label="k: 000000000000005b, b: 64, v: \"\"5b\"\""]
N347 [label="k: 000000000000005c, b: 64, v: \"\"5c\"\""]
N348 [label="k: 000000000000005d, b: 64, v: \"\"5d\"\""]
N349 [label="k: 000000000000005e, b: 64, v: \"\"5e\"\""]
N350 [label="k: 000000000000005f, b: 64, v: \"\"5f\"\""]
N351 [label="k: 0000000000000060, b: 64, v: \"\"60\"\""]
N352 [label="k: 0000000000000061, b: 64, v: \"\"61\"\""]
N353 [label="k: 0000000000000062, b: 64, v: \"\"62\"\""]
N354 [label="k: 0000000000000063, b: 64, v: \"\"63\"\""]
N355 [label="k: 0000000000000064, b: 64, v: \"\"64\"\""]
N356 [label="k: 0000000000000065, b: 64, v: \"\"65\"\""]
N357 [label="k: 0000000000000066, b: 64, v: \"\"66\"\""]
N358 [label="k: 0000000000000067, b: 64, v: \"\"67\"\""]
N359 [label="k: 0000000000000068, b: 64, v: \"\"68\"\""]
N360 [label="k: 0000000000000069, b: 64, v: \"\"69\"\""]
N361 [label="k: 000000000000006a, b: 64, v: \"\"6a\"\""]
N362 [label="k: 000000000000006b, b: 64, v: \"\"6b\"\""]
N363 [label="k: 000000000000006c, b: 64, v: \"\"6c\"\""]
N364 [label="k: 000000000000006d, b: 64, v: \"\"6d\"\""]
N365 [label="k: 000000000000006e, b: 64, v: \"\"6e\"\""]
N366 [label="k: 000000000000006f, b: 64, v: \"\"6f\"\""]
N367 [label="k: 0000000000000070, b: 64, v: \"\"70\"\""]
N368 [label="k: 0000000000000071, b: 64, v: \"\"71\"\""]
N369 [label="k: 0000000000000072, b: 64, v: \"\"72\"\""]
N370 [label="k: 0000000000000073, b: 64, v: \"\"73\"\""]
N371 [label="k: 0000000000000074, b: 64, v: \"\"74\"\""]
N372 [label="k: 0000000000000075, b: 64, v: \"\"75\"\""]
N373 [label="k: 0000000000000076, b: 64, v: \"\"76\"\""]
N374 [label="k: 0000000000000077, b: 64, v: \"\"77\"\""]
N375 [label="k: 0000000000000078, b: 64, v: \"\"78\"\""]
N376 [label="k: 0000000000000079, b: 64, v: \"\"79\"\""]
N377 [label="k: 000000000000007a, b: 64, v: \"\"7a\"\""]
N378 [label="k: 000000000000007b, b: 64, v: \"\"7b\"\""]
N379 [label="k: 000000000000007c, b: 64, v: \"\"7c\"\""]
N380 [label="k: 000000000000007d, b: 64, v: \"\"7d\"\""]
N381 [label="k: 000000000000007e, b: 64, v: \"\"7e\"\""]
N382 [label="k: 000000000000007f, b: 64, v: \"\"7f\"\""]
N383 [label="k: 0000000000000080, b: 64, v: \"\"80\"\""]
N384 [label="k: 0000000000000081, b: 64, v: \"\"81\"\""]
N385 [label="k: 0000000000000082, b: 64, v: \"\"82\"\""]
N386 [label="k: 0000000000000083, b: 64, v: \"\"83\"\""]
N387 [label="k: 0000000000000084, b: 64, v: \"\"84\"\""]
N388 [label="k: 0000000000000085, b: 64, v: \"\"85\"\""]
N389 [label="k: 0000000000000086, b: 64, v: \"\"86\"\""]
N390 [label="k: 0000000000000087, b: 64, v: \"\"87\"\""]
N391 [label="k: 0000000000000088, b: 64, v: \"\"88\"\""]
N392 [label="k: 0000000000000089, b: 64, v: \"\"89\"\""]
N393 [label="k: 000000000000008a, b: 64, v: \"\"8a\"\""]
N394 [label="k: 000000000000008b, b: 64, v: \"\"8b\"\""]
N395 [label="k: 000000000000008c, b: 64, v: \"\"8c\"\""]
N396 [label="k: 000000000000008d, b: 64, v: \"\"8d\"\""]
N397 [label="k: 000000000000008e, b: 64, v: \"\"8e\"\""]
N398 [label="k: 000000000000008f, b: 64, v: \"\"8f\"\""]
N399 [label="k: 0000000000000090, b: 64, v: \"\"90\"\""]
N400 [label="k: 0000000000000091, b: 64, v: \"\"91\"\""]
N401 [label="k: 0000000000000092, b: 64, v: \"\"92\"\""]
N402 [label="k: 0000000000000093, b: 64, v: \"\"93\"\""]
N403 [label="k: 0000000000000094, b: 64, v: \"\"94\"\""]
N404 [label="k: 0000000000000095, b: 64, v: \"\"95\"\""]
N405 [label="k: 0000000000000096, b: 64, v: \"\"96\"\""]
N406 [label="k: 0000000000000097, b: 64, v: \"\"97\"\""]
N407 [label="k: 0000000000000098, b: 64, v: \"\"98\"\""]
N408 [label="k: 0000000000000099, b: 64, v: \"\"99\"\""]
N409 [label="k: 000000000000009a, b: 64, v: \"\"9a\"\""]
N410 [label="k: 000000000000009b, b: 64, v: \"\"9b\"\""]
N411 [label="k: 000000000000009c, b: 64, v: \"\"9c\"\""]
N412 [label="k: 000000000000009d, b: 64, v: \"\"9d\"\""]
N413 [label="k: 000000000000009e, b: 64, v: \"\"9e\"\""]
N414 [label="k: 000000000000009f, b: 64, v: \"\"9f\"\""]
N415 [label="k: 00000000000000a0, b: 64, v: \"\"a0\"\""]
N416 [label="k: 00000000000000a1, b: 64, v: \"\"a1\"\""]
N417 [label="k: 00000000000000a2, b: 64, v: \"\"a2\"\""]
N418 [label="k: 00000000000000a3, b: 64, v: \"\"a3\"\""]
N419 [label="k: 00000000000000a4, b: 64, v: \"\"a4\"\""]
N420 [label="k: 00000000000000a5, b: 64, v: \"\"a5\"\""]
N421 [label="k: 00000000000000a6, b: 64, v: \"\"a6\"\""]
N422 [label="k: 00000000000000a7, b: 64, v: \"\"a7\"\""]
N423 [label="k: 00000000000000a8, b: 64, v: \"\"a8\"\""]
N424 [label="k: 00000000000000a9, b: 64, v: \"\"a9\"\""]
N425 [label="k: 00000000000000aa, b: 64, v: \"\"aa\"\""]
N426 [label="k: 00000000000000ab, b: 64, v: \"\"ab\"\""]
N427 [label="k: 00000000000000ac, b: 64, v: \"\"ac\"\""]
N428 [label="k: 00000000000000ad, b: 64, v: \"\"ad\"\""]
N429 [label="k: 00000000000000ae, b: 64, v: \"\"ae\"\""]
N430 [label="k: 00000000000000af, b: 64, v: \"\"af\"\""]
N431 [label="k: 00000000000000b0, b: 64, v: \"\"b0\"\""]
N432 [label="k: 00000000000000b1, b: 64, v: \"\"b1\"\""]
N433 [label="k: 00000000000000b2, b: 64, v: \"\"b2\"\""]
N434 [label="k: 00000000000000b3, b: 64, v: \"\"b3\"\""]
N435 [label="k: 00000000000000b4, b: 64, v: \"\"b4\"\""]
N436 [label="k: 00000000000000b5, b: 64, v: \"\"b5\"\""]
N437 [label="k: 00000000000000b6, b: 64, v: \"\"b6\"\""]
N438 [label="k: 00000000000000b7, b: 64, v: \"\"b7\"\""]
N439 [label="k: 00000000000000b8, b: 64, v: \"\"b8\"\""]
N440 [label="k: 00000000000000b9, b: 64, v: \"\"b9\"\""]
N441 [label="k: 00000000000000ba, b: 64, v: \"\"ba\"\""]
N442 [label="k: 00000000000000bb, b: 64, v: \"\"bb\"\""]
N443 [label="k: 00000000000000bc, b: 64, v: \"\"bc\"\""]
N444 [label="k: 00000000000000bd, b: 64, v: \"\"bd\"\""]
N445 [label="k: 00000000000000be, b: 64, v: \"\"be\"\""]
N446 [label="k: 00000000000000bf, b: 64, v: \"\"bf\"\""]
N447 [label="k: 00000000000000c0, b: 64, v: \"\"c0\"\""]
N448 [label="k: 00000000000000c1, b: 64, v: \"\"c1\"\""]
N449 [label="k: 00000000000000c2, b: 64, v: \"\"c2\"\""]
N450 [label="k: 00000000000000c3, b: 64, v: \"\"c3\"\""]
N451 [label="k: 00000000000000c4, b: 64, v: \"\"c4\"\""]
N452 [label="k: 00000000000000c5, b: 64, v: \"\"c5\"\""]
N453 [label="k: 00000000000000c6, b: 64, v: \"\"c6\"\""]
N454 [label="k: 00000000000000c7, b: 64, v: \"\"c7\"\""]
N455 [label="k: 00000000000000c8, b: 64, v: \"\"c8\"\""]
N456 [label="k: 00000000000000c9, b: 64, v: \"\"c9\"\""]
N457 [label="k: 00000000000000ca, b: 64, v: \"\"ca\"\""]
N458 [label="k: 00000000000000cb, b: 64, v: \"\"cb\"\""]
N459 [label="k: 00000000000000cc, b: 64, v: \"\"cc\"\""]
N460 [label="k: 00000000000000cd, b: 64, v: \"\"cd\"\""]
N461 [label="k: 00000000000000ce, b: 64, v: \"\"ce\"\""]
N462 [label="k: 00000000000000cf, b: 64, v: \"\"cf\"\""]
N463 [label="k: 00000000000000d0, b: 64, v: \"\"d0\"\""]
N464 [label="k: 00000000000000d1, b: 64, v: \"\"d1\"\""]
N465 [label="k: 00000000000000d2, b: 64, v: \"\"d2\"\""]
N466 [label="k: 00000000000000d3, b: 64, v: \"\"d3\"\""]
N467 [label="k: 00000000000000d4, b: 64, v: \"\"d4\"\""]
N468 [label="k: 00000000000000d5, b: 64, v: \"\"d5\"\""]
N469 [label="k: 00000000000000d6, b: 64, v: \"\"d6\"\""]
N470 [label="k: 00000000000000d7, b: 64, v: \"\"d7\"\""]
N471 [label="k: 00000000000000d8, b: 64, v: \"\"d8\"\""]
N472 [label="k: 00000000000000d9, b: 64, v: \"\"d9\"\""]
N473 [label="k: 00000000000000da, b: 64, v: \"\"da\"\""]
N474 [label="k: 00000000000000db, b: 64, v: \"\"db\"\""]
N475 [label="k: 00000000000000dc, b: 64, v: \"\"dc\"\""]
N476 [label="k: 00000000000000dd, b: 64, v: \"\"dd\"\""]
N477 [label="k: 00000000000000de, b: 64, v: \"\"de\"\""]
N478 [label="k: 00000000000000df, b: 64, v: \"\"df\"\""]
N479 [label="k: 00000000000000e0, b: 64, v: \"\"e0\"\""]
N480 [label="k: 00000000000000e1, b: 64, v: \"\"e1\"\""]
N481 [label="k: 00000000000000e2, b: 64, v: \"\"e2\"\""]
N482 [label="k: 00000000000000e3, b: 64, v: \"\"e3\"\""]
N483 [label="k: 00000000000000e4, b: 64, v: \"\"e4\"\""]
N484 [label="k: 00000000000000e5, b: 64, v: \"\"e5\"\""]
N485 [label="k: 00000000000000e6, b: 64, v: \"\"e6\"\""]
N486 [label="k: 00000000000000e7, b: 64, v: \"\"e7\"\""]
N487 [label="k: 00000000000000e8, b: 64, v: \"\"e8\"\""]
N488 [label="k: 00000000000000e9, b: 64, v: \"\"e9\"\""]
N489 [label="k: 00000000000000ea, b: 64, v: \"\"ea\"\""]
N490 [label="k: 00000000000000eb, b: 64, v: \"\"eb\"\""]
N491 [label="k: 00000000000000ec, b: 64, v: \"\"ec\"\""]
N492 [label="k: 00000000000000ed, b: 64, v: \"\"ed\"\""]
N493 [label="k: 00000000000000ee, b: 64, v: \"\"ee\"\""]
N494 [label="k: 00000000000000ef, b: 64, v: \"\"ef\"\""]
N495 [label="k: 00000000000000f0, b: 64, v: \"\"f0\"\""]
N496 [label="k: 00000000000000f1, b: 64, v: \"\"f1\"\""]
N497 [label="k: 00000000000000f2, b: 64, v: \"\"f2\"\""]
N498 [label="k: 00000000000000f3, b: 64, v: \"\"f3\"\""]
N499 [label="k: 00000000000000f4, b: 64, v: \"\"f4\"\""]
N500 [label="k: 00000000000000f5, b: 64, v: \"\"f5\"\""]
N501 [label="k: 00000000000000f6, b: 64, v: \"\"f6\"\""]
N502 [label="k: 00000000000000f7, b: 64, v: \"\"f7\"\""]
N503 [label="k: 00000000000000f8, b: 64, v: \"\"f8\"\""]
N504 [label="k: 00000000000000f9, b: 64, v: \"\"f9\"\""]
N505 [label="k: 00000000000000fa, b: 64, v: \"\"fa\"\""]
N506 [label="k: 00000000000000fb, b: 64, v: \"\"fb\"\""]
N507 [label="k: 00000000000000fc, b: 64, v: \"\"fc\"\""]
N508 [label="k: 00000000000000fd, b: 64, v: \"\"fd\"\""]
N509 [label="k: 00000000000000fe, b: 64, v: \"\"fe\"\""]
N510 [label="k: 00000000000000ff, b: 64, v: \"\"ff\"\""]
}
`

	TestTree64BigTreeInvertedInsertions = `digraph d {
N0 [label="k: 0000000000000000, b: 0"]
N0 -> { N1 N2 }
N1 [label="k: 0000000000000000, b: 1"]
N1 -> { N3 N4 }
N2 [label="k: 8000000000000000, b: 1"]
N2 -> { N5 N6 }
N3 [label="k: 0000000000000000, b: 2"]
N3 -> { N7 N8 }
N4 [label="k: 4000000000000000, b: 2"]
N4 -> { N9 N10 }
N5 [label="k: 8000000000000000, b: 2"]
N5 -> { N11 N12 }
N6 [label="k: c000000000000000, b: 2"]
N6 -> { N13 N14 }
N7 [label="k: 0000000000000000, b: 3"]
N7 -> { N15 N16 }
N8 [label="k: 2000000000000000, b: 3"]
N8 -> { N17 N18 }
N9 [label="k: 4000000000000000, b: 3"]
N9 -> { N19 N20 }
N10 [label="k: 6000000000000000, b: 3"]
N10 -> { N21 N22 }
N11 [label="k: 8000000000000000, b: 3"]
N11 -> { N23 N24 }
N12 [label="k: a000000000000000, b: 3"]
N12 -> { N25 N26 }
N13 [label="k: c000000000000000, b: 3"]
N13 -> { N27 N28 }
N14 [label="k: e000000000000000, b: 3"]
N14 -> { N29 N30 }
N15 [label="k: 0000000000000000, b: 4"]
N15 -> { N31 N32 }
N16 [label="k: 1000000000000000, b: 4"]
N16 -> { N33 N34 }
N17 [label="k: 2000000000000000, b: 4"]
N17 -> { N35 N36 }
N18 [label="k: 3000000000000000, b: 4"]
N18 -> { N37 N38 }
N19 [label="k: 4000000000000000, b: 4"]
N19 -> { N39 N40 }
N20 [label="k: 5000000000000000, b: 4"]
N20 -> { N41 N42 }
N21 [label="k: 6000000000000000, b: 4"]
N21 -> { N43 N44 }
N22 [label="k: 7000000000000000, b: 4"]
N22 -> { N45 N46 }
N23 [label="k: 8000000000000000, b: 4"]
N23 -> { N47 N48 }
N24 [label="k: 9000000000000000, b: 4"]
N24 -> { N49 N50 }
N25 [label="k: a000000000000000, b: 4"]
N25 -> { N51 N52 }
N26 [label="k: b000000000000000, b: 4"]
N26 -> { N53 N54 }
N27 [label="k: c000000000000000, b: 4"]
N27 -> { N55 N56 }
N28 [label="k: d000000000000000, b: 4"]
N28 -> { N57 N58 }
N29 [label="k: e000000000000000, b: 4"]
N29 -> { N59 N60 }
N30 [label="k: f000000000000000, b: 4"]
N30 -> { N61 N62 }
N31 [label="k: 0000000000000000, b: 5"]
N31 -> { N63 N64 }
N32 [label="k: 0800000000000000, b: 5"]
N32 -> { N65 N66 }
N33 [label="k: 1000000000000000, b: 5"]
N33 -> { N67 N68 }
N34 [label="k: 1800000000000000, b: 5"]
N34 -> { N69 N70 }
N35 [label="k: 2000000000000000, b: 5"]
N35 -> { N71 N72 }
N36 [label="k: 2800000000000000, b: 5"]
N36 -> { N73 N74 }
N37 [label="k: 3000000000000000, b: 5"]
N37 -> { N75 N76 }
N38 [label="k: 3800000000000000, b: 5"]
N38 -> { N77 N78 }
N39 [label="k: 4000000000000000, b: 5"]
N39 -> { N79 N80 }
N40 [label="k: 4800000000000000, b: 5"]
N40 -> { N81 N82 }
N41 [label="k: 5000000000000000, b: 5"]
N41 -> { N83 N84 }
N42 [label="k: 5800000000000000, b: 5"]
N42 -> { N85 N86 }
N43 [label="k: 6000000000000000, b: 5"]
N43 -> { N87 N88 }
N44 [label="k: 6800000000000000, b: 5"]
N44 -> { N89 N90 }
N45 [label="k: 7000000000000000, b: 5"]
N45 -> { N91 N92 }
N46 [label="k: 7800000000000000, b: 5"]
N46 -> { N93 N94 }
N47 [label="k: 8000000000000000, b: 5"]
N47 -> { N95 N96 }
N48 [label="k: 8800000000000000, b: 5"]
N48 -> { N97 N98 }
N49 [label="k: 9000000000000000, b: 5"]
N49 -> { N99 N100 }
N50 [label="k: 9800000000000000, b: 5"]
N50 -> { N101 N102 }
N51 [label="k: a000000000000000, b: 5"]
N51 -> { N103 N104 }
N52 [label="k: a800000000000000, b: 5"]
N52 -> { N105 N106 }
N53 [label="k: b000000000000000, b: 5"]
N53 -> { N107 N108 }
N54 [label="k: b800000000000000, b: 5"]
N54 -> { N109 N110 }
N55 [label="k: c000000000000000, b: 5"]
N55 -> { N111 N112 }
N56 [label="k: c800000000000000, b: 5"]
N56 -> { N113 N114 }
N57 [label="k: d000000000000000, b: 5"]
N57 -> { N115 N116 }
N58 [label="k: d800000000000000, b: 5"]
N58 -> { N117 N118 }
N59 [label="k: e000000000000000, b: 5"]
N59 -> { N119 N120 }
N60 [label="k: e800000000000000, b: 5"]
N60 -> { N121 N122 }
N61 [label="k: f000000000000000, b: 5"]
N61 -> { N123 N124 }
N62 [label="k: f800000000000000, b: 5"]
N62 -> { N125 N126 }
N63 [label="k: 0000000000000000, b: 6"]
N63 -> { N127 N128 }
N64 [label="k: 0400000000000000, b: 6"]
N64 -> { N129 N130 }
N65 [label="k: 0800000000000000, b: 6"]
N65 -> { N131 N132 }
N66 [label="k: 0c00000000000000, b: 6"]
N66 -> { N133 N134 }
N67 [label="k: 1000000000000000, b: 6"]
N67 -> { N135 N136 }
N68 [label="k: 1400000000000000, b: 6"]
N68 -> { N137 N138 }
N69 [label="k: 1800000000000000, b: 6"]
N69 -> { N139 N140 }
N70 [label="k: 1c00000000000000, b: 6"]
N70 -> { N141 N142 }
N71 [label="k: 2000000000000000, b: 6"]
N71 -> { N143 N144 }
N72 [label="k: 2400000000000000, b: 6"]
N72 -> { N145 N146 }
N73 [label="k: 2800000000000000, b: 6"]
N73 -> { N147 N148 }
N74 [label="k: 2c00000000000000, b: 6"]
N74 -> { N149 N150 }
N75 [label="k: 3000000000000000, b: 6"]
N75 -> { N151 N152 }
N76 [label="k: 3400000000000000, b: 6"]
N76 -> { N153 N154 }
N77 [label="k: 3800000000000000, b: 6"]
N77 -> { N155 N156 }
N78 [label="k: 3c00000000000000, b: 6"]
N78 -> { N157 N158 }
N79 [label="k: 4000000000000000, b: 6"]
N79 -> { N159 N160 }
N80 [label="k: 4400000000000000, b: 6"]
N80 -> { N161 N162 }
N81 [label="k: 4800000000000000, b: 6"]
N81 -> { N163 N164 }
N82 [label="k: 4c00000000000000, b: 6"]
N82 -> { N165 N166 }
N83 [label="k: 5000000000000000, b: 6"]
N83 -> { N167 N168 }
N84 [label="k: 5400000000000000, b: 6"]
N84 -> { N169 N170 }
N85 [label="k: 5800000000000000, b: 6"]
N85 -> { N171 N172 }
N86 [label="k: 5c00000000000000, b: 6"]
N86 -> { N173 N174 }
N87 [label="k: 6000000000000000, b: 6"]
N87 -> { N175 N176 }
N88 [label="k: 6400000000000000, b: 6"]
N88 -> { N177 N178 }
N89 [label="k: 6800000000000000, b: 6"]
N89 -> { N179 N180 }
N90 [label="k: 6c00000000000000, b: 6"]
N90 -> { N181 N182 }
N91 [label="k: 7000000000000000, b: 6"]
N91 -> { N183 N184 }
N92 [label="k: 7400000000000000, b: 6"]
N92 -> { N185 N186 }
N93 [label="k: 7800000000000000, b: 6"]
N93 -> { N187 N188 }
N94 [label="k: 7c00000000000000, b: 6"]
N94 -> { N189 N190 }
N95 [label="k: 8000000000000000, b: 6"]
N95 -> { N191 N192 }
N96 [label="k: 8400000000000000, b: 6"]
N96 -> { N193 N194 }
N97 [label="k: 8800000000000000, b: 6"]
N97 -> { N195 N196 }
N98 [label="k: 8c00000000000000, b: 6"]
N98 -> { N197 N198 }
N99 [label="k: 9000000000000000, b: 6"]
N99 -> { N199 N200 }
N100 [label="k: 9400000000000000, b: 6"]
N100 -> { N201 N202 }
N101 [label="k: 9800000000000000, b: 6"]
N101 -> { N203 N204 }
N102 [label="k: 9c00000000000000, b: 6"]
N102 -> { N205 N206 }
N103 [label="k: a000000000000000, b: 6"]
N103 -> { N207 N208 }
N104 [label="k: a400000000000000, b: 6"]
N104 -> { N209 N210 }
N105 [label="k: a800000000000000, b: 6"]
N105 -> { N211 N212 }
N106 [label="k: ac00000000000000, b: 6"]
N106 -> { N213 N214 }
N107 [label="k: b000000000000000, b: 6"]
N107 -> { N215 N216 }
N108 [label="k: b400000000000000, b: 6"]
N108 -> { N217 N218 }
N109 [label="k: b800000000000000, b: 6"]
N109 -> { N219 N220 }
N110 [label="k: bc00000000000000, b: 6"]
N110 -> { N221 N222 }
N111 [label="k: c000000000000000, b: 6"]
N111 -> { N223 N224 }
N112 [label="k: c400000000000000, b: 6"]
N112 -> { N225 N226 }
N113 [label="k: c800000000000000, b: 6"]
N113 -> { N227 N228 }
N114 [label="k: cc00000000000000, b: 6"]
N114 -> { N229 N230 }
N115 [label="k: d000000000000000, b: 6"]
N115 -> { N231 N232 }
N116 [label="k: d400000000000000, b: 6"]
N116 -> { N233 N234 }
N117 [label="k: d800000000000000, b: 6"]
N117 -> { N235 N236 }
N118 [label="k: dc00000000000000, b: 6"]
N118 -> { N237 N238 }
N119 [label="k: e000000000000000, b: 6"]
N119 -> { N239 N240 }
N120 [label="k: e400000000000000, b: 6"]
N120 -> { N241 N242 }
N121 [label="k: e800000000000000, b: 6"]
N121 -> { N243 N244 }
N122 [label="k: ec00000000000000, b: 6"]
N122 -> { N245 N246 }
N123 [label="k: f000000000000000, b: 6"]
N123 -> { N247 N248 }
N124 [label="k: f400000000000000, b: 6"]
N124 -> { N249 N250 }
N125 [label="k: f800000000000000, b: 6"]
N125 -> { N251 N252 }
N126 [label="k: fc00000000000000, b: 6"]
N126 -> { N253 N254 }
N127 [label="k: 0000000000000000, b: 7"]
N127 -> { N255 N256 }
N128 [label="k: 0200000000000000, b: 7"]
N128 -> { N257 N258 }
N129 [label="k: 0400000000000000, b: 7"]
N129 -> { N259 N260 }
N130 [label="k: 0600000000000000, b: 7"]
N130 -> { N261 N262 }
N131 [label="k: 0800000000000000, b: 7"]
N131 -> { N263 N264 }
N132 [label="k: 0a00000000000000, b: 7"]
N132 -> { N265 N266 }
N133 [label="k: 0c00000000000000, b: 7"]
N133 -> { N267 N268 }
N134 [label="k: 0e00000000000000, b: 7"]
N134 -> { N269 N270 }
N135 [label="k: 1000000000000000, b: 7"]
N135 -> { N271 N272 }
N136 [label="k: 1200000000000000, b: 7"]
N136 -> { N273 N274 }
N137 [label="k: 1400000000000000, b: 7"]
N137 -> { N275 N276 }
N138 [label="k: 1600000000000000, b: 7"]
N138 -> { N277 N278 }
N139 [label="k: 1800000000000000, b: 7"]
N139 -> { N279 N280 }
N140 [label="k: 1a00000000000000, b: 7"]
N140 -> { N281 N282 }
N141 [label="k: 1c00000000000000, b: 7"]
N141 -> { N283 N284 }
N142 [label="k: 1e00000000000000, b: 7"]
N142 -> { N285 N286 }
N143 [label="k: 2000000000000000, b: 7"]
N143 -> { N287 N288 }
N144 [label="k: 2200000000000000, b: 7"]
N144 -> { N289 N290 }
N145 [label="k: 2400000000000000, b: 7"]
N145 -> { N291 N292 }
N146 [label="k: 2600000000000000, b: 7"]
N146 -> { N293 N294 }
N147 [label="k: 2800000000000000, b: 7"]
N147 -> { N295 N296 }
N148 [label="k: 2a00000000000000, b: 7"]
N148 -> { N297 N298 }
N149 [label="k: 2c00000000000000, b: 7"]
N149 -> { N299 N300 }
N150 [label="k: 2e00000000000000, b: 7"]
N150 -> { N301 N302 }
N151 [label="k: 3000000000000000, b: 7"]
N151 -> { N303 N304 }
N152 [label="k: 3200000000000000, b: 7"]
N152 -> { N305 N306 }
N153 [label="k: 3400000000000000, b: 7"]
N153 -> { N307 N308 }
N154 [label="k: 3600000000000000, b: 7"]
N154 -> { N309 N310 }
N155 [label="k: 3800000000000000, b: 7"]
N155 -> { N311 N312 }
N156 [label="k: 3a00000000000000, b: 7"]
N156 -> { N313 N314 }
N157 [label="k: 3c00000000000000, b: 7"]
N157 -> { N315 N316 }
N158 [label="k: 3e00000000000000, b: 7"]
N158 -> { N317 N318 }
N159 [label="k: 4000000000000000, b: 7"]
N159 -> { N319 N320 }
N160 [label="k: 4200000000000000, b: 7"]
N160 -> { N321 N322 }
N161 [label="k: 4400000000000000, b: 7"]
N161 -> { N323 N324 }
N162 [label="k: 4600000000000000, b: 7"]
N162 -> { N325 N326 }
N163 [label="k: 4800000000000000, b: 7"]
N163 -> { N327 N328 }
N164 [label="k: 4a00000000000000, b: 7"]
N164 -> { N329 N330 }
N165 [label="k: 4c00000000000000, b: 7"]
N165 -> { N331 N332 }
N166 [label="k: 4e00000000000000, b: 7"]
N166 -> { N333 N334 }
N167 [label="k: 5000000000000000, b: 7"]
N167 -> { N335 N336 }
N168 [label="k: 5200000000000000, b: 7"]
N168 -> { N337 N338 }
N169 [label="k: 5400000000000000, b: 7"]
N169 -> { N339 N340 }
N170 [label="k: 5600000000000000, b: 7"]
N170 -> { N341 N342 }
N171 [label="k: 5800000000000000, b: 7"]
N171 -> { N343 N344 }
N172 [label="k: 5a00000000000000, b: 7"]
N172 -> { N345 N346 }
N173 [label="k: 5c00000000000000, b: 7"]
N173 -> { N347 N348 }
N174 [label="k: 5e00000000000000, b: 7"]
N174 -> { N349 N350 }
N175 [label="k: 6000000000000000, b: 7"]
N175 -> { N351 N352 }
N176 [label="k: 6200000000000000, b: 7"]
N176 -> { N353 N354 }
N177 [label="k: 6400000000000000, b: 7"]
N177 -> { N355 N356 }
N178 [label="k: 6600000000000000, b: 7"]
N178 -> { N357 N358 }
N179 [label="k: 6800000000000000, b: 7"]
N179 -> { N359 N360 }
N180 [label="k: 6a00000000000000, b: 7"]
N180 -> { N361 N362 }
N181 [label="k: 6c00000000000000, b: 7"]
N181 -> { N363 N364 }
N182 [label="k: 6e00000000000000, b: 7"]
N182 -> { N365 N366 }
N183 [label="k: 7000000000000000, b: 7"]
N183 -> { N367 N368 }
N184 [label="k: 7200000000000000, b: 7"]
N184 -> { N369 N370 }
N185 [label="k: 7400000000000000, b: 7"]
N185 -> { N371 N372 }
N186 [label="k: 7600000000000000, b: 7"]
N186 -> { N373 N374 }
N187 [label="k: 7800000000000000, b: 7"]
N187 -> { N375 N376 }
N188 [label="k: 7a00000000000000, b: 7"]
N188 -> { N377 N378 }
N189 [label="k: 7c00000000000000, b: 7"]
N189 -> { N379 N380 }
N190 [label="k: 7e00000000000000, b: 7"]
N190 -> { N381 N382 }
N191 [label="k: 8000000000000000, b: 7"]
N191 -> { N383 N384 }
N192 [label="k: 8200000000000000, b: 7"]
N192 -> { N385 N386 }
N193 [label="k: 8400000000000000, b: 7"]
N193 -> { N387 N388 }
N194 [label="k: 8600000000000000, b: 7"]
N194 -> { N389 N390 }
N195 [label="k: 8800000000000000, b: 7"]
N195 -> { N391 N392 }
N196 [label="k: 8a00000000000000, b: 7"]
N196 -> { N393 N394 }
N197 [label="k: 8c00000000000000, b: 7"]
N197 -> { N395 N396 }
N198 [label="k: 8e00000000000000, b: 7"]
N198 -> { N397 N398 }
N199 [label="k: 9000000000000000, b: 7"]
N199 -> { N399 N400 }
N200 [label="k: 9200000000000000, b: 7"]
N200 -> { N401 N402 }
N201 [label="k: 9400000000000000, b: 7"]
N201 -> { N403 N404 }
N202 [label="k: 9600000000000000, b: 7"]
N202 -> { N405 N406 }
N203 [label="k: 9800000000000000, b: 7"]
N203 -> { N407 N408 }
N204 [label="k: 9a00000000000000, b: 7"]
N204 -> { N409 N410 }
N205 [label="k: 9c00000000000000, b: 7"]
N205 -> { N411 N412 }
N206 [label="k: 9e00000000000000, b: 7"]
N206 -> { N413 N414 }
N207 [label="k: a000000000000000, b: 7"]
N207 -> { N415 N416 }
N208 [label="k: a200000000000000, b: 7"]
N208 -> { N417 N418 }
N209 [label="k: a400000000000000, b: 7"]
N209 -> { N419 N420 }
N210 [label="k: a600000000000000, b: 7"]
N210 -> { N421 N422 }
N211 [label="k: a800000000000000, b: 7"]
N211 -> { N423 N424 }
N212 [label="k: aa00000000000000, b: 7"]
N212 -> { N425 N426 }
N213 [label="k: ac00000000000000, b: 7"]
N213 -> { N427 N428 }
N214 [label="k: ae00000000000000, b: 7"]
N214 -> { N429 N430 }
N215 [label="k: b000000000000000, b: 7"]
N215 -> { N431 N432 }
N216 [label="k: b200000000000000, b: 7"]
N216 -> { N433 N434 }
N217 [label="k: b400000000000000, b: 7"]
N217 -> { N435 N436 }
N218 [label="k: b600000000000000, b: 7"]
N218 -> { N437 N438 }
N219 [label="k: b800000000000000, b: 7"]
N219 -> { N439 N440 }
N220 [label="k: ba00000000000000, b: 7"]
N220 -> { N441 N442 }
N221 [label="k: bc00000000000000, b: 7"]
N221 -> { N443 N444 }
N222 [label="k: be00000000000000, b: 7"]
N222 -> { N445 N446 }
N223 [label="k: c000000000000000, b: 7"]
N223 -> { N447 N448 }
N224 [label="k: c200000000000000, b: 7"]
N224 -> { N449 N450 }
N225 [label="k: c400000000000000, b: 7"]
N225 -> { N451 N452 }
N226 [label="k: c600000000000000, b: 7"]
N226 -> { N453 N454 }
N227 [label="k: c800000000000000, b: 7"]
N227 -> { N455 N456 }
N228 [label="k: ca00000000000000, b: 7"]
N228 -> { N457 N458 }
N229 [label="k: cc00000000000000, b: 7"]
N229 -> { N459 N460 }
N230 [label="k: ce00000000000000, b: 7"]
N230 -> { N461 N462 }
N231 [label="k: d000000000000000, b: 7"]
N231 -> { N463 N464 }
N232 [label="k: d200000000000000, b: 7"]
N232 -> { N465 N466 }
N233 [label="k: d400000000000000, b: 7"]
N233 -> { N467 N468 }
N234 [label="k: d600000000000000, b: 7"]
N234 -> { N469 N470 }
N235 [label="k: d800000000000000, b: 7"]
N235 -> { N471 N472 }
N236 [label="k: da00000000000000, b: 7"]
N236 -> { N473 N474 }
N237 [label="k: dc00000000000000, b: 7"]
N237 -> { N475 N476 }
N238 [label="k: de00000000000000, b: 7"]
N238 -> { N477 N478 }
N239 [label="k: e000000000000000, b: 7"]
N239 -> { N479 N480 }
N240 [label="k: e200000000000000, b: 7"]
N240 -> { N481 N482 }
N241 [label="k: e400000000000000, b: 7"]
N241 -> { N483 N484 }
N242 [label="k: e600000000000000, b: 7"]
N242 -> { N485 N486 }
N243 [label="k: e800000000000000, b: 7"]
N243 -> { N487 N488 }
N244 [label="k: ea00000000000000, b: 7"]
N244 -> { N489 N490 }
N245 [label="k: ec00000000000000, b: 7"]
N245 -> { N491 N492 }
N246 [label="k: ee00000000000000, b: 7"]
N246 -> { N493 N494 }
N247 [label="k: f000000000000000, b: 7"]
N247 -> { N495 N496 }
N248 [label="k: f200000000000000, b: 7"]
N248 -> { N497 N498 }
N249 [label="k: f400000000000000, b: 7"]
N249 -> { N499 N500 }
N250 [label="k: f600000000000000, b: 7"]
N250 -> { N501 N502 }
N251 [label="k: f800000000000000, b: 7"]
N251 -> { N503 N504 }
N252 [label="k: fa00000000000000, b: 7"]
N252 -> { N505 N506 }
N253 [label="k: fc00000000000000, b: 7"]
N253 -> { N507 N508 }
N254 [label="k: fe00000000000000, b: 7"]
N254 -> { N509 N510 }
N255 [label="k: 0000000000000000, b: 64, v: \"\"00\"\""]
N256 [label="k: 0100000000000000, b: 64, v: \"\"01\"\""]
N257 [label="k: 0200000000000000, b: 64, v: \"\"02\"\""]
N258 [label="k: 0300000000000000, b: 64, v: \"\"03\"\""]
N259 [label="k: 0400000000000000, b: 64, v: \"\"04\"\""]
N260 [label="k: 0500000000000000, b: 64, v: \"\"05\"\""]
N261 [label="k: 0600000000000000, b: 64, v: \"\"06\"\""]
N262 [label="k: 0700000000000000, b: 64, v: \"\"07\"\""]
N263 [label="k: 0800000000000000, b: 64, v: \"\"08\"\""]
N264 [label="k: 0900000000000000, b: 64, v: \"\"09\"\""]
N265 [label="k: 0a00000000000000, b: 64, v: \"\"0a\"\""]
N266 [label="k: 0b00000000000000, b: 64, v: \"\"0b\"\""]
N267 [label="k: 0c00000000000000, b: 64, v: \"\"0c\"\""]
N268 [label="k: 0d00000000000000, b: 64, v: \"\"0d\"\""]
N269 [label="k: 0e00000000000000, b: 64, v: \"\"0e\"\""]
N270 [label="k: 0f00000000000000, b: 64, v: \"\"0f\"\""]
N271 [label="k: 1000000000000000, b: 64, v: \"\"10\"\""]
N272 [label="k: 1100000000000000, b: 64, v: \"\"11\"\""]
N273 [label="k: 1200000000000000, b: 64, v: \"\"12\"\""]
N274 [label="k: 1300000000000000, b: 64, v: \"\"13\"\""]
N275 [label="k: 1400000000000000, b: 64, v: \"\"14\"\""]
N276 [label="k: 1500000000000000, b: 64, v: \"\"15\"\""]
N277 [label="k: 1600000000000000, b: 64, v: \"\"16\"\""]
N278 [label="k: 1700000000000000, b: 64, v: \"\"17\"\""]
N279 [label="k: 1800000000000000, b: 64, v: \"\"18\"\""]
N280 [label="k: 1900000000000000, b: 64, v: \"\"19\"\""]
N281 [label="k: 1a00000000000000, b: 64, v: \"\"1a\"\""]
N282 [label="k: 1b00000000000000, b: 64, v: \"\"1b\"\""]
N283 [label="k: 1c00000000000000, b: 64, v: \"\"1c\"\""]
N284 [label="k: 1d00000000000000, b: 64, v: \"\"1d\"\""]
N285 [label="k: 1e00000000000000, b: 64, v: \"\"1e\"\""]
N286 [label="k: 1f00000000000000, b: 64, v: \"\"1f\"\""]
N287 [label="k: 2000000000000000, b: 64, v: \"\"20\"\""]
N288 [label="k: 2100000000000000, b: 64, v: \"\"21\"\""]
N289 [label="k: 2200000000000000, b: 64, v: \"\"22\"\""]
N290 [label="k: 2300000000000000, b: 64, v: \"\"23\"\""]
N291 [label="k: 2400000000000000, b: 64, v: \"\"24\"\""]
N292 [label="k: 2500000000000000, b: 64, v: \"\"25\"\""]
N293 [label="k: 2600000000000000, b: 64, v: \"\"26\"\""]
N294 [label="k: 2700000000000000, b: 64, v: \"\"27\"\""]
N295 [label="k: 2800000000000000, b: 64, v: \"\"28\"\""]
N296 [label="k: 2900000000000000, b: 64, v: \"\"29\"\""]
N297 [label="k: 2a00000000000000, b: 64, v: \"\"2a\"\""]
N298 [label="k: 2b00000000000000, b: 64, v: \"\"2b\"\""]
N299 [label="k: 2c00000000000000, b: 64, v: \"\"2c\"\""]
N300 [label="k: 2d00000000000000, b: 64, v: \"\"2d\"\""]
N301 [label="k: 2e00000000000000, b: 64, v: \"\"2e\"\""]
N302 [label="k: 2f00000000000000, b: 64, v: \"\"2f\"\""]
N303 [label="k: 3000000000000000, b: 64, v: \"\"30\"\""]
N304 [label="k: 3100000000000000, b: 64, v: \"\"31\"\""]
N305 [label="k: 3200000000000000, b: 64, v: \"\"32\"\""]
N306 [label="k: 3300000000000000, b: 64, v: \"\"33\"\""]
N307 [label="k: 3400000000000000, b: 64, v: \"\"34\"\""]
N308 [label="k: 3500000000000000, b: 64, v: \"\"35\"\""]
N309 [label="k: 3600000000000000, b: 64, v: \"\"36\"\""]
N310 [label="k: 3700000000000000, b: 64, v: \"\"37\"\""]
N311 [label="k: 3800000000000000, b: 64, v: \"\"38\"\""]
N312 [label="k: 3900000000000000, b: 64, v: \"\"39\"\""]
N313 [label="k: 3a00000000000000, b: 64, v: \"\"3a\"\""]
N314 [label="k: 3b00000000000000, b: 64, v: \"\"3b\"\""]
N315 [label="k: 3c00000000000000, b: 64, v: \"\"3c\"\""]
N316 [label="k: 3d00000000000000, b: 64, v: \"\"3d\"\""]
N317 [label="k: 3e00000000000000, b: 64, v: \"\"3e\"\""]
N318 [label="k: 3f00000000000000, b: 64, v: \"\"3f\"\""]
N319 [label="k: 4000000000000000, b: 64, v: \"\"40\"\""]
N320 [label="k: 4100000000000000, b: 64, v: \"\"41\"\""]
N321 [label="k: 4200000000000000, b: 64, v: \"\"42\"\""]
N322 [label="k: 4300000000000000, b: 64, v: \"\"43\"\""]
N323 [label="k: 4400000000000000, b: 64, v: \"\"44\"\""]
N324 [label="k: 4500000000000000, b: 64, v: \"\"45\"\""]
N325 [label="k: 4600000000000000, b: 64, v: \"\"46\"\""]
N326 [label="k: 4700000000000000, b: 64, v: \"\"47\"\""]
N327 [label="k: 4800000000000000, b: 64, v: \"\"48\"\""]
N328 [label="k: 4900000000000000, b: 64, v: \"\"49\"\""]
N329 [label="k: 4a00000000000000, b: 64, v: \"\"4a\"\""]
N330 [label="k: 4b00000000000000, b: 64, v: \"\"4b\"\""]
N331 [label="k: 4c00000000000000, b: 64, v: \"\"4c\"\""]
N332 [label="k: 4d00000000000000, b: 64, v: \"\"4d\"\""]
N333 [label="k: 4e00000000000000, b: 64, v: \"\"4e\"\""]
N334 [label="k: 4f00000000000000, b: 64, v: \"\"4f\"\""]
N335 [label="k: 5000000000000000, b: 64, v: \"\"50\"\""]
N336 [label="k: 5100000000000000, b: 64, v: \"\"51\"\""]
N337 [label="k: 5200000000000000, b: 64, v: \"\"52\"\""]
N338 [label="k: 5300000000000000, b: 64, v: \"\"53\"\""]
N339 [label="k: 5400000000000000, b: 64, v: \"\"54\"\""]
N340 [label="k: 5500000000000000, b: 64, v: \"\"55\"\""]
N341 [label="k: 5600000000000000, b: 64, v: \"\"56\"\""]
N342 [label="k: 5700000000000000, b: 64, v: \"\"57\"\""]
N343 [label="k: 5800000000000000, b: 64, v: \"\"58\"\""]
N344 [label="k: 5900000000000000, b: 64, v: \"\"59\"\""]
N345 [label="k: 5a00000000000000, b: 64, v: \"\"5a\"\""]
N346 [label="k: 5b00000000000000, b: 64, v: \"\"5b\"\""]
N347 [label="k: 5c00000000000000, b: 64, v: \"\"5c\"\""]
N348 [label="k: 5d00000000000000, b: 64, v: \"\"5d\"\""]
N349 [label="k: 5e00000000000000, b: 64, v: \"\"5e\"\""]
N350 [label="k: 5f00000000000000, b: 64, v: \"\"5f\"\""]
N351 [label="k: 6000000000000000, b: 64, v: \"\"60\"\""]
N352 [label="k: 6100000000000000, b: 64, v: \"\"61\"\""]
N353 [label="k: 6200000000000000, b: 64, v: \"\"62\"\""]
N354 [label="k: 6300000000000000, b: 64, v: \"\"63\"\""]
N355 [label="k: 6400000000000000, b: 64, v: \"\"64\"\""]
N356 [label="k: 6500000000000000, b: 64, v: \"\"65\"\""]
N357 [label="k: 6600000000000000, b: 64, v: \"\"66\"\""]
N358 [label="k: 6700000000000000, b: 64, v: \"\"67\"\""]
N359 [label="k: 6800000000000000, b: 64, v: \"\"68\"\""]
N360 [label="k: 6900000000000000, b: 64, v: \"\"69\"\""]
N361 [label="k: 6a00000000000000, b: 64, v: \"\"6a\"\""]
N362 [label="k: 6b00000000000000, b: 64, v: \"\"6b\"\""]
N363 [label="k: 6c00000000000000, b: 64, v: \"\"6c\"\""]
N364 [label="k: 6d00000000000000, b: 64, v: \"\"6d\"\""]
N365 [label="k: 6e00000000000000, b: 64, v: \"\"6e\"\""]
N366 [label="k: 6f00000000000000, b: 64, v: \"\"6f\"\""]
N367 [label="k: 7000000000000000, b: 64, v: \"\"70\"\""]
N368 [label="k: 7100000000000000, b: 64, v: \"\"71\"\""]
N369 [label="k: 7200000000000000, b: 64, v: \"\"72\"\""]
N370 [label="k: 7300000000000000, b: 64, v: \"\"73\"\""]
N371 [label="k: 7400000000000000, b: 64, v: \"\"74\"\""]
N372 [label="k: 7500000000000000, b: 64, v: \"\"75\"\""]
N373 [label="k: 7600000000000000, b: 64, v: \"\"76\"\""]
N374 [label="k: 7700000000000000, b: 64, v: \"\"77\"\""]
N375 [label="k: 7800000000000000, b: 64, v: \"\"78\"\""]
N376 [label="k: 7900000000000000, b: 64, v: \"\"79\"\""]
N377 [label="k: 7a00000000000000, b: 64, v: \"\"7a\"\""]
N378 [label="k: 7b00000000000000, b: 64, v: \"\"7b\"\""]
N379 [label="k: 7c00000000000000, b: 64, v: \"\"7c\"\""]
N380 [label="k: 7d00000000000000, b: 64, v: \"\"7d\"\""]
N381 [label="k: 7e00000000000000, b: 64, v: \"\"7e\"\""]
N382 [label="k: 7f00000000000000, b: 64, v: \"\"7f\"\""]
N383 [label="k: 8000000000000000, b: 64, v: \"\"80\"\""]
N384 [label="k: 8100000000000000, b: 64, v: \"\"81\"\""]
N385 [label="k: 8200000000000000, b: 64, v: \"\"82\"\""]
N386 [label="k: 8300000000000000, b: 64, v: \"\"83\"\""]
N387 [label="k: 8400000000000000, b: 64, v: \"\"84\"\""]
N388 [label="k: 8500000000000000, b: 64, v: \"\"85\"\""]
N389 [label="k: 8600000000000000, b: 64, v: \"\"86\"\""]
N390 [label="k: 8700000000000000, b: 64, v: \"\"87\"\""]
N391 [label="k: 8800000000000000, b: 64, v: \"\"88\"\""]
N392 [label="k: 8900000000000000, b: 64, v: \"\"89\"\""]
N393 [label="k: 8a00000000000000, b: 64, v: \"\"8a\"\""]
N394 [label="k: 8b00000000000000, b: 64, v: \"\"8b\"\""]
N395 [label="k: 8c00000000000000, b: 64, v: \"\"8c\"\""]
N396 [label="k: 8d00000000000000, b: 64, v: \"\"8d\"\""]
N397 [label="k: 8e00000000000000, b: 64, v: \"\"8e\"\""]
N398 [label="k: 8f00000000000000, b: 64, v: \"\"8f\"\""]
N399 [label="k: 9000000000000000, b: 64, v: \"\"90\"\""]
N400 [label="k: 9100000000000000, b: 64, v: \"\"91\"\""]
N401 [label="k: 9200000000000000, b: 64, v: \"\"92\"\""]
N402 [label="k: 9300000000000000, b: 64, v: \"\"93\"\""]
N403 [label="k: 9400000000000000, b: 64, v: \"\"94\"\""]
N404 [label="k: 9500000000000000, b: 64, v: \"\"95\"\""]
N405 [label="k: 9600000000000000, b: 64, v: \"\"96\"\""]
N406 [label="k: 9700000000000000, b: 64, v: \"\"97\"\""]
N407 [label="k: 9800000000000000, b: 64, v: \"\"98\"\""]
N408 [label="k: 9900000000000000, b: 64, v: \"\"99\"\""]
N409 [label="k: 9a00000000000000, b: 64, v: \"\"9a\"\""]
N410 [label="k: 9b00000000000000, b: 64, v: \"\"9b\"\""]
N411 [label="k: 9c00000000000000, b: 64, v: \"\"9c\"\""]
N412 [label="k: 9d00000000000000, b: 64, v: \"\"9d\"\""]
N413 [label="k: 9e00000000000000, b: 64, v: \"\"9e\"\""]
N414 [label="k: 9f00000000000000, b: 64, v: \"\"9f\"\""]
N415 [label="k: a000000000000000, b: 64, v: \"\"a0\"\""]
N416 [label="k: a100000000000000, b: 64, v: \"\"a1\"\""]
N417 [label="k: a200000000000000, b: 64, v: \"\"a2\"\""]
N418 [label="k: a300000000000000, b: 64, v: \"\"a3\"\""]
N419 [label="k: a400000000000000, b: 64, v: \"\"a4\"\""]
N420 [label="k: a500000000000000, b: 64, v: \"\"a5\"\""]
N421 [label="k: a600000000000000, b: 64, v: \"\"a6\"\""]
N422 [label="k: a700000000000000, b: 64, v: \"\"a7\"\""]
N423 [label="k: a800000000000000, b: 64, v: \"\"a8\"\""]
N424 [label="k: a900000000000000, b: 64, v: \"\"a9\"\""]
N425 [label="k: aa00000000000000, b: 64, v: \"\"aa\"\""]
N426 [label="k: ab00000000000000, b: 64, v: \"\"ab\"\""]
N427 [label="k: ac00000000000000, b: 64, v: \"\"ac\"\""]
N428 [label="k: ad00000000000000, b: 64, v: \"\"ad\"\""]
N429 [label="k: ae00000000000000, b: 64, v: \"\"ae\"\""]
N430 [label="k: af00000000000000, b: 64, v: \"\"af\"\""]
N431 [label="k: b000000000000000, b: 64, v: \"\"b0\"\""]
N432 [label="k: b100000000000000, b: 64, v: \"\"b1\"\""]
N433 [label="k: b200000000000000, b: 64, v: \"\"b2\"\""]
N434 [label="k: b300000000000000, b: 64, v: \"\"b3\"\""]
N435 [label="k: b400000000000000, b: 64, v: \"\"b4\"\""]
N436 [label="k: b500000000000000, b: 64, v: \"\"b5\"\""]
N437 [label="k: b600000000000000, b: 64, v: \"\"b6\"\""]
N438 [label="k: b700000000000000, b: 64, v: \"\"b7\"\""]
N439 [label="k: b800000000000000, b: 64, v: \"\"b8\"\""]
N440 [label="k: b900000000000000, b: 64, v: \"\"b9\"\""]
N441 [label="k: ba00000000000000, b: 64, v: \"\"ba\"\""]
N442 [label="k: bb00000000000000, b: 64, v: \"\"bb\"\""]
N443 [label="k: bc00000000000000, b: 64, v: \"\"bc\"\""]
N444 [label="k: bd00000000000000, b: 64, v: \"\"bd\"\""]
N445 [label="k: be00000000000000, b: 64, v: \"\"be\"\""]
N446 [label="k: bf00000000000000, b: 64, v: \"\"bf\"\""]
N447 [label="k: c000000000000000, b: 64, v: \"\"c0\"\""]
N448 [label="k: c100000000000000, b: 64, v: \"\"c1\"\""]
N449 [label="k: c200000000000000, b: 64, v: \"\"c2\"\""]
N450 [label="k: c300000000000000, b: 64, v: \"\"c3\"\""]
N451 [label="k: c400000000000000, b: 64, v: \"\"c4\"\""]
N452 [label="k: c500000000000000, b: 64, v: \"\"c5\"\""]
N453 [label="k: c600000000000000, b: 64, v: \"\"c6\"\""]
N454 [label="k: c700000000000000, b: 64, v: \"\"c7\"\""]
N455 [label="k: c800000000000000, b: 64, v: \"\"c8\"\""]
N456 [label="k: c900000000000000, b: 64, v: \"\"c9\"\""]
N457 [label="k: ca00000000000000, b: 64, v: \"\"ca\"\""]
N458 [label="k: cb00000000000000, b: 64, v: \"\"cb\"\""]
N459 [label="k: cc00000000000000, b: 64, v: \"\"cc\"\""]
N460 [label="k: cd00000000000000, b: 64, v: \"\"cd\"\""]
N461 [label="k: ce00000000000000, b: 64, v: \"\"ce\"\""]
N462 [label="k: cf00000000000000, b: 64, v: \"\"cf\"\""]
N463 [label="k: d000000000000000, b: 64, v: \"\"d0\"\""]
N464 [label="k: d100000000000000, b: 64, v: \"\"d1\"\""]
N465 [label="k: d200000000000000, b: 64, v: \"\"d2\"\""]
N466 [label="k: d300000000000000, b: 64, v: \"\"d3\"\""]
N467 [label="k: d400000000000000, b: 64, v: \"\"d4\"\""]
N468 [label="k: d500000000000000, b: 64, v: \"\"d5\"\""]
N469 [label="k: d600000000000000, b: 64, v: \"\"d6\"\""]
N470 [label="k: d700000000000000, b: 64, v: \"\"d7\"\""]
N471 [label="k: d800000000000000, b: 64, v: \"\"d8\"\""]
N472 [label="k: d900000000000000, b: 64, v: \"\"d9\"\""]
N473 [label="k: da00000000000000, b: 64, v: \"\"da\"\""]
N474 [label="k: db00000000000000, b: 64, v: \"\"db\"\""]
N475 [label="k: dc00000000000000, b: 64, v: \"\"dc\"\""]
N476 [label="k: dd00000000000000, b: 64, v: \"\"dd\"\""]
N477 [label="k: de00000000000000, b: 64, v: \"\"de\"\""]
N478 [label="k: df00000000000000, b: 64, v: \"\"df\"\""]
N479 [label="k: e000000000000000, b: 64, v: \"\"e0\"\""]
N480 [label="k: e100000000000000, b: 64, v: \"\"e1\"\""]
N481 [label="k: e200000000000000, b: 64, v: \"\"e2\"\""]
N482 [label="k: e300000000000000, b: 64, v: \"\"e3\"\""]
N483 [label="k: e400000000000000, b: 64, v: \"\"e4\"\""]
N484 [label="k: e500000000000000, b: 64, v: \"\"e5\"\""]
N485 [label="k: e600000000000000, b: 64, v: \"\"e6\"\""]
N486 [label="k: e700000000000000, b: 64, v: \"\"e7\"\""]
N487 [label="k: e800000000000000, b: 64, v: \"\"e8\"\""]
N488 [label="k: e900000000000000, b: 64, v: \"\"e9\"\""]
N489 [label="k: ea00000000000000, b: 64, v: \"\"ea\"\""]
N490 [label="k: eb00000000000000, b: 64, v: \"\"eb\"\""]
N491 [label="k: ec00000000000000, b: 64, v: \"\"ec\"\""]
N492 [label="k: ed00000000000000, b: 64, v: \"\"ed\"\""]
N493 [label="k: ee00000000000000, b: 64, v: \"\"ee\"\""]
N494 [label="k: ef00000000000000, b: 64, v: \"\"ef\"\""]
N495 [label="k: f000000000000000, b: 64, v: \"\"f0\"\""]
N496 [label="k: f100000000000000, b: 64, v: \"\"f1\"\""]
N497 [label="k: f200000000000000, b: 64, v: \"\"f2\"\""]
N498 [label="k: f300000000000000, b: 64, v: \"\"f3\"\""]
N499 [label="k: f400000000000000, b: 64, v: \"\"f4\"\""]
N500 [label="k: f500000000000000, b: 64, v: \"\"f5\"\""]
N501 [label="k: f600000000000000, b: 64, v: \"\"f6\"\""]
N502 [label="k: f700000000000000, b: 64, v: \"\"f7\"\""]
N503 [label="k: f800000000000000, b: 64, v: \"\"f8\"\""]
N504 [label="k: f900000000000000, b: 64, v: \"\"f9\"\""]
N505 [label="k: fa00000000000000, b: 64, v: \"\"fa\"\""]
N506 [label="k: fb00000000000000, b: 64, v: \"\"fb\"\""]
N507 [label="k: fc00000000000000, b: 64, v: \"\"fc\"\""]
N508 [label="k: fd00000000000000, b: 64, v: \"\"fd\"\""]
N509 [label="k: fe00000000000000, b: 64, v: \"\"fe\"\""]
N510 [label="k: ff00000000000000, b: 64, v: \"\"ff\"\""]
}
`

	TestTree64EmptyTree = `digraph d {
N0 [label="nil"]
}
`

	TestTree64WithDeletedChildNode = `digraph d {
N0 [label="k: aaaaaaaa00000000, b: 9, v: \"\"test\"\""]
}
`

	TestTree64WithDeletedChildAndNonLeafNodes = `digraph d {
N0 [label="k: a800000000000000, b: 6"]
N0 -> { N1 N2 }
N1 [label="k: a8aaaaaa00000000, b: 9, v: \"\"L2.1\"\""]
N2 [label="k: aaaaaaaa00000000, b: 7, v: \"\"L1\"\""]
N2 -> { N3 N4 }
N3 [label="k: aaaa000000000000, b: 15"]
N3 -> { N5 N6 }
N4 [label="k: abaaaaaa00000000, b: 9, v: \"\"L2.2\"\""]
N5 [label="k: aaaaaaaa00000000, b: 18, v: \"\"L3\"\""]
N6 [label="k: aaabaaaa00000000, b: 24, v: \"\"L5\"\""]
}
`

	TestTree64WithDeletedTwoChildrenAndNonLeafNodes = `digraph d {
N0 [label="k: a800000000000000, b: 6"]
N0 -> { N1 N2 }
N1 [label="k: a8aaaaaa00000000, b: 9, v: \"\"L2.1\"\""]
N2 [label="k: aaaaaaaa00000000, b: 7, v: \"\"L1\"\""]
N2 -> { N3 N4 }
N3 [label="k: aaabaaaa00000000, b: 24, v: \"\"L5\"\""]
N4 [label="k: abaaaaaa00000000, b: 9, v: \"\"L2.2\"\""]
}
`
)

var inv64 = []uint64{
	0x0, 0x80, 0x40, 0xc0, 0x20, 0xa0, 0x60, 0xe0, 0x10, 0x90, 0x50, 0xd0, 0x30, 0xb0, 0x70, 0xf0,
	0x8, 0x88, 0x48, 0xc8, 0x28, 0xa8, 0x68, 0xe8, 0x18, 0x98, 0x58, 0xd8, 0x38, 0xb8, 0x78, 0xf8,
	0x4, 0x84, 0x44, 0xc4, 0x24, 0xa4, 0x64, 0xe4, 0x14, 0x94, 0x54, 0xd4, 0x34, 0xb4, 0x74, 0xf4,
	0xc, 0x8c, 0x4c, 0xcc, 0x2c, 0xac, 0x6c, 0xec, 0x1c, 0x9c, 0x5c, 0xdc, 0x3c, 0xbc, 0x7c, 0xfc,
	0x2, 0x82, 0x42, 0xc2, 0x22, 0xa2, 0x62, 0xe2, 0x12, 0x92, 0x52, 0xd2, 0x32, 0xb2, 0x72, 0xf2,
	0xa, 0x8a, 0x4a, 0xca, 0x2a, 0xaa, 0x6a, 0xea, 0x1a, 0x9a, 0x5a, 0xda, 0x3a, 0xba, 0x7a, 0xfa,
	0x6, 0x86, 0x46, 0xc6, 0x26, 0xa6, 0x66, 0xe6, 0x16, 0x96, 0x56, 0xd6, 0x36, 0xb6, 0x76, 0xf6,
	0xe, 0x8e, 0x4e, 0xce, 0x2e, 0xae, 0x6e, 0xee, 0x1e, 0x9e, 0x5e, 0xde, 0x3e, 0xbe, 0x7e, 0xfe,
	0x1, 0x81, 0x41, 0xc1, 0x21, 0xa1, 0x61, 0xe1, 0x11, 0x91, 0x51, 0xd1, 0x31, 0xb1, 0x71, 0xf1,
	0x9, 0x89, 0x49, 0xc9, 0x29, 0xa9, 0x69, 0xe9, 0x19, 0x99, 0x59, 0xd9, 0x39, 0xb9, 0x79, 0xf9,
	0x5, 0x85, 0x45, 0xc5, 0x25, 0xa5, 0x65, 0xe5, 0x15, 0x95, 0x55, 0xd5, 0x35, 0xb5, 0x75, 0xf5,
	0xd, 0x8d, 0x4d, 0xcd, 0x2d, 0xad, 0x6d, 0xed, 0x1d, 0x9d, 0x5d, 0xdd, 0x3d, 0xbd, 0x7d, 0xfd,
	0x3, 0x83, 0x43, 0xc3, 0x23, 0xa3, 0x63, 0xe3, 0x13, 0x93, 0x53, 0xd3, 0x33, 0xb3, 0x73, 0xf3,
	0xb, 0x8b, 0x4b, 0xcb, 0x2b, 0xab, 0x6b, 0xeb, 0x1b, 0x9b, 0x5b, 0xdb, 0x3b, 0xbb, 0x7b, 0xfb,
	0x7, 0x87, 0x47, 0xc7, 0x27, 0xa7, 0x67, 0xe7, 0x17, 0x97, 0x57, 0xd7, 0x37, 0xb7, 0x77, 0xf7,
	0xf, 0x8f, 0x4f, 0xcf, 0x2f, 0xaf, 0x6f, 0xef, 0x1f, 0x9f, 0x5f, 0xdf, 0x3f, 0xbf, 0x7f, 0xff}