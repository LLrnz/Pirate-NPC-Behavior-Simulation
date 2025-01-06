# IA04 - Simulation multi-agents

## Sujet

### Qu'aurait dû faire la France pour se propulser au sommet de la richesse pendant l'Âge d'or de la piraterie (1716 - 1726) ?

Notre simulation se déroule lors de l'Âge d'or de la piraterie, entre 1716 et 1726.
Trois états parcourent les mers, cherchant à accumuler le plus de richesses : les royaumes de France, d'Espagne, et d'Angleterre.

TODO - conclusion sur nos itérations de simulations

## Lancer le projet
TODO 


## Structure
### Agents
TODO

### Objets
Les ports sont des objets avec lesquels les agents navires peuvent interagir. Les navires peuvent accoster dans un port pour y faire du commerce (vendre et acheter des ressources), peuvent se réparer (si le nombre de points de vie sont inférieur au nombre de points de vie maximum), et s'y réfugier. Lorsque les navires sont à quais, ils ne peuvent attaquer, ni être attaqués (que ce soit par d'autres navires, ou bien par le kraken).

### Autres
Marché

## Point forts
Les points forts de notre projet sont les suivants :
- marché évolutif : les prix des ressources évoluent en fonction de l'offre et de la demande
- le projet répond à la problématique : de par sa structure et ses éléments variables, le projet apporte une réponse à la problématique de base ; vaut-il mieux mettre l'accent sur le nombre de bateaux marchands, les navires de l'armée, ...?
- algorithme A* : l'itinéraire des agents est calculé grâce à l'algorithme de recherche de chemin A* (dérivé de l'algorithme de Djikstra). C'est un algorithme notamment utilisé dans les jeux vidéos pour les déplacements dynamiques de PNJ (ennemis cherchant le plus court chemin pour arriver au niveau du joueur et l'attaquer, par exemple)
- les visuels : tous les assets graphiques ont été créés spécifiquement pour le projet
- ça fonctionne : le projet fonctionne


## Améliorations possibles
Les points d'amélioration de notre projet sont les suivants :
- niveaux : originellement, nous souhaitions ajouter des niveaux aux agents. Un navire pourrait augmenter de niveau en payant une certaine somme. Une augmentation de niveau mènerait à un navire aux capacités plus importantes (vie, attaque, vitesse plus grandes).
- constantes : les constantes pourraient être définies avec plus de précision en réalisant 
- communication inter-agents : nous avons pensé (la veille du rendu) à mettre en place des échanges de messages entre les agents navires : un navire marchand attaqué pourrait joindre le navire de l'armée de sa nationalité le plus proche (sous réserve que celui-ci ne soit pas en train d'attaquer un autre agent), afin qu'il puisse lui venir en aide.