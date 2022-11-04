package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var normalMode = os.FileMode(0644)

func main() {
	log.SetFlags(log.Lshortfile)
	cf, err := os.Create("run.go")
	if err != nil {
		log.Fatal(err)
	}
	defer cf.Close()
	var data runData

	files, err := filepath.Glob("day[0-9][0-9]/part[1-2]*.go")
	if err != nil {
		panic(err)
	}
	sort.Strings(files)
	var previousDay = -1
	for _, file := range files {
		day, err := strconv.Atoi(file[3:5])
		if err != nil {
			panic(err)
		}
		part, err := strconv.Atoi(file[10:11])
		if err != nil {
			panic(err)
		}
		i := strings.Index(file[11:], ".")
		variant := file[11 : 11+i]
		if day != previousDay {
			previousDay = day
			data.Days = append(data.Days, runDataDay{Day: day})
		}
		currentDay := &data.Days[len(data.Days)-1]
		currentPart := runDataPart{Part: part, Variant: variant}
		currentDay.Parts = append(currentDay.Parts, currentPart)
	}
	t := template.Must(template.New("run").Parse(runTemplate))
	if err := t.Execute(cf, data); err != nil {
		log.Fatal(err)
	}
	// This is so that download input command does not fail on missing days	/
	// TODO: This really does not belong here and should be move with the downloading code
	for i := 1; i <= data.Days[len(data.Days)-1].Day; i++ {
		if err := os.Mkdir(fmt.Sprintf("day%02d", i), normalMode); err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
	}
}
