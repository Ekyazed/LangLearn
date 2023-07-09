package main

import (
	"encoding/json"
	"os"
	"time"
)

type JsonSaveModel struct {
	Element           string     `json:"element"`
	Reponse           string     `json:"reponse"`
	Niveau            string     `json:"niveau"`
	ProchaineEcheance *time.Time `json:"prochaine_echeance"`
}

func GetDataFromSave() []JsonSaveModel {
	readData, err := os.ReadFile("save/data.json")
	if err != nil {
		panic(err)
	}

	var data []JsonSaveModel

	err = json.Unmarshal(readData, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func SaveData(data []JsonSaveModel) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("save/data.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func HandleAnswerUpLevel(item JsonSaveModel) JsonSaveModel {
	switch item.Niveau {
	case FRESH:
		item.Niveau = APPRENTISSAGE1
		item.ProchaineEcheance = APPR1()
	case APPRENTISSAGE1:
		item.Niveau = APPRENTISSAGE2
		item.ProchaineEcheance = APPR2()
	case APPRENTISSAGE2:
		item.Niveau = APPRENTISSAGE3
		item.ProchaineEcheance = APPR3()
	case APPRENTISSAGE3:
		item.Niveau = APPRENTISSAGE4
		item.ProchaineEcheance = APPR4()
	case APPRENTISSAGE4:
		item.Niveau = CONNAISSEUR1
		item.ProchaineEcheance = CONN1()
	case CONNAISSEUR1:
		item.Niveau = CONNAISSEUR2
		item.ProchaineEcheance = CONN2()
	case CONNAISSEUR2:
		item.Niveau = MAITRE1
		item.ProchaineEcheance = MTRE1()
	case MAITRE1:
		item.Niveau = MAITRE2
		item.ProchaineEcheance = MTRE2()
	case MAITRE2:
		item.Niveau = EXPERT1
		item.ProchaineEcheance = EXP1()
	case EXPERT1:
		item.Niveau = EXPERT2
		item.ProchaineEcheance = EXP2()
	case EXPERT2:
		item.Niveau = BURNED
		item.ProchaineEcheance = nil

	}

	return item
}

func HandleAnswerDownLevel(item JsonSaveModel) JsonSaveModel {
	switch item.Niveau {
	case APPRENTISSAGE2:
		item.Niveau = APPRENTISSAGE1
		item.ProchaineEcheance = APPR1()
	case APPRENTISSAGE3:
		item.Niveau = APPRENTISSAGE2
		item.ProchaineEcheance = APPR2()
	case APPRENTISSAGE4:
		item.Niveau = APPRENTISSAGE3
		item.ProchaineEcheance = APPR3()
	case CONNAISSEUR1:
		item.Niveau = APPRENTISSAGE4
		item.ProchaineEcheance = APPR4()
	case CONNAISSEUR2:
		item.Niveau = CONNAISSEUR1
		item.ProchaineEcheance = CONN1()
	case MAITRE1:
		item.Niveau = CONNAISSEUR2
		item.ProchaineEcheance = CONN2()
	case MAITRE2:
		item.Niveau = MAITRE1
		item.ProchaineEcheance = MTRE1()
	case EXPERT1:
		item.Niveau = MAITRE2
		item.ProchaineEcheance = MTRE2()
	case EXPERT2:
		item.Niveau = EXPERT1
		item.ProchaineEcheance = EXP1()

	}

	return item
}
