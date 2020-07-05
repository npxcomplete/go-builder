# go-builder
A builder generation tool for go structs.

# Install

```bash
go install github.com/npxcomplete/go-builder/src/cmd/go-builder
```

# Usage Example 

##### Declared code:

```go
// go:generate go-builder -struct=myPrivateStruct -interface=MyInterface -package=my_package -input my_struct.go
package my_package

type MyInterface interface {
  Data() string
}

type myPrivateStruct struct {
  data string `builder:"datum"`
}
```

##### generated code (my_struct_builder_generated.go)

```go
package my_package

type MyInterfaceBuilder struct {
	datum string `yaml:"datum" json:"datum" form:"datum"`
}

func (my myPrivateStruct) Data() string {
	return my.data
}

func (builder *MyInterfaceBuilder) Datum(datum string) *MyInterfaceBuilder {
	builder.datum = datum
	return builder
}

func (builder *MyInterfaceBuilder) From(source MyInterface) *MyInterfaceBuilder {
	*builder = MyInterfaceBuilder {
		datum: source.Data(),
	}
	return builder
}

func (builder *MyInterfaceBuilder) Copy() *MyInterfaceBuilder {
	return &MyInterfaceBuilder {
		datum: builder.datum,
	}
}

func (builder *MyInterfaceBuilder) Build() MyInterface {
	return myPrivateStruct {
		data: builder.datum,
	}
}
```

# Planned features

* Omit generated methods if present in input file.
* Copy tags from struct
* Special handling & deep copy for pointers / slices / maps
