package examples

import "fmt"

func TestRange() {
	ints := []int{1, 2, 3}
	sum := 0

	for _, num := range ints {
		sum += num
	}
	fmt.Println(sum)

	kvs := make(map[string]string)
	kvs["testKey1"] = "testValue1"
	kvs["testKey2"] = "testValue2"
	for k, v := range kvs {
		fmt.Printf("%s--->%s\n", k, v)
	}
}
