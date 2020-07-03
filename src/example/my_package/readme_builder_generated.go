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
	*builder = MyInterfaceBuilder{
		datum: source.Data(),
	}
	return builder
}

func (builder *MyInterfaceBuilder) Copy() *MyInterfaceBuilder {
	return &MyInterfaceBuilder{
		datum: builder.datum,
	}
}

func (builder *MyInterfaceBuilder) Build() MyInterface {
	return myPrivateStruct{
		data: builder.datum,
	}
}
