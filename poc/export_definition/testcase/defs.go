package testcase

type PublicStruct struct{}
type privateStruct struct{}

type PublicInterface interface{}
type privateInterface interface{}

type PublicTypeAlias = int
type privateTypeAlias = int

var PublicVar = 0
var privateVar = 0

const PublicConst = 0
const privateConst = 0

func PublicFunc()  {}
func privateFunc() {}

func (PublicStruct) PublicMethodOfPublicStruct()  {}
func (PublicStruct) privateMethodOfPublicStruct() {}

func (privateStruct) PublicMethodOfPrivateStruct()  {}
func (privateStruct) privateMethodOfPrivateStruct() {}

type PublicIF interface {
	Method()
}
type privateIF interface {
	Method()
}
type Impl struct{}
type impl struct{}

func (Impl) Method() {}
func (impl) Method() {}
