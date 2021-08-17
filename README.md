# dnode

`dnode` is a CLI tool to search and delete any `node_module` folders.

- [1. Instillation](#1-instillation)
  - [1.1. Via Go CLI](#11-via-go-cli)
  - [1.2. Build it yourself](#12-build-it-yourself)
  - [1.3. Download prebuilt binaries](#13-download-prebuilt-binaries)
- [2. Usage](#2-usage)
  - [2.1. list](#21-list)
  - [2.2. delete](#22-delete)

## 1. Instillation

There are three ways to install the binary

### 1.1. Via Go CLI

Open your terminal and type in

```sh
go install github.com/akshaybabloo/dnode@latest
```

### 1.2. Build it yourself

```bash
git clone github.com/akshaybabloo/dnode
cd dnode
go install
```

### 1.3. Download prebuilt binaries

You can download OS dependent binary from the [releases](https://github.com/akshaybabloo/dnode/releases/latest) page.

## 2. Usage

There are two commands available

### 2.1. list

Lists all `node_module` folders in the current directory tree recursively.

**Example**

```sh
> dnode list
                   PATH                  | DIRECTORY SIZE
-----------------------------------------+-----------------
  ./gollahalli.com/node_modules          |     33 MB
  ./keertana.gollahalli.com/node_modules |     34 MB
-----------------------------------------+-----------------
                  TOTAL                  |     67 MB
-----------------------------------------+-----------------
```

### 2.2. delete

Deletes all `node_module` folders in the current directory tree recursively.

```sh
> dnode delete
Deleted ./gollahalli.com/node_modules
Deleted ./keertana.gollahalli.com/node_modules
67 MB freed
```