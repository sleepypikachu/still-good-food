package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type NutritionInfo struct {
	Kcal      string
	Fat       string
	Saturates string
	Carbs     string
	Sugars    string
	Fibre     string
	Protein   string
	Salt      string
}

type Recipe struct {
	Name        string
	Ingredients []string
	Steps       []string
	Yield       string
	Difficulty  string
	Preparation string
	Cook        string
	Nutrition   NutritionInfo
}

func Scrape(recipe string) (Recipe, error) {
	doc, err := goquery.NewDocument(recipe)
	if err != nil {
		return Recipe{}, err
	}

	name := doc.Find(".recipe-header__title").First().Text()

	var ingredients []string
	doc.Find(".ingredients-list__item").Each(func(i int, s *goquery.Selection) {
		s.Find("span").Remove()
		ingredients = append(ingredients, strings.TrimSpace(s.Text()))
	})

	var steps []string
	doc.Find(".method__item").Each(func(i int, s *goquery.Selection) {
		steps = append(steps, strings.TrimSpace(s.Text()))
	})

	yield := extract("recipeYield", doc)

	difficulty := strings.TrimSpace(doc.Find("section.recipe-details__item--skill-level").Text())

	preparationSpan := doc.Find(".recipe-details__cooking-time-prep")
	preparationSpan.Find("strong").Remove()

	preparation := strings.TrimSpace(preparationSpan.Text())

	cookSpan := doc.Find(".recipe-details__cooking-time-cook")
	cookSpan.Find("strong").Remove()

	cook := strings.TrimSpace(cookSpan.Text())

	nutrition := NutritionInfo{
		Kcal:      extract("calories", doc),
		Fat:       extract("fatContent", doc),
		Saturates: extract("saturatedFatContent", doc),
		Sugars:    extract("sugarContent", doc),
		Fibre:     extract("fiberContent", doc),
		Protein:   extract("proteinContent", doc),
		Salt:      extract("sodiumContent", doc),
	}

	r := Recipe{
		Name:        name,
		Ingredients: ingredients,
		Steps:       steps,
		Yield:       yield,
		Difficulty:  difficulty,
		Preparation: preparation,
		Cook:        cook,
		Nutrition:   nutrition,
	}

	return r, nil
}

func extract(itemprop string, doc *goquery.Document) string {
	return strings.TrimSpace(doc.Find(fmt.Sprintf("span[itemprop='%s']", itemprop)).Text())
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <url>\n", os.Args[0])
		return
	}

	r, err := Scrape(os.Args[1])
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)
}
