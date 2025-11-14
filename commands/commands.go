package commands

import (
	"web/commands/generate"
	"web/commands/gorm"
	"web/commands/migrate"
	"web/commands/worker"

	"github.com/urfave/cli/v2"
)

func All() []*cli.Command {
	commands := []*cli.Command{
		migrate.Command(),
		generate.Command(),
		gorm.Command(),
		worker.Command(),
	}
	return commands
}
