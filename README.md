# gosh
**Gosh** is a shell written in Go.

### Rationale behind the shell
The idea was to mimic the behaviour of a shell like `bash` or `zsh` at least to a certain degree. The behaviour involved 
understanding how a shell runs. If you know about `REPL (Read Eval Print Loop)`, you probably know what I'm talking about.
So, I wanted this command line tool to run in a loop which keeps running until it's either stopped via a command like 
`exit` or interrupted by a signal like `SIGINT`

The execution starts from the `cmd/main.go` where `Run()` is called, the responsibility is then passed to the 
`CommandParse()` function. The input to the function is the command which is to be executed. The command is parsed 
(extra white spaces are removed, arguments splitted), the corresponding command is then invoked (only the ones that the 
tool supports). Furthermore, the splitted arguments are then passed to the `CommandExecutor()` where the corresponding 
command is executed.
In `CommandExecutor()`, the commands are executed using `os/exec` package.


> Note: The tool works for `darwin` only

The list of commands that the tool currently supports are:
1. pwd
2. ls (with flags)
3. exit

Examples of command snippets are:

Example One:

![](shellfiles/Screenshot%202024-02-28%20at%2012.54.15%E2%80%AFAM.png)

Example Two:

![](shellfiles/Screenshot%202024-02-28%20at%201.10.43%E2%80%AFAM.png)
