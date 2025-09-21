# goa-boilerplate

A boilerplate for starting a new project using [goa](https://goa.design/).

## Features

- Basic project structure with folders for API, services, and models.
- Pre-configured `Makefile` for common tasks like code generation, building, and running tests.
- Example goa design file to get you started quickly.
- Dockerfile for containerizing your application.
- GitHub Actions workflow for CI/CD.

## Getting Started

- 1. Install copier with pipx:

```bash
  pipx install copier
```

- 2. Use copier to create a new project from this template:

```bash
  copier copy gh:applinh/goa-boilerplate . --trust -x .git
```

- 3. Answer the prompts to customize your project, and you're ready to go!