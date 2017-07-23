package tloader

import (
  "errors"
  "path"
  "strings"
  "os"
)

func CollectTemplatePaths(basename string, basepath string, extensions []string, recursive bool, names []string, paths []string) ([]string, []string, error) {

  cwd, err := os.Open(basepath)
  if err != nil {
    return nil, nil, errors.New("Error: Could not open directory at " + basepath + ".")
  }

  files_all, err := cwd.Readdir(0)
  if err != nil {
    return nil, nil, errors.New("Error: Could not read directory at " + basepath + ".")
  }

	for _ , f := range files_all {
    if f.IsDir() {
      if recursive {
        basename_new := ""
        if basename == "" {
          basename_new = f.Name()
        } else {
          basename_new = basename+"/"+f.Name()
        }
        names, paths, err = CollectTemplatePaths(basename_new, basepath+"/"+f.Name(), extensions, recursive, names, paths)
        if err != nil {
          return nil, nil, err
        }
      }
    } else {
      filename := path.Base(f.Name())
      valid := false
      for _ , ext := range extensions {
        valid = strings.HasSuffix(filename, "."+ext)
        if valid {
          break
        }
      }

      if valid {
        if basename == "" {
          names = append(names, f.Name())
        } else {
          names = append(names, basename+"/"+f.Name())
        }
        if basepath == "" {
          paths = append(paths, f.Name())
        } else {
          paths = append(paths, basepath+"/"+f.Name())
        }
      }
    }
  }

  return names, paths, err
}
