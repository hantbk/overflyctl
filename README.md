# Overfly Cloud Command Line Tool

## Overview

Overfly Cloud Command Line is a powerful tool designed to interact with the Overfly Cloud API. It provides a seamless interface for managing Linux servers, executing remote commands, and handling SSH connections.

## Features

### 🖥️ Server Management

- **Add Servers**: Easily add new Linux servers to your database.
- **List Servers**: View all your registered servers at a glance.
- **Delete Servers**: Remove servers from your database when no longer needed.

### 🔐 SSH Operations

- **Connect**: Establish secure SSH connections to your servers.
- **Execute Commands**: Run commands on remote servers effortlessly.

### 🔑 Key Management

- **Automatic Key Copying**: Seamlessly copy SSH keys to your servers for passwordless authentication.

## Commands

### `add`

Add a new Linux server to your database.

```bash
overfly add --name <server_name> --ip <ip_address> --username <ssh_username> --password <ssh_password> --ssh-key <path_to_ssh_key>
```

### `list`

Display all registered Linux servers.

```bash
overfly list
```

### `delete`

Remove a server from the database.

```bash
overfly delete --id <server_id>
```

### `connect`

Establish an SSH connection to a server.

```bash
overfly connect --server <server_name> --username <ssh_username>
```

### `execute`

Execute a command on a remote server.

```bash
overfly execute --server <server_name> --command <command_to_execute>
```

## Getting Started

1. Install the Overfly Cloud Command Line tool.
2. Use the `add` command to register your first server.
3. Explore other commands to manage and interact with your servers.

## Support

For any issues or feature requests, please open an issue on our GitHub repository.

## License

This project is licensed under the Apache License 2.0. See the LICENSE file for details.

---

Empower your cloud management with Overfly Cloud Command Line Tool! 🚀



