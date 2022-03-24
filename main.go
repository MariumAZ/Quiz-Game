package main

import (
    "encoding/csv"
	"os"
    "log"
    "fmt"
    "io"
)

func main() {
     f, err := os.Open("data.csv")
     if err != nil {
        log.Fatal(err)
     }
      // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    for {
        rec, err := csvReader.Read()
        if len(rec) > 1 {
            Q:= rec[:1]
            A := rec[len(rec)-1]
            fmt.Printf("%+v\n", Q)
            fmt.Printf("%+v\n", A)
        }
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
    }
}


