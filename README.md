# 🕒 WebSocket Clock com Go

Este é um projeto simples que demonstra como usar **WebSocket** com **Go (Golang)** para enviar mensagens em tempo real para o front-end. Neste exemplo, o servidor envia a **hora atual (hh:mm:ss)** a cada 5 segundos para o cliente via WebSocket.

## 🚀 Funcionalidades

- Servidor WebSocket implementado com Go
- Página HTML simples que recebe e exibe mensagens em tempo real
- Atualização automática da hora sem precisar recarregar a página

---

## 📁 Estrutura do Projeto

.
├── main.go # Servidor WebSocket em Go
└── index.html # Página HTML que consome o WebSocket


---

## ▶️ Como rodar localmente

### Pré-requisitos

- Go instalado (versão 1.16+)
- Navegador moderno (Chrome, Firefox, Edge, etc.)

### Passo a passo

1. Clone ou copie os arquivos `main.go` e `index.html` no mesmo diretório.

2. Execute o servidor:

   ```bash
   go run main.go

3. Acesse no navegador: http://localhost:8080

Você verá a hora sendo atualizada a cada 5 segundos no navegador.
