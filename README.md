### Требования

- **Go** ≥ `1.24`
- Docker
- **Taskfile CLI** → [инструкция по установке](https://taskfile.dev/#/installation)

### CI/CD

Проект использует GitHub Actions. Основные workflow:

- **CI** (`.github/workflows/ci.yml`) - проверяет код при каждом push и pull request
  - Запуск линтера golangci-lint

