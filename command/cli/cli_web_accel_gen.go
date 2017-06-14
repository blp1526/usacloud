// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cli

import (
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/completion"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	deleteCacheParam := params.NewDeleteCacheWebAccelParam()

	cliCommand := &cli.Command{
		Name:  "web-accel",
		Usage: "A manage commands of WebAccel",
		Subcommands: []*cli.Command{
			{
				Name:      "delete-cache",
				Aliases:   []string{"purge"},
				Usage:     "DeleteCache WebAccel",
				ArgsUsage: "[URLs]...",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "assumeyes",
						Aliases: []string{"y"},
						Usage:   "Assume that the answer to any question which would be asked is yes",
					},
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.StringFlag{
						Name:    "output-type",
						Aliases: []string{"out"},
						Usage:   "Output type [json/csv/tsv]",
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:    "quiet",
						Aliases: []string{"q"},
						Usage:   "Only display IDs",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"fmt"},
						Usage:   "Output format(see text/template package document for detail)",
					},
					&cli.StringFlag{
						Name:  "format-file",
						Usage: "Output format from file(see text/template package document for detail)",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					command.GlobalOption.Validate(false)

					// build command context
					ctx := command.NewContext(c, realArgs, deleteCacheParam)

					// Set option values
					if c.IsSet("assumeyes") {
						deleteCacheParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						deleteCacheParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						deleteCacheParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("output-type") {
						deleteCacheParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						deleteCacheParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						deleteCacheParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						deleteCacheParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						deleteCacheParam.FormatFile = c.String("format-file")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.WebAccelDeleteCacheCompleteArgs(ctx, deleteCacheParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.WebAccelDeleteCacheCompleteArgs(ctx, deleteCacheParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.WebAccelDeleteCacheCompleteFlags(ctx, deleteCacheParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.WebAccelDeleteCacheCompleteArgs(ctx, deleteCacheParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					deleteCacheParam.ParamTemplate = c.String("param-template")
					deleteCacheParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(deleteCacheParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewDeleteCacheWebAccelParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(deleteCacheParam, p)
					}

					// Set option values
					if c.IsSet("assumeyes") {
						deleteCacheParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						deleteCacheParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						deleteCacheParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("output-type") {
						deleteCacheParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						deleteCacheParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						deleteCacheParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						deleteCacheParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						deleteCacheParam.FormatFile = c.String("format-file")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := deleteCacheParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), deleteCacheParam)

					// confirm
					if !deleteCacheParam.Assumeyes && !command.ConfirmContinue("delete-cache") {
						return nil
					}

					// Run command with params
					return funcs.WebAccelDeleteCache(ctx, deleteCacheParam)

				},
			},
		},
	}

	// build Category-Resource mapping
	AppendResourceCategoryMap("web-accel", &schema.Category{
		Key:         "saas",
		DisplayName: "Other services",
		Order:       80,
	})

	// build Category-Command mapping

	AppendCommandCategoryMap("web-accel", "delete-cache", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})

	// build Category-Param mapping

	AppendFlagCategoryMap("web-accel", "delete-cache", "assumeyes", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("web-accel", "delete-cache", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("web-accel", "delete-cache", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("web-accel", "delete-cache", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("web-accel", "delete-cache", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("web-accel", "delete-cache", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("web-accel", "delete-cache", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("web-accel", "delete-cache", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
