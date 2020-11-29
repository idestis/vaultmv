# VAULTMV

**VAULTMV** is a simple go application which allows you to move or extend your vault path with ease. 

Probably once you created a path that was good at the moment. This path already contains a lot of sensitive data and the next day you faced the major issue, that you need to move all secrets for one level deeper to divide logical things. A manual copy of a huge JSON to the file and past it back is a regular method. But we are engineers, let's automate the routine. 

## Usage

> Refer to Install block for getting latest `vaultmv` binary

To use `vaultmv` at your local machine, download the latest binary for your operating system from the project [release page](https://github.com/idestis/vaultmv/releases)

```bash
$ vaultmv do --source secret/data/foo/bar --dest secret/data/foo/bar/add
```

## Bulk

Bulk option allow you to run this tool over `.csv` file.

```csv
secret/data/from,secret/data/to,true
```

First two fields is to define **secret path movements**, the third field stands for **permanent** delete of the source. **TBD** in version 0.0.2

### **Examples:**

In example below we are trying to resolve two issues with Vault:

- We need to move **secret/foo/bar** to **secret/bar/foo**
- We need to extend path **secret/services/awesome** to **secret/services/awesome/fonts** to be able to use space in **secret/services/awesome** with additonal paths.

**KV Secrets Engine - Version 1 (API)**
https://www.vaultproject.io/api-docs/secret/kv/kv-v1

```csv
secrets/foo/bar,secrets/bar/foo,true
secrets/services/awesome,secrets/services/awesome/fonts,true
```

**KV Secrets Engine - Version 2 (API)**
https://www.vaultproject.io/api-docs/secret/kv/kv-v2

```csv
secret/data/foo/bar,secret/data/bar/foo,true
secret/data/services/awesome,secret/data/services/awesome/fonts,true
```

## Install

#### **Binary**

Binary are available for download on the project [release page](https://github.com/idestis/vaultmv/releases)

However, you also able to change the source code and build your `vaultmv` for yourself

```bash
$ go build ./...
```

## Contribute

Refer to CONTRIBUTE.md

## Future updates

Star the repo to stay tunned.

- Complete deletion of the source secret path.
- Go routines to complete bulk changes quicker.