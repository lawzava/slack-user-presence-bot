# slack-user-presence-bot

The app has 2 functions:
1. Tracks users presence (online/offline)
2. Reports users presence statistics (average daily/weekly/monthly; total daily/weekly/monthly)

## Workflow:
- Uses SQLite as data storage
- Checks for user status every 10 minutes and inserts only active times
- Reports average/total daily/weekly/monthly activity in hours (for custom timespan)

## How to use:
### On 64bit linux machine:
Download binary from /bin/linux and run with **SLACK_TOKEN** as environmental variable.

### Any other machine:
Download code, compile and as noted above.


## Commands
*Date format: 2018-01-01*
/average [daily/weekly/monthly] [start date] [end date]
/total [start date] [end date]

### Note:
It will give list of all users, sorted by their presence, descending.

## TODO
- Pretify responses
- Error reporting
- Tests
- Auto-revive
