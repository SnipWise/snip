package tools

import (
	"strings"
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"math/rand"
)

type DiceRollInput struct {
	NumDice  int `json:"num_dice"`
	NumFaces int `json:"num_faces"`
}

type DiceRollResult struct {
	Rolls []int `json:"rolls"`
	Total int   `json:"total"`
}

type CharacterNameInput struct {
	Race string `json:"race"`
}

type CharacterNameResult struct {
	Name string `json:"name"`
	Race string `json:"race"`
}

// [NOTE]: this is a work-in-progress
func Catalog(g *genkit.Genkit) []ai.ToolRef {
	// Define tools
	diceRollTool := genkit.DefineTool(g, "roll_dice", "Roll n dice with n faces each",
		func(ctx *ai.ToolContext, input DiceRollInput) (DiceRollResult, error) {
			return rollDice(input.NumDice, input.NumFaces), nil
		},
	)

	characterNameTool := genkit.DefineTool(g, "generate_character_name", "Generate a D&D character name for a specific race",
		func(ctx *ai.ToolContext, input CharacterNameInput) (CharacterNameResult, error) {
			return generateCharacterName(input.Race), nil
		},
	)

	return []ai.ToolRef{diceRollTool, characterNameTool}
}

func rollDice(numDice, numFaces int) DiceRollResult {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rolls := make([]int, numDice)
	total := 0

	for i := 0; i < numDice; i++ {
		roll := r.Intn(numFaces) + 1
		rolls[i] = roll
		total += roll
	}

	return DiceRollResult{
		Rolls: rolls,
		Total: total,
	}
}

func generateCharacterName(race string) CharacterNameResult {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	namesByRace := map[string][]string{
		"elf":      {"Aerdrie", "Ahvonna", "Aramil", "Aranea", "Berrian", "Caelynn", "Carric", "Dayereth", "Enna", "Galinndan"},
		"dwarf":    {"Adrik", "Baern", "Darrak", "Eberk", "Fargrim", "Gardain", "Harbek", "Kildrak", "Morgran", "Thorek"},
		"human":    {"Aerdrie", "Aramil", "Berris", "Cithreth", "Dayereth", "Enna", "Galinndan", "Hadarai", "Immeral", "Lamlis"},
		"halfling": {"Alton", "Ander", "Bernie", "Bobbin", "Cade", "Callus", "Corrin", "Dannad", "Garret", "Lindal"},
		"orc":      {"Gash", "Gell", "Henk", "Holg", "Imsh", "Keth", "Krusk", "Mhurren", "Ront", "Shump"},
		"tiefling": {"Akmenos", "Amnon", "Barakas", "Damakos", "Ekemon", "Iados", "Kairon", "Leucis", "Melech", "Mordai"},
	}

	raceLower := strings.ToLower(race)
	names, exists := namesByRace[raceLower]
	if !exists {
		names = namesByRace["human"] // Default to human names
	}

	selectedName := names[r.Intn(len(names))]

	return CharacterNameResult{
		Name: selectedName,
		Race: race,
	}
}
