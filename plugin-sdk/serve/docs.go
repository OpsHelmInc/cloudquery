package serve

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/docs"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/plugin"
)

const (
	pluginDocShort = "Generate documentation for tables"
	pluginDocLong  = `Generate documentation for tables

If format is markdown, a destination directory will be created (if necessary) containing markdown files.
Example:
doc ./output 

If format is JSON, a destination directory will be created (if necessary) with a single json file called __tables.json.
Example:
doc --format json .
`
)

func (s *PluginServe) newCmdPluginDoc() *cobra.Command {
	format := newEnum([]string{"json", "markdown"}, "markdown")
	cmd := &cobra.Command{
		Use:   "doc <directory>",
		Short: pluginDocShort,
		Long:  pluginDocLong,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := s.plugin.Init(cmd.Context(), nil, plugin.NewClientOptions{
				NoConnection: true,
			}); err != nil {
				return err
			}
			tables, err := s.plugin.Tables(cmd.Context(), plugin.TableOptions{
				Tables: []string{"*"},
			})
			if err != nil {
				return err
			}
			g := docs.NewGenerator(s.plugin.Name(), tables)
			f := docs.FormatMarkdown
			if format.Value == "json" {
				f = docs.FormatJSON
			}
			return g.Generate(args[0], f)
		},
	}
	cmd.Flags().Var(format, "format", fmt.Sprintf("output format. one of: %s", strings.Join(format.Allowed, ",")))
	return cmd
}
