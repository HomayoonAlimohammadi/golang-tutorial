package cache

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("app not found")

type App struct {
	Name   string
	source string
}

type CacheLayer interface {
	Get(key string) (App, error)
}

type cacheLayerImpl1 struct{}

func (cacheLayerImpl1) Get(packageName string) (App, error) {
	delay := time.Duration(rand.Intn(10)+40) * time.Millisecond
	time.Sleep(delay)

	if rand.Intn(100) > 10 {
		return App{}, ErrNotFound
	}

	return App{
		Name:   "App-1",
		source: "cache-1",
	}, nil
}

type cacheLayerImpl2 struct{}

func (cacheLayerImpl2) Get(packageName string) (App, error) {
	delay := time.Duration(rand.Intn(50)+100) * time.Millisecond
	time.Sleep(delay)

	if rand.Intn(100) > 70 {
		return App{}, ErrNotFound
	}
	return App{
		Name:   "App-1",
		source: "cache-2",
	}, nil
}

type cacheLayerImpl3 struct{}

func (cacheLayerImpl3) Get(packageName string) (App, error) {
	delay := time.Duration(rand.Intn(5000)+500) * time.Millisecond
	time.Sleep(delay)

	return App{
		Name:   "App-1",
		source: "cache-3",
	}, nil
}

type DataProvider struct {
	CacheLayers []CacheLayer
}

func NewDataProvider() *DataProvider {
	return &DataProvider{
		CacheLayers: []CacheLayer{
			&cacheLayerImpl1{},
			&cacheLayerImpl2{},
			&cacheLayerImpl3{},
		},
	}
}

func (d *DataProvider) GetData(packageName string) (App, error) {
	resultChan := make(chan App, 1)
	for _, cacheLayer := range d.CacheLayers {
		go func(c CacheLayer) {
			app, err := c.Get(packageName)
			if err == nil {
				resultChan <- app
				return
			}
		}(cacheLayer)
	}
	result := <-resultChan
	return result, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dataProvider := NewDataProvider()
	startTime := time.Now()
	result, err := dataProvider.GetData("ir.divar")
	delay := -time.Until(startTime)
	fmt.Printf("Result: %v\nErr: %v\nDelay: %v", result, err, delay)
}
