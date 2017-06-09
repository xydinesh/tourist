package pkg

import (
	log "github.com/sirupsen/logrus"
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
		log.WithFields(log.Fields{"expected": 29,
			"result": d}).Error("Expected value does not match with return value")
	}

	if tsp_instance.Name != "wi29" {
		log.WithField("Name", tsp_instance.Name).Error("Return value is different from expected value of wi29")
	}

	if tsp_instance.Type != "TSP" {
		log.WithField("TSP", tsp_instance.Name).Error("Return value is different from expected value of TSP")
	}

	if tsp_instance.EdgeWeightType != "EUC_2D" {
		log.WithField("Edge weight type", tsp_instance.Name).Error("Return value is different from expected value of EUC_2D")
	}

	if len(tsp_instance.Nodes) != tsp_instance.Dimension {
		log.WithFields(log.Fields{"NodeCount": len(tsp_instance.Nodes),
			"Dimension": tsp_instance.Dimension}).Error("Node count is not equal to dimension of the instance")
	}

	node := tsp_instance.Nodes[0]
	if node.Id != "1" || node.X != 20833.3333 || node.Y != 17100.0000 {
		log.WithFields(log.Fields{"Id": node.Id, "X": node.X, "Y": node.Y}).Error("Node does not have expected values")
	}

	node = tsp_instance.Nodes[28]
	if node.Id != "29" || node.X != 27462.5000 || node.Y != 12992.2222 {
		log.WithFields(log.Fields{"Id": node.Id, "X": node.X, "Y": node.Y}).Error("Node does not have expected values")
	}

	node = tsp_instance.Nodes[10]
	if node.Id != "11" || node.X != 23700.0000 || node.Y != 15933.3333 {
		log.WithFields(log.Fields{"Id": node.Id, "X": node.X, "Y": node.Y}).Error("Node does not have expected values")
	}

}
