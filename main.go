package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	goccy_json "github.com/goccy/go-json"
	"github.com/jaswdr/faker"
	jsonutil "github.com/multiprocessio/go-json"
	"github.com/xitongsys/parquet-go/writer"
	"github.com/xuri/excelize/v2"
)

var preferredParallelism = runtime.NumCPU() * 2

// faker.Lorem().Word() only seems to have like 180 unique entries.
// So we brought in our own dictionary so that the words are unique
// but not just completely random collections of letters which are not
// realistic in the real world and destroy sorting algorithms
// unfairly.
func randomWord() string {
	return WORDS[rand.Intn(len(WORDS))]
}

var (
	yearsAhead50 = time.Now().Add(time.Hour * 8760 * 50)
	yearsAgo50   = time.Now().Add(-1 * time.Hour * 8760 * 50)
)

func jsonEncode(out io.Writer, generator chan map[string]any, sep string) error {
	isArray := sep == ",\n"
	if isArray {
		encoder := jsonutil.NewGenericStreamEncoder(out, goccy_json.Marshal, isArray)
		defer encoder.Close()

		for row := range generator {
			err := encoder.EncodeRow(row)
			if err != nil {
				return err
			}
		}

		return nil
	}

	w := goccy_json.NewEncoder(out)
	first := false
	for row := range generator {
		if !first {
			_, err := out.Write([]byte(sep))
			if err != nil {
				return err
			}
		}

		err := w.Encode(row)
		if err != nil {
			return err
		}
	}

	return nil
}

func csvEncode(out io.Writer, generator chan map[string]any, sep rune) error {
	w := csv.NewWriter(out)
	defer w.Flush()
	w.Comma = sep
	first := true
	var columns []string
	var rowStrings []string
	for row := range generator {
		if first {
			for col := range row {
				columns = append(columns, col)
				rowStrings = append(rowStrings, "")
			}
			sort.Strings(columns)

			err := w.Write(columns)
			if err != nil {
				return nil
			}

			first = false
		}

		for i, col := range columns {
			s := ""
			switch t := row[col].(type) {
			case time.Time:
				s = t.Format(time.RFC3339)
			default:
				s = fmt.Sprintf("%v", t)
			}
			rowStrings[i] = s
		}

		err := w.Write(rowStrings)
		if err != nil {
			return err
		}
	}

	return nil
}

var errOutMustBeSet = errors.New("--out X, -o X output file setting must be set when using this format")

func excelEncode(out string, generator chan map[string]any) error {
	if out == "" {
		return errOutMustBeSet
	}

	xlsx := excelize.NewFile()
	streamWriter, err := xlsx.NewStreamWriter("Sheet1")
	if err != nil {
		return err
	}

	i := 0
	var columns []string
	var rowValues []any
	for row := range generator {
		i++
		cell, err := excelize.CoordinatesToCellName(1, i)
		if err != nil {
			return err
		}

		if i == 1 {
			for col := range row {
				columns = append(columns, col)
				rowValues = append(rowValues, "")
			}
			sort.Strings(columns)

			// Excelize requires an array of any
			var colsInt []any
			for _, col := range columns {
				colsInt = append(colsInt, col)
			}

			err = streamWriter.SetRow(cell, colsInt)
			if err != nil {
				return nil
			}

			// Increment cell for first non-header row
			i++
			cell, err = excelize.CoordinatesToCellName(1, i)
			if err != nil {
				return err
			}
		}

		for i, col := range columns {
			s := ""
			switch t := row[col].(type) {
			case time.Time:
				s = t.Format(time.RFC3339)
			default:
				s = fmt.Sprintf("%v", t)
			}
			rowValues[i] = s
		}

		err = streamWriter.SetRow(cell, rowValues)
		if err != nil {
			return err
		}
	}

	err = streamWriter.Flush()
	if err != nil {
		return err
	}

	return xlsx.SaveAs(out)
}

func orcEncode(out io.Writer, generator chan map[string]any) error {
	return errors.New("ORC unimplemented")
}

