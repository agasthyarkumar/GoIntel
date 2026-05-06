# GoIntel

> Concurrent File Intelligence & Security Toolkit built with Go

GoIntel is a high-performance command-line utility built using Go that combines:

- ⚡ Concurrent file scanning
- 🔍 SHA256 hashing
- 📂 Duplicate file detection
- 📊 Storage analytics
- 🔐 AES-256 file encryption/decryption
- 📦 ZIP compression

The project is designed to demonstrate real-world systems programming concepts using:

- Goroutines
- Channels
- Worker Pools
- Mutexes
- Concurrent Pipelines
- Filesystem Engineering
- Cryptography

---

# Features

## 🔍 File Scanning
- Recursively scans directories
- Processes files concurrently
- Extracts metadata

---

## 🔐 AES-256 Encryption
- Encrypt files securely using AES-GCM
- Generate encrypted `.enc` files

---

## 🔓 AES-256 Decryption
- Restore encrypted files using the same AES key

---

## 📂 Duplicate Detection
- Detect duplicate files using SHA256 hashing

---

## 📊 Storage Analytics
- Largest files
- Total storage usage
- Total file count

---

## 📦 Compression
- Compress files or folders into `.zip`

---

# Tech Stack

- Go (Golang)
- Cobra CLI Framework
- AES-GCM Cryptography
- Goroutines & Channels

---

# Project Structure

```bash
GoIntel/
│
├── cmd/
│   ├── root.go
│   ├── scan.go
│   ├── encrypt.go
│   ├── decrypt.go
│   └── compress.go
│
├── internal/
│   ├── analytics/
│   ├── checksum/
│   ├── compression/
│   ├── crypto/
│   ├── models/
│   ├── scanner/
│   └── worker/
│
├── main.go
├── go.mod
├── go.sum
└── README.md
```

---

# Installation

# 1. Clone Repository

```bash
git clone https://github.com/agasthyarkumar/GoIntel.git
```

```bash
cd GoIntel
```

---

# 2. Install Dependencies

```bash
go mod tidy
```

---

# 3. Build Binary

## Linux / macOS

```bash
go build -o goi
```

## Windows

```powershell
go build -o goi.exe
```

---

# Running the Tool

## Linux / macOS

```bash
./goi --help
```

## Windows

```powershell
.\goi.exe --help
```

---

# Commands

| Command | Alias | Description |
|---|---|---|
| scan | s | Scan directory |
| encrypt | e | Encrypt file |
| decrypt | d | Decrypt file |
| compress | c | Compress file/folder |
| help | h | Show help |

---

# Usage Examples

# 🔍 Scan Files

```bash
./goi scan ~/Desktop
```

OR

```bash
./goi s ~/Desktop
```

---

# 🔐 Encrypt File

```bash
./goi encrypt secret.txt
```

OR

```bash
./goi e secret.txt
```

Output:

```text
secret.txt.enc
```

---

# 🔓 Decrypt File

```bash
./goi decrypt secret.txt.enc
```

OR

```bash
./goi d secret.txt.enc
```

Output:

```text
secret.txt.enc.dec
```

---

# 📦 Compress Folder

```bash
./goi compress ~/Desktop/project
```

OR

```bash
./goi c ~/Desktop/project
```

Output:

```text
project.zip
```

---

# Sample Output

```text
🚀 Starting workers: 8

================================
File: /home/dell/Desktop/test.txt
Size: 120 bytes
SHA256: 8f434346...

==============================
🔥 DUPLICATE FILES
==============================
✅ No duplicate files found

==============================
📊 STORAGE ANALYTICS
==============================
Total Files: 42
Total Size: 120394 bytes

🔥 Largest Files
--------------------------------
File: dataset.zip
Size: 928394 bytes
```

---

# Using GoIntel as a Global CLI Tool

# Linux / macOS

After building:

```bash
sudo mv goi /usr/local/bin/
```

Now you can use:

```bash
goi s ~/Desktop
```

from anywhere.

---

# Windows

## Option 1 — Add Folder to PATH

Move `goi.exe` into a folder like:

```text
C:\GoIntel\
```

Then:

### Open:
- Windows Search
- "Environment Variables"
- Edit System Environment Variables
- Environment Variables
- Select PATH
- Add:

```text
C:\GoIntel\
```

Now restart terminal and run:

```powershell
goi.exe s C:\Users\YourName\Desktop
```

---

# Concurrency Architecture

GoIntel uses:

- Goroutines
- Channels
- Worker Pools
- Mutex synchronization

for concurrent file processing.

---

# Security

GoIntel currently uses:

- AES-256 GCM encryption
- SHA256 hashing

---

# Current Limitations

- Encryption key is hardcoded
- Compression stores absolute paths
- No password-based encryption yet

---

# Planned Features

- Password-derived encryption
- Concurrent folder encryption
- Progress bars
- File organization engine
- Storage visualization
- Checksum verification
- Smart duplicate cleanup

---

# Learning Concepts Demonstrated

This project demonstrates:

- Concurrent Programming
- Systems Programming
- Cryptography
- CLI Tool Engineering
- Worker Pool Architecture
- Mutex Synchronization
- File Compression
- Filesystem Traversal

---

# Contributing

Pull requests are welcome.

For major changes:
- Open an issue first
- Discuss proposed changes

---

# License

MIT License

---

# Author

Agasthya R

Built using Go ❤️