# Martin-solving multi-tools CLI

### Installation avec Homebrew uniquement sur Mac :

Tapez :

```
brew tap Lachignol/homebrew-Lachignol 
```

```
brew install lachignol/lachignol/martin-solving
```


## Premiere fonctionnalité la todo-list.


### Interface graphique:

https://github.com/user-attachments/assets/29f73125-7062-4019-845f-0dcbe904e502

- [x] Ajout d'une nouvelle tâche.
- [x] Possibilité de marquer la tâche comme "complétée".
- [x] Lister les tâches sous forme d'un tableau graphique.
- [x] Suppression d'une tâche.
- [x] Mise en place d'un moyen de stockage persistant et accessible depuis n'importe quel endroit dans votre terminal (BDD SQlite):

      
      
<img width="1133" alt="installation:todocmd screenshot" src="https://github.com/Lachignol/cli-app/assets/110435478/a4a2c0b0-66cd-4d9f-ac3d-1bbda79c7fd5">
Toutes ces commandes sont aussi exécutables en simple ligne de commande .

## Deuxieme fonctionnalité actionOnFile.
Elle permet de remplacer un motif spécifique dans un fichier d'entrée et de créer un nouveau fichier en sortie contenant le contenu modifié. Elle est particulièrement utile pour les utilisateurs qui souhaitent effectuer des remplacements de texte dans des fichiers sans modifier l'original.


### Interface graphique:

- [x] Choix du fichier à modifier en entrée.
- [x] Choix du pattern à remplacer/choix du pattern de remplacement.
- [x] Création du fichier en sortie.
- [x] Affichage de la commande CLI avec les flags pour une utilisation ultérieure.

### Commande cli :

#### Description des flags :

<img width="1434" alt="Capture d’écran 2024-08-26 à 01 07 41" src="https://github.com/user-attachments/assets/d2b6336f-ba61-4f96-b359-9cffa62fdaba">


#### Exemple :

```
martin-solving actionOnFile -i cheminDuFichier/nomDufichier.txt -o cheminDuFichier/nomDufichier.txt -r 'patternARemplacer' -b 'patternDeRemplacement'
```







## Uninstall afin de pouvoir supprimer proprement l'application.
      
<img width="1133" alt="uninstall screenshot" src="https://github.com/Lachignol/cli-app/assets/110435478/cf43a513-d0e7-48b0-ae64-f18b0db4530d">