func parquetEncode(out io.Writer, schema map[string]columnKind, generator chan map[string]any) error {
	var columns []string
	for col := range schema {
		columns = append(columns, col)
	}
	sort.Strings(columns)

	parquetSchema := "{\"Tag\": \"name=root\", \"Fields\": [\n"
	for i, col := range columns {
		if i > 0 {
			parquetSchema += ",\n"
		}

		kind := schema[col]
		parquetInfo := map[columnKind]string{
			boolColumn:   "type=BOOLEAN",
			stringColumn: "type=BYTE_ARRAY, convertedtype=UTF8",
			intColumn:    "type=INT64",
			floatColumn:  "type=FLOAT",
			timeColumn:   "type=INT64, convertedtype=TIMESTAMP_MILLIS",
		}[kind]
		col = strings.ReplaceAll(col, "-", "")
		parquetSchema += fmt.Sprintf(`{"Tag": "name=%s, %s"}`, col, parquetInfo)
	}
	parquetSchema += "\n]}"

	w, err := writer.NewParquetWriterFromWriter(out, parquetSchema, int64(preferredParallelism))
	if err != nil {
		return err
	}
	for row := range generator {
		for col, kind := range schema {
			if kind == timeColumn {
				t := row[col].(time.Time)
				row[col] = t.UnixNano() / int64(time.Millisecond)
			}
		}
		err := w.Write(row)
		if err != nil {
			return nil
		}
	}

	return w.WriteStop()
}

func odsEncode(out io.Writer, generator chan map[string]any) error {
	return errors.New("ods unimplemented")
}

type columnKind = uint

const (
	boolColumn columnKind = iota
	stringColumn
	intColumn
	floatColumn
	timeColumn
)

type columnSchema struct {
	generator func() any
	kind      columnKind
}

var nullFrequency = .01

func timeForNull() bool {
	// Don't trust float equality
	if nullFrequency < .00001 {
		return false
	}

	times := 1 / nullFrequency * 1000
	r := rand.Intn(int(math.Ceil(times)))
	n := times - times*nullFrequency
	return float64(r) > n
}

func makeGenerator(rows, cols int) (map[string]columnKind, chan map[string]any) {
	faker := faker.New()

	// Populate all the column names
	schema := map[string]func() any{}
	types := []columnSchema{
		{
			func() any {
				return faker.Lorem().Paragraph(rand.Intn(3) + 1)
			},
			stringColumn,
		},
		{
			func() any {
				return int(faker.Int32())
			},
			intColumn,
		},
		{
			func() any {
				return faker.Float32(6, -100000, 100000)
			},
			floatColumn,
		},
		{
			func() any {
				return faker.Bool()
			},
			boolColumn,
		},
		{
			func() any {
				return faker.Time().TimeBetween(yearsAgo50, yearsAhead50)
			},
			timeColumn,
		},
		{
			func() any {
				return strings.Join(faker.Lorem().Words(rand.Intn(10)+1), " ")
			},
			stringColumn,
		},
		{
			func() any {
				return int(faker.UInt16())
			},
			intColumn,
		},
		{
			func() any {
				return randomWord()
			},
			stringColumn,
		},
	}

	failures := 0
	schemaTypes := map[string]columnKind{}
	for len(schema) < cols {
		// TODO: handle when more than 370k columns
		// requested. Should probably try to combine words
		// instead of reusing them or failing here.
		column := randomWord()
		_, exists := schema[column]
		if exists {
			failures += 1
			if failures > len(WORDS)/10 {
				log.Fatal("Running out of unique entries.")
			}
			continue
		}
		columnSchema := types[rand.Intn(len(types))]
		schema[column] = columnSchema.generator
		schemaTypes[column] = columnSchema.kind
	}

	out := make(chan map[string]any)

	go func() {
		defer close(out)
		for i := 0; i < rows; i += 1 {
			row := map[string]any{}
			for col, generator := range schema {
				if timeForNull() {
					row[col] = nil
				} else {
					row[col] = generator()
				}
			}

			out <- row
		}
	}()

	return schemaTypes, out
}

