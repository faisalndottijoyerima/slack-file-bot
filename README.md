# Slack File Bot

A Go utility that uploads a local file to a Slack channel.

## Requirements

- Go 1.26 or newer
- A Slack bot token with permission to upload files
- A Slack channel ID
- The file to upload in the project directory

## Configuration

Set these environment variables before running:

```text
SLACK_BOT_TOKEN=your-slack-bot-token
CHANNEL_ID=your-channel-id
```

Update the file path in `main.go` to point to the file you want to upload. Keep credentials out of source control.

## Run

```bash
go mod download
go run .
```

The program prints the uploaded file title and Slack file ID when successful.
