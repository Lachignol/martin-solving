/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
var selectedNew = false
var selectedEdit string
var selectedToggle = -1

var selectedDel = -1

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240")).
	MarginTop(30).Align()

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
		case "t":
			index, err := strconv.Atoi(m.table.SelectedRow()[0])
			if err != nil {
				fmt.Println(err)
			}
			if m.table.SelectedRow()[2] == "✅" {
				m.table.SelectedRow()[2] = "❌"
				m.table.SelectedRow()[4] = ""
			} else {
				currentTime := time.Now()
				m.table.SelectedRow()[2] = "✅"
				m.table.SelectedRow()[4] = currentTime.Format("02 January 2006 15:04:05")
			}
			note.Toggle(index)
			m.table.MoveDown(1)
			return m, nil
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
		case "n":
			selectedNew = true
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
	return baseStyle.Render(m.table.View()) + "\n" +
		HelpStyle.Render("[ Tapez q ou ctrl+c pour quitter ]") + " " +
		HelpStyle.Render("[ Naviguer avec ⬆ et ⬇ ]") + " " +
		HelpStyle.Render("[ Tapez t completer/décompleter la tache ]") + " " +
		HelpStyle.Render("[ Tapez n pour ajouter une tache ]") + " " +
		HelpStyle.Render("[ Tapez d pour supprimer la tache ]") + " "
	// HelpStyle.Render("[ Tapez e pour editer la tache ]")
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
			{Title: "Titre", Width: 109},
			{Title: "Completed", Width: 9},
			{Title: "Created_at", Width: 22},
			{Title: "Completed_at", Width: 22},
		}
		notes := note.RecupNotes()
		var rows = []table.Row{}
		count := 1
		for _, note := range notes {
			completed := note.Completed
			completedAt := ""
			var iscompleted string
			if !completed {
				iscompleted = "❌"
			} else {
				iscompleted = "✅"
			}
			if note.Completed_at != nil {

				completedAt = note.Completed_at.Format("02 January 2006 15:04:05")
			}

			rows = append(rows, table.Row{
				strconv.FormatInt(int64(count), 10),
				note.Title,
				iscompleted,
				note.Created_at.Format("02 January 2006 15:04:05"),
				completedAt,
			})
			count++
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithHeight(20),
		)

		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(true)
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
		if selectedToggle != -1 {
			err := note.Toggle(selectedToggle)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("status de la tache correctement changé")
			}
		}
		if selectedNew {
			newCmd.Run(cmd, []string{})
			// if err != nil {
			// 	fmt.Println(err)
			// }

		}
	}}

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
