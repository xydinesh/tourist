package pkg

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
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

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func ReadDataFile(filename string) (TSPInstance, error) {
	log.WithFields(log.Fields{"filename": filename}).Debug("Reading file for data")

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
					node.X = float32(f)

					f, err = strconv.ParseFloat(data[2], 64)
					if err != nil {
						log.WithError(err).Error("Float parse failed")
					}
					node.Y = float32(f)
					tsp.Nodes = append(tsp.Nodes, node)
				}
			}
		}
	}

	return tsp, nil
}
