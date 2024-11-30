package main

import (
	"fmt"
	"hash/fnv"
	"time"
)

// 179717770 op/s
func main() {
	c := NewEngine(16)

	var (
		stop  bool = true
		index int
	)

	go func() {
		<-time.Tick(30 * time.Second)

		stop = false
	}()

	for stop {
		c.Insert(fmt.Sprintf("%d", index))
		index++
	}

	fmt.Println(index)
}

type Engine struct {
	sendCh        []chan string
	shard         []map[string]struct{}
	numberOfShard uint32
}

func (e *Engine) Insert(data string) {
	hash := fnv.New32a()
	hash.Write([]byte(data))

	shardIndex := int(hash.Sum32() % uint32(e.numberOfShard))

	e.sendCh[shardIndex] <- data
}

func NewEngine(capacity uint32) *Engine {
	if capacity < 1 {
		capacity = 1
	}

	e := Engine{
		sendCh:        make([]chan string, capacity),
		shard:         make([]map[string]struct{}, capacity),
		numberOfShard: capacity,
	}

	for index := 0; index < int(capacity); index++ {
		e.sendCh[index] = make(chan string, 100) // Buffered channel
		e.shard[index] = make(map[string]struct{})

		go e.Receive(index)
	}

	return &e
}

func (e *Engine) Receive(index int) {
	for data := range e.sendCh[index] {
		e.shard[index][data] = struct{}{}
	}
}

func Insert(str map[string]struct{}, data []byte) {
	str[string(data)] = struct{}{}
}
