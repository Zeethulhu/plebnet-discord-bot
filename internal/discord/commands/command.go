package commands

import "github.com/bwmarrin/discordgo"

// Command represents a bot command implementation.
type Command interface {
	// Name returns the command's name used to invoke it.
	Name() string
	// Description returns a short description shown in help output.
	Description() string
	// Execute runs the command with the provided arguments.
	Execute(args []string, s *discordgo.Session, m *discordgo.MessageCreate)
}

var registry = make(map[string]Command)

// Register adds a command to the registry. Typically called from init().
func Register(cmd Command) {
	registry[cmd.Name()] = cmd
}

// Get retrieves a command by name.
func Get(name string) (Command, bool) {
	cmd, ok := registry[name]
	return cmd, ok
}

// All returns all registered commands.
func All() []Command {
	cmds := make([]Command, 0, len(registry))
	for _, c := range registry {
		cmds = append(cmds, c)
	}
	return cmds
}
