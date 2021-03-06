package pkg

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"testing"
	"time"
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

func TestGetDistance(t *testing.T) {
	tsp_instance, err := ReadDataFile("../data/wi29.tsp")
	if err != nil {
		t.Errorf("An error happened, %s", err)
	}

	n0 := tsp_instance.Nodes[0]
	n1 := tsp_instance.Nodes[1]
	n27 := tsp_instance.Nodes[27]
	if d := GetDistance(&n0, &n1); d != 74.535614 {
		log.WithFields(log.Fields{"expected": 74.535614,
			"result": d}).Error("Expected value does not match with return value")
	}

	if d := GetDistance(&n0, &n27); d != 8102.468759 {
		log.WithFields(log.Fields{"expected": 8102.468759,
			"result": d}).Error("Expected value does not match with return value")
	}
}

func TestGenerateRandomRoute(t *testing.T) {
	r := GenerateRandomRoute(29)
	if len(r.NodeOrder) != 29 {
		log.WithField("Nodes in radom route", 29).Error("Expected value does not match with return value")
	}
	c := make([]int, 29)
	for _, n := range r.NodeOrder {
		c[n] = 1
	}
	if !IsRouteReady(c) {
		log.Error("Route is not valid")
	}

	r = GenerateRandomRoute(2900)
	if len(r.NodeOrder) != 2900 {
		log.WithField("Nodes in radom route", 2900).Error("Expected value does not match with return value")
	}
	c = make([]int, 2900)
	for _, n := range r.NodeOrder {
		c[n] = 1
	}
	if !IsRouteReady(c) {
		log.Error("Route is not valid")
	}
}

func TestGetRouteCost(t *testing.T) {
	tsp_instance, err := ReadDataFile("../data/wi29.tsp")
	if err != nil {
		t.Errorf("An error happened, %s", err)
	}
	r := GenerateRandomRoute(tsp_instance.Dimension)
	cost := tsp_instance.GetRouteCost(&r)
	log.WithField("cost", cost).Info("Cost for route")
	if cost < 27603 {
		log.WithFields(log.Fields{"cost": cost, "optimal": 27603}).Error("Solution better than known optimal is found")
	}

	nr := GenerateNeighborRoute(&r)
	cost = tsp_instance.GetRouteCost(&nr)
	log.WithField("cost", cost).Info("Cost for neighbor route")
}

func TestComputeOptimalRoute(t *testing.T) {
	rand.Seed(5)
	tsp_instance, err := ReadDataFile("../data/ex1.tsp")
	if err != nil {
		t.Errorf("An error happened, %s", err)
	}
	r := GenerateRandomRoute(tsp_instance.Dimension)

	s := StopConditon{Goal: 20.0, Iterations: 50, Output: 10}
	nr := tsp_instance.ComputeOptimalRoute(&r, 1.0, 0.99981, &s)
	cost := tsp_instance.GetRouteCost(&nr)
	if cost > s.Goal {
		log.WithFields(log.Fields{"cost": cost, "goal": s.Goal}).Error("Didn't reach the goal")
	} else {
		log.WithFields(log.Fields{"cost": cost, "goal": s.Goal}).Info("Reach the goal")
	}
}

func TestComputeOptimalRoute2(t *testing.T) {
	tsp_instance, err := ReadDataFile("../data/wi29.tsp")
	if err != nil {
		t.Errorf("An error happened, %s", err)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	r := GenerateRandomRoute(tsp_instance.Dimension)
	s := StopConditon{Goal: 27610.0, Iterations: 250000, Output: 10000}
	// r.NodeOrder = []int{2, 3, 7, 4, 0, 1, 5, 9, 10, 11, 14, 18, 17, 21, 22, 20, 28, 27, 25, 24, 26, 23, 15, 19, 16, 13, 12, 8, 6}
	// r.NodeOrder = []int{0, 16, 19, 27, 20, 10, 15, 13, 4, 2, 17, 11, 24, 7, 25, 12, 14, 18, 23, 9, 1, 5, 28, 8, 6, 21, 3, 22, 26}
	// [12 13 15 23 26 24 19 25 27 28 22 21 20 16 17 18 14 11 10 9 5 1 0 4 7 3 2 6 8]
	nr := tsp_instance.ComputeOptimalRoute(&r, 16000.0, 0.99981, &s)
	cost := tsp_instance.GetRouteCost(&nr)
	log.WithField("r", r).Debug("r")
	log.WithField("nr", nr).Debug("nr")
	if cost > s.Goal {
		log.WithFields(log.Fields{"cost": cost, "goal": s.Goal}).Error("Didn't reach the goal")
	} else {
		log.WithFields(log.Fields{"cost": cost, "goal": s.Goal}).Info("Reach the goal")
	}
}

func TestComputeOptimalRoute3(t *testing.T) {
	tsp_instance, err := ReadDataFile("../data/uy734.tsp")
	if err != nil {
		t.Errorf("An error happened, %s", err)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	r := GenerateRandomRoute(tsp_instance.Dimension)
	s := StopConditon{Goal: 79200.0, Iterations: 5250000, Output: 25000}
	nr := tsp_instance.ComputeOptimalRoute(&r, 10.0, 0.9999787, &s)
	cost := tsp_instance.GetRouteCost(&nr)
	if cost > s.Goal {
		log.WithFields(log.Fields{"cost": cost, "goal": s.Goal}).Error("Didn't reach the goal")
	} else {
		log.WithFields(log.Fields{"cost": cost, "goal": s.Goal}).Info("Reach the goal")
	}
}
