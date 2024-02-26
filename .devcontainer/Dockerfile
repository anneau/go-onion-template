FROM mcr.microsoft.com/devcontainers/go:1-1.22-bullseye

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends postgresql-client \
    && apt-get autoremove -y \
    && apt-get clean

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install github.com/google/wire/cmd/wire@v0.6.0 \
    && go install github.com/rubenv/sql-migrate/sql-migrate@v1.4.0

# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
