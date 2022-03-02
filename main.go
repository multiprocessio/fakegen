package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jaswdr/faker"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	faker := faker.New()

	var rows, cols int
	for i, arg := range os.Args {
		var err error
		if arg == "--rows" {
			rows, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				log.Fatal("--rows must be an int. e.g. --rows 10000: " + err.Error())
			}

			i += 1
			continue
		}

		if arg == "--cols" {
			cols, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				log.Fatal("--cols must be an int. e.g. --cols 10000: " + err.Error())
			}

			i += 1
			continue
		}
	}

	if rows < 1 {
		log.Fatal("--rows must be at least 1")
	}

	if cols < 1 {
		log.Fatal("--cols must be at least 1")
	}

	// Populate all the column names
	schema := map[string]func() interface{}{}
	types := []func() interface{}{
		func() interface{} {
			return faker.Lorem().Paragraph(rand.Intn(3) + 1)
		},
		func() interface{} {
			return faker.Int32()
		},
		func() interface{} {
			return strings.Join(faker.Lorem().Words(rand.Intn(10)+1), " ")
		},
		func() interface{} {
			return faker.UInt16()
		},
	}
	for i := 0; i < cols; i += 1 {
		schema[fmt.Sprintf("col%d", i)] = types[rand.Intn(len(types))]
	}

	_, err := os.Stdout.Write([]byte("["))
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	row := map[string]interface{}{}
	for i := 0; i < rows; i += 1 {
		if i > 0 {
			_, err := os.Stdout.Write([]byte(",\n"))
			if err != nil {
				log.Fatal(err)
			}
		}
		for col, generator := range schema {
			row[col] = generator()
		}

		err = encoder.Encode(row)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = os.Stdout.Write([]byte("]"))
	if err != nil {
		log.Fatal(err)
	}
}
