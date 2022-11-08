## Mapper
golang 数据深拷贝的类库，支持数据自动映射。 map to struct, struct to map, struct to struct.


### Install 
```go
go get -u github.com/rentiansheng/mapper
```

### Getting Started

#### struct tag 顺序
tag order: copy > json > gorm

example: [example/struct_tag/main.go](/example/base/main.go)

#### example

| name             | code                                                      |
|------------------|-----------------------------------------------------------|
| base             | [example/base/main.go](/example/base/main.go)             |
| struct tag order | [example/struct_tag/main.go](/example/struct_tag/main.go) |
| struct validate  | [example/validator/main.go](/example/validator/main.go)   |


### Features

- 支持struct私有字段自动映射
- 支持slice 自动映射
- 支持按照字段名自动映射
- 支持按照tag 自动映射
- 支持struct 到map 自动映射
- 支持map 到 struct 自动映射
- 支持[]byte to string 
- 数据类型自动识别
- 支持 数据 to interface 自动映射
- 实现[]*Type to []Type
- 实现[]Type to []*Type 
- 支持struct 多种类型tag。(copy,json,gorm)，
- 支持对struct 按照tag 规则校验 [参考go-playground/validator](https://github.com/go-playground/validator#baked-in-validations)
- 支持 integer 和 unsigned integer 自动映射
- 支持 json.Number to integer 自动映射
