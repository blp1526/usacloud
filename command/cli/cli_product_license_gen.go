// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cli

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/completion"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	listParam := params.NewListProductLicenseParam()
	readParam := params.NewReadProductLicenseParam()

	cliCommand := &cli.Command{
		Name:    "product-license",
		Aliases: []string{"license-info"},
		Usage:   "A manage commands of ProductLicense",
		Action: func(c *cli.Context) error {
			comm := c.App.Command("list")
			if comm != nil {
				return comm.Run(c)
			}
			return cli.ShowSubcommandHelp(c)
		},
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls", "find"},
				Usage:   "List ProductLicense",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "name",
						Usage: "set filter by name(s)",
					},
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:        "from",
						Aliases:     []string{"offset"},
						Usage:       "set offset",
						Destination: &listParam.From,
					},
					&cli.IntFlag{
						Name:        "max",
						Aliases:     []string{"limit"},
						Usage:       "set limit",
						Destination: &listParam.Max,
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
					&cli.StringFlag{
						Name:        "output-type",
						Aliases:     []string{"out"},
						Usage:       "Output type [json/csv/tsv]",
						Destination: &listParam.OutputType,
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:        "quiet",
						Aliases:     []string{"q"},
						Usage:       "Only display IDs",
						Destination: &listParam.Quiet,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"fmt"},
						Usage:       "Output format(see text/template package document for detail)",
						Destination: &listParam.Format,
					},
					&cli.StringFlag{
						Name:        "format-file",
						Usage:       "Output format from file(see text/template package document for detail)",
						Destination: &listParam.FormatFile,
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
					ctx := command.NewContext(c, realArgs, listParam)

					// Set option values for slice
					listParam.Name = c.StringSlice("name")
					listParam.Id = c.Int64Slice("id")
					listParam.Sort = c.StringSlice("sort")
					listParam.Column = c.StringSlice("column")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.ProductLicenseListCompleteArgs(ctx, listParam, cur, prev, commandName)
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
										completion.ProductLicenseListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.ProductLicenseListCompleteFlags(ctx, listParam, name, cur)
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
							completion.ProductLicenseListCompleteArgs(ctx, listParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					listParam.Name = c.StringSlice("name")
					listParam.Id = c.Int64Slice("id")
					listParam.Sort = c.StringSlice("sort")
					listParam.Column = c.StringSlice("column")

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return funcs.ProductLicenseList(ctx, listParam)

				},
			},
			{
				Name:      "read",
				Usage:     "Read ProductLicense",
				ArgsUsage: "<ID>",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "assumeyes",
						Aliases:     []string{"y"},
						Usage:       "assume that the answer to any question which would be asked is yes",
						Destination: &readParam.Assumeyes,
					},
					&cli.StringFlag{
						Name:        "output-type",
						Aliases:     []string{"out"},
						Usage:       "Output type [json/csv/tsv]",
						Destination: &readParam.OutputType,
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:        "quiet",
						Aliases:     []string{"q"},
						Usage:       "Only display IDs",
						Destination: &readParam.Quiet,
					},
					&cli.StringFlag{
						Name:        "format",
						Aliases:     []string{"fmt"},
						Usage:       "Output format(see text/template package document for detail)",
						Destination: &readParam.Format,
					},
					&cli.StringFlag{
						Name:        "format-file",
						Usage:       "Output format from file(see text/template package document for detail)",
						Destination: &readParam.FormatFile,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &readParam.Id,
						Hidden:      true,
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
					ctx := command.NewContext(c, realArgs, readParam)

					// Set option values for slice
					readParam.Column = c.StringSlice("column")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.ProductLicenseReadCompleteArgs(ctx, readParam, cur, prev, commandName)
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
										completion.ProductLicenseReadCompleteArgs(ctx, readParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.ProductLicenseReadCompleteFlags(ctx, readParam, name, cur)
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
							completion.ProductLicenseReadCompleteArgs(ctx, readParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					readParam.Column = c.StringSlice("column")

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					if c.NArg() == 0 {
						return fmt.Errorf("ID argument is required")
					}
					c.Set("id", c.Args().First())

					// Validate specific for each command params
					if errors := readParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), readParam)

					// confirm
					if !readParam.Assumeyes && !command.ConfirmContinue("read") {
						return nil
					}

					// Run command with params
					return funcs.ProductLicenseRead(ctx, readParam)

				},
			},
		},
	}

	// build Category-Resource mapping
	AppendResourceCategoryMap("product-license", &schema.Category{
		Key:         "information",
		DisplayName: "Service/Product informations",
		Order:       90,
	})

	// build Category-Command mapping

	AppendCommandCategoryMap("product-license", "list", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})
	AppendCommandCategoryMap("product-license", "read", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})

	// build Category-Param mapping

	AppendFlagCategoryMap("product-license", "list", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "list", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "list", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "list", "from", &schema.Category{
		Key:         "limit-offset",
		DisplayName: "Limit/Offset options",
		Order:       2147483597,
	})
	AppendFlagCategoryMap("product-license", "list", "id", &schema.Category{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       2147483587,
	})
	AppendFlagCategoryMap("product-license", "list", "max", &schema.Category{
		Key:         "limit-offset",
		DisplayName: "Limit/Offset options",
		Order:       2147483597,
	})
	AppendFlagCategoryMap("product-license", "list", "name", &schema.Category{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       2147483587,
	})
	AppendFlagCategoryMap("product-license", "list", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "list", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "list", "sort", &schema.Category{
		Key:         "sort",
		DisplayName: "Sort options",
		Order:       2147483607,
	})
	AppendFlagCategoryMap("product-license", "read", "assumeyes", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("product-license", "read", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "read", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "read", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "read", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	AppendFlagCategoryMap("product-license", "read", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-license", "read", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
