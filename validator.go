package mapper

import (
	valid "github.com/go-playground/validator/v10"
)

/***************************
    @author: tiansheng.ren
    @date: 2022/11/1
    @desc:

***************************/

type Validator interface {
	Struct(s interface{}) error
}

type validate struct {
	v *valid.Validate
}

func (v *validate) Struct(s interface{}) error {
	return v.v.Struct(s)
}

func newValidateStruct() Validator {
	v := valid.New()
	//v.SetTagName(tagName)
	return &validate{
		v: v,
	}
}
