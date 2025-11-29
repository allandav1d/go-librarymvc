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
- **Tailwind CSS** - Framework CSS utility-first com design system shadcn-like
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
├── web/
│   └── controller/          # Controllers web para templates
├── templates/               # Templates HTML
│   ├── layout.html
│   ├── dashboard.html
│   ├── books.html
│   ├── users.html
│   └── loans.html
├── static/
│   └── css/
│       ├── input.css        # CSS Tailwind (source)
│       └── output.css       # CSS compilado (gerado)
├── .air.toml                # Configuração do Air
├── tailwind.config.js       # Configuração do Tailwind
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

**Alternativa: Criar um alias para o Air**

Se preferir usar um alias ao invés de adicionar ao PATH:
```bash
# Para ZSH (macOS padrão)
echo 'alias air="$HOME/go/bin/air"' >> ~/.zshrc
source ~/.zshrc

# Para Bash
echo 'alias air="$HOME/go/bin/air"' >> ~/.bashrc
source ~/.bashrc
```

4. Configure o Tailwind CSS:

Baixe o Tailwind CLI para macOS:
```bash
# Para macOS ARM (M1/M2/M3)
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 tailwindcss

# Para macOS Intel
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64
chmod +x tailwindcss-macos-x64
mv tailwindcss-macos-x64 tailwindcss

# Para Ubuntu/Linux
# Para Linux ARM64
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-arm64
chmod +x tailwindcss-linux-arm64
mv tailwindcss-linux-arm64 tailwindcss

# Para Linux x64
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
```

Compile o CSS:
```bash
./tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify
```

> **Nota:** Se você usar o Air para desenvolvimento, o Tailwind CSS será compilado automaticamente antes de cada build.

5. Execute a aplicação:

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

## 🎨 Design System

O projeto utiliza Tailwind CSS v4 com um design system inspirado no **shadcn/ui**, incluindo:

### Componentes Estilizados
- **Buttons**: `.btn`, `.btn-primary`, `.btn-secondary`, `.btn-destructive`, `.btn-success`, `.btn-outline`, `.btn-ghost`
- **Cards**: `.card` com suporte a hover effects e shadows
- **Badges**: `.badge-success`, `.badge-warning`, `.badge-destructive`, `.badge-info`
- **Forms**: `.form-input`, `.form-select`, `.form-label` com focus states
- **Alerts**: `.alert-success`, `.alert-destructive`, `.alert-warning`, `.alert-info`

### Paleta de Cores (HSL)
- **Primary**: `hsl(222.2 47.4% 11.2%)`
- **Secondary**: `hsl(210 40% 96.1%)`
- **Destructive**: `hsl(0 84.2% 60.2%)`
- **Success**: `hsl(142.1 76.2% 36.3%)`
- **Warning**: `hsl(38 92% 50%)`
- **Info**: `hsl(199 89% 48%)`

### Responsividade
O design é totalmente responsivo com breakpoints:
- `sm`: 640px
- `md`: 768px
- `lg`: 1024px

Para customizar o design, edite:
- `static/css/input.css` - Componentes e estilos customizados
- `tailwind.config.js` - Configuração do Tailwind

## 🎓 Curso

Este projeto foi desenvolvido durante o curso da Rocketseat sobre desenvolvimento backend com Go.

## 📄 Licença

Este projeto foi criado para fins educacionais.

