#!/bin/sh

# Используйте двойные кавычки и проверку на пустоту
if [ "${DEBUG:-0}" = "1" ]; then 
  RED='\033[0;31m'
  NC='\033[0m' # No Color
  echo "${RED}DEBUG mode is ON${NC}"
fi

cd /app

# export GIN_BUILD_ARGS="-race"

gin -i main.go