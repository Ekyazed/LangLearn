package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func InitApplication() []JsonSaveModel {

	var path string
	var list []string
	var lastString string
	var jsonList []JsonSaveModel

	fmt.Print("Renseigner le chemin de votre fichier (appuyer sur \"entrer\" pour passer) \n")
	fmt.Scanln(&path)

	jsonList = GetDataFromSave()

	if path != "" {
		fmt.Println("\nRécupération de la liste depuis votre fichier")
		file, _ := os.Open(path)
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			list = append(list, scanner.Text())
		}
		if len(jsonList) == 0 {
			for index, line := range list {
				if index%2 == 0 {
					lastString = line
				} else {
					newSave := JsonSaveModel{
						Element:           lastString,
						Reponse:           line,
						Niveau:            FRESH,
						ProchaineEcheance: time.Now().AddDate(0, 0, -1),
					}
					jsonList = append(jsonList, newSave)
				}
			}
		} else {
			for _, save := range jsonList {
				for index, line := range list {
					if index%2 == 0 {
						if save.Element == line {
							continue
						} else {
							lastString = line
						}
					} else {
						newSave := JsonSaveModel{
							Element:           lastString,
							Reponse:           line,
							Niveau:            FRESH,
							ProchaineEcheance: time.Now().AddDate(0, 0, -1),
						}
						jsonList = append(jsonList, newSave)
					}
				}
			}
		}

	}
	return FilterForSelection(jsonList)

}
