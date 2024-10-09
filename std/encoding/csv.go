package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func main() {
	b := bytes.NewBuffer(nil)
	w := csv.NewWriter(b)
	w.Write([]string{"user1", "18", `data: "example"`})
	w.Write([]string{"user2", "20", `data:"test xxx"`})
	w.Flush()
	fmt.Println(b)
	r := csv.NewReader(b)
	recs, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	for _, rec := range recs {
		fmt.Println(strings.Join(rec, "|"))
	}
}
