package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	count := FilterForSelection(jsonList)
	fmt.Printf("\nNombre d'éléments sauvegarder: %s, nombre d'éléments sauvegarder prêt pour le passage: %s", strconv.Itoa(len(jsonList)), strconv.Itoa(len(count)))

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
			for index, line := range list {
				if index%2 == 0 {
					lastString = line
				} else {
					var exist = false
					for _, item := range jsonList {
						if item.Element == lastString {
							exist = true
							break
						}
					}
					if !exist {
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
