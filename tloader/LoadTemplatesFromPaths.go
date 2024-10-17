package tloader

import (
	"errors"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/ttacon/chalk"
)

func LoadTemplatesFromPaths(pathsToTemplates string, templateExtensions []string, templateFunctions template.FuncMap, verbose bool) (*template.Template, error) {

	template_names := make([]string, 0)
	template_paths := make([]string, 0)
	for _, x := range strings.Split(pathsToTemplates, ":") {
		collected_names, collected_paths, err := CollectTemplatePaths("", x, templateExtensions, true, []string{}, []string{})
		if verbose {
			log.Println(chalk.Green, "Names: ", collected_names, chalk.Reset)
			log.Println(chalk.Green, "Paths: ", collected_paths, chalk.Reset)
		}
		if err != nil {
			log.Println(chalk.Red, "Error: Could Not Collect templates files from ", x, chalk.Reset)
			return nil, err
		}
		template_names = append(template_names, collected_names...)
		template_paths = append(template_paths, collected_paths...)
	}

	tmpl := template.New("blank.tpl.html").Funcs(templateFunctions)
	for i := 0; i < len(template_names); i++ {
		template_content, err := os.ReadFile(template_paths[i])
		if err != nil {
			return nil, errors.New("Error: Could not parse template at path " + template_paths[i] + " .")
		}
		tmpl, err = tmpl.New(template_names[i]).Parse(string(template_content))
		if err != nil {
			return tmpl, err
		}
	}

	return tmpl, nil
}
