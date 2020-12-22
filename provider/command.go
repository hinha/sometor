package provider

// Command CLI command bearer of coronator
type Command interface {
	InjectCommand(scaffold ...CommandScaffold)
}

// CommandScaffold use for standard of creating CLI command
type CommandScaffold interface {
	Use() string
	Example() string
	Short() string
	Run(args []string)
}
