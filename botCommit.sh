#!/bin/bash

# Atualiza o repositório local com as alterações do repositório remoto
git pull origin main

# Adiciona todas as alterações
git add .

# Faz o commit com uma mensagem vazia
git commit -a --allow-empty-message -m ""

# Envia as alterações para o repositório remoto
git push origin main
