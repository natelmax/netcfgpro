package generators

import (
	"fmt"
	"io"
	"netcfgpro-go/internal/models"
	"netcfgpro-go/internal/templates"
	"text/template"
)

// Generate creates a network device configuration from a DeviceConfig struct.
// It uses the embedded templates to generate the configuration and writes it to the provided io.Writer.
func Generate(config *models.DeviceConfig, w io.Writer) error {
	if config == nil {
		return fmt.Errorf("device config cannot be nil")
	}

	// Validate the configuration before generating.
	if err := config.Validate(); err != nil {
		return fmt.Errorf("invalid device configuration: %w", err)
	}

	templateName := fmt.Sprintf("%s.tmpl", config.Vendor)

	t, err := template.ParseFS(templates.TemplateFS, templateName)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templateName, err)
	}

	err = t.Execute(w, config)
	if err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	return nil
}
