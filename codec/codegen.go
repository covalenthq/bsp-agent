/* Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

/* Code Generation Tool for Go-Avro
codegen allows to automatically create Go structs based on defined Avro schema.Usage:

go run codegen.go --schema foo.avsc --schema bar.avsc --out foo.go
Command line flags:
--schema - absolute or relative path to Avro schema file. Multiple of those are allowed but at least one is required.
--out - absolute or relative path to output file. All directories will be created if necessary. Existing file will be truncated.
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/elodina/go-avro"
)

type schemas []string

func (i *schemas) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *schemas) Set(value string) error {
	*i = append(*i, value)

	return nil
}

var schema schemas
var output = flag.String("out", "", "Output file name.")

func main() {
	parseAndValidateArgs()

	var schemas []string
	for _, schema := range schema {
		contents, err := ioutil.ReadFile(schema)
		checkErr(err)
		schemas = append(schemas, string(contents))
	}

	gen := avro.NewCodeGenerator(schemas)
	code, err := gen.Generate()
	checkErr(err)

	createDirs()
	err = ioutil.WriteFile(*output, []byte(code), 0664)
	checkErr(err)
}

func parseAndValidateArgs() {
	flag.Var(&schema, "schema", "Path to avsc schema file.")
	flag.Parse()

	if len(schema) == 0 {
		fmt.Println("At least one --schema flag is required.")
		os.Exit(1)
	}

	if *output == "" {
		fmt.Println("--out flag is required.")
		os.Exit(1)
	}
}

func createDirs() {
	index := strings.LastIndex(*output, "/")
	if index != -1 {
		path := (*output)[:index]
		err := os.MkdirAll(path, 0777)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
