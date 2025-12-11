# workflow-test

A simple test repository to validate CI/CD release workflows on both **GitHub** and **GitLab**.

## Purpose

This project is used to test:
- GitHub Actions release workflow
- GitLab CI/CD release pipeline
- Cross-platform builds (Linux, Windows, macOS)
- Package generation (.deb, .rpm)
- Automatic release creation

## Building

```bash
go build -o workflow-test ./cmd/app
```

## With Version

```bash
go build -ldflags="-X main.appVersion=1.0.0" -o workflow-test ./cmd/app
```

## Repositories

- **GitHub**: https://github.com/cantalupo555/workflow-test
- **GitLab**: https://gitlab.com/cantalupo555-group/workflow-test
