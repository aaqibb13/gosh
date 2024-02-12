package ui

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"gosh/main/parse"
)

func PromptUIWrapper() {

	// this is a template example for PromptUI example
	// validate := func(input string) error {
	// 	_, err := strconv.ParseInt(input, 10, 64)
	// 	fmt.Println("testing validate")
	// 	return nil
	// }

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		//Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "gosh",
		Templates: templates,
		// Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Interrupt called %v\n", err)
		return
	}

	fmt.Printf("You answered %s\n", result)

	// Here you can parse the sent command along with flags and pass them to the corresponding commands.
	parse.CommandParse(result)
}
