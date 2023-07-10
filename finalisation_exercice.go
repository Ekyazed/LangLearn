package main

func FinalisationExercice(list []JsonSaveModel) {
	var elem JsonSaveModel
	initialList := GetDataFromSave()
	if len(initialList) != 0 {
		for index, initial := range initialList {
			var found = false
			for _, fromExercice := range list {
				if initial.Element == fromExercice.Element {
					initialList[index] = fromExercice
					found = true
				} else {
					elem = fromExercice
				}
			}
			if !found {
				initialList = append(initialList, elem)
			}
		}
	} else {
		for _, item := range list {
			initialList = append(initialList, item)
		}
	}

	err := SaveData(initialList)
	if err != nil {
		println(err.Error())
	}

}
