package compile

import "co.insou/gobyteme/compile"

type Blah struct {}

func (blah *Blah) DoBlah()  {
	compile.Compile(nil)
}
