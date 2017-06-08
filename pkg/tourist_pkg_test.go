package pkg

import (
	"testing"
)

func TestNodeStruct(t *testing.T) {
	t.Log("Testing Node structure")
	node := Node{}
	if id := node.Id; id != "" {
		t.Errorf("Expected id is null, got %s \n", id)
	}

	if x := node.X; x != 0 {
		t.Errorf("Expected X is 0, got %f \n", x)
	}

	if y := node.Y; y != 0 {
		t.Errorf("Expected id is 0, got %f \n", y)
	}

}

func TestReadDataFile(t *testing.T) {
	t.Log("Testing ReadDataFile function")
	tsp_instance, err := ReadDataFile("../data/wi29.tsp")
	if err != nil {
		t.Errorf("An error happened, %s", err)
	}

	if d := tsp_instance.Dimension; d != 29 {
		t.Errorf("Expected 29, got %d", d)
	}
}
