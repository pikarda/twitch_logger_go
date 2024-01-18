package main

import "github.com/charmbracelet/lipgloss"

func styledText(color string, prompt string) string {
	text := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color(color)).
		Padding(0, 1)

	return text.Render(prompt)
}

func styledUser(prompt string) string {
	text := lipgloss.NewStyle().
		Bold(true)

	return text.Render(prompt)
}

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
