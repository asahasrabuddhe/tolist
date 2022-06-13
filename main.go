package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type List struct {
	ApiVersion string           `yaml:"apiVersion"`
	Kind       string           `yaml:"kind"`
	Items      []map[string]any `yaml:"items"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	in, err := io.ReadAll(reader)
	if err != nil {
		log.Panicln(err)
	}

	manifests := strings.Split(string(in), "---")

	list := &List{
		ApiVersion: "v1",
		Kind:       "List",
	}

	list.Items = make([]map[string]any, len(manifests))

	for i, manifest := range manifests {
		var item map[string]any
		err := yaml.Unmarshal([]byte(manifest), &item)
		if err != nil {
			log.Panicln(err)
		}
		list.Items[i] = item
	}

	out, err := yaml.Marshal(list)
	if err != nil {
		log.Panicln(err)
	}

	_, _ = fmt.Fprintln(os.Stdout, string(out))
}
