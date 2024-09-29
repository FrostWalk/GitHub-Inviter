#!/bin/bash

# Nome del file Go che desideri compilare
SOURCE_FILE="main.go"

# Nome del binario risultante
OUTPUT_NAME="inviter"
OUTPUT_PATH="./build"

# Funzione per compilare l'applicazione
build() {
    local os=$1
    local arch=$2
    local output=$3

    echo "Compilazione per $os/$arch..."

    # Impostiamo le variabili di ambiente per la compilazione
    GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -ldflags="-s -w" -o "$OUTPUT_PATH/$output" $SOURCE_FILE

    if [ $? -ne 0 ]; then
        echo "Errore durante la compilazione per $os/$arch"
        exit 1
    fi

    echo "Compilazione completata: $output"
}

# Compilazione per Windows (amd64)
build "windows" "amd64" "${OUTPUT_NAME}_windows_amd64.exe"

# Compilazione per Linux (amd64)
build "linux" "amd64" "${OUTPUT_NAME}_linux_amd64"

# Compilazione per Linux (arm64)
build "linux" "arm64" "${OUTPUT_NAME}_linux_arm64"

# Compilazione per macOS (Intel)
build "darwin" "amd64" "${OUTPUT_NAME}_macos_amd64"

# Compilazione per macOS (ARM - serie M)
build "darwin" "arm64" "${OUTPUT_NAME}_macos_arm64"

echo "Tutte le compilazioni sono completate."
