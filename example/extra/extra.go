package main

import (
	"context"
	"fmt"
	"github.com/rentiansheng/mapper"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/11
    @desc:

***************************/

var (
	ctx context.Context
)

func main() {
	sliceStruct()
	sliceMap()
	MapKey()
}

func sliceStruct() {
	items := []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}{
		{ID: 1, Name: "name1"}, {ID: 2, Name: "name2"},
	}
	idValues, nameValues := make([]int, 0), make([]string, 0)
	if err := mapper.ObjectsField(ctx, "id", items, &idValues); err != nil {
		fmt.Println("extractor struct field error.", err.Error())
		return
	}
	if err := mapper.ObjectsField(ctx, "name", items, &nameValues); err != nil {
		fmt.Println("extractor struct field error. ", err.Error())
		return
	}

	fmt.Println("struct field id", idValues)
	fmt.Println("struct field name", nameValues)
}

func sliceMap() {
	items := []map[string]int{
		{"id": 1, "name": 1}, {"id": 2, "name": 2},
	}
	idValues, nameValues := make([]int, 0), make([]int64, 0)
	if err := mapper.ObjectsField(ctx, "id", items, &idValues); err != nil {
		fmt.Println("extractor struct field error.", err.Error())
		return
	}
	if err := mapper.ObjectsField(ctx, "name", items, &nameValues); err != nil {
		fmt.Println("extractor struct field error. ", err.Error())
		return
	}
	fmt.Println("map field id", idValues)
	fmt.Println("map field name", nameValues)
}

func MapKey() {
	items := map[string]struct{}{
		"1": {},
		"2": {},
	}
	keys := make([]string, 0)
	if err := mapper.MapKeys(ctx, items, &keys); err != nil {
		fmt.Println("extractor struct field error. ", err.Error())
		return
	}
	fmt.Println("map keys", keys)

}
