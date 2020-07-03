package main

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_s(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	debugRelativePath()

	str := generate("myPrivateStruct", "MyInterface", "my_package", "../../example/my_package/readme.go")
	expected, err := ioutil.ReadFile("../../example/my_package/readme_builder_generated.go")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), str)
}
