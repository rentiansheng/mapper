## Mapper

golang  deep copy library，automatic data mapping。 map to struct, struct to map, struct to struct.
[中文文档](/README-zh-cn.md)

### Install
```go
go get -u github.com/rentiansheng/mapper
```

### Getting Started

#### struct tag order 
tag order: copy > json > gorm
example: [example/struct_tag/main.go](/example/base/main.go)

#### example

| name                                        | code                                                       |
|---------------------------------------------|------------------------------------------------------------|
| base                                        | [example/base/main.go](/example/base/main.go)              |
| struct tag order                            | [example/struct_tag/main.go](/example/struct_tag/main.go)  |
| struct validate                             | [example/validator/main.go](/example/validator/main.go)    |
| extractor value from map or struct by field | [example/extra/extra.go](example/extra/extra.go)           |                        |                                                        |
| extractor keys from map                     | [example/extra/extra.go](example/extra/extra.go)           |                       |                                                        |






### Features

- struct private field automatic mapping
- slice automatic mapping
- automatic mapping by field name
- automatic mapping by field tag
- struct to map automatic mapping
- map to struct automatic mapping
- []byte to string automatic mapping
- data type automatic mapping 
-  any data type to interface data type
- []*Type to []Type automatic mapping
- []Type to []*Type  automatic mapping
- copy struct support multiple tag name,(copy,json,gorm)，
- validate data by struct tag role [rule detail go-playground/validator](https://github.com/go-playground/validator#baked-in-validations)
- automatic mapping integer and unsigned integer
- automatic mapping json.Number to integer 
- extractor value from map or struct by field
- extractor keys/values from map 