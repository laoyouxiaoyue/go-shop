package config

import (
	"fmt"
	"strings"
	"text/template"
)

type TemplateManager interface {
	GetTemplate(templateID string) (TemplateDetail, error)
	GetAllTemplates() map[string]TemplateDetail
	RenderContent(templateID, code string) (string, error)
}
type TemplateDetail struct {
	ID               string `mapstructure:"id"`
	CodeLength       int    `mapstructure:"code_length"`
	ExpireSeconds    int    `mapstructure:"expire_seconds"`
	MaxDailySends    int    `mapstructure:"max_daily_sends"`
	ProviderTemplate string `mapstructure:"provider_template"`
	ContentTemplate  string `mapstructure:"content_template"`
}

type YamlTemplateManager struct {
	Templates map[string]TemplateDetail `mapstructure:"templates"`
}

func (ym *YamlTemplateManager) GetTemplate(templateID string) (TemplateDetail, error) {
	if template, exists := ym.Templates[templateID]; exists {
		return template, nil
	}
	return TemplateDetail{}, fmt.Errorf("template not found: %s", templateID)
}

func (ym *YamlTemplateManager) GetAllTemplates() map[string]TemplateDetail {
	return ym.Templates
}
func (ym *YamlTemplateManager) RenderContent(templateID, code string) (string, error) {
	templateDetail, err := ym.GetTemplate(templateID)
	if err != nil {
		return "", err
	}
	expireSeconds := templateDetail.ExpireSeconds
	data := struct {
		Code          string
		ExpireMinutes int
		ExpireSeconds int
	}{
		Code:          code,
		ExpireMinutes: expireSeconds / 60,
		ExpireSeconds: expireSeconds,
	}

	tmpl, err := template.New("sms").Parse(templateDetail.ContentTemplate)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, data); err != nil {
		return "", err
	}

	return result.String(), nil
}
