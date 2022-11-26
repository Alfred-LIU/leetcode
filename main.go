package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//construct version store
	versionStore := Constructor()
	in := bufio.NewReader(os.Stdin)
	for {
		s, err := in.ReadString('\n')
		if err != nil {
			// io.EOF is expected, anything else
			// should be handled/reported
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		// Do something with the line of text
		// in string variable s.
		callCmd(&versionStore, s)
	}
	fmt.Println("LMSLMS", versionStore.Get("key2"))
	fmt.Println("LMSLMS", versionStore)

}

func callCmd(versionStore *VersionStore, cmd string) {
	cmdElement := strings.Split(cmd, " ")
	if cmdElement[0] == put {
		value, _ := strconv.Atoi(cmdElement[2])
		versionStore.Put(cmdElement[1], value)
	} else if cmdElement[0] == get {
		if len(cmdElement) == 2 {
			fmt.Println(versionStore.Get(cmdElement[1]))
		} else if len(cmdElement) == 3 {
			version, _ := strconv.Atoi(cmdElement[2])
			fmt.Println(versionStore.GetVersion(cmdElement[1], version))
		}
	}
}

const (
	put = "PUT"
	get = "GET"
)

// gloabal version
var version = 1

type elem struct {
	version int
	value   int
}

type VersionStore struct {
	m map[string][]elem
}

func Constructor() VersionStore {
	return VersionStore{m: make(map[string][]elem)}
}

func (this *VersionStore) Put(key string, value int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = []elem{}
	}
	this.m[key] = append(this.m[key], elem{version, value})
	version++
}

func (this *VersionStore) Get(key string) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}
	elems := this.m[key]
	return elems[len(elems)-1].value
}

func (this *VersionStore) GetVersion(key string, version int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}

	index := findFirst(this.m[key], version)

	if index < 0 {
		return -1
	}

	return this.m[key][index].value
}

func findFirst(input []elem, version int) int {
	left, right := 0, len(input)-1
	for left <= right {
		mid := (right-left)/2 + left
		if input[mid].version == version {
			return mid
		} else if input[mid].version < version {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}