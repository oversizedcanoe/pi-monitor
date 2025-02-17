# pi-monitor
A script to monitor and track various metrics on a Raspberry Pi (_but actually works on most machines_).

### Method
This script uses [gopsutil](github.com/shirou/gopsutil/v4) to retrieve several metrics about the machine it is running on, and stores them to a local SQLite database.

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

### Intended Functionality / TODO
- Send emails when thresholds are met
- Create UI to view data (Web or CLI)
- Convert script to not run forever and instead use cron job
