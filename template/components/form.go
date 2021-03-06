package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type FormAttribute struct {
	Name         string
	Header       template.HTML
	Content      []types.Form
	GroupContent [][]types.Form
	GroupHeaders []string
	Footer       template.HTML
	Url          string
	Method       string
	PrimaryKey   string
	InfoUrl      string
	CSRFToken    string
	Title        template.HTML
	Prefix       string
	types.Attribute
}

func (compo *FormAttribute) SetHeader(value template.HTML) types.FormAttribute {
	compo.Header = value
	return compo
}

func (compo *FormAttribute) SetPrimaryKey(value string) types.FormAttribute {
	compo.PrimaryKey = value
	return compo
}

func (compo *FormAttribute) SetContent(value []types.Form) types.FormAttribute {
	compo.Content = value
	return compo
}

func (compo *FormAttribute) SetGroupContent(value [][]types.Form) types.FormAttribute {
	compo.GroupContent = value
	return compo
}

func (compo *FormAttribute) SetGroupHeaders(value []string) types.FormAttribute {
	compo.GroupHeaders = value
	return compo
}

func (compo *FormAttribute) SetFooter(value template.HTML) types.FormAttribute {
	compo.Footer = value
	return compo
}

func (compo *FormAttribute) SetPrefix(value string) types.FormAttribute {
	compo.Prefix = value
	return compo
}

func (compo *FormAttribute) SetUrl(value string) types.FormAttribute {
	compo.Url = value
	return compo
}

func (compo *FormAttribute) SetInfoUrl(value string) types.FormAttribute {
	compo.InfoUrl = value
	return compo
}

func (compo *FormAttribute) SetMethod(value string) types.FormAttribute {
	compo.Method = value
	return compo
}

func (compo *FormAttribute) SetTitle(value template.HTML) types.FormAttribute {
	compo.Title = value
	return compo
}

func (compo *FormAttribute) SetToken(value string) types.FormAttribute {
	compo.CSRFToken = value
	return compo
}

func (compo *FormAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "form",
		"form/default", "form/file", "form/textarea",
		"form/selectbox", "form/text", "form/radio",
		"form/password", "form/select", "form/singleselect",
		"form/richtext", "form/iconpicker", "form/datetime", "form/number",
		"form/email", "form/url", "form/ip", "form/color", "form/currency", "form_components")
}
