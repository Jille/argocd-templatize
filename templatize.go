// Binary templatize finds ./*.json and ./*.yaml and passes them through text/template with the environment variables named TPZ_*.
// TPZ_SOME_THING is available as {{.SomeThing}} in the template.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var (
	prefix = flag.String("prefix", "TPZ_", "Include all environment variables starting with this prefix")
)

func main() {
	flag.Parse()
	if err := do(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func getVariables() map[string]string {
	vars := map[string]string{}
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		k := kv[0]
		if !strings.HasPrefix(k, *prefix) {
			continue
		}
		k = strings.TrimPrefix(k, *prefix)
		k = strcase.ToCamel(strings.ToLower(k))
		vars[k] = kv[1]
	}
	return vars
}

func do(out io.Writer) error {
	vars := getVariables()

	dh, err := os.Open(".")
	if err != nil {
		return fmt.Errorf("Failed to open directory .: %v", err)
	}
	names, err := dh.Readdirnames(0)
	if err != nil {
		return fmt.Errorf("Failed to read directory .: %v", err)
	}
	dh.Close()
	for _, n := range names {
		if !strings.HasSuffix(n, ".json") && !strings.HasSuffix(n, ".yaml") {
			continue
		}
		t, err := template.ParseFiles(n)
		if err != nil {
			return fmt.Errorf("Failed to parse %q: %v", n, err)
		}
		t = t.Option("missingkey=error")
		if err := t.Execute(out, vars); err != nil {
			return fmt.Errorf("Failed to execute %q: %v", n, err)
		}
		fmt.Fprintf(out, "\n\n---\n")
	}
	return nil
}
