//go:generate go-builder -struct=myPrivateStruct -interface=MyInterface -package=my_package -input ./readme.go

package my_package

type MyInterface interface {
	Data() string
}

type myPrivateStruct struct {
	data string `builder:"datum"`
}
