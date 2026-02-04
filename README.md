# BMS - Business Management System

A professional Business Management System API built with Go and PostgreSQL.

## 📋 Project Information


## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 14+
- Git

### Installation
```bash
# Clone repository
git clone https://github.com/shiiqmohamed/bms.git
cd bms

# Setup environment
cp .env.example .env
# Edit .env with your database credentials

# Setup database
./scripts/db-setup.sh

# Install dependencies
go mod download

# Run application
go run cmd/main.go