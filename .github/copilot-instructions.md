# Copilot Instructions for pa-api

## Project Context
- This repository is a Go library for Amazon Creators API (migration from PA-API v5).
- Package import path and package name compatibility are important for users.
- Prefer backward-compatible changes unless explicitly asked for breaking changes.

## Coding Guidelines
- Keep changes minimal and focused on the requested task.
- Avoid unrelated refactors and broad formatting-only edits.
- Preserve exported API behavior and signatures when possible.
- Add concise comments only when logic is not obvious.
- Write all source-code comments in English.
- Prefer clear error handling over panic paths in library code.

## API and Migration Safety
- Creators API uses OAuth2 and x-marketplace header routing.
- Be careful with legacy compatibility methods and deprecated fields.
- If a legacy option is now ignored, document it clearly in README and comments.

## Tests and Validation
- Run tests after code changes:
  - task
- If behavior changes, add or update regression tests.
- Keep test names descriptive and focused on one behavior.

## Documentation
- Update README when behavior visible to users changes.
- Call out migration caveats explicitly.
- Keep examples aligned with current API behavior.

## Commit Hygiene
- Group related edits into a single logical commit.
- Do not include unrelated local changes in the same commit.
- Use short, descriptive commit messages.
