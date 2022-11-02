package main

import (
	"context"
	"fmt"
	"github.com/rentiansheng/mapper"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/2
    @desc:

***************************/

type Copy struct {
	F string `json:"json_f"  gorm:"column:gorm_f" copy:"copy_f"`
}
type CopyJSON struct {
	F string `json:"json_f,copy=json_copy_f"  gorm:"column:gorm_f"`
}
type JSON struct {
	F string `json:"json_f"  gorm:"column:gorm_f"`
}
type GORM struct {
	F string `gorm:"column:gorm_f"`
}
type RawField struct {
	F string
}

func main() {
	ctx := context.Background()
	src := map[string]string{
		"copy_f":      "copy_f",
		"json_copy_f": "json_copy_f",
		"json_f":      "json_f",
		"gorm_f":      "gorm_f",
		"F":           "F",
	}

	resultCopy := Copy{}
	err := mapper.Mapper(ctx, src, &resultCopy)
	if err != nil {
		fmt.Println(err)
		return
	}

	resultJSONCopy := CopyJSON{}
	err = mapper.Mapper(ctx, src, &resultJSONCopy)
	if err != nil {
		fmt.Println(err)
		return
	}

	resultJSON := JSON{}
	err = mapper.Mapper(ctx, src, &resultJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	resultGORM := GORM{}
	err = mapper.Mapper(ctx, src, &resultGORM)
	if err != nil {
		fmt.Println(err)
		return
	}

	resultRawField := RawField{}
	err = mapper.Mapper(ctx, src, &resultRawField)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("copy_f      ", "copy_f" == resultCopy.F)
	fmt.Println("json_copy_f ", "json_copy_f" == resultJSONCopy.F)
	fmt.Println("json_f      ", "json_f" == resultJSON.F)
	fmt.Println("gorm_f      ", "gorm_f" == resultGORM.F)
	fmt.Println("F           ", "F" == resultRawField.F)
}
