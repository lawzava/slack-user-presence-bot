# slack-user-presence-bot

The app has 2 functions:
1. Tracks users presence (online/offline)
2. Reports users presence statistis

## Workflow:
- Uses SQLite as data storage
- Checks for user status every 10 minutes and inserts only active times
- Reports total activity in hours (for custom timespan)

## How to use:
### On 64bit linux machine:
Download binary from /bin/linux and run with **SLACK_TOKEN** and **PORT** as environmental variables.

### Any other machine:
Download code, compile and as noted above.

After that configure your slack slash commands to endpoints of: /help, /total.

## Commands
*Date format: 2018-01-01*
/total [start date] [end date]
/?

### Note:
It will give list of all users, sorted by their presence, descending.

## TODO
- Pretify responses
- Error reporting
- Tests
- Auto-revive
- Optimize database
