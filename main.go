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

//var source = "Partis politiques"

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

	//fmt.Println("sites : ", sites)

	formatSitesNewLine(sites, "es")

}

// Fetches the webistes matching cc (country code), formats them on "line by line" format, and prints them on a new file (inside dir "/out" )
func formatSitesNewLine(sites []string, cc string) {
	fmt.Println("path" , filepath.Join(Path + "/out/" + cc + ".txt"))
	file, err := os.Create(filepath.Join(Path + "/out/" + cc + ".txt"))
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
