package shell

import (
	"fmt"
	"testing"
)

type Handler struct {
}

func (h *Handler) HandleEcho(token ShellToken, echo string) {
	fmt.Printf("ECHO ==> token: %d, echo: %s\n", token, echo)
}
func (h *Handler) HandleError(token ShellToken, exitCode int, state string) {
	fmt.Printf("ERROR ==> token: %d, code: %d, state: %s\n", token, exitCode, state)
}
func (h *Handler) HandleSuccess(token ShellToken) {
	fmt.Printf("SUCCESS ==> tokend: %d\n", token)
}

func TestExecShell(t *testing.T) {
	state, s := ExecShell("ls")
	t.Logf("State: %s\nRes: %s\n", state.String(), s)
}

func TestExecShellAsync(t *testing.T) {
	h := Handler{}
	ExecShellAsync("ls /", &h, ShellToken(1))
}

func TestExecShellAdmin(t *testing.T) {
	h := Handler{}
	ExecShellAdmin("ls /", &h, ShellToken(1))
}
