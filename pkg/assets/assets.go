package assets

import "embed"

//go:embed web
//go:embed web/_next/static
//go:embed web/_next/static/chunks/pages/*.js
//go:embed web/_next/static/*/*.js
//go:embed web/_next/static/*/*.css
var Storage embed.FS
