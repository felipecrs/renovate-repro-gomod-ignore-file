# Reproduction for PR #12345

The preset enables all dependencies from go datasource, but the config inside the `repositories` array disables updates for `project-b/go.mod`.

When Renovate runs, it still creates PRs for dependencies inside `project-b/go.mod` from go datasource, which is unexpected behavior.

This means the preset configuration is being applied after the repository-specific configuration, causing the ignore rule to be overridden.

That should not happen and does not happen in other scenarios like when defining the config directly in `renovate.json` inside the repository.

## Reproduction Steps

1. Renovate must be ran with something like:

    ```js
    // config.js
    // ...

    repositories: [
        {
        repository: 'felipecrs/renovate-repro-gomod-ignore-file',
        enabledManagers: ['gomod'],
        postUpdateOptions: ["gomodTidy"],
        packageRules: [
            {
            description: "Disable updates for project-b/go.mod",
            matchFileNames: ["project-b/go.mod"],
            enabled: false
            }
        ],
        extends: [
            'config:recommended',
            'https://raw.githubusercontent.com/felipecrs/renovate-repro-gomod-ignore-file/refs/heads/master/preset.json'
        ],
        }
    ],
    // ...
    ```

2. Note that the dependencies from `project-b/go.mod` from go datasource are still being updated.
