package main

import (
	"sync"
	"time"
)

func main() {
	tracker := NewTracker()

	bob, mary, lisa, jane := NewNode("bob"), NewNode("mary"), NewNode("lisa"), NewNode("jane")

	tracker.AddNode(bob)
	time.Sleep(time.Second * 2)

	tracker.AddNode(mary)
	time.Sleep(time.Second * 2)

	tracker.AddNode(lisa)
	time.Sleep(time.Second * 2)

	tracker.AddNode(jane)
	time.Sleep(time.Second * 2)

	bob.Inc("a", 10)
	mary.Inc("a", 5)
	lisa.Inc("a", 20)
	jane.Inc("a", 2)

	a := NewQuery("a", 2)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	bob.Query(a)

	wg.Wait()
}