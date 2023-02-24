package set

import (
	"github.com/sailpoint-oss/sailpoint-cli/internal/config"
	"github.com/spf13/cobra"
)

func newSearchTemplateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "searchTemplates",
		Short:   "configure the custom search template file path",
		Long:    "configure the custom search template file path",
		Example: "sail set search /path/to/search/templates",
		Aliases: []string{"search"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			err := config.InitConfig()
			if err != nil {
				return err
			}

			config.SetCustomSearchTemplatePath(args[0])

			err = config.SaveConfig()
			if err != nil {
				return err
			}

			return nil
		},
	}
	return cmd

}
