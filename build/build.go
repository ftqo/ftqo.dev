package build

import _ "embed"

//go:embed files.zip
var F []byte

var TmpDir = "tmp"
var AssetsDir = "assets"
var BuildDir = "build"
var TemplateDir = "templates"
var StaticBundleName = "files.zip"
