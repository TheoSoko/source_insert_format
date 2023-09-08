package main

import (
	"net/http"
	"time"	
	"theosoko/ossInserter/db"
)

// Gets all the country codes and souce type available, and creates an OSS index for each source type in each country.
// Then populates the indices with corresponding webistes url (fetched from DB)
func ossCreateAll() {
	countryCodes := db.GetAllCountryCodes()

	for _, country := range countryCodes {
		for source, sourceId := range sourceTypes {
			indexName := country + " - " + string(source)
			ossCreateIndex(indexName)

			websites := db.GetWebSites(country, sourceId)

			ossInsertWebsites(indexName, websites)
		}
	}
}

func ossCreateIndex(indexName string) {
	client := http.Client{Timeout: time.Second * 11}

	client.Get("http://uadumtwwaat")
}

func ossInsertWebsites(indexName string, websites []string) {
	client := http.Client{Timeout: time.Second * 11}

	client.Get("http://idratherdoanythihng")
}
