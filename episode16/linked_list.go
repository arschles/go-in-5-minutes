package main

type node struct {
	val  string
	next *node
}

func newLinkedList(vals ...string) *node {
	if len(vals) == 0 {
		return nil
	}
	ret := &node{}
	cur := ret
	for _, val := range vals {
		cur.val = val
		nxt := &node{}
		cur.next = nxt
		cur = nxt
	}
	return ret
}

func (n *node) len() int {
	if n == nil {
		return 0
	}
	i := 1
	cur := n
	for {
		if cur.next == nil {
			break
		}
		i++
		cur = cur.next

	}
	return i
}

func (n *node) find(s string) bool {
	if n == nil {
		return false
	}
	for {
		if n.val == s {
			return true
		}
		if n.next == nil {
			return false
		}
		n = n.next
	}
}

func (n *node) append(s string) {
	last := n
	//find true last node
	for {
		if last.next == nil {
			break
		}
		last = last.next
	}
	last.next = &node{val: s, next: nil}
}
