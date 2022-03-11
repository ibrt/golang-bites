package internal

import (
	_ "embed" // embed
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
)
