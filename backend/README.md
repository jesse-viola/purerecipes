# backend

## Poetry commands

Once poetry is installed using this script 

- https://python-poetry.org/docs/master/#installing-with-the-official-installer

CD into the backend directory and run the command

- This will install all dependencies from the `poetry.lock` file

    > poetry install

- Each poetry project (has a poetry lock file) will also create a virutalenv automatically. Here is the location of my current cache dir (MacOS)

    > /Users/jesseviola/Library/Caches/pypoetry/virtualenvs/

Find this setting in the VSCode perferences `Python: Venv Path`. Now when a python file is opened in the project you should be able to select the auto generated virtualenv per that projects lock file. This is how it will look on the command line terminal when it is activated.

`(backend-Yd2aQhFP-py3.9) username@MacBook-Pro purerecipes %`

### Updating private git dependencies

Navigate to the pyproject.toml and remove the git dependency. Then run `poetry install && poetry update`. Re-add the line and run the install and update again. This is a workaround for a bug in poetry currently

## Debugging

- Start the backend container as bash as PID1

    > docker compose run backend bash

- Attach VSCode using extension for Docker

- Once in new VSCode window make sure to select the venv python interpreter
    > /opt/pysetup/.venv/bin/python3.9

- Create a new run/debug configuration and select Current File for python and navigate to main.py


