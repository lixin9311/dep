package dep

import "lixin.io/hello/src"

func Say() *hello.English {
	return hello.Say()
}

func Translate(eng *hello.English) {
	hello.Translate(eng)
}
