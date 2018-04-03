// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListPriceParam is input parameters for the sacloud API
type ListPriceParam struct {
	Name              []string `json:"name"`
	Id                []int64  `json:"id"`
	From              int      `json:"from"`
	Max               int      `json:"max"`
	Sort              []string `json:"sort"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
}

// NewListPriceParam return new ListPriceParam
func NewListPriceParam() *ListPriceParam {
	return &ListPriceParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListPriceParam) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []int64{0}
	}
	if isEmpty(p.From) {
		p.From = 0
	}
	if isEmpty(p.Max) {
		p.Max = 0
	}
	if isEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}

}

// Validate checks current values in model
func (p *ListPriceParam) Validate() []error {
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
		validator := define.Resources["Price"].Commands["list"].Params["id"].ValidateFunc
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
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
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

func (p *ListPriceParam) GetResourceDef() *schema.Resource {
	return define.Resources["Price"]
}

func (p *ListPriceParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListPriceParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListPriceParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListPriceParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListPriceParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListPriceParam) SetName(v []string) {
	p.Name = v
}

func (p *ListPriceParam) GetName() []string {
	return p.Name
}
func (p *ListPriceParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListPriceParam) GetId() []int64 {
	return p.Id
}
func (p *ListPriceParam) SetFrom(v int) {
	p.From = v
}

func (p *ListPriceParam) GetFrom() int {
	return p.From
}
func (p *ListPriceParam) SetMax(v int) {
	p.Max = v
}

func (p *ListPriceParam) GetMax() int {
	return p.Max
}
func (p *ListPriceParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListPriceParam) GetSort() []string {
	return p.Sort
}
func (p *ListPriceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListPriceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListPriceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListPriceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListPriceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListPriceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListPriceParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListPriceParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListPriceParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListPriceParam) GetColumn() []string {
	return p.Column
}
func (p *ListPriceParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListPriceParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListPriceParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListPriceParam) GetFormat() string {
	return p.Format
}
func (p *ListPriceParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListPriceParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListPriceParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListPriceParam) GetQuery() string {
	return p.Query
}
