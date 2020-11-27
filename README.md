# VAULTMV

**VAULTMV** is a simple go application which allows you to move or extend your vault path with ease. 

Probably once you created a path that was good at the moment. This path already contains a lot of sensitive data and the next day you faced the major issue, that you need to move all secrets for one level deeper to divide logical things. A manual copy of a huge JSON to the file and past it back is a regular method. But we are engineers, let's automate the routine. 

## Usage

> Refer to Install block for getting latest `vaultmv` binary

To use `vaultmv` at your local machine, download the latest binary for your operating system from the project [release page](https://github.com/idestis/vaultmv/releases)


## Bulk

Bulk option allow you to run this tool over `.csv` file.

```csv
secret/data/from,secret/data/to,true
```

First two fields is to define **secret path movements**, the third one is designed for **permanent** delete of previous location.

### Examples:

**KV API V1**

```csv
secrets/foo/bar,secrets/foo/bar/environment,true
secrets/services/awesome,secrets/services/awesome/fonts,true
```

**KV API V2** require us to write addional information to the path.

```csv
secret/data/foo/bar,secret/data/foo/bar/environment,true
secret/data/services/awesome,secret/data/services/awesome/fonts,true
```