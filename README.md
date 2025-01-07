# IA04 - Simulation multi-agents

## Sujet

### Qu'aurait dû faire la France pour se propulser au sommet de la richesse pendant l'Âge d'or de la piraterie (1716 - 1726) ?

Notre simulation se déroule lors de l'Âge d'or de la piraterie, entre 1716 et 1726.
Trois états parcourent les mers, cherchant à accumuler le plus de richesses : les royaumes de France, d'Espagne, et d'Angleterre.

#### Résultats de nos simulations 
Pour pouvoir répondre à la question posée, nous avons réalisé diverses simulations.

- Première Simulation : 
Nous avons observé que lorsque nous avons comme conditions 1 Kraken et 50 marchands pour chacune des 3 nationalités, il est largement préférable d'avoir uniquement des Privateer dans le cas où les autres nations ont un nombre supérieur ou égal d'Army par rapport à leur nombre de Privateer.  

- Deuxième Simulation : 
Nous avons observé que dans le cas où il n'y a que 2 nations et un nombre de kraken égal à 1 et de marchands égal à 50 dans chaque nations, il est beaucoup plus rentable d'avoir beaucoup de Privateer et quelques bateaux Army que d'avoir uniquement des Privateer.

- Troisième simulation : 
Nous avons également observé, que lorsque uniquement deux nations s'affrontent sachant qu'il y a un nombre élevé de Kraken (supérieur ou égal à 50) il est largement préférable que le pays se dote de beaucoup plus de Army que de Privateer, car les Army ont une bien meilleure capacité à combattre les Kraken.

## Lancer le projet
Pour lancer le projet, il suffit d'utiliser la commande suivante :
`go run .`

## Structure
### Agents
Les agents du projet interagissent entre eux et se déplacent dans l'environnement.

- Navires marchands : les navires marchands sont affiliés à l'un des trois états (France, Espagne ou Angleterre) et ont pour objectif principal de faire du commerce (vente et achat de ressources). L'argent possédé par ces navires provient des marchandises qu'ils vendent. Lorsqu'ils repèrent un danger (navires pirates, navires corsaires d'une autre nationalité que la leur, Zoro, Rachkam le Rouge, kraken), les navires marchands fuient pour se réfugier dans un port. Ces navires n'attaquent pas d'autres agents.
- Navires de l'armée : les navires de l'armée sont affiliés à l'un des trois états. Ils ne pratiquent pas de commerce (ils n'achètent ni ne vendent de ressources). Leur argent est directement issu du trésor du pays auquel ils sont affiliés. Les navires de l'armée attaquent les ennemis qu'ils perçoivent (navires pirates, navires corsaires d'une autre nationalité, Zoro, Rackham le Rouge et kraken). Ces navires ont une attaque et un champ de vue plus haut que les autres navires.
- Navires pirates : les navires pirates ne sont affiliés à aucun état. Ces navires pratiquent le commerce et peuvent acheter et vendre des ressources. Ces navires fuient face au kraken et aux bateaux de l'armée (qu'importe la nationalité). Ils peuvent attaquer tous les bateaux marchands, pirates et corsaires qu'ils perçoivent. Si les navires pirates attaquent et réussissent à battre un navire, ils peuvent récupérer 60% de ses biens ; le reste est considéré perdu lors de l'attaque...
- Navires corsaires : le comportement des navires corsaires est sensiblement similaire à celui des navires pirates, à l'exception que ces derniers sont affiliés à un pays, et n'attaquent pas les navires marchands et corsaires de la même nationalité qu'eux.
- Zoro : Zoro s'est perdu. Il se déplace complètement au hasard sur la carte, et son champ de vision est extrêmement réduit. Il attaque tout agent dans son champ de vision. Son attaque est bien plus élevée que celles des autres agents.
- Rachkam le Rouge : Rackham le Rouge est très fier et très cruel. Il n'hésite pas à attaquer tout navire entrant dans son champ de vision. Cela dit, son attaque est très basse, donc il ne fait pratiquement pas de dégâts... Mais il a la détermination. C'est au moins cela.
- Kraken : le kraken est une créature mythique que peu peuvent se targuer d'avoir rencontrée (et d'être encore en vie pour le raconter). Il arpente les mers, et attaque tout navire entrant dans son champ de vision.

