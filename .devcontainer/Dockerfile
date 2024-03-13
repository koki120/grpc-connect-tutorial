FROM --platform=linux/x86_64 mcr.microsoft.com/vscode/devcontainers/go

ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# bufをインストール
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# grpcurlをインストール
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# protoc-gen-goをインストール
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# protoc-gen-connect-goをインストール
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
