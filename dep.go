package dep

import "github.com/lixin9311/hello"

func Say() *hello.English {
	return hello.Say()
}

func Translate(eng *hello.English) {
	hello.Translate(eng)
}
