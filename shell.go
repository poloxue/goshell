package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CmdSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 如果已经到了数据的结尾，但没有数据，则返回
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	inSingleQuote := false
	inDoubleQuote := false
	for i, b := range data {
		switch b {
		case '\'':
			if !inDoubleQuote {
				inSingleQuote = !inSingleQuote
			}
		case '"':
			if !inSingleQuote {
				inDoubleQuote = !inDoubleQuote
			}
		case '\n':
			if !inSingleQuote && !inDoubleQuote {
				// 返回遇到的第一个未被引号包裹的换行符
				return i + 1, data[:i], nil
			}
		}
	}

	// 如果我们处于文件末尾并且还有数据，返回剩余的数据
	if atEOF {
		return len(data), data, nil
	}

	// 如果没有找到换行符，请求更多的数据
	return 0, nil, nil
}

type Shell struct {
	scanner     *bufio.Scanner
	builtinCmds map[string]BuiltinCmder
}

func NewShell() *Shell {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(CmdSplitFunc)

	return &Shell{
		scanner:     scanner,
		builtinCmds: make(map[string]BuiltinCmder),
	}
}

func (s *Shell) RegisterBuiltinCmd(cmdName string, cmd BuiltinCmder) {
	s.builtinCmds[cmdName] = cmd
}

func (s *Shell) PrintPrompt() {
	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		// 如果无法获取工作目录，打印错误并使用默认提示符
		fmt.Println("Error getting current directory:", err)
		fmt.Print("$ ")
		return
	}

	// 获取当前用户的HOME目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		fmt.Print("$ ")
		return
	}

	// 如果当前工作目录以HOME目录开头，则用'~'替换掉HOME目录部分
	if strings.HasPrefix(cwd, homeDir) {
		cwd = strings.Replace(cwd, homeDir, "~", 1)
	}

	// 打印包含当前工作目录的提示符
	fmt.Printf("[%s]$ ", cwd)
}

func (s *Shell) ReadInput() (string, error) {
	if s.scanner.Scan() {
		return s.scanner.Text(), nil
	}
	if err := s.scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

type CmdRequest struct {
	Name string
	Args []string
}

func (s *Shell) ParseInput(input string) []*CmdRequest {
	subInputs := strings.Split(input, ";")

	cmdRequests := make([]*CmdRequest, 0, len(subInputs))
	fmt.Println(subInputs)
	for _, subInput := range subInputs {
		subInput = strings.Trim(subInput, " ")
		subInput = strings.TrimSuffix(subInput, "\n")
		subInput = strings.TrimSuffix(subInput, "\r")
		args := strings.Split(subInput, " ")
		cmdRequests = append(cmdRequests, &CmdRequest{Name: args[0], Args: args[1:]})
	}

	return cmdRequests
}

func (s *Shell) ExecuteCmd(cmdName string, cmdArgs []string) error {
	if cmd, ok := s.builtinCmds[cmdName]; ok {
		return cmd.Execute(cmdArgs...)
	}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func (s *Shell) RunAndListen() error {
	for {
		s.PrintPrompt()

		input, err := s.ReadInput()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		cmdRequests := s.ParseInput(input)
		for _, cmdRequest := range cmdRequests {
			cmdName := cmdRequest.Name
			cmdArgs := cmdRequest.Args

			if cmdName == "exit" {
				return nil
			}

			if err := s.ExecuteCmd(cmdName, cmdArgs); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
		}
	}
}
