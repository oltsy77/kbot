/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "Telegram bot for controlling traffic light signals",
	Long: `A Telegram bot that allows controlling traffic light signals through GPIO pins.
The bot accepts commands to toggle red, amber, and green lights on/off.
Usage:
  /start - Show control panel with buttons
  /help - Display available commands
  /red - Toggle red light on/off
  /amber - Toggle amber light on/off
  /green - Toggle green light on/off
  /status - Show current state of all lights
  /reset - Turn off all lights`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started\n", appVersion)

		if TeleToken == "" {
			log.Fatal("TELE_TOKEN environment variable is not set")
		}

		fmt.Println("Telegram bot started...")

		// Map to store the current state of each light
		lightStatus := make(map[string]bool)
		lightStatus["red"] = false
		lightStatus["amber"] = false
		lightStatus["green"] = false

		// Create a new bot instance
		kbot, err := telebot.NewBot(telebot.Settings{
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Failed to create bot: %s", err)
		}

		// Create buttons for the keyboard
		menu := &telebot.ReplyMarkup{ResizeKeyboard: true}

		btnRed := menu.Text("🔴 Red")
		btnAmber := menu.Text("🟠 Amber")
		btnGreen := menu.Text("🟢 Green")
		btnStatus := menu.Text("📊 Status")
		btnReset := menu.Text("🔄 Reset")

		menu.Reply(
			menu.Row(btnRed, btnAmber, btnGreen),
			menu.Row(btnStatus, btnReset),
		)

		// Handle /start command
		kbot.Handle("/start", func(m telebot.Context) error {
			welcome := fmt.Sprintf("Hello! I'm Traffic Light Bot %s!\nUse the buttons below to control the traffic light:", appVersion)
			return m.Send(welcome, menu)
		})

		// Handle /help command
		kbot.Handle("/help", func(m telebot.Context) error {
			helpText := `Traffic Light Bot Commands:
/start - Show control panel
/status - Show current light status
/red - Toggle red light
/amber - Toggle amber light
/green - Toggle green light
/reset - Turn off all lights
/help - Show this help message`
			return m.Send(helpText)
		})

		// Handle button and command for red light
		kbot.Handle(&btnRed, func(m telebot.Context) error {
			return toggleLight(m, "red", lightStatus)
		})
		kbot.Handle("/red", func(m telebot.Context) error {
			return toggleLight(m, "red", lightStatus)
		})

		// Handle button and command for amber light
		kbot.Handle(&btnAmber, func(m telebot.Context) error {
			return toggleLight(m, "amber", lightStatus)
		})
		kbot.Handle("/amber", func(m telebot.Context) error {
			return toggleLight(m, "amber", lightStatus)
		})

		// Handle button and command for green light
		kbot.Handle(&btnGreen, func(m telebot.Context) error {
			return toggleLight(m, "green", lightStatus)
		})
		kbot.Handle("/green", func(m telebot.Context) error {
			return toggleLight(m, "green", lightStatus)
		})

		// Handle status button and command
		kbot.Handle(&btnStatus, func(m telebot.Context) error {
			return showStatus(m, lightStatus)
		})
		kbot.Handle("/status", func(m telebot.Context) error {
			return showStatus(m, lightStatus)
		})

		// Handle reset button and command
		kbot.Handle(&btnReset, func(m telebot.Context) error {
			return resetLights(m, lightStatus)
		})
		kbot.Handle("/reset", func(m telebot.Context) error {
			return resetLights(m, lightStatus)
		})

		// Handle text messages (replacing your original handler)
		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Printf("Received message: %s", m.Text())

			text := strings.ToLower(m.Text())

			// Handle commands with or without slash
			switch {
			case text == "hello":
				return m.Send(fmt.Sprintf("Hello, I'm Traffic Light Bot %s!", appVersion), menu)
			case text == "red" || text == "/red":
				return toggleLight(m, "red", lightStatus)
			case text == "amber" || text == "/amber":
				return toggleLight(m, "amber", lightStatus)
			case text == "green" || text == "/green":
				return toggleLight(m, "green", lightStatus)
			case text == "status" || text == "/status":
				return showStatus(m, lightStatus)
			case text == "reset" || text == "/reset":
				return resetLights(m, lightStatus)
			case text == "help" || text == "/help":
				return m.Send(`Available commands:
- red - Toggle red light
- amber - Toggle amber light
- green - Toggle green light
- status - Show current status
- reset - Turn off all lights
- help - Show this help message`)
			default:
				return m.Send("Unknown command. Type 'help' to see available commands.", menu)
			}
		})

		fmt.Println("Bot is ready to accept commands...")
		kbot.Start()
	},
}

// Function to toggle a light
func toggleLight(m telebot.Context, color string, status map[string]bool) error {
	status[color] = !status[color]

	var emoji, stateText string
	if status[color] {
		stateText = "ON"
	} else {
		stateText = "OFF"
	}

	switch color {
	case "red":
		emoji = "🔴"
	case "amber":
		emoji = "🟠"
	case "green":
		emoji = "🟢"
	}

	return m.Send(fmt.Sprintf("%s %s light turned %s", emoji, strings.Title(color), stateText))
}

// Function to show the status of all lights
func showStatus(m telebot.Context, status map[string]bool) error {
	message := "Current Traffic Light Status:\n"

	if status["red"] {
		message += "🔴 Red: ON\n"
	} else {
		message += "⚫ Red: OFF\n"
	}

	if status["amber"] {
		message += "🟠 Amber: ON\n"
	} else {
		message += "⚫ Amber: OFF\n"
	}

	if status["green"] {
		message += "🟢 Green: ON"
	} else {
		message += "⚫ Green: OFF"
	}

	return m.Send(message)
}

// Function to reset all lights
func resetLights(m telebot.Context, status map[string]bool) error {
	status["red"] = false
	status["amber"] = false
	status["green"] = false

	return m.Send("🔄 All lights have been turned OFF")
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
