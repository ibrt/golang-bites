package internal

import (
	"embed"
)

const (
	// ExampleDirPathPrefix is the path prefix for ExampleDirAssetFS.
	ExampleDirPathPrefix = "assets/example-dir"
)

// Embedded assets.
var (
	//go:embed assets/enum.go.gotpl
	EnumGoTemplateAsset string

	//go:embed assets/enum_test.go.gotpl
	EnumTestGoTemplateAsset string

	//go:embed assets/numeric.go.gotmpl
	NumericGoTemplateAsset string

	//go:embed assets/numeric_test.go.gotmpl
	NumericTestGoTemplateAsset string

	//go:embed assets/example-dir
	ExampleDirAssetFS embed.FS
)
