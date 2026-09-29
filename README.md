# Go Keeper

Go Keeper organizes files in your `Downloads` and `Desktop` folders and removes
old files according to its configuration.

> [!WARNING]
> Cleanup commands can permanently delete expired files. Review your
> configuration with `gokeeper config show` before running them.

## Build and run

This project requires Go 1.26.3 or later.

Run a command directly from the source code:

```sh
go run . <command>
```

Or build a local executable:

```sh
go build -o gokeeper .
./gokeeper <command>
```

If the executable is installed in a directory on your `PATH`, omit `./` and
run it as `gokeeper`.

## Commands

### `gokeeper config show`

Displays the active retention settings without cleaning any files:

```sh
./gokeeper config show
```

The default settings are 90 days for code, 30 days for documents, 30 days for
installers, 60 days for media, 30 days for other downloads, and 30 days for
screenshots.

### `gokeeper config path`

Prints the location of the JSON configuration file:

```sh
./gokeeper config path
```

Go Keeper creates this file with the default settings the first time it runs.
Edit the file to change how many days files are retained. Setting a download
category to `null` disables moving and deleting files in that category.

### `gokeeper downloads`

Organizes and cleans the current user's `~/Downloads` folder:

```sh
./gokeeper downloads
```

It creates these folders under `~/Downloads`:

- `Code Related`
- `Documents Related`
- `Installers`
- `Media Related`
- `Others`

Files in the root of `~/Downloads` are classified by extension. Files younger
than their configured retention period are moved into the matching folder.
Expired files are permanently deleted, including expired files already inside
the category folders.

### `gokeeper desktop`

Organizes and cleans the current user's `~/Desktop` folder:

```sh
./gokeeper desktop
```

It creates `~/Desktop/Screenshots`, moves image files from the Desktop into
that folder, and moves other Desktop files into `~/Downloads`. Inside the
Screenshots folder, files whose names begin with `Screenshot` are permanently
deleted when they exceed `screenshotsExpirationDays`.

Existing directories on the Desktop are left unchanged. If a non-image file
already exists in Downloads with the same name, that Desktop file is skipped.

## Running without building

Every command has an equivalent `go run` form:

```sh
go run . config show
go run . config path
go run . downloads
go run . desktop
```
