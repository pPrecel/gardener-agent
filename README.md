# Cloud Agent

[![license](https://img.shields.io/badge/License-MIT-brightgreen.svg?style=for-the-badge)](https://github.com/pPrecel/cloudagent/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/pPrecel/cloudagent?style=for-the-badge)](https://goreportcard.com/report/github.com/pPrecel/cloudagent)
[![build](https://img.shields.io/github/workflow/status/pPrecel/cloudagent/build?style=for-the-badge)](https://github.com/pPrecel/cloudagent/actions/workflows/build.yml)
[![Coverage](https://img.shields.io/coveralls/github/pPrecel/cloudagent?style=for-the-badge)](https://coveralls.io/github/pPrecel/cloudagent)

The simple and easy-to-use program is designed to watch user activity and possible orphan clusters for Cloud Providers:

- Gardener
- GCP (work in progress)

This application is created with a view to using it as [the tmux](https://github.com/tmux/tmux) status. To fulfill this criterion the procedure of getting resources from the gardener is separated and is in the second service which serves the UNIX socket that is used by the first one. This architecture allows not to block the main tmux process during calling the right gardener endpoint.

## Installation

### Github Release

Visit the [releases page](https://github.com/pPrecel/cloudagent/releases) to download one of the pre-built binaries for your platform.

### Homebrew

1. Use Homebrew to install `cloudagent`:

    ```bash
    brew install pPrecel/tap/cloudagent
    ```

    or

    ```bash
    brew tap pPrecel/tap
    brew install cloudagent
    ```

2. Extend configuration ([see also](./docs/configuration-file.md)) in the `~/.cloudagent.conf.yaml` location.

3. Start the `cloudagent` service:

    ```bash
    brew services start cloudagent
    ```

### Local development

1. Verify and build the program:

    ```bash
    make verify
    make build
    ```

2. Extend configuration ([see also](./docs/configuration-file.md)) in the `~/.cloudagent.conf.yaml` location.

3. Add the program to PATH and install it as a system agent:

    ```bash
    make ln-to-path
    make install-agent
    ```

    > **NOTE:** for local development or need to get more information from the agent you can pass more arguments to the `make install-agent` command like: `other_flags=--agentVerbose`.

4. Check if the program works by getting its logs:

    ```bash
    tail /tmp/cloudagent.stdout
    ```

5. After waiting ~60 seconds for the first iteration of the watcher you can get the cluster state:

    ```bash
    cloudagent state
    ```

## Integration with tmux

To add this application to tmux put the line below in the `~/.tmux.conf` file:

```text
set -ag status-right ' #(cloudagent state --createdBy <OWNER_NAME> -o text) '
```
