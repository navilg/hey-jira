# hey-jira
A simple command line tool to work with Jira

## Plan

1. `hey-jira configure`

This command should configure username and token of the default user to use for authentication.
This will store the server url, username and token in home directory of user in json format.
Token/Password will be base64 encoded before storage.
Configuration will also  include default project and proxy to use.
It will create profile. Json file name will be `.hey-jira/<profile>/config.json`. Later command can be run by passing --profile as well to use specific authentication out of many configured.
By default, default profile will be configured.

2. `hey-jira --profile <profile-name> close-issues --status=<status> --since=Nd

This closes the issues which is in status <status> since last N days.