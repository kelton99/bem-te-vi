#!/bin/sh

user="abacateiro"
pass="abacateiro"
db="abacateiro"
host="localhost"
port="5432"

# Função para verificar se o PostgreSQL está pronto
wait_for_postgres() {
  echo "Aguardando o PostgreSQL ficar pronto..."
  echo "pg_isready -h $host -p $port -U $user"
  
  while true; do
    response=$(pg_isready -h "$host" -p "$port" -U "$user")
    if [ "$(echo $response | grep 'accepting connections')" ]; then
      break
    fi
    sleep 1
  done
  
  echo "PostgreSQL está pronto!"
}

# Executa as migrações
run_migrations() {
  echo ""
  echo "*** INICIANDO MIGRAÇÕES ***"
  # echo "migrate -path=\"/docker-entrypoint-initdb.d/migrations\" -database=\"postgres://$user:$pass@$host:$port/$db?sslmode=disable\" up"
  migrate -path="/docker-entrypoint-initdb.d/migrations" -database="postgres://$user:$pass@$host:$port/$db?sslmode=disable" up
  echo "*** MIGRAÇÕES CONCLUÍDAS ***"
}

# Executa as funções
wait_for_postgres
run_migrations