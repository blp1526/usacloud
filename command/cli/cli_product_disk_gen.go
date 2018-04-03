// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/completion"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
)

func init() {
	listParam := params.NewListProductDiskParam()
	readParam := params.NewReadProductDiskParam()

	cliCommand := &cli.Command{
		Name:    "product-disk",
		Aliases: []string{"disk-plan"},
		Usage:   "A manage commands of ProductDisk",
		Action: func(c *cli.Context) error {
			comm := c.App.Command("list")
			if comm != nil {
				return comm.Action(c)
			}
			return cli.ShowSubcommandHelp(c)
		},
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
				Name:    "from",
				Aliases: []string{"offset"},
				Usage:   "set offset",
			},
			&cli.IntFlag{
				Name:    "max",
				Aliases: []string{"limit"},
				Usage:   "set limit",
			},
			&cli.StringSliceFlag{
				Name:  "sort",
				Usage: "set field(s) for sort",
			},
			&cli.StringFlag{
				Name:  "param-template",
				Usage: "Set input parameter from string(JSON)",
			},
			&cli.StringFlag{
				Name:  "param-template-file",
				Usage: "Set input parameter from file",
			},
			&cli.BoolFlag{
				Name:  "generate-skeleton",
				Usage: "Output skelton of parameter JSON",
			},
			&cli.StringFlag{
				Name:    "output-type",
				Aliases: []string{"out"},
				Usage:   "Output type [table/json/csv/tsv]",
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
			&cli.StringFlag{
				Name:  "query",
				Usage: "JMESPath query(using when '--output-type' is json only)",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls", "find"},
				Usage:   "List ProductDisk (default)",
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
						Name:    "from",
						Aliases: []string{"offset"},
						Usage:   "set offset",
					},
					&cli.IntFlag{
						Name:    "max",
						Aliases: []string{"limit"},
						Usage:   "set limit",
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
					&cli.StringFlag{
						Name:    "output-type",
						Aliases: []string{"out"},
						Usage:   "Output type [table/json/csv/tsv]",
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
					&cli.StringFlag{
						Name:  "query",
						Usage: "JMESPath query(using when '--output-type' is json only)",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
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

					// set default output-type
					// when params have output-type option and have empty value
					var outputTypeHolder interface{} = listParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// build command context
					ctx := command.NewContext(c, realArgs, listParam)

					// Set option values
					if c.IsSet("name") {
						listParam.Name = c.StringSlice("name")
					}
					if c.IsSet("id") {
						listParam.Id = c.Int64Slice("id")
					}
					if c.IsSet("from") {
						listParam.From = c.Int("from")
					}
					if c.IsSet("max") {
						listParam.Max = c.Int("max")
					}
					if c.IsSet("sort") {
						listParam.Sort = c.StringSlice("sort")
					}
					if c.IsSet("param-template") {
						listParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						listParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						listParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						listParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						listParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						listParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						listParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						listParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("query") {
						listParam.Query = c.String("query")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.ProductDiskListCompleteArgs(ctx, listParam, cur, prev, commandName)
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
										completion.ProductDiskListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.ProductDiskListCompleteFlags(ctx, listParam, name, cur)
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
							completion.ProductDiskListCompleteArgs(ctx, listParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					listParam.ParamTemplate = c.String("param-template")
					listParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(listParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewListProductDiskParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(listParam, p)
					}

					// Set option values
					if c.IsSet("name") {
						listParam.Name = c.StringSlice("name")
					}
					if c.IsSet("id") {
						listParam.Id = c.Int64Slice("id")
					}
					if c.IsSet("from") {
						listParam.From = c.Int("from")
					}
					if c.IsSet("max") {
						listParam.Max = c.Int("max")
					}
					if c.IsSet("sort") {
						listParam.Sort = c.StringSlice("sort")
					}
					if c.IsSet("param-template") {
						listParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						listParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						listParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						listParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						listParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						listParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						listParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						listParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("query") {
						listParam.Query = c.String("query")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = listParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Generate skeleton
					if listParam.GenerateSkeleton {
						listParam.GenerateSkeleton = false
						listParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(listParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return funcs.ProductDiskList(ctx, listParam)

				},
			},
			{
				Name:      "read",
				Usage:     "Read ProductDisk",
				ArgsUsage: "<ID>",
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
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
					&cli.StringFlag{
						Name:    "output-type",
						Aliases: []string{"out"},
						Usage:   "Output type [table/json/csv/tsv]",
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
					&cli.StringFlag{
						Name:  "query",
						Usage: "JMESPath query(using when '--output-type' is json only)",
					},
					&cli.Int64Flag{
						Name:   "id",
						Usage:  "[Required] set resource ID",
						Hidden: true,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
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

					// set default output-type
					// when params have output-type option and have empty value
					var outputTypeHolder interface{} = readParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// build command context
					ctx := command.NewContext(c, realArgs, readParam)

					// Set option values
					if c.IsSet("assumeyes") {
						readParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						readParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						readParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						readParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						readParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						readParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						readParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						readParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						readParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("query") {
						readParam.Query = c.String("query")
					}
					if c.IsSet("id") {
						readParam.Id = c.Int64("id")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.ProductDiskReadCompleteArgs(ctx, readParam, cur, prev, commandName)
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
										completion.ProductDiskReadCompleteArgs(ctx, readParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.ProductDiskReadCompleteFlags(ctx, readParam, name, cur)
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
							completion.ProductDiskReadCompleteArgs(ctx, readParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					readParam.ParamTemplate = c.String("param-template")
					readParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(readParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewReadProductDiskParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(readParam, p)
					}

					// Set option values
					if c.IsSet("assumeyes") {
						readParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						readParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						readParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						readParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						readParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						readParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						readParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						readParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						readParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("query") {
						readParam.Query = c.String("query")
					}
					if c.IsSet("id") {
						readParam.Id = c.Int64("id")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = readParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Generate skeleton
					if readParam.GenerateSkeleton {
						readParam.GenerateSkeleton = false
						readParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(readParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					if c.NArg() == 0 {
						return fmt.Errorf("ID argument is required")
					}
					c.Set("id", c.Args().First())
					readParam.SetId(c.Int64("id"))

					// Validate specific for each command params
					if errors := readParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), readParam)

					// confirm
					if !readParam.Assumeyes {
						if !isTerminal() {
							return fmt.Errorf("When using redirect/pipe, specify --assumeyes(-y) option")
						}
						if !command.ConfirmContinue("read") {
							return nil
						}
					}

					// Run command with params
					return funcs.ProductDiskRead(ctx, readParam)

				},
			},
		},
	}

	// build Category-Resource mapping
	AppendResourceCategoryMap("product-disk", &schema.Category{
		Key:         "information",
		DisplayName: "Service/Product informations",
		Order:       90,
	})

	// build Category-Command mapping

	AppendCommandCategoryMap("product-disk", "list", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})
	AppendCommandCategoryMap("product-disk", "read", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})

	// build Category-Param mapping

	AppendFlagCategoryMap("product-disk", "list", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "list", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "list", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "list", "from", &schema.Category{
		Key:         "limit-offset",
		DisplayName: "Limit/Offset options",
		Order:       2147483597,
	})
	AppendFlagCategoryMap("product-disk", "list", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("product-disk", "list", "id", &schema.Category{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       2147483587,
	})
	AppendFlagCategoryMap("product-disk", "list", "max", &schema.Category{
		Key:         "limit-offset",
		DisplayName: "Limit/Offset options",
		Order:       2147483597,
	})
	AppendFlagCategoryMap("product-disk", "list", "name", &schema.Category{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       2147483587,
	})
	AppendFlagCategoryMap("product-disk", "list", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "list", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("product-disk", "list", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("product-disk", "list", "query", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "list", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "list", "sort", &schema.Category{
		Key:         "sort",
		DisplayName: "Sort options",
		Order:       2147483607,
	})
	AppendFlagCategoryMap("product-disk", "read", "assumeyes", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("product-disk", "read", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "read", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "read", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "read", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("product-disk", "read", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	AppendFlagCategoryMap("product-disk", "read", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "read", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("product-disk", "read", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("product-disk", "read", "query", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("product-disk", "read", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
