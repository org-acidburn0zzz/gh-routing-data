package main

import (
	"log"
	"strings"
)

// alphabetical order not suitable, so, list explicitly
var regionGroups = []string{"Europe", "Northern Europe", "North America", "Asia", "Other"}
var continentCodeToRegion = map[string]string{"EU": "Europe", "AS": "Asia", "AF": "Other"}

type GraphHopperVersionToRegions struct {
	GraphHopperVersion string `json:"ghVersion"`

	localRootDir  string
	remoteRootDir string

	Regions []*Region `json:"regions"`
}

type Region struct {
	Name           string `json:"name"`
	Title          string `json:"title"`
	Group          string `json:"group"`
	TotalSize      int64  `json:"totalSize"`
	TotalSizeHuman string `json:"totalSizeHuman"`

	DirUrl   string `json:"dirUrl"`
	LocusUrl string `json:"locusUrl"`

	Parts []PartInfo `json:"parts"`
}

type PartInfo struct {
	FileName string `json:"fileName"`
	Index    int    `json:"index"`
	Size     int64  `json:"size"`
}

func getRegionTitle(name string) string {
	switch name {
	case "us-midwest":
		return "US Midwest"
	case "us-northeast":
		return "US Northeast"
	case "us-pacific":
		return "US Pacific"
	case "us-south":
		return "US South"
	case "us-west":
		return "US West"

	case "de-at-ch":
		return "Germany, Austria and Switzerland"
	case "portugal-spain":
		return "Portugal and Spain"
	case "estonia-latvia-lithuania":
		return "Estonia, Latvia and Lithuania"
	case "finland-norway-sweden":
		return "Finland, Norway and Sweden"
	case "al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si":
		return "Albania, Bosnia-Herzegovina, Bulgaria, Croatia, Hungary, Kosovo, Macedonia, Moldova, Montenegro, Romania, Serbia, Slovakia and Slovenia"

	case "bayern-at-cz":
		return "Bayern (Germany), Austria, Czech Republic"
	case "ireland-and-northern-ireland":
		return "Ireland and Northern Ireland"
	}

	index := strings.IndexRune(name, '-')
	if index > 0 {
		return strings.ToUpper(string(name[0])) + name[1:index] + " " + strings.ToUpper(string(name[index+1])) + name[index+2:]
	} else {
		return strings.ToUpper(string(name[0])) + name[1:]
	}
}

var northernEuropeRegions = map[string]string{"iceland": "", "great-britain": "", "sweden": "", "norway": "", "denmark": "", "ireland-and-northern-ireland": ""}
var asiaRegions = map[string]string{"japan": "", "india": "", "china": "", "indonesia": "", "thailand": "", "taiwan": ""}

func getRegionScopeName(name string) string {
	if strings.HasPrefix(name, "us-") || name == "canada" {
		return "North America"
	} else if name == "australia" || name == "new-zealand" || name == "africa" || name == "south-america" || name == "brazil" || name == "central-america" {
		return "Other"
	} else if _, ok := asiaRegions[name]; ok {
		return "Asia"
	} else if _, ok := northernEuropeRegions[name]; ok || strings.HasPrefix(name, "finland") {
		return "Northern Europe"
	} else {
		continent, err := getContinentByCountryName(name)
		if err != nil {
			log.Fatal(err)
		}

		if continent == "" {
			return "Europe"
		}

		result, ok := continentCodeToRegion[continent]
		if !ok {
			log.Fatal("unknown continent: " + continent)
		}
		return result
	}
}