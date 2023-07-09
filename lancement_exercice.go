package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func LancementExercice(list []JsonSaveModel) []JsonSaveModel {
	var itemPassed = 0
	var completedItem []JsonSaveModel
	listLength := len(list)
	fmt.Printf("\nLes données du fichier ont été récupérer !, %s éléments récupérés \n", strconv.Itoa(listLength))

	fmt.Printf("\nC'est parti !\n")
	fmt.Println("")

	errHandle := make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)

	for listLength > itemPassed {
		index := rand.Intn(len(list) - 0)
		var answer string

		fmt.Printf("%s : %s \n", strconv.Itoa(itemPassed+1)+"/"+strconv.Itoa(listLength), list[index].Element)
		fmt.Println("Votre réponse ?")
		//keyboard.GetSingleKey()
		//fmt.Printf("Réponse : %s \n", list[index+1])
		scanner.Scan()
		answer = scanner.Text()

		indexItems := strings.Split(list[index].Reponse, ",")
		var goodAnswer = false
		for _, item := range indexItems {
			if strings.ToLower(item) == strings.ToLower(answer) {
				goodAnswer = true
				break
			}
		}

		if goodAnswer {
			if value, exist := errHandle[index]; !exist {
				list[index] = HandleAnswerUpLevel(list[index])
				color.Green("Bien jouer. Cette question est maintenant au niveau " + list[index].Niveau)
			} else if value == 1 {
				color.HiYellow("Bien jouer. Vous avez raté " + strconv.Itoa(value) + " fois. vous restez au niveau " + list[index].Niveau)
			}
			completedItem = append(completedItem, list[index])
			list = append(list[:index], list[index+1:]...)
			itemPassed++
		} else {
			value, exist := errHandle[index]
			if exist && (value+1 == 2) {
				list[index] = HandleAnswerDownLevel(list[index])
				color.Red("Raté. la bonne réponse est " + list[index].Reponse + ". Vous êtes tombé au niveau " + list[index].Niveau)
				completedItem = append(completedItem, list[index])
				list = append(list[:index], list[index+1:]...)
				itemPassed++
			} else if !exist {
				errHandle[index] = 1
				color.Red("Dommage, réessayez plus tard")
			}
		}

	}

	fmt.Printf("La session est terminer ! appuyer sur entré pour fermer la fenêtre")
	keyboard.GetSingleKey()

	fmt.Println("\nFinalisation de la mise a jour de la liste")

	return completedItem
}
