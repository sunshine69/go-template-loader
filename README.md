# go-template-loader

# Description

**go-template-loader** is a loader for `html/template` templates that supports [Django](https://www.djangoproject.com/)-like subfolders and loading from a [go.rice](https://github.com/GeertJohan/go.rice) box.

# Installation

```
go get github.com/spatialcurrent/go-template-loader
```

# Usage

**LoadTemplatesFromPaths**

```
import (
  "html/template"
)

import (
  "github.com/spatialcurrent/go-template-loader/tloader"
)

pathsToTemplates := "~/templates" // relative or absolute path to your templates
templateExtensions := []string{"html", "yml", "json", "md"}
templateFunctions := template.FuncMap{} // the functions to be used in the templates
verbose := True
tmpl, err := tloader.LoadTemplatesFromPaths(templates_uri, templateExtensions, templateFunctions, verbose)

```

**LoadTemplatesFromBinary**

This functions load from a [go.rice](https://github.com/GeertJohan/go.rice) box.  Be sure to include `rice.FindBox("...")` in your main package, so go.rice will add the templates to the box.

```
import (
  "html/template"
)

import (
  "github.com/spatialcurrent/go-template-loader/tloader"
)

templatesBox, err := rice.FindBox("templates") // be sure to put this line in your main package
templateFunctions := template.FuncMap{} // the functions to be used in the templates
verbose := True
tmpl, err := tloader.LoadTemplatesFromBinary(templatesBox, templateFunctions, verbose)

```

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-template-loader/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
