FROM node:18

# Instalar dependências necessárias com apt-get, se houver
RUN apt-get update && apt-get install -y --no-install-recommends \
    bash \
    git \
    python3 \
    python3-pip \
    make \
    g++ \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Instalar Yarn globalmente
RUN corepack enable && corepack prepare yarn@stable --activate

# Instalar as dependências do projeto
# É uma boa prática instalar as dependências diretamente no volume
# Se você não usa o volume, você pode manter o COPY package.json e yarn.lock aqui
COPY package.json yarn.lock ./

# Instalar as dependências do projeto
RUN yarn install

# Instalar Vue CLI e Quasar CLI globalmente usando npm
RUN npm install -g @vue/cli @quasar/cli

# Copiar o restante do código do projeto
COPY . .

# Expor a aplicação na porta 9000
EXPOSE 9000

# Iniciar o servidor de desenvolvimento
CMD ["yarn", "dev"]
