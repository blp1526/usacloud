// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// DeleteCacheWebAccelParam is input parameters for the sacloud API
type DeleteCacheWebAccelParam struct {
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
}

// NewDeleteCacheWebAccelParam return new DeleteCacheWebAccelParam
func NewDeleteCacheWebAccelParam() *DeleteCacheWebAccelParam {
	return &DeleteCacheWebAccelParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *DeleteCacheWebAccelParam) FillValueToSkeleton() {
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
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

}

// Validate checks current values in model
func (p *DeleteCacheWebAccelParam) Validate() []error {
	errors := []error{}

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

func (p *DeleteCacheWebAccelParam) GetResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *DeleteCacheWebAccelParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete-cache"]
}

func (p *DeleteCacheWebAccelParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteCacheWebAccelParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteCacheWebAccelParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteCacheWebAccelParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteCacheWebAccelParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteCacheWebAccelParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteCacheWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteCacheWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteCacheWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteCacheWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteCacheWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteCacheWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteCacheWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteCacheWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteCacheWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteCacheWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteCacheWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteCacheWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteCacheWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteCacheWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *DeleteCacheWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteCacheWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
