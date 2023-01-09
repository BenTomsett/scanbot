# Scanbot
CLI tool for interacting with and processing ImunifyAV malware scans

## Description
This tool allows you to interact with ImunifyAV using the command line. ImunifyAV has command line tools,
but these can be clunky to use. `scanbot` provides a simple list of commands for the most common tasks, and
allows you to present and export the results in machine-readable formats.

`scanbot` interacts directly with ImunifyAV's database, which is stored in SQLite format. Currently, the
tool looks for the database in `/var/imunify360/imunify360.db`, where it is stored on cPanel servers.
In the future, this may be configurable.

`scanbot` is basic, by design. New features will be added in the future :)

## Commands

### `scanbot list [flags]`
Return a list of the latest malware scans, with the most recent first.

#### Flags

```
  -c, --count    The number of results to return. (default 10)
```

<br/>

### `scanbot report <scan ID>`

If provided with a scan ID, this command will generate a CSV report for the specified scan. If no scan ID is provided, the user will be prompted to choose a scan from the latest 10 scans.


## Future features

- Testing
- Add a `scan` command to start a scan and email the results as a CSV attachment
- Add flag to specify the path to the Imunify database


## License
This project is licensed under the GNU General Public License v3.0 License - see the [LICENSE](https://github.com/BenTomsett/scanbot/blob/main/LICENSE) file for details
