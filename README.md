# pi-monitor
A script to monitor and track various metrics on a Raspberry Pi

- Create a loop that lasts forever.
- At the end of the loop, sleep for a set amount of time
- Possibly could end up not doing a loop/sleep and setting up a cron job to execute this script once, but the structure stand
- Capture system values: Temperature, CPU%, GPU%, Ram usage, Storage usage, other
- Save to SQLite database
- Send emails when threshholds are met.
- Allow user to view data somehow: Web UI, daily emails, CLI
