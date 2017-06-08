package pkg

import (
	"fmt"
	"io/ioutil"
)

type Node struct {
	Id string
	X  float32
	Y  float32
}

type TSPInstance struct {
	Name           string
	Comments       []string
	Type           string
	Dimension      int
	EdgeWeightType string
	Nodes          []Node
}

func ReadDataFile(filename string) (TSPInstance, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return TSPInstance{}, err
	}
	fmt.Printf("%v", b)
	return TSPInstance{}, nil
}
