package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"theosoko/ossInserter/db"

	"theosoko/ossInserter/ressources"

	_ "github.com/go-sql-driver/mysql"
)

var (
	_, file, _, _ = runtime.Caller(0)
	Path          = filepath.Dir(file)
)

func main() {
	db.DbAccess()
	if db.Db == nil {
		panic("wat, db not init")
	}

	defer db.Db.Close()

	cc := db.GetAvailCountryCodes()
	/*for _, code := range cc {
		code = strings.ToLower(code)
	*/

	sourceTypeIds := db.GetAvailSourceTypes()
	type SourceType struct {
		name string
		id   int
	}
	sourceTypes := []SourceType {}

	for _, strID := range sourceTypeIds {
		id, _ := strconv.Atoi(strID)
		name, ok := ressources.SourceTypesReversedEn[id]
		if ok {
			sourceTypes = append(sourceTypes, SourceType{name: name, id: id})
		}
	}

	for _, code := range cc {
		for _, s := range sourceTypes {
			sites := db.GetWebSites(code, s.id)

			formatSitesNewLine(sites, code, s.name)
		}
	}

}

// Takes a list of websites and a country_code, prints the sites on "line by line" format in  new file (inside dir "/out" )
func formatSitesNewLine(sites []string, cc string, source string) {
	if len(sites) == 0 {
		return
	}

	fileToWrite := filepath.Join(Path + "/out/" + cc + "_" + source + ".txt")
	fmt.Println("path", fileToWrite)

	file, err := os.Create(fileToWrite)
	if err != nil {
		log.Fatal("err at create file", err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, site := range sites {
		_, err := w.WriteString(site + "\n")
		if err != nil {
			fmt.Println("err at file write for "+site+". \n", err.Error())
		}
	}

	w.Flush()
}
