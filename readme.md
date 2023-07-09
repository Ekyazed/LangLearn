# <u>LangLearn</u>

## <u>Comment fonctionne l'application</u>

<span style="font-size: 18px">

Pour alimenter l'application, il faut créer un fichier texte (exemple: Bloc Note) qui devra être rempli de la manière suivante : 

	question
	liste de réponse

Exemple (français/anglais):
    
    Bonjour
    Hello, good morning, ...

Dans la liste de proposition des réponses, chaque éléments devra être séparer par ", ".     
Un fichier sera générer dans un fichier "save.json" une liste d'éléments pour la lecture
pour l'execice en cours.

</span>

## <u>Un système d'apprentissage sur le long terme</u>

<span style="font-size: 18px">
    Cette application utilise une méthode d'apprentissage sur le long terme: <br />
    
```text
    Apprentissage 1 (4 heures pour le prochain passage)
    Apprentissage 2 (8 heures pour le prochain passage) 
    Apprentissage 3 (1 jour pour le prochain passage) 
    Apprentissage 4 (4 jours pour le prochain passage) 
    Connaisseur 1 (1 semaine pour le prochain passage) 
    Connaisseur 2 (2 semaines pour le prochain passage) 
    Maitre 1 (1 mois pour le prochain passage)
    Maitre 2 (2 mois pour le prochain passage)
    Expert 1 (4 mois pour le prochain passage)
    Expert 2 (12 mois pour le prochain passage)
    Burned (ne tombera plus dans la liste des passage)
```

</span>