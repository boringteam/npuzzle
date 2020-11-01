package algo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/boringteam/npuzzle/src/checker"
)

func TestTabInSlice(t *testing.T) {
	result := checker.BuildCorrectResult(3)
	open_node1 := createNode(nil, []int16{8, 2, 5, 1, 6, 3, 0, 4, 7}, result, 0, "manhattan")
	open_node2 := createNode(nil, []int16{2, 8, 5, 1, 6, 3, 0, 4, 7}, result, 0, "manhattan")
	open_node3 := createNode(nil, []int16{5, 2, 8, 1, 6, 3, 0, 4, 7}, result, 0, "manhattan")
	tab := []int16{8, 2, 5, 1, 6, 3, 0, 4, 7}
	openList := []*node{open_node1, open_node2, open_node3}
	res := tabInSlice(tab, openList)
	if !reflect.DeepEqual(res.tab, tab) {
		t.Error("Error: TestTabInSlice")
	}
	tab = []int16{8, 2, 5, 1, 6, 3, 0, 7, 4}
	res = tabInSlice(tab, openList)
	if res != nil {
		t.Error("Error: TestTabInSlice")
	}
}

func TestAddToList(t *testing.T) {
	open_node1 := node{parent: nil, F: 5, G: 0, H: 0, tab: []int16{8, 2, 5, 1, 6, 3, 0, 4, 7}}
	open_node2 := node{parent: nil, F: 2, G: 0, H: 0, tab: []int16{2, 8, 5, 1, 6, 3, 0, 4, 7}}
	open_node3 := node{parent: nil, F: 3, G: 0, H: 0, tab: []int16{5, 2, 8, 1, 6, 3, 0, 4, 7}}
	openList := []*node{&open_node2, &open_node1}
	finalList := addToList(&open_node3, openList)
	if !reflect.DeepEqual(finalList, []*node{&open_node2, &open_node3, &open_node1}) {
		t.Error("Error: TestAddToList")
	}
	open_node4 := node{parent: nil, F: 1, G: 0, H: 0, tab: []int16{5, 2, 8, 1, 6, 3, 0, 4, 7}}
	finalList = addToList(&open_node4, finalList)
	if !reflect.DeepEqual(finalList, []*node{&open_node4, &open_node2, &open_node3, &open_node1}) {
		t.Error("Error: TestAddToList2")
	}
}

func TestRemoveFromList(t *testing.T) {
	open_node1 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{8, 2, 5, 1, 6, 3, 0, 4, 7}}
	open_node2 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{2, 8, 5, 1, 6, 3, 0, 4, 7}}
	open_node3 := node{parent: nil, F: 0, G: 0, H: 0, tab: []int16{5, 2, 8, 1, 6, 3, 0, 4, 7}}
	openList := []*node{&open_node1, &open_node2, &open_node3}
	finalList := removeFromList(&open_node2, openList)
	if !reflect.DeepEqual(finalList, []*node{&open_node1, &open_node3}) {
		t.Error("Error: TestRemoveFromList")
	}
	openList = []*node{&open_node1}
	finalList = removeFromList(&open_node1, openList)
	if fmt.Sprint(finalList) != fmt.Sprint([]*node{}) {
		t.Error("Error: TestRemoveFromList")
	}
}
