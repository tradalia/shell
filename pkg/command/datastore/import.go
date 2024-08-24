//=============================================================================
/*
Copyright © 2022 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package datastore

import (
	"bufio"
	"github.com/spf13/cobra"
	"log"
	"os"
)

//=============================================================================

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports a file into the datastore",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		importFile()
	},
}

//=============================================================================

var source string
var symbol string
var file   string
var ftype  string

//=============================================================================

func init() {
	importCmd.Flags().StringVarP(&source, "source", "S", "", "Target data source where data will be imported (required)")
	importCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Target symbol where data will be imported (required)")
	importCmd.Flags().StringVarP(&file,   "file",   "f", "", "File containing prices for symbols (required)")
	importCmd.Flags().StringVarP(&ftype,  "type",   "t", "", "File type. Can be 'ts' for TradeStation CSV format  (required)")
	importCmd.MarkFlagRequired("source")
	importCmd.MarkFlagRequired("symbol")
	importCmd.MarkFlagRequired("file")
	importCmd.MarkFlagRequired("type")
}

//=============================================================================

func importFile() {
	log.Println("Source : " + source)
	log.Println("Symbol : " + symbol)
	log.Println("File   : " + file)
	log.Println("Type   : " + ftype)

	f, err := os.Open(file)

	if err != nil {
		log.Println("Cannot open data file for reading --> " + err.Error())
		return
	}

	defer f.Close()

	scanner   := bufio.NewScanner(f)
	firstLine := true

	for scanner.Scan() {
		if firstLine {
			err = handleHeader(scanner.Text())
			firstLine = false
		} else {
			err = handleData(scanner.Text())
		}

		if err != nil {
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Cannot scan file: " + file + " (cause is: " + err.Error() + " )")
	} else {
		//tool.DoPut(url, &data, &data)
		//fmt.Println("Instrument added with id=", data.Id)
	}
}

//=============================================================================

func handleHeader(line string) error {
	//tokens := strings.Split(line, "|")
	return nil
}

//=============================================================================

func handleData(line string) error {
	//tokens := strings.Split(line, "|")
	return nil
}

//=============================================================================
