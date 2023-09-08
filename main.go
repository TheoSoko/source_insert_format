package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"theosoko/ossInserter/db"

	_ "github.com/go-sql-driver/mysql"
)


var (
	_, file, _, _ = runtime.Caller(0)
	Path         = filepath.Dir(file)
)

func main() {
	db.DbAccess()
	if db.Db == nil {
		panic("wat, db not init")
	}

	defer db.Db.Close()

	sites := db.GetWebsitesByCountry("es")

	formatSitesNewLine(sites, "es")

}

// Takes a list of websites and a country_code, prints the sites on "line by line" format in  new file (inside dir "/out" )
func formatSitesNewLine(sites []string, cc string) {
	fileToWrite := filepath.Join(Path + "/out/" + cc + ".txt")
	fmt.Println("path" , fileToWrite)
	
	file, err := os.Create(fileToWrite)
	if err != nil {
		log.Fatal(err)
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
