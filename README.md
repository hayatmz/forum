# 🗨️ Forum
Il s'agit d'un forum web full-stack entièrement développé en **Go**, avec une base de données **SQLite**, sans aucun framework frontend externe.
> Projet collaboratif réalisé dans le cadre de la formation [Zone01](https://campus-saint-marc.com/zone-01/).

## 🚀 Fonctionnalités principales
- 🔐 **Authentification** : Inscription et connexion utilisateur avec sessions cookies.
- 📬 **Posts** : Création de posts vivibles par tous.
- 💬 **Commentaires** : Ajout de commentaires à des posts.
- 👍👎 **Likes/Dislikes** : Pour les posts et commentaires.
- 🗂️ **Catégories** : Association d'une ou plusieurs catégories à un post avec des #Hashtags.
- 🔍 **Barre de recherche** : Suggestions dynamiques en tapant dans la barre de recherche.
- 🧼 **Filtrage** : Par catégories, posts aimés ou crées.
- 📦 **Docker** : Projet intégralement dockerisé pour une exécution isolée.

## 🔧 Utilisation
1. Assure toi d'avoir **Go** et **Docker** installés sur ta machine :
```
go version
docker -v
```
si ce n'est pas le cas, [installe Golang](https://go.dev/doc/install) et [installe Docker](https://docs.docker.com/get-started/get-docker/).

2. **Clone le dépôt** :
```
git clone https://github.com/hayatmz/forum
cd forum
```

3. **Lancer automatiquement avec le script ```forumDocker.sh```** :
Un script bash est fourni pour **supprimer l'ancien conteneur/image** (si existants), **reconstruire** le projet, et **lancer** l'application proprement :
```
cd docker
chmod +x forumDocker.sh
./forumDocker.sh
```

L'application sera alors accessible à l'adresse :<br>
[http://localhost:15040](http://localhost:15040)
