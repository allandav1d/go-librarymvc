# Library MVC

Sistema de gerenciamento de biblioteca desenvolvido em Go seguindo o padrão MVC (Model-View-Controller).

Este projeto foi criado durante o curso da Rocketseat sobre desenvolvimento backend com Go.

## 📋 Sobre o Projeto

Sistema de biblioteca que permite gerenciar:
- **Usuários**: Cadastro e gerenciamento de usuários
- **Livros**: Cadastro e gerenciamento de livros
- **Empréstimos**: Controle de empréstimos de livros

## 🚀 Tecnologias

- **Go** 1.24.3
- **Gin** - Framework web para Go
- **Gin Validator** - Validação de dados
- **Air** - Hot reload para desenvolvimento (opcional)

## 📁 Estrutura do Projeto

```
go-librarymvc/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
├── internal/
│   ├── books/               # Módulo de livros
│   ├── loans/               # Módulo de empréstimos
│   └── users/               # Módulo de usuários
├── go.mod
└── go.sum
```

## 🔧 Como Executar

### Pré-requisitos

- Go 1.24.3 ou superior instalado

### Instalação

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
cd go-librarymvc
```

2. Instale as dependências:
```bash
go mod download
```

3. (Opcional) Instale o Air para desenvolvimento com hot reload:
```bash
go install github.com/air-verse/air@latest
```

Adicione o Go bin ao seu PATH (se ainda não estiver configurado):
```bash
# Para ZSH (macOS padrão)
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

# Para Bash
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Verifique a instalação:
```bash
air -v
```

4. Execute a aplicação:

**Modo de desenvolvimento (com hot reload):**
```bash
air
```
> O Air irá monitorar mudanças nos arquivos `.go` e `.html` e automaticamente reconstruir e reiniciar a aplicação.

**Modo normal:**
```bash
go run cmd/api/main.go
```

A aplicação estará rodando em `http://localhost:8080`

## 📝 Endpoints

A aplicação possui rotas para:
- Gerenciamento de usuários
- Gerenciamento de livros
- Gerenciamento de empréstimos

## 🎓 Curso

Este projeto foi desenvolvido durante o curso da Rocketseat sobre desenvolvimento backend com Go.

## 📄 Licença

Este projeto foi criado para fins educacionais.

