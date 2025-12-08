package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	treeMap := make(map[int]*Node)
	IDSlice, err := validateRecords(records)
	if err != nil {
		return nil, err
	}
	for _, v := range IDSlice {
		treeMap[v] = New()
		treeMap[v].ID = v
	}

	for _, p := range records {
		if p.ID == 0 {
			continue
		}
		treeMap[p.Parent].Children = append(treeMap[p.Parent].Children, treeMap[p.ID])
	}
	for _, v := range treeMap {
		sort.SliceStable(v.Children, func(i int, j int) bool {
			return v.Children[i].ID < v.Children[j].ID
		})
	}

	return treeMap[0], nil
}

func New() *Node {
	return &Node{}
}

func validateRecords(records []Record) ([]int, error) {
	IDSlice := make([]int, 0)
	for _, v := range records {
		IDSlice = append(IDSlice, v.ID)
	}
	sort.SliceStable(IDSlice, func(i int, j int) bool {
		return IDSlice[i] < IDSlice[j]
	})
	if IDSlice[0] != 0 {
		return nil, errors.New("invalid records")
	}
	for i := 1; i < len(IDSlice); i++ {
		if IDSlice[i]-IDSlice[i-1] != 1 {
			return nil, errors.New("duplicate or missing record")
		}
	}
	for _, v := range records {
		if v.ID < 0 {
			return nil, errors.New("invalid records id")
		}
		if v.Parent < 0 {
			return nil, errors.New("invalid records parent")
		}
		if v.ID >= len(records) {
			return nil, errors.New("invalid tree ")
		}
		if v.Parent >= len(records) {
			return nil, errors.New("invalid tree ")
		}
		if v.ID == 0 && v.Parent != 0 {
			return nil, errors.New("invalid root")
		}
		if v.ID <= v.Parent && v.ID != 0 {
			return nil, errors.New("invalid parentship")
		}
	}

	return IDSlice, nil
}
