# Groupie Tracker

## 1) Introduction
Nous avions à notre disposition une api composée de nombreux artistes et groupes de musique avec des informations sur chacun d'eux. Notre but était de créer un site web afin d'organiser de manière esthétique et fonctionnelle chaque information. 

Nous avons donc réalisé cet objectif et nous vous présentons les différentes 
fonctionnalités du site dans ce document.
<br>
<br>
## 2) Lancement du site web
Afin de lancer notre site web, il faut tout d'abord rentrer la commande "go run main/main.go" dans le terminal et ensuite ouvrir le navigateur web au localhost 5500 en tapant "http://localhost:5500" dans la barre de recherche.
<br>
<br>

## 3) Les fonctionnalités
### a) Barre de Recherche
Quand on ouvre le site web, nous avons accès à tous les groupes de musiques disponibles avec leurs informations sous forme de carte.
<br>
<br>
La barre de recherche, situé sur la droite du site, permet à l'utilisateur de renter lui-même une date de création, le nom d'un groupe ou le nom d'un membre. A chaque caractère entré, la page d'accueil du site se met à jour en enlevant petit à petit les groupes qui ne font pas parti des critères rentrés dans la barre de recherche. Par exemple, si on écrit 1970 dans la barre de recherche, il n'y aura d'afficher que les groupes créé en 1970.
En plus de cela, au fur et à mesure que nous rentrons des caractères dans la barre de recherche, celle-ci nous fait des propositions logique en fonction de ce que nous écrivons.
<br>
<br>

### b) Les filtres
Sur la gauche de notre site web, nous avons également ajouter la possibilité à l'utilisateur de filter les groupes en fonction du nombre de leurs membres, de leur date de création, de la date de leur premier album ou encore du lieu de leurs concerts. Une fois les filtres rentrés, il ne suffit plus que de cliquer sur le bouton "Apply filters" pour que la page d'accueil se mette à jour avec les groupes cochant les bons critères.
<br>
<br>
Pour les membres, ce sont 7 checkbox allant de 1 à 7 qui sont utilisées. L'utilisateur n'a plus qu'à cocher sur l'une d'entre elles.
<br>
En ce qui concerne les dates de création des groupes ou des premiers albums, nous utilisons des "slider", ce qui permet à l'utilisateur de mettre le curseur sur une année précise parmi une soixantaine d'année. Si on décide de mettre le curseur à 1970, après actualisation le site web affichera tous les groupes ayant été créé avant 1970 inclus. C'est le même système pour la date des premiers albums.
<br>
<br>

###

