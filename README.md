# KBot - Telegram Traffic Light Control Bot

A Telegram bot built with Go that allows controlling traffic light signals. The bot enables users to toggle red, amber, and green lights through Telegram commands and a user-friendly interface with buttons.

## Link to Bot

You can access the bot at: t.me/kbot_bot

## Features

- Toggle traffic lights (red, amber, green) on/off
- Simple control panel with buttons for easy interaction
- Status monitoring of all lights
- Reset functionality to turn off all lights at once

## Installation

### Prerequisites

- Go 1.24.1 or later
- Telegram Bot Token (get from [@BotFather](https://t.me/BotFather))

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/oltsy77/kbot.git
   cd kbot
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up your Telegram bot token as an environment variable:
   ```bash
   export TELE_TOKEN="your_telegram_bot_token_here"
   ```

4. Build the application:
   ```bash
   go build
   ```

5. Run the bot:
   ```bash
   ./kbot start
   ```

## Usage

### Bot Commands

- `/start` - Show the control panel with buttons
- `/help` - Display available commands
- `/red` - Toggle red light on/off
- `/amber` - Toggle amber light on/off
- `/green` - Toggle green light on/off
- `/status` - Show current state of all lights
- `/reset` - Turn off all lights

### Starting the Bot

```bash
./kbot start
```

### Example Command Usage

1. Starting the bot:
   ```
   /start
   ```
   This command displays a welcome message and the control panel with buttons.

2. Toggling a light:
   ```
   /red
   ```
   This command toggles the red light on if it's off, or off if it's on.

3. Checking status:
   ```
   /status
   ```
   This command displays the current status of all lights (ON/OFF).

4. Resetting all lights:
   ```
   /reset
   ```
   This command turns off all lights.

## Development

### Build Configuration

You can customize the build by setting environment variables:

```bash
TARGETOS=linux    # Target OS (linux, darwin, windows)
TARGETARCH=arm64  # Target architecture (amd64, arm64)
```

### Project Structure

```
kbot/
├── cmd/
│   ├── kbot.go    # Main bot implementation and traffic light control
│   ├── root.go    # Root command configuration
│   └── version.go # Version command implementation
├── main.go        # Entry point
├── go.mod
├── go.sum
└── README.md
```

### Building from Source

```bash
go build -o kbot main.go
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Telebot](https://github.com/tucnak/telebot) - Telegram bot framework
