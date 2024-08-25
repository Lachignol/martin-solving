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
      
https://github.com/user-attachments/assets/7a293494-44c4-4cd8-9693-e76cd2505a66
      
- [x] Création du fichier en sortie.
- [x] Affichage de la commande CLI avec les flags pour une utilisation ultérieure.
      
https://github.com/user-attachments/assets/4086b518-dd42-4453-873f-1bb6f45fe9c6

### Commande cli :

#### Exemple :

```
martin-solving actionOnFile -i cheminDuFichier/nomDufichier.txt -o cheminDuFichier/nomDufichier.txt -r 'patternARemplacer' -b 'patternDeRemplacement'
```

https://github.com/user-attachments/assets/31481b79-379c-4c33-86a6-7325e0a0420c


#### Description des flags :

<img width="1434" alt="Capture d’écran 2024-08-26 à 01 07 41" src="https://github.com/user-attachments/assets/d2b6336f-ba61-4f96-b359-9cffa62fdaba">


## Uninstall afin de pouvoir supprimer proprement l'application.

<img width="1434" alt="uninstall-snapshot" src="https://github.com/user-attachments/assets/69461641-ebcb-4f24-9431-6d6b016d6baa">
