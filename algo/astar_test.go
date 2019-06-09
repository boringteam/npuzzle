package algo

import (
	"fmt"
	"testing"
)

func TestTabInSlice(t *testing.T) {
	open_node1 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{8, 2, 5, 1, 6, 3, 0, 4, 7}}
	open_node2 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{2, 8, 5, 1, 6, 3, 0, 4, 7}}
	open_node3 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{5, 2, 8, 1, 6, 3, 0, 4, 7}}
	tab := []int16{8, 2, 5, 1, 6, 3, 0, 4, 7}
	openList := []*node{&open_node1, &open_node2, &open_node3}
	res := tabInSlice(tab, openList)
	if fmt.Sprint(res.tab) != fmt.Sprint(tab) {
		t.Error("Error: TestTabInSlice")
	}
	tab = []int16{8, 2, 5, 1, 6, 3, 0, 7, 4}
	res = tabInSlice(tab, openList)
	if res != nil {
		t.Error("Error: TestTabInSlice")
	}
}

func TestRemoveFromList(t *testing.T) {
	open_node1 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{8, 2, 5, 1, 6, 3, 0, 4, 7}}
	open_node2 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{2, 8, 5, 1, 6, 3, 0, 4, 7}}
	open_node3 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{5, 2, 8, 1, 6, 3, 0, 4, 7}}
	openList := []*node{&open_node1, &open_node2, &open_node3}
	finalList := removeFromList(&open_node2, openList)
	if fmt.Sprint(finalList) != fmt.Sprint([]*node{&open_node1, &open_node3}) {
		t.Error("Error: TestRemoveFromList")
	}
	openList = []*node{&open_node1}
	finalList = removeFromList(&open_node1, openList)
	fmt.Println(finalList)
	if fmt.Sprint(finalList) != fmt.Sprint([]*node{}) {
		t.Error("Error: TestRemoveFromList")
	}
}
