package dep

import "hello"

func Say() *hello.English {
	return hello.Say()
}

func Translate(eng *hello.English) {
	hello.Translate(eng)
}
