# gosh
**Gosh** is a Go based shell written in Go.

#### Rationale behind the shell
The idea was to mimic the behaviour of an shell like bash or zsh at least to a certain degree. It was good way to see 
what was needed to create the something similar in Go.

The execution flow of the tool starts with promptUI as the whole logic is wrapped up by the promptUI. The command is 
parsed and then passed to the next layer which directs the commands to the `cmd/` directory where the corresponding 
command is invoked. In each command, the final responsibility is handled by the `CommandExecutor` which executes the 
commands using `os/exec` package.

PromptUI is used to mimic the behaviour of a shell (which is the interactiveness of a shell)


The list of commands that the tool currently supports are:
1. pwd
2. ls
