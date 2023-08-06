package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/extism/extism"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Monster struct {
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Message   string `json:"message"`
	Action    string `json:"action"`
	direction Direction
}

type NonPlayerCharacter struct {
}

type Game struct {
	title              string
	bords              Bords
	player             Player
	monsters           map[string]Monster
	nonPlayerCharacter map[string]NonPlayerCharacter
}

type model struct {
	game Game
}

func initialModel(game Game) model {
	return model{
		game: game,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.game.player.Move(North, m.game.bords)

		case "down":
			m.game.player.Move(South, m.game.bords)

		case "right":
			m.game.player.Move(East, m.game.bords)

		case "left":
			m.game.player.Move(West, m.game.bords)

		case "enter", " ":
			m.game.title = "👋 Hello World 🌍"
		}
	}

	return m, nil
}

func (m model) View() string {
	s := m.game.title + "\n\n"

	for row := 0; row < m.game.bords.ground.height; row++ {
		s += fmt.Sprintf("%02d ", row)
		for col := 0; col < m.game.bords.ground.width; col++ {

			if m.game.player.x == col && m.game.player.y == row {
				s += m.game.player.avatar
			} else {
				objectBoxe := m.game.bords.objects.boxes[row][col]
				groundBoxe := m.game.bords.ground.boxes[row][col]

				if objectBoxe == Nothing {
					s += string(groundBoxe)
				} else {
					s += string(objectBoxe)
				}
			}
			//s += m.playGround.ground[i][j]
		}

		s += "\n"
	}
	s += "\n"

	s += fmt.Sprintf("🧭 %s x:%02d y:%02d 👀 N:%s S:%s E:%s W:%s [%s]",
		m.game.player.direction,
		m.game.player.x,
		m.game.player.y,
		m.game.player.LookAt(North, m.game.bords),
		m.game.player.LookAt(South, m.game.bords),
		m.game.player.LookAt(East, m.game.bords),
		m.game.player.LookAt(West, m.game.bords),
		m.game.player.LookBelow(m.game.bords),
	)
	s += "\nPress q to quit.\n"

	return s
}

var game = Game{
	title: "👽 Forbidden Planet 🚀",
	bords: Bords{
		ground: Board{
			boxes: [][]Boxe{
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟦", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟦", "🟦", "🟦", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟦", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟦", "🟦", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟦", "🟦", "🟦", "🟦", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟦", "🟦", "🟦", "🟦", "🟦", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟦", "🟦", "🟦", "🟦", "🟦", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟦", "🟦", "🟦", "🟩", "🟩", "🟩", "🟩"},
				{"🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩", "🟩"},
			},
			width:  15,
			height: 15,
		},
		objects: Board{
			boxes: [][]Boxe{
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "🎃", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "🔥", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "🌸", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "🍔", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
				{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
			},
			width:  15,
			height: 15,
		},
	},

	player: Player{
		avatar: "🤖",
		name:   "Robby",
		x:      2,
		y:      3,
	},
}

func main() {

	wasmFilePath := os.Args[1:][0]

	ctx := context.Background()
	config := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
	}
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFilePath},
		},
	}

	display := extism.HostFunction{
		Name:      "hostDisplay",
		Namespace: "env",
		Callback: func(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

			// Read function arguments
			offset := stack[0]
			bufferInput, err := plugin.ReadBytes(offset)

			if err != nil {
				fmt.Println("🥵", err.Error())
				panic(err)
			}

			monsterData := string(bufferInput)

			fmt.Println("🟠 " + monsterData)

			// Return data
			plugin.Free(offset)

			offset, err = plugin.WriteBytes([]byte(""))

			if err != nil {
				fmt.Println("😡", err.Error())
				panic(err)
			}

			stack[0] = offset
		},
		Params:  []api.ValueType{api.ValueTypeI64},
		Results: []api.ValueType{api.ValueTypeI64},
	}


	sendMessage := extism.HostFunction{
		Name:      "hostSendMessage",
		Namespace: "env",
		Callback: func(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

			// Read function arguments
			offset := stack[0]
			bufferInput, err := plugin.ReadBytes(offset)

			if err != nil {
				fmt.Println("🥵", err.Error())
				panic(err)
			}

			//monsterData := string(bufferInput)
			monsterData := bufferInput

			monster := Monster{}
			json.Unmarshal([]byte(monsterData), &monster)

			fmt.Println("🟢 Monster:", monster)
			fmt.Println("🟩 Action:", monster.Action)

			// Return data
			plugin.Free(offset)

			switch monster.Action {
			case "toctoc":
				fmt.Println("👋 Action:", monster.Action)
				offset, err = plugin.WriteBytes([]byte("TOCTOC"))
			case "yo":
				offset, err = plugin.WriteBytes([]byte("YO"))
			default:
				fmt.Println("🎃 Action: default")
				offset, err = plugin.WriteBytes([]byte(""))
			}

			if err != nil {
				fmt.Println("😡", err.Error())
				panic(err)
			}

			stack[0] = offset
		},
		Params:  []api.ValueType{api.ValueTypeI64},
		Results: []api.ValueType{api.ValueTypeI64},
	}


	hostFunctions := []extism.HostFunction{
		display,
		sendMessage,
	}

	pluginMonster, err := extism.NewPlugin(ctx, manifest, config, hostFunctions) // new

	/*
	getName := func() string {
		_, monsterName, err := pluginMonster.Call("getName", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return string(monsterName)
	}
	*/

	_, monsterName, err := pluginMonster.Call("getName", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*
	getAvatar := func() string {
		_, monsterAvatar, err := pluginMonster.Call("getAvatar", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return string(monsterAvatar)
	}
	*/

	_, monsterAvatar, err := pluginMonster.Call("getAvatar", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//_, out, err := pluginMonster.Call("hello", []byte("monster"))

	fmt.Println("🔥 Monster ->", string(monsterAvatar)+" "+string(monsterName))
	//fmt.Println("🔥 Monster ->", getAvatar()+" "+getName())


	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

	sayHey := func() {
		pluginMonster.Call("hey", nil)
	}
	sayHey()

	

	p := tea.NewProgram(initialModel(game))
	if _, err := p.Run(); err != nil {
		fmt.Printf("😡, there's been an error: %v", err)
		os.Exit(1)
	}
}
