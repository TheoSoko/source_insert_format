package main

import (
	"fmt"
	"theosoko/ossInserter/db"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"bufio"
)

//var source = "Partis politiques"

func main() {
	db.DbAccess()
	if db.Db == nil {
		panic("wat, db not init")
	}

	defer db.Db.Close()

	sites := db.GetWebsitesByCountry("fr")

	//fmt.Println("sites : ", sites)

	formatSitesNewLine(sites, "fr")

}

// Fetches the webistes matching cc (country code), formats them on "line by line" format, and prints them on a new file (inside dir "/out" )
func formatSitesNewLine(sites []string, cc string) {
	file, err := os.Create("C:/Users/theos/Documents/Zemus Projects/sources d'actus/ossInserter/out/"+cc+".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()


	w := bufio.NewWriter(file)
	for _, site := range sites {
		_, err := w.WriteString(site+"\n")
		if err != nil {
			fmt.Println("err at file write for " + site + ". \n", err.Error())
		}
	}

	w.Flush()
}


