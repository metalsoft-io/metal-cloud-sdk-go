//This helper tool generates the exports functions (not all necessarily) by looking at all the functions with the first param with an 'id' and
//generates two functions: a normal function that gets an int as first param  and a byLabel that gets a label
//if there is any reference to an id in the rest of the params, int will be used

// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type funcDef struct {
	Comment        string
	FunctionName   string
	FirstParamName string
	RestParams     string
	ReturnType     string
}

const funcTemplate = `
//{{capitalize .FunctionName}} {{.Comment}}
func (c *Client) {{capitalize .FunctionName}}({{.FirstParamName}} int{{replaceIdWithInt .RestParams}}) {{.ReturnType}}{
	return c.{{.FunctionName}}({{.FirstParamName}}{{extractParamNames .RestParams}})
}

//{{capitalize .FunctionName}}ByLabel {{.Comment}}
func (c *Client) {{capitalize .FunctionName}}ByLabel({{removeIDSuffix .FirstParamName}}Label string{{replaceIdWithInt .RestParams}}) {{.ReturnType}}{
	return c.{{.FunctionName}}({{removeIDSuffix .FirstParamName}}Label{{extractParamNames .RestParams}})
}
`
const funcRegex = `\/\/[a-zA-Z0-9_]+\s(.*)\nfunc\s\(c\s\*Client\)\s([a-zA-Z0-9_]+)\(([a-zA-Z0-9_]+)\sid([a-zA-Z0-9,\s_\*\[\]]*)\)\s*(.*)\s*{`

func extractParamNames(s string) string {
	paramNames := []string{}
	for _, pair := range strings.Split(s, ",") {
		paramNames = append(paramNames, strings.Split(strings.Trim(pair, " "), " ")[0])
	}
	return strings.Join(paramNames, ",")
}

func replaceIdWithInt(s string) string {
	return strings.ReplaceAll(s, " id", " int")
}

func capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

func removeIDSuffix(s string) string {
	return s[:len(s)-2]
}

func genFuncDefFromTemplate(f funcDef) (string, error) {
	var buf bytes.Buffer

	funcMap := map[string]interface{}{
		"capitalize":        capitalize,
		"removeIDSuffix":    removeIDSuffix,
		"extractParamNames": extractParamNames,
		"replaceIdWithInt":  replaceIdWithInt,
	}

	tmpl, err := template.New("intFunc").Funcs(funcMap).Parse(funcTemplate)
	if err != nil {
		return "", err
	}

	tmpl.Execute(&buf, f)
	return buf.String(), nil
}

//returns the function with an int as an ID instead of id param
func generateFuncDef(s string) (string, error) {

	r := regexp.MustCompile(funcRegex)

	var sb strings.Builder

	ml := r.FindAllStringSubmatch(s, -1)

	for _, m := range ml {

		f := funcDef{
			Comment:        m[1], //the comment
			FunctionName:   m[2], //the function name e.g. DriveArray
			FirstParamName: m[3], //the param name e.g infrastructureID
			RestParams:     m[4], //the rest of the params of the function
			ReturnType:     m[5], //the return type e.g. (infrastructure, error)

		}

		s, err := genFuncDefFromTemplate(f)
		if err != nil {
			return "", err
		}

		sb.WriteString(s)

	}
	return sb.String(), nil
}

func main() {

	var inputFilePath = flag.String("input", os.Getenv("GOFILE"), "Input file path")
	var fileSuffix = flag.String("suffix", "_exports", "Suffix to append to file name before .go")
	var outputFilePath = fmt.Sprintf("%s%s.go", strings.TrimRight(*inputFilePath, ".go"), *fileSuffix)
	var packageToUse = flag.String("package", os.Getenv("GOPACKAGE"), "package to use in generated file")

	flag.Parse()

	s, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(-1)
	}

	g, err := generateFuncDef(string(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(-1)
	}

	var sb strings.Builder
	sb.WriteString("// Code generated by gen_exports.go DO NOT EDIT\n\n")
	sb.WriteString(fmt.Sprintf("package %s\n", *packageToUse))

	sb.WriteString(g)

	ioutil.WriteFile(outputFilePath, []byte(sb.String()), 0644)

}
