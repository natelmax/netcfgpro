// package templates provides access to the embedded template files.
package templates

import "embed"

// TemplateFS holds the embedded template files for the application.
// The go:embed directive loads the files matching *.tmpl into this variable.
//
//go:embed *.tmpl
var TemplateFS embed.FS
