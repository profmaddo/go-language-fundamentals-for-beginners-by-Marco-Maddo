# INITIAL SETUP – Go Programming Environment

This guide will help you install **Go**, **Visual Studio Code**, and **Git** on your system.  
Instructions are provided for **macOS**, **Linux**, and **Windows**, in **English and Portuguese**.

---

## 🖥️ macOS

### ✅ Step-by-step in English

#### 1. Install Go
- Go to the official site: [https://go.dev/dl](https://go.dev/dl)
- Download the `.pkg` file for macOS and run it.
- After installation, open the Terminal and type:
```bash
go version
```

#### 2. Install Visual Studio Code
- Visit: [https://code.visualstudio.com](https://code.visualstudio.com)
- Download the macOS version (`.zip`).
- Move the VS Code app to the `/Applications` folder.
- Open it and install the **Go extension**:
  - Press `Cmd+Shift+X`, search for **Go**, and install it.

#### 3. Install Git (optional)
- Open Terminal and type:
```bash
git --version
```
- If Git is not installed, install via Homebrew:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
brew install git
```

---

### ✅ Passo a passo em Português

#### 1. Instalar o Go
- Acesse: [https://go.dev/dl](https://go.dev/dl)
- Baixe o arquivo `.pkg` para macOS e execute-o.
- Após a instalação, abra o Terminal e digite:
```bash
go version
```

#### 2. Instalar o Visual Studio Code
- Visite: [https://code.visualstudio.com](https://code.visualstudio.com)
- Baixe a versão para macOS (`.zip`).
- Mova o aplicativo para a pasta `/Applications`.
- Abra o VS Code e instale a extensão **Go**:
  - Pressione `Cmd+Shift+X`, procure por **Go** e instale.

#### 3. Instalar o Git (opcional)
- Abra o Terminal e digite:
```bash
git --version
```
- Se o Git não estiver instalado, instale com o Homebrew:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
brew install git
```

---

## 🐧 Linux (Ubuntu/Debian)

### ✅ Step-by-step in English

#### 1. Install Go
```bash
sudo apt update
sudo apt install golang-go
go version
```

#### 2. Install Visual Studio Code
```bash
sudo apt update
sudo apt install wget gpg
wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
sudo install -o root -g root -m 644 microsoft.gpg /etc/apt/trusted.gpg.d/
sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list'
sudo apt update
sudo apt install code
```

- Open VS Code and install the Go extension (`Ctrl+Shift+X`, search for Go)

#### 3. Install Git
```bash
sudo apt install git
git --version
```

---

### ✅ Passo a passo em Português

#### 1. Instalar o Go
```bash
sudo apt update
sudo apt install golang-go
go version
```

#### 2. Instalar o Visual Studio Code
```bash
sudo apt update
sudo apt install wget gpg
wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
sudo install -o root -g root -m 644 microsoft.gpg /etc/apt/trusted.gpg.d/
sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list'
sudo apt update
sudo apt install code
```

- Abra o VS Code e instale a extensão Go (`Ctrl+Shift+X`, procure Go)

#### 3. Instalar o Git
```bash
sudo apt install git
git --version
```

---

## 🪟 Windows

### ✅ Step-by-step in English

#### 1. Install Go
- Visit [https://go.dev/dl](https://go.dev/dl)
- Download the Windows `.msi` installer and run it.
- Open `Command Prompt` or `PowerShell` and type:
```cmd
go version
```

#### 2. Install Visual Studio Code
- Visit [https://code.visualstudio.com](https://code.visualstudio.com)
- Download the installer and install it.
- Open VS Code, press `Ctrl+Shift+X`, search for **Go**, and install the extension.

#### 3. Install Git
- Visit: [https://git-scm.com](https://git-scm.com)
- Download and run the installer.
- Confirm installation with:
```cmd
git --version
```

---

### ✅ Passo a passo em Português

#### 1. Instalar o Go
- Acesse [https://go.dev/dl](https://go.dev/dl)
- Baixe o instalador `.msi` do Windows e execute.
- Abra o `Prompt de Comando` ou `PowerShell` e digite:
```cmd
go version
```

#### 2. Instalar o Visual Studio Code
- Acesse: [https://code.visualstudio.com](https://code.visualstudio.com)
- Baixe o instalador e instale.
- Abra o VS Code, pressione `Ctrl+Shift+X`, procure por **Go** e instale a extensão.

#### 3. Instalar o Git
- Visite: [https://git-scm.com](https://git-scm.com)
- Baixe e execute o instalador.
- Verifique com:
```cmd
git --version
```

---

✅ You're now ready to begin writing your first Go programs!  
✅ Agora você está pronto para começar a programar em Go!
