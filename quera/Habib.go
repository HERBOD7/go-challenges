package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	seasons := map[string]map[string]string{
		"SPRING": {
			"COAT":   "may",
			"SHIRT":  "definitely",
			"PANTS":  "definitely",
			"CAP":    "may",
			"JACKET": "",
		},
		"SUMMER": {
			"COAT":   "",
			"SHIRT":  "definitely",
			"PANTS":  "definitely",
			"CAP":    "definitely",
			"JACKET": "",
		},
		"FALL": {
			"COAT":   "may",
			"SHIRT":  "definitely",
			"PANTS":  "definitely",
			"CAP":    "may",
			"JACKET": "",
		},
		"WINTER": {
			"COAT":   "may",
			"SHIRT":  "definitely",
			"PANTS":  "definitely",
			"CAP":    "",
			"JACKET": "may",
		},
	}

	collections := make(map[string][]string)
	collections["COAT"] = []string{}
	collections["SHIRT"] = []string{}
	collections["PANTS"] = []string{}
	collections["CAP"] = []string{}
	collections["JACKET"] = []string{}
	currentSeason := ""
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 6; i++ {
		scanner.Scan()
		if i < 5 {
			value := scanner.Text()
			result := strings.Split(value, ":")
			color := strings.Trim(result[1], " ")
			collectionColor := strings.Split(color, " ")
			collections[result[0]] = collectionColor
		} else {
			currentSeason = scanner.Text()
		}
	}

	availableCloth := seasons[currentSeason]

	shirtsAndPantsOutcomes := []string{}
	shirtsAndPantsAndCapOutcomes := []string{}
	for _, shirt := range collections["SHIRT"] {
		for _, pants := range collections["PANTS"] {
			if availableCloth["CAP"] == "definitely" {
				for _, cap := range collections["CAP"] {
					result := ""
					result = "SHIRT: " + shirt + " PANTS: " + pants + " CAP: " + cap
					shirtsAndPantsOutcomes = append(shirtsAndPantsOutcomes, result)
				}
			} else if availableCloth["CAP"] == "may" {
				result := ""
				result = "SHIRT: " + shirt + " PANTS: " + pants
				shirtsAndPantsOutcomes = append(shirtsAndPantsOutcomes, result)
				for _, cap := range collections["CAP"] {
					capResult := ""
					capResult = "SHIRT: " + shirt + " PANTS: " + pants + " CAP: " + cap
					shirtsAndPantsAndCapOutcomes = append(shirtsAndPantsAndCapOutcomes, capResult)
				}
			} else {
				result := ""
				result = "SHIRT: " + shirt + " PANTS: " + pants
				shirtsAndPantsOutcomes = append(shirtsAndPantsOutcomes, result)
			}
		}
	}

	jacketOutcomes := []string{}
	if availableCloth["JACKET"] == "may" {
		for _, shirtAndPants := range shirtsAndPantsOutcomes {
			for _, jacket := range collections["JACKET"] {
				result := ""
				result = shirtAndPants + " JACKET: " + jacket
				jacketOutcomes = append(jacketOutcomes, result)
			}
		}
	}

	coatOutcomes := []string{}
	if availableCloth["COAT"] == "may" {
		for _, shirtAndPants := range shirtsAndPantsOutcomes {
			for _, coat := range collections["COAT"] {
				isForbiddenYellow := coat == "yellow" && currentSeason == "FALL"
				isForbiddenOrange := coat == "orange" && currentSeason == "FALL"
				if !isForbiddenYellow && !isForbiddenOrange {
					result := ""
					result = "COAT: " + coat + " " + shirtAndPants
					coatOutcomes = append(coatOutcomes, result)
					if availableCloth["CAP"] == "may" {
						for _, shirtAndPantsAndCap := range collections["CAP"] {
							shirtAndPantsAndCapResult := ""
							shirtAndPantsAndCapResult = "COAT: " + coat + " " + shirtAndPants + " " + "CAP: " + shirtAndPantsAndCap
							coatOutcomes = append(coatOutcomes, shirtAndPantsAndCapResult)
						}
					}
				}
			}
		}
	}

	finalResultList := []string{}
	if availableCloth["CAP"] == "definitely" || currentSeason == "FALL" || currentSeason == "SPRING" {
		finalResultList = append(finalResultList, shirtsAndPantsOutcomes...)
	}
	finalResultList = append(finalResultList, shirtsAndPantsAndCapOutcomes...)
	finalResultList = append(finalResultList, jacketOutcomes...)
	finalResultList = append(finalResultList, coatOutcomes...)

	for _, v := range finalResultList {
		fmt.Println(v)
	}
}
