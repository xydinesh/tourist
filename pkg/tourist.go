package pkg

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	Id string
	X  float64
	Y  float64
}

type TSPInstance struct {
	Name           string
	Comments       []string
	Type           string
	Dimension      int
	EdgeWeightType string
	Nodes          []Node
}

type Route struct {
	NodeOrder []int
	Size      int
}

type StopConditon struct {
	Goal           float32
	MinTemperature float64
	Iterations     int
	Output         int
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	rand.Seed(time.Now().UTC().UnixNano())

}

func IsRouteReady(n []int) bool {
	for i, v := range n {
		if v == 0 {
			log.WithFields(log.Fields{"i": i, "v": v}).Error("Node is not set properly")
			return false
		}
	}
	return true
}

func (tsp *TSPInstance) ComputeOptimalRoute(r *Route, t0 float64, b float64, s *StopConditon) Route {
	currentIter := 0
	t := t0
	cost := float32(s.Goal) + 100.0
	var optRoute Route
	for currentIter < s.Iterations && s.Goal < cost && 1e-200 < t {
		nr := GenerateNeighborRoute(r)
		c0 := tsp.GetRouteCost(r)
		c1 := tsp.GetRouteCost(&nr)
		delta := c1 - c0
		u := c0 - float32(t*math.Log(rand.Float64()))
		cost = c0
		if delta < 0 {
			r = &nr
			cost = c1
			// log.WithFields(log.Fields{"c0": c0, "c1": c1, "delta": c1 - c0, "t": t, "i": currentIter, "u": u}).Debug("New route picked")
		} else if c1 <= u {
			r = &nr
			cost = c1
			// log.WithFields(log.Fields{"c0": c0, "c1": c1, "delta": c1 - c0, "t": t, "i": currentIter, "u": u}).Debug("New route picked")

		}
		t = b * t
		if currentIter%s.Output == 0 {
			log.WithFields(log.Fields{"c0": c0, "c1": c1, "delta": c1 - c0, "t": t, "i": currentIter, "u": u}).Info("Opimization")
		}

		optRoute = *r
		currentIter += 1
	}
	return optRoute
}

func (tsp *TSPInstance) GetRouteCost(r *Route) float32 {
	cost := float32(0.0)
	for i := 0; i < tsp.Dimension-1; i++ {
		d := GetDistance(&tsp.Nodes[r.NodeOrder[i]], &tsp.Nodes[r.NodeOrder[i+1]])
		cost += d
	}
	// Calculate cost to return
	cost += GetDistance(&tsp.Nodes[r.NodeOrder[tsp.Dimension-1]], &tsp.Nodes[r.NodeOrder[0]])
	return cost
}

func GenerateNeighborRoute(r *Route) Route {
	nr := Route{}
	nr.Size = r.Size
	nr.NodeOrder = make([]int, nr.Size)
	copy(nr.NodeOrder, r.NodeOrder)
	i := rand.Intn(nr.Size)
	j := rand.Intn(nr.Size)
	for i == j {
		i = rand.Intn(nr.Size)
		j = rand.Intn(nr.Size)
	}

	ni := nr.NodeOrder[i]
	nr.NodeOrder[i] = nr.NodeOrder[j]
	nr.NodeOrder[j] = ni
	return nr
}

func GenerateRandomRoute(n int) Route {
	r := Route{}
	r.Size = n
	r.NodeOrder = make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		rn := rand.Intn(n)
		for c[rn] == 1 {
			rn = rand.Intn(n)
		}
		// log.WithFields(log.Fields{"i": i, "rn": rn}).Debug("Node order")
		r.NodeOrder[i] = rn
		c[rn] = 1
	}

	if !IsRouteReady(c) {
		return Route{}
	}

	// log.WithField("order", r.NodeOrder).Debug("Node order")
	return r
}

func GetDistance(n1 *Node, n2 *Node) float32 {
	// log.WithFields(log.Fields{"x": n1.X, "y": n1.Y, "id": n1.Id}).Debug("Node 1")
	// log.WithFields(log.Fields{"x": n2.X, "y": n2.Y, "id": n2.Id}).Debug("Node 2")
	xd := float64(n1.X - n2.X)
	yd := float64(n1.Y - n2.Y)
	d := math.Sqrt(xd*xd + yd*yd)
	// log.WithFields(log.Fields{"n1": n1.Id, "n2": n2.Id, "distance": d}).Debug("Distance between Nodes")
	return float32(d)
}

func ReadDataFile(filename string) (TSPInstance, error) {
	f, err := os.Open(filename)
	if err != nil {
		return TSPInstance{}, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	parserState := 0
	tsp := TSPInstance{}
	for s.Scan() {
		switch parserState {
		case 0:
			{
				data := strings.Split(s.Text(), ":")
				if len(data) > 1 {
					key := strings.TrimSpace(data[0])
					value := strings.TrimSpace(data[1])
					switch key {
					case "NAME":
						tsp.Name = value
						log.WithField(key, value).Debug("Name")
					case "COMMENT":
						tsp.Comments = append(tsp.Comments, value)
						log.WithField(key, value).Debug("Comment")
					case "TYPE":
						tsp.Type = value
						log.WithField(key, value).Debug("Type")
					case "DIMENSION":
						d, err := strconv.Atoi(value)
						if err != nil {
							log.WithError(err).Error("String to Int conversion failed")
						}
						tsp.Dimension = d
						log.WithField(key, value).Debug("Dimension")
					case "EDGE_WEIGHT_TYPE":
						tsp.EdgeWeightType = value
						log.WithField(key, value).Debug("EdgeWeightType")
					default:
						log.WithField(key, value).Warning("Didn't find any matching fielts")
					}
				} else {
					parserState = 1
				}
			}
		case 1:
			{
				data := strings.Split(s.Text(), " ")
				if len(data) >= 3 {
					node := Node{}

					node.Id = data[0]
					f, err := strconv.ParseFloat(data[1], 64)
					if err != nil {
						log.WithError(err).Error("Float parse failed")
					}
					node.X = f

					f, err = strconv.ParseFloat(data[2], 64)
					if err != nil {
						log.WithError(err).Error("Float parse failed")
					}
					node.Y = f
					tsp.Nodes = append(tsp.Nodes, node)
				}
			}
		}
	}

	return tsp, nil
}
