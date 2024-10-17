package tloader

import (
  "errors"
  "log"
  "html/template"
  "sort"
  "strings"
)

import (
  "github.com/GeertJohan/go.rice"
)

func LoadPathsFromBox(templatesBox *rice.Box, current *rice.File, parent string, paths []string) ([]string, error) {

  current_stat, err := current.Stat()
  if err != nil {
    return nil, errors.New("Error: Could not stat current directory.")
  }

  basepath := current_stat.Name()
  if len(parent) > 0 {
    basepath = parent + "/" + basepath
  }

  infos, err := current.Readdir(0)
  current.Close()
  if err != nil {
    return nil, errors.New("Error: Could not read directory names from rice box for templates.")
  }

  for _, info := range infos {
    if info.IsDir() {
      next, err := templatesBox.Open(basepath+"/"+info.Name())
      if err != nil {
        return nil, err
      }
      paths, err = LoadPathsFromBox(templatesBox, next, basepath, paths)
      if err != nil {
        return nil, err
      }
    } else {
      if basepath == "" || basepath == "/" {
        paths = append(paths, info.Name())
      } else {
        paths = append(paths, basepath+"/"+info.Name())
      }
    }
  }

  return paths, nil
}

func LoadTemplatesFromBinary(templatesBox *rice.Box, templateFunctions template.FuncMap, verbose bool) (* template.Template, error) {

  tmpl := template.New("blank.tpl.html").Funcs(templateFunctions)

  root, err := templatesBox.Open("/")
  if err != nil {
    return nil, errors.New("Error: Could not open root of rice box for templates.")
  }

  stat, err := root.Stat()
  if err != nil {
  	return nil, errors.New("Error: Could not stat root of rice box for templates.")
  }

	if !stat.IsDir() {
		return nil, errors.New("Error: Root of rice box for templates is not a directory.")
	}

  paths, err := LoadPathsFromBox(templatesBox, root, "", []string{})
  if err != nil {
    return nil, errors.New("Error: Could not load paths.")
  }
  sort.Strings(paths)

  if verbose {
    log.Println("Paths: "+strings.Join(paths, "; "))
  }

  for _, path := range paths {
    content, err := templatesBox.String(path)
    if err != nil {
      return nil, errors.New("Error: Could not read content from template at '"+path+"'.")
    }
    if verbose {
      log.Println("Registering template for path '"+path+"'.")
    }
    tmpl, err = tmpl.New(path).Parse(content)
    if err != nil {
      return nil, errors.New("Error: Could not parse content from template at '"+path+"'.")
    }
  }

  return tmpl, err
}
