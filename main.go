package main

import (
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"
)

var (
	m = map[int]int{
		1: 1,
		2: 2,
	}
	// Z Expose to outside world
	Z = 42
)

func main() {
	fmt.Println("Hello, playground")

	i := 27
	fmt.Printf("%v, %T n", m, m)

	//convert
	j := float32(i)
	fmt.Printf("%v, %T n", j, j)

	s := strconv.Itoa(i)
	fmt.Printf("%v, %T n", s, s)

	var n bool = true
	fmt.Printf("%v, %T n", n, n)

	var x uint16 = 42
	x = x << 3
	fmt.Printf("%v, %T n", x, x)

	var s1 = "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T n", s1[2], s1[2])
	fmt.Printf("%v, %T n", b, b)

	var r rune = 'a'
	fmt.Printf("%v, %T n", r, r)

	//fmt.Println(strings.ReplaceAll("{\\\"_source\\\":false,\\\"aggregations\\\":{\\\"attrAggLvl1\\\":{\\\"aggregations\\\":{\\\"attrAggLvl2\\\":{\\\"composite\\\":{\\\"size\\\":1000,\\\"sources\\\":[{\\\"attribute_name\\\":{\\\"terms\\\":{\\\"field\\\":\\\"attributes.name\\\"}}},{\\\"attribute_value\\\":{\\\"terms\\\":{\\\"field\\\":\\\"attributes.value\\\"}}}]}}},\\\"nested\\\":{\\\"path\\\":\\\"attributes\\\"}},\\\"categoryAggrLvl1\\\":{\\\"aggregations\\\":{\\\"categoryAggrLvl2\\\":{\\\"terms\\\":{\\\"field\\\":\\\"collections.cid\\\",\\\"size\\\":1000}}},\\\"nested\\\":{\\\"path\\\":\\\"collections\\\"}},\\\"maxPrice\\\":{\\\"aggregations\\\":{\\\"priceMaxAggr\\\":{\\\"max\\\":{\\\"field\\\":\\\"variations.price\\\"}}},\\\"nested\\\":{\\\"path\\\":\\\"variations\\\"}},\\\"minPrice\\\":{\\\"aggregations\\\":{\\\"priceMinAggr\\\":{\\\"min\\\":{\\\"field\\\":\\\"variations.price\\\"}}},\\\"nested\\\":{\\\"path\\\":\\\"variations\\\"}},\\\"varAttrAggLvl1\\\":{\\\"aggregations\\\":{\\\"varAttrAggLvl2\\\":{\\\"composite\\\":{\\\"size\\\":1000,\\\"sources\\\":[{\\\"attribute_name\\\":{\\\"terms\\\":{\\\"field\\\":\\\"variations.attributes.name\\\"}}},{\\\"attribute_value\\\":{\\\"terms\\\":{\\\"field\\\":\\\"variations.attributes.value\\\"}}}]}}},\\\"nested\\\":{\\\"path\\\":\\\"variations.attributes\\\"}}},\\\"from\\\":0,\\\"query\\\":{\\\"bool\\\":{\\\"must\\\":{\\\"bool\\\":{\\\"should\\\":{\\\"bool\\\":{\\\"must\\\":{\\\"bool\\\":{\\\"must\\\":{\\\"wildcard\\\":{\\\"title\\\":{\\\"value\\\":\\\"*helloween*\\\"}}}}}}}}}}},\\\"size\\\":500,\\\"sort\\\":[{\\\"id\\\":{\\\"order\\\":\\\"desc\\\"}}]}\"", "\\", ""))

	sliceTest := make([]int, 100)
	sliceTest = append(sliceTest, 1)

	//string
	for i := range "abc" {
		//i is byte
		fmt.Printf("LMSLMS %v, %T n", i, i)
	}

	for _, v := range "abc" {
		//v is int 32
		fmt.Printf("LMSLMS %v, %T n", v, v)
	}

	iii := 3 / 2
	fmt.Printf("abc %v", iii)

	jsonSize := ""
	encodeSize, _ := EncodeRedisData(jsonSize)
	fmt.Printf("a: %v, %d\n", encodeSize, len(encodeSize))

	//constant
	const myConst int = 1
	const (
		myConst1 = 1
	)

	const (
		a = iota + 5
		bc
		cc
	)

	//array
	var students [...]string
	students[0] = ""

	//total copy
	newStudents := students
	newStudents[0] = ""

	//slice
	aslice := []int{1, 2, 3}
	aslice = append(aslice, 4)

	bslice := make([]int, 10)
	bslice = append(bslice, 11)

}

func EncodeRedisData(data interface{}) (string, error) {
	var sb strings.Builder
	encoder := gob.NewEncoder(&sb)
	err := encoder.Encode(data)
	return sb.String(), err
}

func ESQuery() {
	jsonSize := "{\\\"_source\\\":false,\\\"from\\\":0,\\\"query\\\":{\\\"function_score\\\":{\\\"boost_mode\\\":\\\"replace\\\",\\\"functions\\\":[{\\\"script_score\\\":{\\\"script\\\":{\\\"params\\\":{\\\"calcTime\\\":1666828800000,\\\"gmvAdjust\\\":7000,\\\"gmvWeight\\\":1,\\\"timeWeight\\\":1},\\\"source\\\":\\\"double getDouble(def doc_v) { if (doc_v.size()==0) {return 0.0} else {return doc_v.value}}\\\\n\\\\t\\\\t\\\\t\\\\tdef timeScore=params.timeWeight * 1000 / (1 + Math.exp(0.02*( (params.calcTime/1000 - doc['created_at'].value.toEpochSecond())/86400-7)));\\\\n\\\\t\\\\t\\\\t\\\\tdef gmvScore=params.gmvWeight * 1000 / (1 + Math.exp(-0.0003*(getDouble(doc['gmv'])-params.gmvAdjust)));\\\\n\\\\t\\\\t\\\\t \\\\treturn gmvScore+timeScore\\\"}}}],\\\"query\\\":{\\\"bool\\\":{\\\"must\\\":{\\\"bool\\\":{\\\"should\\\":{\\\"bool\\\":{\\\"must\\\":{\\\"bool\\\":{\\\"must\\\":[{\\\"bool\\\":{\\\"filter\\\":{\\\"nested\\\":{\\\"path\\\":\\\"mfps\\\",\\\"query\\\":{\\\"bool\\\":{\\\"filter\\\":[{\\\"term\\\":{\\\"mfps.type\\\":1}},{\\\"term\\\":{\\\"mfps.event_id\\\":\\\"633dd2e64aa65c8f49f7e29c\\\"}},{\\\"range\\\":{\\\"mfps.min_discount_percentage\\\":{\\\"from\\\":20,\\\"include_lower\\\":true,\\\"include_upper\\\":true,\\\"to\\\":null}}}]}}}}}},{\\\"bool\\\":{\\\"filter\\\":{\\\"nested\\\":{\\\"path\\\":\\\"wish_category_path_to_root\\\",\\\"query\\\":{\\\"term\\\":{\\\"wish_category_path_to_root.wish_category_id\\\":\\\"1495\\\"}}}}}}]}}}}}}}},\\\"score_mode\\\":\\\"sum\\\"}},\\\"size\\\":1000,\\\"timeout\\\":\\\"5000ms\\\"}"
	fmt.Printf(strings.ReplaceAll(jsonSize, "\\", ""))
}
