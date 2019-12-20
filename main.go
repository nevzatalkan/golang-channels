package main

import (
	"fmt"
	"sync"
)

type Child struct {
	ID    int
	Value string
}

func prepareChildren(childcount int) []Child {

	children := make([]Child, childcount)
	for cInd := 0; cInd < childcount; cInd++ {
		children[cInd] = Child{ID: cInd, Value: "Child" + fmt.Sprint(cInd)}
	}

	return children
}

var childArr []Child
var wg sync.WaitGroup

func init() {
	childArr = prepareChildren(10)
	//fmt.Println("%v", childArr)
}

func processChildren(veletler []Child, c chan []Child) {
	defer wg.Done()
	for index := 0; index < len(veletler); index++ {
		veletler[index].Value = "İşlendi" + veletler[index].Value
	}
	c <- veletler
}

func processChildrenLoop(veletler []Child) []Child {

	for index := 0; index < len(veletler); index++ {
		veletler[index].Value = "İşlendi" + veletler[index].Value
	}
	return veletler
}

func parallel2() []Child {
	veletler := childArr
	//donenveletler := make([]Child, 10)

	wg.Add(10)

	var chans [10]chan []Child
	for i := range chans {
		chans[i] = make(chan []Child)
	}

	for i := 0; i < 10; i++ {
		go processChildren(veletler[i:i+1], chans[i])
	}

	donenveletler := make([]Child, 0)
	for i := 0; i < 10; i++ {
		donenTemp := <-chans[i]
		// fmt.Println(donenTemp)
		donenveletler = append(donenveletler, donenTemp...)
	}

	wg.Wait()

	return donenveletler
}

func forloop() {
	veletler := childArr
	for i := 0; i < 10; i++ {
		//processChildrenLoop(veletler[i:i+1])
		for index := 0; index < len(veletler); index++ {
			veletler[index].Value = "İşlendi" + veletler[index].Value
		}
	}
}

func main() {
	sonuc := parallel2()
	fmt.Println(sonuc)
}
