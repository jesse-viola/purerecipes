# purerecipes-compose

## Development tools

https://python-poetry.org/ for backend package management

All submodules have both Dockerfile.dev which is called from the docker-compose.dev.yml.

### What commands do I need to run dev mode containers?

- Build, run and detach

  > docker-compose -f $(compose_yaml) up -d

- Shutdown service containers

  > docker-compose -f $(compose_yaml) down

- Destroy

  > docker-compose -f (docker.compose yaml) down -v

- Logs (Or view inside of docker desktop using GUI)

  > docker-compose -f docker-compose.yml logs --tail=100 -f $(container_name)

## Dev mode containers should hot reload on both backend and frontend changes.
