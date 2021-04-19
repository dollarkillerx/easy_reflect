package easy_reflect

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

type user struct {
	Name     string `json:"name"`
	nickname string `json:"nickname"`
	Age      int    `json:"age"`
}

func TestStruct(t *testing.T) {
	u := user{
		Name:     "name1",
		nickname: "name2",
		Age:      18,
	}

	r := NewReflect(&u)
	fmt.Println(r.Type())
	fmt.Println(r.Elem().Type())
	fmt.Println(r.Elem().NumField())
	field, err := r.Elem().NumField()
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < field; i++ {
		index, err := r.Elem().FieldByIndex(i)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(index.Tag.Get("json"))
		fmt.Println(index.Name)

		fmt.Println()

		if index.CanSet() {
			switch index.Kind() {
			case reflect.String:
				index.SetString("xxx")
			case reflect.Int:
				index.SetInt(4646)
			}
		}
	}

	fmt.Println(u)
}

//// list 绑定
//func BindAllEdges(resultSet []map[string]interface{}, v interface{}) error {
//	refType := reflect.TypeOf(v)
//	refVal := reflect.ValueOf(v)
//	if refType.Kind() != reflect.Ptr {
//		return errors.New("类型错误 应该为&[]")
//	}
//
//	// 解引用看内部类型
//	if refType.Elem().Kind() != reflect.Slice {
//		return errors.New("类型错误 应该为&[]")
//	}
//
//	elem := refType.Elem().Elem() // 内部具体的类型
//	sliceVal := refVal.Elem()     // 具体slice
//
//	// 建立一个新数组
//	newArr := make([]reflect.Value, 0)
//
//	// 建立一个item
//	for idx := range resultSet {
//		index := idx
//		edge, err := newEdge(resultSet[index], elem)
//		if err != nil {
//			return errors.WithStack(err)
//		}
//
//		newArr = append(newArr, edge.Elem())
//	}
//
//	// 重写
//	resArr := reflect.Append(sliceVal, newArr...)
//	sliceVal.Set(resArr)
//
//	return nil
//}
//
//// 无中生有
//func newEdge(record map[string]interface{}, elemType reflect.Type) (resVal reflect.Value, err error) {
//	defer func() {
//		if error := recover(); error != nil {
//			err = errors.Errorf("%s", error)
//		}
//	}()
//	// 1. 获取struct所有的 json tag
//	// 2. 对应ColNames
//	// 3. 类型转换
//	// 4. 填数据
//	refVal := reflect.New(elemType)
//	refTypeElem := elemType
//	refValElem := refVal.Elem()
//
//	for i := 0; i < refTypeElem.NumField(); i++ {
//		typeElem := refTypeElem.Field(i)
//		structElem := refValElem.Field(i)
//		if !structElem.CanSet() {
//			continue
//		}
//
//		// 取tag
//		tag := typeElem.Tag.Get("json")
//		if tag == "" {
//			tag = typeElem.Name
//		}
//
//		if tag == "-" {
//			continue
//		}
//
//		col, ex := record[tag]
//		if !ex {
//			continue
//		}
//
//		switch typeElem.Type.Kind() {
//		case reflect.Float64:
//			f, ex := col.(float64)
//			if !ex {
//				continue
//			}
//			structElem.SetFloat(f)
//		case reflect.String:
//			s, ex := col.(string)
//			if !ex {
//				continue
//			}
//			s = strings.TrimSpace(s)
//			s = strings.Replace(s, "\"", "", -1)
//			structElem.SetString(s)
//		}
//	}
//
//	return refVal, err
//}
