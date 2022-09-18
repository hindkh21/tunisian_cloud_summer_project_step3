package main

import (
	
	"log"
	"fmt"
	"github.com/gocolly/colly"
)

type Secteur_Et_URL struct {
	Secteur_Full_URL string `json:"secteur_url"`
	Secteur_Nom      string `json:"secteur_nom"`
}

type Domaine_Et_URL struct {
	domaine_Full_URL string `json:"domaine_url"`
	domaine_Nom      string `json:"domaine_nom"`
}

func main() {

	var DData1 Secteur_Et_URL
	var DData2 Domaine_Et_URL
	list_of_Secteur_Et_URL := make([]Secteur_Et_URL, 0)
	list_of_domain := make([]Domaine_Et_URL, 0)

	collector_secteur := colly.NewCollector()
	

	collector_secteur.OnHTML("a.stretched-link.text-center", func(element *colly.HTMLElement) {
		DData1.Secteur_Nom = element.Text
		list_of_Secteur_Et_URL = append(list_of_Secteur_Et_URL, DData1)

	})
	i := 0
	collector_secteur.OnHTML("h3", func(element *colly.HTMLElement) {
		DData1.Secteur_Full_URL= "https://www.goafricaonline.com" + element.ChildAttr("a", "href")
		list_of_Secteur_Et_URL[i].Secteur_Full_URL = "https://www.goafricaonline.com" + element.ChildAttr("a", "href")
		i++
	})

	collector_secteur.OnError(func(r *colly.Response, err error) {
		log.Fatal("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	collector_domaine.OnError(func(r *colly.Response, err error) {
		log.Fatal("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	
	collector_domaine.OnHTML("h3", func(element *colly.HTMLElement) {
		DData2.domaine_Nom = element.Text
		list_of_domain = append(list_of_domain, DData2)

	})
	j := 0
	collector_domaine.OnXML("/html/body/div[6]/div[2]/div/div[5]/div/div[*]", func(element *colly.XMLElement) {
		list_of_domain[j].domaine_Full_URL = element.ChildAttr("a", "href")
		j++
	})

	collector_secteur.Visit("https://www.goafricaonline.com/annuaire")
	
	collector_domaine.Visit(list_of_Secteur_Et_URL[0].Secteur_Full_URL)
	
	for k:=0;k<len(list_of_Secteur_Et_URL);k++{
		fmt.Printf("le nom du secteur est %s",list_of_Secteur_Et_URL[k].Secteur_Nom)
		fmt.Println()
		fmt.Printf("l'url du secteur est %s",list_of_Secteur_Et_URL[k].Secteur_Full_URL)
		fmt.Println()
		fmt.Println("=====================================================")
//##############################################################################################################




//##############################################################################################################3

	}

}
