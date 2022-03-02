package main

import (
	"os"
	"time"
	"encoding/json"
	"log"
	"rand"

	"github.com/jaswdr/faker"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var rows, columns int
	for i, arg := range os.Args {
		var err error
		if arg == "--rows" {
			rows, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				log.Fatal("--rows must be an int. e.g. --rows 10_000")
			}

			if rows <= 1 {
				log.Fatal("--rows must be at least 1")
			}

			i += 1
			continue
		}

		if arg == "--cols" {
			cols, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				log.Fatal("--cols must be an int. e.g. --cols 10_000")
			}

			if cols <= 1 {
				log.Fatal("--cols must be at least 1")
			}

			i += 1
			continue
		}
	}

	// Populate all the column names
	var schema map[string]func()interface{}
	types := []func()interface{}{
		faker.Lorem().Paragraph,
		faker.Int32,
		faker.Lorem().Word,
		faker.Lorem().Image,
		faker.Lorem().Food,
		faker.Lorem().Color,
		faker.Lorem().Car,
	}
outer:
	for {
		if len(schema) == cols {
			break
		}

		w := faker.Lorem().Word()
		_, exists := keys[w]
		if !exists {
			schema[w] = types[rand.Intn(len(types))]
		}
	}

	_, err := os.Stdout.Write([]byte("]"))
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.Encode(os.Stdout)
	row := map[string]interface{}{}
	for i := 0; i < rows; i += 1 {
		if i > 0 {
			_, err := os.Stdout.Write([]byte("]"))
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

	_, err := os.Stdout.Write([]byte("]"))
	if err != nil {
		log.Fatal(err)
	}
}
