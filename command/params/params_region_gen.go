// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListRegionParam is input parameters for the sacloud API
type ListRegionParam struct {
	Name              []string
	Id                []int64
	From              int
	Max               int
	Sort              []string
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
}

// NewListRegionParam return new ListRegionParam
func NewListRegionParam() *ListRegionParam {
	return &ListRegionParam{}
}

// Validate checks current values in model
func (p *ListRegionParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Region"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListRegionParam) GetResourceDef() *schema.Resource {
	return define.Resources["Region"]
}

func (p *ListRegionParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListRegionParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListRegionParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListRegionParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListRegionParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListRegionParam) GetOutputFormat() string {
	return "table"
}

func (p *ListRegionParam) SetName(v []string) {
	p.Name = v
}

func (p *ListRegionParam) GetName() []string {
	return p.Name
}
func (p *ListRegionParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListRegionParam) GetId() []int64 {
	return p.Id
}
func (p *ListRegionParam) SetFrom(v int) {
	p.From = v
}

func (p *ListRegionParam) GetFrom() int {
	return p.From
}
func (p *ListRegionParam) SetMax(v int) {
	p.Max = v
}

func (p *ListRegionParam) GetMax() int {
	return p.Max
}
func (p *ListRegionParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListRegionParam) GetSort() []string {
	return p.Sort
}
func (p *ListRegionParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListRegionParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListRegionParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListRegionParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListRegionParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListRegionParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListRegionParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListRegionParam) GetColumn() []string {
	return p.Column
}
func (p *ListRegionParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListRegionParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListRegionParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListRegionParam) GetFormat() string {
	return p.Format
}
func (p *ListRegionParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListRegionParam) GetFormatFile() string {
	return p.FormatFile
}

// ReadRegionParam is input parameters for the sacloud API
type ReadRegionParam struct {
	Assumeyes         bool
	ParamTemplate     string
	ParamTemplateFile string
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Id                int64
}

// NewReadRegionParam return new ReadRegionParam
func NewReadRegionParam() *ReadRegionParam {
	return &ReadRegionParam{}
}

// Validate checks current values in model
func (p *ReadRegionParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Region"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadRegionParam) GetResourceDef() *schema.Resource {
	return define.Resources["Region"]
}

func (p *ReadRegionParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadRegionParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadRegionParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadRegionParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadRegionParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadRegionParam) GetOutputFormat() string {
	return "table"
}

func (p *ReadRegionParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadRegionParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadRegionParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadRegionParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadRegionParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadRegionParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadRegionParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadRegionParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadRegionParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadRegionParam) GetColumn() []string {
	return p.Column
}
func (p *ReadRegionParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadRegionParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadRegionParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadRegionParam) GetFormat() string {
	return p.Format
}
func (p *ReadRegionParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadRegionParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadRegionParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadRegionParam) GetId() int64 {
	return p.Id
}