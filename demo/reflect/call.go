package reflect

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

/*
	reflect实现的一种，通过函数名统一调用函数的方式
*/

type Methods struct {
	Name string
	Age  int
}

func (methods Methods) CallMethod(name string, params ...reflect.Value) []reflect.Value {
	methodsRef := reflect.ValueOf(methods)

	methodName := Camel(name)

	providerMethod := methodsRef.MethodByName(methodName)
	result := providerMethod.Call(params)
	return result
}

func (methods Methods) SayHello() {
	fmt.Println("hello world!")
}

func (methods Methods) SayName(name string) {
	fmt.Println("my name is:", name)
}

func (methods Methods) GetAge() (int, int) {
	return methods.Age, methods.Age + 1
}

/*
func main() {
	fmt.Println("start")
	methods := &reflect_go.Methods{
		Age: 18,
	}
	methods.CallMethod("say_hello")
	methods.CallMethod("say_name", reflect.ValueOf("xiao ming"))
	result := methods.CallMethod("get_age")
	fmt.Println(result, len(result), result[0].Interface())
	fmt.Println("done!")
}
*/

/*将字符转换为camel风格*/
func Camel(strRaw string) string {
	strList := strings.Split(strRaw, "_")
	var charBuffer bytes.Buffer
	for _, str := range strList {
		str = strings.Title(str)
		charBuffer.WriteString(str)
	}
	return charBuffer.String()
}
