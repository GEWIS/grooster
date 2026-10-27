# The frontend depends on the local generated client (file:../packages/grooster-client),
# which lives outside this build context. It is passed in as the named build context
# "grooster-client" (see .github/workflows/docker-build.yml):
#   docker build --build-context grooster-client=packages/grooster-client frontend
FROM node:22.23.4-alpine AS client

WORKDIR /packages/grooster-client

COPY --from=grooster-client ./package.json ./
COPY --from=grooster-client ./package-lock.json ./

RUN npm ci

COPY --from=grooster-client ./tsconfig.json ./
COPY --from=grooster-client ./src ./src

RUN npm run build

FROM node:22.23.4-alpine AS build

# Matches the relative path in package.json, as seen from WORKDIR /app
COPY --from=client /packages/grooster-client /packages/grooster-client

WORKDIR /app

COPY ./package.json ./
COPY ./package-lock.json ./

RUN npm install

COPY . ./

RUN npm run build

#Sever
FROM caddy:2-alpine
COPY --from=build /app/dist /srv
COPY ./caddy/Caddyfile /etc/caddy/Caddyfile
EXPOSE 8080
