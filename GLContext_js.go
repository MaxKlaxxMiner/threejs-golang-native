//go:build js

package three

import "syscall/js"

type GLContext struct {
	js.Value
}
