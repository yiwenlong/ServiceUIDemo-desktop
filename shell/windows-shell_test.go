package shell

import (
	"fmt"
	"testing"
)

func TestWindowsShell_Exec(t *testing.T) {
	ws := WindowsShell{}
	state, res := ws.Exec("dir")
	fmt.Println(res)
	fmt.Println(state.String())
}

func TestWindowsShell_ExecAsync(t *testing.T) {
	ws := WindowsShell{}
	ws.ExecAsync("dir", &Handler{}, 0)
}
