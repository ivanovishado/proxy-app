package middlewares

import (
	"fmt"

	"github.com/kataras/iris"

	utils "github.com/ivanovishado/proxy-app/api/utils"
)

type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Que declaration
var Que []*Queue

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

func (q *Queue) Read() []*Queue {
	return MockQueue()
}

// MockQueue should mock an array of Queues
func MockQueue() []*Queue {
	return []*Queue{
		{
			Domain:   "alpha",
			Weight:   7,
			Priority: 7,
		},
		{
			Domain:   "omega",
			Weight:   1,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weight:   1,
			Priority: 1,
		},
	}
}

// InitQueue should return an array of data
func InitQueue() {
	Que = append(Que, &Queue{})
}

// ProxyMiddleware determines in which order the requests should be processed
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	// Could use an if here to prevent the initialization of the empty Queue
	var repo Repository
	repo = &Queue{}
	fmt.Println("From header", domain)

	for _, row := range repo.Read() {
		fmt.Printf("From source: Domain: %s, Weight: %d, Priority: %d\n", row.Domain, row.Weight, row.Priority)

		switch determinePriorityLevel(row.Weight, row.Priority) {
		case utils.High:
			fmt.Println("High")
			break
		case utils.Medium:
			fmt.Println("Medium")
			break
		case utils.Low:
			fmt.Println("Low")
			break
		}
	}

	c.Next()
}

func determinePriorityLevel(weight int, priority int) utils.PriorityLevel {
	if weight > 5 && priority > 5 {
		return utils.High
	} else if weight < 5 && priority < 5 {
		return utils.Low
	}

	return utils.Medium
}
