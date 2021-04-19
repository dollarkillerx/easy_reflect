# easy_reflect
golang reflect utils,  Make the use of reflection more convenient

## use 
`go get github.com/dollarkillerx/easy_reflect`

### set struct
```go 
type user struct {
	Name     string `json:"name"`
	nickname string `json:"nickname"`
	Age      int    `json:"age"`
}

u := user{
    Name:     "name1",
    nickname: "name2",
    Age:      18,
}
r := NewReflect(&u)

fmt.Println(r.Type())   // ptr 
fmt.Println(r.Elem().Type())  // struct
fmt.Println(r.Elem().NumField()) // 3

for i := 0; i < field; i++ {
    index, err := r.Elem().FieldByIndex(i)
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(index.Tag.Get("json"))
    fmt.Println(index.Name)

    if index.CanSet() {
        switch index.Kind() {
        case reflect.String:
            index.SetString("xxx")
        case reflect.Int:
            index.SetInt(4646)
        }
    }
}
```