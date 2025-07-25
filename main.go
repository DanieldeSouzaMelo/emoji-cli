package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/loudercake/emoji-cli/utils"
	"github.com/atotto/clipboard"
)

const TABLE_HEIGHT = 5
const TABLE_WIDTH = 8

type model struct {
	emojis []utils.Emoji
	cursor int
	offset int
}

var selected_style = lipgloss.NewStyle().Background(lipgloss.Color("7"))

func initialModel() model {
	emoji_list := utils.GetEmojis()
	if (os.Args[1] != "") {
		emoji_list = utils.SearchEmojis(&emoji_list, os.Args[1])
	} 
	return model{
		emojis: emoji_list,
		cursor: 0,
		offset: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func processScroll(cursor int, offset *int) {
	if (cursor < (*offset * TABLE_WIDTH) ) {
		*offset--
	} else if (cursor > (TABLE_WIDTH * TABLE_HEIGHT) + (*offset * TABLE_WIDTH) - 1) {
		*offset++
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "up", "k":
			if (m.cursor - TABLE_WIDTH > 0) {
				m.cursor -= TABLE_WIDTH
				processScroll(m.cursor, &m.offset)
			}
		case "down", "j":
			if (m.cursor + TABLE_WIDTH < len(m.emojis)) {
				m.cursor += TABLE_WIDTH
				processScroll(m.cursor, &m.offset)
			}
		case "left", "h":
			if (m.cursor > 0) {
				m.cursor--
				processScroll(m.cursor, &m.offset)
			}
		case "right", "l":
			if (m.cursor + 1 < len(m.emojis)) {
				m.cursor++
				processScroll(m.cursor, &m.offset)
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			clipboard.WriteAll(m.emojis[m.cursor].Emoji)
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	// s := fmt.Sprintf("cursor: %d, offset: %d\n", m.cursor, m.offset)
	s := ""
	table_range_start := (m.offset * TABLE_WIDTH)
	table_range_end := (TABLE_WIDTH * TABLE_HEIGHT) + (m.offset * TABLE_WIDTH)
	i := table_range_start
	width_counter := 0
	for i < table_range_end {
		if (i > len(m.emojis) - 1) {
			break
		}
		if (m.cursor == i) {
			s += selected_style.Render(m.emojis[i].Emoji)	
		} else {
			s += m.emojis[i].Emoji
		}
		width_counter++
		if (width_counter == TABLE_WIDTH) {
		s += "\n" 
		width_counter = 0
		}
		i++
	}
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
