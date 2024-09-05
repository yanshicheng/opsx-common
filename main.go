package main

import (
	"encoding/json"
	"fmt"
	"github.com/yanshicheng/opsx-common/handler/errorx/types"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"reflect"
)

type Student struct {
	Name string
	Age  int
	Sex  string
}

type ResponseStructSlice struct {
	Page    int
	Total   int
	PageNum int
	Data    []Student
}

func main() {
	fmt.Println("Hello World")
	s1 := Student{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}
	s2 := Student{
		Name: "李四",
		Age:  19,
		Sex:  "女",
	}
	s3 := Student{
		Name: "王五",
		Age:  20,
		Sex:  "男",
	}
	res := ResponseStructSlice{
		Page:    1,
		Total:   3,
		PageNum: 1,
		Data:    []Student{s1, s2, s3},
	}
	stype := types.Status{
		Code:    0,
		Data:    res,
		Message: "ok",
	}
	// json 序列化
	marData, _ := protojson.Marshal(&stype)
	fmt.Println(string(marData))
}

// ConvertToAnyMap 接收一个 map[string]interface{} 类型的 data，并返回一个 map[string]*anypb.Any
func ConvertToAnyMap(data map[string]interface{}) (map[string]*anypb.Any, error) {
	result := make(map[string]*anypb.Any)

	for key, value := range data {
		var anyValue *anypb.Any
		var err error

		switch v := value.(type) {
		case proto.Message:
			// 如果值是 proto.Message 类型，直接封装为 *anypb.Any
			anyValue, err = anypb.New(v)
		case string:
			// 如果值是字符串，封装为自定义的 StringValue 类型
			stringValue := &types.StringValue{Value: v}
			anyValue, err = anypb.New(stringValue)
		default:
			// 处理其他类型，包括 struct 和 slice
			rv := reflect.ValueOf(value)
			switch rv.Kind() {
			case reflect.Struct:
				// 处理 struct 类型
				anyValue, err = handleStruct(v)
			case reflect.Slice, reflect.Array:
				// 处理 slice 或 array 类型
				anyValue, err = handleSlice(rv)
			default:
				err = fmt.Errorf("unsupported data type: %T", value)
			}
		}

		if err != nil {
			return nil, err
		}

		result[key] = anyValue
	}

	return result, nil
}

// handleStruct 处理 struct 类型并返回 *anypb.Any
func handleStruct(value interface{}) (*anypb.Any, error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	structMessage := &GenericStruct{}
	err = json.Unmarshal(jsonData, structMessage)
	if err != nil {
		return nil, err
	}
	return anypb.New(structMessage)
}

// handleSlice 处理 slice 或 array 类型并返回 *anypb.Any
func handleSlice(rv reflect.Value) (*anypb.Any, error) {
	var anyItems []*anypb.Any
	for i := 0; i < rv.Len(); i++ {
		item := rv.Index(i).Interface()
		anyItem, err := MarshalToAny(item)
		if err != nil {
			return nil, err
		}
		anyItems = append(anyItems, anyItem)
	}
	anyList := &types.AnyList{Items: anyItems}
	return anypb.New(anyList)
}

// 定义一个通用的 Proto message 用于封装 Struct
type GenericStruct struct {
	Data map[string]interface{} `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}
