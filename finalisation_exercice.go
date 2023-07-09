package main

func FinalisationExercice(list []JsonSaveModel) {

	initialList := GetDataFromSave()
	if len(initialList) != 0 {
		for index, initial := range initialList {
			for _, fromExercice := range list {
				if initial.Element == fromExercice.Element {
					initialList[index] = fromExercice
				} else {
					initialList = append(initialList, fromExercice)
				}
			}
		}
	}
	for _, item := range list {
		initialList = append(initialList, item)
	}

	err := SaveData(initialList)
	if err != nil {
		println(err.Error())
	}

}
