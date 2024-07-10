/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Lachignol/martin-solving/note"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type modelarray struct {
	table table.Model
}

var selectedChoice string
var selectedEdit string
var selectedDel = -1

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m modelarray) Init() tea.Cmd { return nil }

func (m modelarray) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "e":
			selectedEdit = string(m.table.SelectedRow()[0])
			return m, tea.Quit
		case "d":
			index, err := strconv.Atoi(m.table.SelectedRow()[0])
			if err != nil {
				fmt.Println(err)
			}
			selectedDel = index
			return m, tea.Quit
		case "enter":
			selectedChoice = string(m.table.SelectedRow()[0])

			return m, tea.Quit

		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m modelarray) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		columns := []table.Column{
			{Title: "Id", Width: 4},
			{Title: "Name", Width: 50},
			{Title: "Description", Width: 60},
		}
		notes := note.RecupNotes()
		var rows = []table.Row{}
		count := 1
		for _, note := range notes {
			rows = append(rows, table.Row{
				strconv.FormatInt(int64(count), 10),
				note.Name,
				note.Description})
			count++
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithHeight(7),
		)

		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false)
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false)
		t.SetStyles(s)

		m := modelarray{
			table: t,
		}
		if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}

		fmt.Println(selectedChoice)
		if selectedDel != -1 {
			note.DeleteNote(selectedDel)
		}

	},
}

func init() {
	noteCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
