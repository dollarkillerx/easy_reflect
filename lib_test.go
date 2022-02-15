package easy_reflect

import (
	"fmt"
	"testing"
)

type user struct {
	Name     string `json:"name"`
	NickName string `json:"nick_name" xid:"pcr"`
	Age      int    `json:"age"`
}

func TestStruct(t *testing.T) {
	var u = user{
		Name:     "dollarkiler",
		NickName: "dollarkiler",
		Age:      16,
	}

	ref := NewReflect(&u)
	reflectStruct, err := ref.Struct()
	if err != nil {
		panic(err)
	}

	reflectStruct.BuildTagIndex("json")
	reflectStruct.BuildTagIndex("xid")
	for k, v := range reflectStruct.TagIndex {
		fmt.Println(k, "  ", v)
	}

	_, err = reflectStruct.GetFieldByTag("pcr")
	if err != nil {
		panic(err)
	}
	_, err = reflectStruct.GetFieldByTag("pc2r")
	if err != nil {
		panic(err)
	}
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
