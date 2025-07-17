# Dockerfile optimizado para devcontainer con Go 1.24 (solo backend)
FROM mcr.microsoft.com/vscode/devcontainers/go:1.24-bullseye

# Instalar herramientas adicionales para desarrollo backend
RUN apt-get update && apt-get install -y \
    git \
    curl \
    wget \
    unzip \
    findutils \
    tree \
    htop \
    postgresql-client \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Instalar herramientas de Go para 1.24
RUN go install github.com/air-verse/air@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install golang.org/x/tools/gopls@latest

# Configurar workspace
WORKDIR /workspace

# Crear directorios para cache
RUN mkdir -p /go/pkg/mod \
    && mkdir -p /home/vscode/.cache/go-build \
    && mkdir -p /home/vscode/.vscode-server/extensions

# Configurar permisos para el usuario vscode
RUN chown -R vscode:vscode /go \
    && chown -R vscode:vscode /home/vscode

# Configurar Git para el contenedor
RUN git config --global --add safe.directory /workspace

# Configurar variables de entorno para Go 1.24
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV GOPATH=/go
ENV PATH=$PATH:/go/bin
ENV GO111MODULE=on

USER vscode

# Configurar shell para el usuario vscode
RUN echo 'export PATH=$PATH:/go/bin' >> ~/.bashrc \
    && echo 'export PATH=$PATH:/go/bin' >> ~/.zshrc \
    && echo 'export GO111MODULE=on' >> ~/.bashrc \
    && echo 'export GO111MODULE=on' >> ~/.zshrc

# Configurar aliases útiles para backend
RUN echo 'alias ll="ls -la"' >> ~/.bashrc \
    && echo 'alias la="ls -la"' >> ~/.bashrc \
    && echo 'alias ..="cd .."' >> ~/.bashrc \
    && echo 'alias gst="git status"' >> ~/.bashrc \
    && echo 'alias gco="git checkout"' >> ~/.bashrc \
    && echo 'alias backend="cd /workspace/backend"' >> ~/.bashrc \
    && echo 'alias gomod="go mod tidy && go mod download"' >> ~/.bashrc \
    && echo 'alias gofmt="gofmt -w ."' >> ~/.bashrc \
    && echo 'alias gorun="go run cmd/api/main.go"' >> ~/.bashrc \
    && echo 'alias air="air"' >> ~/.bashrc