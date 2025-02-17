# pi-monitor
A script to monitor and track various metrics on a Raspberry Pi.

### Functionality
This script uses [gopsutil](github.com/shirou/gopsutil/v4) to retrieve several metrics about the machine it is running on, and stores them to a local SQLite database.

It can be configured to send out warning emails when threshholds are exceeded.

> [!NOTE]
> Because [gopsutil](github.com/shirou/gopsutil/v4) works on most OSs, this application is not specific to Raspberry Pis. However, it may need to be run as an Administrator.

### Metrics
Currently, the following metrics are stored:
| Metric | Unit |
| ------ | ---- |
| Temperature | º Celcius |
| CPU Usage | % |
| Disk 1 Usage | % |
| Disk 2 Usage | % |
| Memory (Ram) Usage | % |
| Uptime | Seconds |

### Installation
- Install Go 1.24
- Rename or copy `config.env.example` to `config.env` and replace all email-related config values
- Optionally, edit the threshhold values located in the config file
- Run `go run .`, or `go build` and run the resulting application

### Intended Functionality / TODO
- Create UI to view data (Web or CLI)
- Convert script to not run forever and instead use cron job
