#!/bin/bash

# Obtient le répertoire actuel
OPENDIR=$(pwd)

# Lancement de l'API User
gnome-terminal --tab --title="API User" --working-directory="$OPENDIR" -- bash -c "cd ./Middleware-API-User/ && go run ./cmd/main.go"

# Lancement de l'API Song
gnome-terminal --tab --title="API Song" --working-directory="$OPENDIR" -- bash -c "cd ./Middleware-API-Song/ && go run ./cmd/main.go"

# Lancement de l'API Rating
gnome-terminal --tab --title="API Rating" --working-directory="$OPENDIR" -- bash -c "cd ./Middleware-API-Rating/ && go run ./cmd/main.go"

# Lancement de l'API Flask
# Note : On change de répertoire avant le lancement pour garantir que le terminal reste ouvert
cd ./Middleware-API-Flask/ && PYTHONPATH=$PYTHONPATH:$(pwd) gnome-terminal --tab --title="API Flask" -- bash -c "python3 src/app.py"

# Lancement du FRONT
gnome-terminal --tab --title="Front" --working-directory="$OPENDIR" -- bash -c "cd ./Middleware-Front/ && npm run dev"
