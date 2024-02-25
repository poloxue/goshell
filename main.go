package main

func main() {
	s := NewShell()

	s.RegisterBuiltinCmd("cd", &ChangeDirCommand{})

	_ = s.RunAndListen()
}
