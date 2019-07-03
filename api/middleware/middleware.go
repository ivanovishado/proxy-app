package middleware

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kataras/iris"

	utils "github.com/ivanovishado/proxy-app/api/utils"
)

// Request is the data that we will receive
type Request struct {
	Domain   string
	Weight   int
	Priority int
}

// AppRequests should contain the sorted requests
var AppRequests []string

var priorityLevels map[string]utils.PriorityLevel

// Repository should implement common methods
type Repository interface {
	Read() []*Request
}

func (q *Request) Read() []*Request {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}

	var queue []*Request
	var domain string
	var weight int
	var priority int

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "":
			queue = append(queue, &Request{domain, weight, priority})
			break
		case "alpha", "omega", "beta":
			domain = text
			break
		}
		if strings.Contains(text, "weight") {
			weight, _ = strconv.Atoi(strings.Split(text, ":")[1])
			continue
		}
		if strings.Contains(text, "priority") {
			priority, _ = strconv.Atoi(strings.Split(text, ":")[1])
			continue
		}
	}

	return queue
}

// SetPriorityLevels adds the priority levels to the map based on the configuration
func SetPriorityLevels() {
	var repo Repository
	repo = &Request{}
	priorityLevels = make(map[string]utils.PriorityLevel)

	for _, row := range repo.Read() {
		priorityLevels[row.Domain] = determinePriorityLevel(row.Weight, row.Priority)
	}
}

// ProxyMiddleware determines in which order the requests should be processed
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "domain error"})
		return
	}

	switch priorityLevels[domain] {
	case utils.High:
		utils.ShiftArray(&AppRequests, 0, domain)
		break
	case utils.Medium:
		utils.ShiftArray(&AppRequests, (len(AppRequests)+1)/2, domain)
		break
	case utils.Low:
		utils.ShiftArray(&AppRequests, len(AppRequests), domain)
		break
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
