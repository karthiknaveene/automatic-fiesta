# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a CI/CD demonstration repository showcasing workflows across multiple platforms:
- **GitHub Actions** (`.github/workflows/`)
- **CloudBees Platform** (`.cloudbees/workflows/`)
- **Jenkins** (`Jenkinsfile`)

The repository is used for testing and learning CI/CD concepts, including artifact registration, workflow automation, and failure scenarios.

## Workflow Organization

### GitHub Actions Workflows

Located in `.github/workflows/`:

- **test-artifacts.yaml**: Registers build artifacts with CloudBees Platform using the `cloudbees-io-gha/register-build-artifact` action. Creates versioned Docker artifacts with format `2.0.$RUN_NUMBER-$RUN_ATTEMPT`.

- **failure.yaml**: Demonstrates intentional workflow failure (exits with status 1). Useful for testing error handling and observability.

- **prodgh.yaml, test1.yaml, prod-new.yml, gh1edited.yaml, gh2kn.yaml**: Simple test workflows that run on push or workflow_dispatch. These print start messages and are used for basic workflow validation.

### CloudBees Workflows

Located in `.cloudbees/workflows/`:

- **my-workflow.yaml**: CloudBees automation workflow using their API format (`apiVersion: automation.cloudbees.io/v1alpha1`). Runs a simple "hello world" in a Go Docker container on all branch pushes.

### Jenkins Pipeline

The `Jenkinsfile` defines a multi-stage pipeline:
1. **Build** stage (nested: Compile → Package)
2. **Registering build artifact** - Uses `registerBuildArtifactMetadata()` function
3. **Test** stage (Unit + Integration tests)
4. **Deploy** stage
5. **Final Status (Aborted)** - Intentionally times out to demonstrate abort handling

## Testing Workflows

### Triggering GitHub Actions
```bash
# Push to trigger workflows configured with 'on: push'
git push

# Manually trigger workflows with workflow_dispatch
gh workflow run <workflow-name>
```

### Validating Workflow Syntax
```bash
# Use actionlint to validate GitHub Actions workflows
actionlint .github/workflows/*.yaml

# Validate YAML syntax
yamllint .github/workflows/
```

### Testing Jenkins Pipeline
The Jenkinsfile can be tested locally using:
```bash
# Jenkins CLI (requires Jenkins setup)
java -jar jenkins-cli.jar -s http://localhost:8080/ build <job-name>
```

## CloudBees Integration

The repository integrates with CloudBees Platform for artifact registration:
- Uses action `cloudbees-io-gha/register-build-artifact@v3.0.2`
- Registers Docker artifacts with metadata (name, version, URL, type)
- Requires `id-token: write` permission for OIDC authentication
- Artifacts are tagged with semantic versioning based on GitHub run metadata

## Key Patterns

- Most workflows trigger on both `push` and `workflow_dispatch` events
- Test workflows use `ubuntu-latest` runners
- Sleep commands in Jenkinsfile simulate long-running build/test operations
- Workflows demonstrate various CI/CD scenarios: success, failure, timeout/abort
