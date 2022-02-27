# Groupie Tracker

## 1) Introduction
Nous avions à notre disposition une API composée de nombreux artistes et groupes de musique avec des informations sur chacun d'eux. Notre but était de créer un site web afin d'organiser de manière esthétique et fonctionnelle chaque information. 

Nous avons donc réalisé cet objectif.
Nous vous présenterons les différentes fonctionnalités du site dans ce document.
<br>
<br>
## 2) Lancement du site web
Afin de lancer notre site web, il faut tout d'abord rentrer la commande "go run main/main.go" dans le terminal et ensuite ouvrir le navigateur web au localhost 5500 en tapant "http://localhost:5500/" dans la barre de recherche.
<br>
<br>

## 3) Les fonctionnalités
Quand on ouvre le site web, nous avons accès à tous les groupes de musique disponible avec leurs informations sous forme de carte. Nous avons également accès à une barre de recherche et un bouton Filtre.
<br>
<br>
### b) Les filtres
Sur la gauche de notre site web, nous avons également ajouté la possibilité à l'utilisateur de filtrer les groupes de musique en fonction du nombre de leurs membres, de leur date de création, de la date de leur premier album ou encore du lieu de leurs concerts. Une fois les filtres rentrés, il ne suffit plus que de cliquer sur le bouton "Apply filters" pour que la page d'accueil se mette à jour avec les groupes cochant les bons critères.
<br>
<br>
Pour les membres, ce sont 7 checkbox allant de 1 à 8 qui sont utilisées. L'utilisateur n'a plus qu'à cocher sur l'une d'entre elles. Si la case 1 est cochée, alors seulement les groupes composés d'un seul artiste seront affiché.
<br>
En ce qui concerne les dates de création des groupes ou des premiers albums, nous utilisons des "sliders", ce qui permet à l'utilisateur de mettre le curseur sur une année précise parmis une soixantaine d'année. Si on décide de mettre le curseur à 1970, après actualisation le site web affichera tous les groupes ayant été créé en 1970. C'est le même système pour la date des premiers albums.
<br>
Enfin, nous pouvons également filtrer en fonction des lieux des différents concerts. Un volet roulant affiche tous les pays disponible et il n'y a plus qu'à en choisir un. Par exemple, si on choisit "Argentine", il n'y aura plus que les groupes ayant des concerts prévu en Argentine sur la page d'accueil.
<br>
<br>
On peut bien évidemment cumuler les filtres.
<br>
<br>

### c) Les cartes
Quand on est sur la page d'accueil et qu'un groupe nous plaît, nous avons la possibilité de cliquer sur la carte du groupe. Faire cela, nous permet d'avoir accès à toutes les informations concernant ce groupe : nom du groupe, membres, date et lieu des concerts, date de création, date du premier album.
<br>
Nous avons en haut à gauche de la page la possibilité de cliquer sur "Home" pour revenir sur la page d'accueil avec tous les groupes disponible.

<br>
<br>

 Il ne reste plus qu'à tester tout ça !