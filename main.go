package main

import (
	"os"
	"bufio"
	"strconv"
	"flag"
	"fmt"
	"github.com/ptzianos/gods/trees/redblacktree"
)

func main() {
	filePath := flag.String("intfile", "ints.txt", "file with integers")
	debug := flag.Bool("debug", false, "Show debug messages")
	flag.Parse()
	file, err := os.Open(*filePath)
	if err != nil {
		panic(fmt.Sprintf("Error when opening %s: %v", *filePath, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	integerRegistry := redblacktree.NewWithIntComparator()
	resultRegistry := make(map[int]int)
	defer integerRegistry.Clear()
	var it *redblacktree.RangedIterator = nil
	targetsFound := 0

	// put all the integers in the map and every time use the RangedIterator to
	// check whether or not there are pairs whose sum falls within the desired
	// range
	for i:=0 ; scanner.Scan() ; i++ {
		nextInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(fmt.Sprintf("Could not read integer from file %s: %v", filePath, err))
		}
		integerRegistry.Put(nextInt, 1)
		it, _ = integerRegistry.IteratorWithin(-10000 - nextInt, 10000 - nextInt)
		it.Begin()
		for it.Next() {
			_, ok := resultRegistry[nextInt + it.Key().(int)]
			if !ok {
				if *debug {
					fmt.Println(nextInt, it.Key(), " = ", nextInt + it.Key().(int))
				}
				targetsFound++
				resultRegistry[nextInt + it.Key().(int)] = 1
			}
		}
		if *debug && i > 0 && i % 10000 == 0 {
			fmt.Println("processed 10 thousand ints")
		}
	}
	fmt.Println("There are", targetsFound, "targets")
}