func formatOrFromFileExtension(format, out string, formatDefault bool) string {
	if out == "" {
		return format
	}

	if !formatDefault {
		return format
	}

	ext := filepath.Ext(out)
	if ext[0] == '.' {
		ext = ext[1:]
	}
	return ext
}

var Version = "latest"

var HELP = `fakegen (Version ` + Version + `) - Single binary CLI for generating a random schema of M columns to populate N rows of data

Usage: fakegen --rows N --cols M > testdata.json
       fakegen --rows N --cols M --out testdata.json
       fakegen --rows N --cols M --format csv > testdata.csv

Flags:

  -h, --help		Dump help text and exit
  -v, --version		Dump version and exit
  -o, --out		Specify file to write to. If empty, writes to stdout
  -f, --format		Specify output format. Inferred by --out flag if present and --format flag empty
  -r, --rows		Number of rows to generate
  -c, --cols		Number of columns to generate

Supported formats:

  json		Array of JSON objects
  jsonl		Newline separated JSON objects
  cjson		Concatenated JSON objects
  xlsx
  csv
  tsv

See the repo for more details: https://github.com/multiprocessio/fakegen.`

func _main() error {
	log.SetFlags(0)
	rand.Seed(time.Now().UnixNano())

	var rows, cols int
	format := "json"
	formatDefault := true
	out := ""
	for i, arg := range os.Args {
		var err error
		if arg == "--rows" || arg == "-r" {
			rows, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				return errors.New("--rows must be an int. e.g. --rows 10000: " + err.Error())
			}

			i++
			continue
		}

		if arg == "--cols" || arg == "-c" {
			cols, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				return errors.New("--cols must be an int. e.g. --cols 10000: " + err.Error())
			}

			i++
			continue
		}

		if arg == "-h" || arg == "--help" {
			log.Println(HELP)
			return nil
		}

		if arg == "-v" || arg == "--version" {
			log.Println("fakegen " + Version)
			return nil
		}

		if arg == "-f" || arg == "--format" {
			format = os.Args[i+1]
			formatDefault = false
			i++
			continue
		}

		if arg == "-o" || arg == "--out" {
			out = os.Args[i+1]
			i++
			continue
		}

		if arg == "-n" || arg == "--null-frequency" {
			nullFrequency, err = strconv.ParseFloat(os.Args[i+1], 32)
			if err != nil {
				return errors.New("-n, --null-frequency must be a float. e.g --null-frequency .001: " + err.Error())
			}

			i++
			continue
		}
	}

	if rows < 1 {
		return errors.New("--rows must be at least 1")
	}

	if cols < 1 {
		return errors.New("--cols must be at least 1")
	}

	schema, generator := makeGenerator(rows, cols)

	format = formatOrFromFileExtension(format, out, formatDefault)
	w := os.Stdout
	if out != "" {
		// Truncates the file if it exists, creates directory if it doesn't exist
		base := filepath.Dir(out)
		err := os.Mkdir(base, os.ModePerm)
		// Other comparisons like os.IsNotExist(err) and errors.Is(err, os.ErrNotExist) don't work.
		if err != nil && !strings.HasSuffix(err.Error(), ": file exists") {
			return err
		}
		w, err = os.OpenFile(out, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		defer w.Close()
	}
	switch format {
	case "json":
		return jsonEncode(w, generator, ",\n")
	case "jsonl":
		return jsonEncode(w, generator, "")
	case "csv":
		return csvEncode(w, generator, rune(','))
	case "tsv":
		return csvEncode(w, generator, rune('\t'))
	case "parquet":
		return parquetEncode(w, schema, generator)
	case "orc":
		return orcEncode(w, generator)
	case "xlsx":
		w.Close()
		return excelEncode(out, generator)
	case "ods":
		return odsEncode(w, generator)
	}

	return fmt.Errorf("Unimplemented format: " + format)
}

func main() {
	err := _main()
	if err != nil {
		log.Fatal(err)
	}
}