Note : Lorsqu'ils détectent une cible à attaquer, les agents peuvent poursuivre ces-dernières avant de les attaquer lorsqu'ils en sont suffisamment proches.

### Objets
Les ports sont des objets avec lesquels les agents navires peuvent interagir. Les navires peuvent accoster dans un port pour y faire du commerce (vendre et acheter des ressources), peuvent se réparer (si le nombre de points de vie est inférieur au nombre de points de vie maximum, et s'ils ont suffisamment d'argent pour effectuer lesdites réparations), et s'y réfugier. Lorsque les navires sont à quai, ils ne peuvent attaquer, ni être attaqués (que ce soit par d'autres navires, ou bien par le kraken).

Lorsque les navires vendent des ressources dans un port, ils subissent une taxe. 50% du prix de vente est reversé à l'Etat auquel est affilié le navire qui effectue la vente (cette somme est perdue si le navire n'est affilié à aucun pays (cas des bateaux pirates)). Le reste est récupéré par le navire (cela constitue son trésor individuel, qu'il peut utiliser pour acheter d'autres ressources ou effectuer des réparations si besoin).

### Autres
#### Ressources 
Voici la liste des ressources disponibles dans les ports : 
- Or
- Argent 
- Café
- Epices 
- Sucre 
- Rhum 
- Tabac 
- Soie 
- Cacao 
- Perles
- Cotton 
- Ivoire 
- Armes 
- **Tipiak** 

#### Marché 
Ces différents types de ressources sont régis par les lois d’un marché qui vise à reproduire le fonctionnement du système d’offre et de demande tel que nous le connaissons aujourd’hui. Il a été simplifié par une fonction de décroissance hyperbolique. Plus un produit est vendu, plus sa valeur dimunue avec une mémoire de 2 ans.

## Point forts
Les points forts de notre projet sont les suivants :
- Marché évolutif : les prix des ressources évoluent en fonction de l'offre et de la demande
- le projet répond à la problématique : de par sa structure et ses éléments variables, le projet apporte une réponse à la problématique de base ; vaut-il mieux mettre l'accent sur le nombre de bateaux marchands, les navires de l'armée, ...?
- algorithme A* : l'itinéraire des agents est calculé grâce à l'algorithme de recherche de chemin A* (dérivé de l'algorithme de Dijkstra). C'est un algorithme notamment utilisé dans les jeux vidéos pour les déplacements dynamiques de PNJ (ennemis cherchant le plus court chemin pour arriver au niveau du joueur et l'attaquer, par exemple)
- les visuels : tous les assets graphiques ont été créés spécifiquement pour le projet
- ça fonctionne : le projet fonctionne


## Améliorations possibles
Les points d'amélioration de notre projet sont les suivants :
- niveaux : originellement, nous souhaitions ajouter des niveaux aux agents. Un navire pourrait augmenter de niveau en payant une certaine somme. Une augmentation de niveau mènerait à un navire aux capacités plus importantes (vie, attaque, vitesse plus grandes).
- constantes : les constantes pourraient être définies avec plus de précision en réalisant un plus grand nombre de simulations, afin de pouvoir obtenir une simulation des plus réaliste. 
- communication inter-agents : nous avons pensé (la veille du rendu) à mettre en place des échanges de messages entre les agents navires : un navire marchand attaqué pourrait joindre le navire de l'armée de sa nationalité le plus proche (sous réserve que celui-ci ne soit pas en train d'attaquer un autre agent), afin qu'il puisse lui venir en aide.