package main

import "fmt"

func main() {
	nums := [7]int{8, 5, 10, 4, 6, 3, 2}
	fmt.Println(IndexOf(nums[:], 10))

	names := []string{"foo", "bar", "baz"}
	fmt.Println(IndexOf(names, "ss"))

	list()
}

func IndexOf[T comparable](s []T, x T) int {
	for i, y := range s {
		if x == y {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func list() {
	head := &List[int]{nil, 5}
	list_append(head, 6)
	list_append(head, 7)
	list_append(head, 8)
	list_append(head, 9)
	list_print(head)
	list_delete(&head, 7)
	list_print(head)

	foo := &List[int]{nil, 5}
	list_print(foo)
	list_delete(&foo, 5)
	list_print(foo)
}

func list_append[T any](l *List[T], v T) {
	curr := l

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &List[T]{nil, v}
}

func list_delete[T comparable](l **List[T], v T) {
	curr := *l
	var prev *List[T]

	found := false

	for curr != nil {
		if v == curr.val {
			found = true
			break
		}
		prev = curr
		curr = curr.next
	}

	if found {
		if prev == nil {
			*l = nil
			return
		}
		prev.next = curr.next
	}
}

func list_print[T any](l *List[T]) {
	out := ""
	curr := l

	for curr != nil {
		out += fmt.Sprintf("%v -> ", curr.val)
		curr = curr.next
	}

	out += "<nil>"
	fmt.Println(out)
}
