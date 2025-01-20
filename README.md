# Tablecheck Pinger

This project is a simple Go application that queries a specified tablecheck.com URL and checks for available slots in the returned JSON response. If available slots are found, it sends a notification to a specified Telegram channel.

## Using the binary

### Prerequisites

- A Telegram bot API token

### Running the binary

- Download the binary: [`bin/linux`](./bin/linux/tablecheck-pinger) or [`darwin`](./bin/darwin/tablecheck-pinger)
- Execute it: `TELEGRAM_BOT_API_TOKEN=<token> ./tablecheck-pinger <url> <channelID>`

### Example

```
TELEGRAM_BOT_API_TOKEN=123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11 ./tablecheck-pinger https://www.tablecheck.com/en/shops/my-restaurant/menu_items?num_people=2&date=2025-01-01 -1234567890
```

## Building from source

### Prerequisites

- Go 1.20 or later
- A Telegram bot API token

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/amanske/tablecheck-pinger-go
   cd tablecheck-pinger-go
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

### Building the Application

To build the application, run:
```
make build
```

### Running the Application

To run the application, use:
```
TELEGRAM_BOT_API_TOKEN=<your-telegram-bot-api-token> ./bin/tablecheck-pinger-go <url> <channel_username>
```

### Example

```
TELEGRAM_BOT_API_TOKEN=123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11 ./bin/tablecheck-pinger-go https://www.tablecheck.com/en/shops/my-restaurant/menu_items?num_people=2&date=2025-01-01 -1234567890
```
