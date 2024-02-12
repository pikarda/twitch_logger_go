package main

import "github.com/charmbracelet/lipgloss"

// add some colors to streamer username when logging (in console)
func styledText(color string, prompt string) string {
	text := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color(color)).
		Padding(0, 1)

	return text.Render(prompt)
}

// add boldness to chatter username when logging (in console)
func styledUser(prompt string) string {
	text := lipgloss.NewStyle().
		Bold(true)

	return text.Render(prompt)
}

// styling of text that pops when app started (in console)
func styledStartApp(prompt string) string {
	text := lipgloss.NewStyle().
		Width(50).
		Height(3).
		PaddingTop(1).
		Align(lipgloss.Center).
		Bold(true).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color("63")).
		BorderTop(true).
		BorderLeft(true).
		BorderBottom(true).
		BorderRight(true).
		Background(lipgloss.Color("63")).
		MarginBottom(1)

	return text.Render(prompt)
}
