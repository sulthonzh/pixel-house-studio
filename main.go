package main

import (
	"fmt"
	"strconv"
	"strings"
)

func q1(n int) int {
	temp := 0
	for i := 0; i < n+1; i++ {
		temp += i
	}
	return temp + 1
}

func q2(n int) (result string) {
	resultTemp := []string{}
	temp := 0
	for i := 0; i < n; i++ {
		temp += i
		resultTemp = append(resultTemp, strconv.Itoa(temp+1))
	}
	return strings.Join(resultTemp, "-")

}

type Node struct {
	name   string
	childs []*Node
}

func (parent *Node) Child(depth int) (result string) {
	resultTemp := []string{}
	if depth > 0 {
		resultTemp = append(resultTemp, fmt.Sprintln())
		if parent != nil {
			if parent.childs != nil {
				for _, v := range parent.childs {
					pathString := ""
					for d := 0; d < depth; d++ {
						pathString = pathString + "-"
					}
					pathString = pathString + v.name
					resultTemp = append(resultTemp, pathString)
					res := v.Child(depth + 1)
					resultTemp = append(resultTemp, res)
				}
			}
		}
	} else {
		resultTemp = append(resultTemp, parent.name)
		res := parent.Child(depth + 1)
		resultTemp = append(resultTemp, res)
	}
	return strings.Join(resultTemp, "")
}

func (parent *Node) AddChild(childs ...*Node) {
	parent.childs = childs
}

func (parent *Node) ChildIsExist(childs []string) (ok bool) {
	if childs[0] != parent.name {
		return
	}

	if parent.childs != nil {
		childCount := len(childs)
		haveChild := false
		if childCount > 1 {
			var childTemp *Node
			for _, child := range parent.childs {
				if childs[1] == child.name {
					haveChild = true
					childTemp = child
					break
				}
			}
			if haveChild {
				ok = childTemp.ChildIsExist(childs[1:])
			}
		} else {
			ok = true
		}
	} else {
		ok = true
	}
	return
}

func (parent *Node) FindChild(parentPath []string) (childs []string, err error) {
	if parentPath[0] != parent.name {
		err = fmt.Errorf("Unknown parent")
		return
	}

	if parent.childs != nil {
		parentPathCount := len(parentPath)
		haveChild := false
		if parentPathCount > 1 {
			var childTemp *Node
			for _, child := range parent.childs {
				if parentPath[1] == child.name {
					haveChild = true
					childTemp = child
					break
				}
			}
			if haveChild {
				return childTemp.FindChild(parentPath[1:])
			}
		} else {
			for _, child := range parent.childs {
				childs = append(childs, child.name)
			}
		}
	} else {
		err = fmt.Errorf("Unknown child")
	}
	return
}

func tree() Node {
	b := Node{"B", nil}
	b.AddChild(&Node{"D", nil}, &Node{"E", nil})
	c := Node{"C", nil}
	c.AddChild(&Node{"F", nil}, &Node{"G", nil}, &Node{"H", nil})
	a := Node{"A", nil}
	a.AddChild(&b, &c)
	return a
}

func q3() string {
	a := tree()
	return a.Child(0)
}

func q4(path string) bool {
	pathAray := strings.Split(path, "-")
	tree := tree()
	return tree.ChildIsExist(pathAray)
}

func q5(path string) string {
	pathAray := strings.Split(path, "-")
	tree := tree()
	childs, err := tree.FindChild(pathAray)
	if err != nil {
		return err.Error()
	} else {
		return strings.Join(childs, "-")
	}
}

func main() {
	// q1
	fmt.Println(q1(1))
	fmt.Println(q1(3))
	fmt.Println(q1(7))
	fmt.Println(q1(20))
	fmt.Println(q1(50))

	// q2
	fmt.Println(q2(7))

	// q3
	fmt.Println(q3())

	// q4
	fmt.Println(q4("A-C-F"))
	fmt.Println(q4("A-B-D"))
	fmt.Println(q4("A-B-F"))
	fmt.Println(q4("A"))

	// q5
	fmt.Println(q5("A-C"))
	fmt.Println(q5("A-B"))
	fmt.Println(q5("A-B-D"))
}
