package mapper

/***************************
    @author: tiansheng.ren
    @date: 2022/11/1
    @desc:

***************************/

type Option struct {
	copyPrivate    *bool
	validateStruct *bool
}

var (
	optionTrue  = true
	optionFalse = false
)

func mergeOption(options ...Option) Option {
	ro := DefaultOption()
	for _, option := range options {

		if option.validateStruct != nil {
			ro.validateStruct = option.validateStruct
		}
		if option.copyPrivate != nil {
			ro.copyPrivate = option.copyPrivate
		}
	}
	return ro
}

func DefaultOption() Option {

	return Option{
		copyPrivate:    &optionFalse,
		validateStruct: &optionFalse,
	}
}

func (o Option) CopyPrivate() Option {
	o.copyPrivate = &optionTrue
	return o
}

func (o Option) ValidateStruct() Option {
	o.validateStruct = &optionTrue
	return o
}

func OptionCopyPrivate() Option {
	o := DefaultOption()
	o.copyPrivate = &optionTrue
	return o
}

func OptionValidateStruct() Option {
	o := DefaultOption()
	o.validateStruct = &optionTrue
	return o
}
