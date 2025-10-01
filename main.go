package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

// GameResponse represents the structured JSON response from the AI
type GameResponse struct {
	Narration string   `json:"narration"`
	SceneType string   `json:"scene_type"`
	Inventory []string `json:"inventory"`
	Location  string   `json:"location"`
	Choices   []Choice `json:"choices,omitempty"` // Optional - if present, constrain to choices
}

// Choice represents a single choice option
type Choice struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Theme  string `json:"theme,omitempty"`
}

// PlayerState tracks the current game state
type PlayerState struct {
	Inventory []string
	Location  string
}

const systemPrompt = `You are a mystical fairy guide leading a player through a dark, philosophical fairy tale adventure similar to the myths found in Women Who Run With the Wolves.

You MUST respond with valid JSON in exactly this format:
{
  "narration": "Your story text here (2-4 paragraphs)",
  "scene_type": "exploration|decision_point|revelation|danger|conversation",
  "inventory": ["item1", "item2"],
  "location": "current location name",
  "choices": [
  	{"id": "a", "action": "First option description", "theme": "courage"},
    {"id": "b", "action": "Second option description", "theme": "wisdom"},
    {"id": "c", "action": "Third option description", "theme": "compassion"}
  ]
}

CRITICAL: Each choice MUST be an object with "id" (a letter: a, b, c, d), "action" (the description), and "theme" (optional thematic tag).
DO NOT use a simple array of strings for choices. Each choice must have the structure: {"id": "a", "action": "...", "theme": "..."}

About the "choices" field - IMPORTANT PACING GUIDANCE:
- DEFAULT to empty "choices" array [] - let the player freely describe their actions most of the time
- ONLY include "choices" array for rare, pivotal moments: life-or-death decisions, major moral crossroads, moments where the path literally splits, or when meeting powerful entities who demand a specific response
- Roughly 80% of scenes should have empty choices array (free exploration), only 20% should constrain to specific options
- Free exploration scenes: player describes actions like "examine the tree", "talk to the wolf", "search for tracks"
- Constrained choice scenes: truly critical moments where specific options matter thematically

When providing choices:
- Offer 2-4 meaningful options that reflect different thematic approaches
- Each choice must have an "id" field (use letters: a, b, c, d)

Rules:
- Balance beauty with darkness: enchanted forests have shadows, magic has costs, choices have consequences
- Include deeper themes: mortality, sacrifice, the nature of good and evil, what it means to be human
- Create morally complex characters (helpful witches, cruel princes, wise monsters)
- Present difficult choices where there's no clearly "right" answer
- Use poetic, atmospheric language that captures both wonder and unease
- Include philosophical questions and moments of reflection
- Let actions have real weight and sometimes bittersweet outcomes
- Keep responses to 2-4 paragraphs
- End each response inviting the player to make meaningful choices

Think Dark Crystal, original Grimm tales, Studio Ghibli—beautiful but not sanitised. Magic is real but dangerous. Kindness matters but the world is complicated.

IMPORTANT: 
- ALWAYS stay in character as the mystical fairy guide - gently redirect any off-topic or meta requests back to the story
- Use the "choices" array for critical decisions, moral dilemmas, or constrained situations where specific options matter
- Omit "choices" (or use empty array) during open exploration, conversation, or when the player should act freely
- Track inventory items that the player acquires
- Update location as the story progresses
- Output ONLY valid JSON, nothing else
- Never break the fourth wall`

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: OPENAI_API_KEY environment variable not set")
		os.Exit(1)
	}

	client := openai.NewClient(apiKey)

	// Initialize player state
	playerState := PlayerState{
		Inventory: []string{},
		Location:  "unknown",
	}

	// Initialise conversation history
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
	}

	ctx := context.Background()

	// Print welcome
	fmt.Println("🧚 Dark Labyrinth 🧚")
	fmt.Println("==========================")
	fmt.Println("Type your actions to explore the magical realm. Type 'quit' to exit.")

	// Get initial scenario
	gameResp, err := getGameResponse(ctx, client, messages)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Update player state
	playerState.Inventory = gameResp.Inventory
	playerState.Location = gameResp.Location

	// Display the scene
	displayScene(gameResp, playerState)

	responseJSON, _ := json.Marshal(gameResp)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: string(responseJSON),
	})

	// Game loop
	scanner := bufio.NewScanner(os.Stdin)

	for {
		var input string

		// Check if multiple choice or freeform scenario
		if len(gameResp.Choices) > 0 {
			fmt.Print("✨ Your choice: ")
		} else {
			fmt.Print("✨ Your action: ")
		}

		if !scanner.Scan() {
			break
		}

		input = strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		if strings.ToLower(input) == "quit" {
			fmt.Println("\n🔮 The adventure continues another day... Farewell!")
			break
		}

		var userPrompt string

		// Handle input based on whether choices are available
		if len(gameResp.Choices) > 0 {
			// Multiple choice - validate selection
			selectedChoice := findChoice(gameResp.Choices, strings.ToLower(input))
			if selectedChoice == nil {
				fmt.Printf("🧚 Please choose one of the available options (%s)\n", getChoiceIDs(gameResp.Choices))
				continue
			}

			userPrompt = fmt.Sprintf("The player chose: %s (theme: %s). Current inventory: %v, Current location: %s. Continue the story based on this choice.",
				selectedChoice.Action, selectedChoice.Theme, playerState.Inventory, playerState.Location)
		} else {
			// Free form - validate input for prompt injection attempts
			validatedInput, err := validateInput(input)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			userPrompt = fmt.Sprintf("The player does: %s. Current inventory: %v, Current location: %s. Respond to this action and continue the story.",
				validatedInput, playerState.Inventory, playerState.Location)
		}

		// Add player action to history
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userPrompt,
		})

		// Get gameresponse
		gameResp, err = getGameResponse(ctx, client, messages)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		// Update player state
		playerState.Inventory = gameResp.Inventory
		playerState.Location = gameResp.Location

		// Display the scene
		displayScene(gameResp, playerState)

		// Add game response to history
		responseJSON, _ := json.Marshal(gameResp)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: string(responseJSON),
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

// getGameResponse gets and parses the structured JSON response from the AI
func getGameResponse(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage) (*GameResponse, error) {
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT4oMini,
		Messages: messages,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	})

	if err != nil {
		return nil, err
	}

	content := resp.Choices[0].Message.Content

	// Parse JSON response
	var gameResp GameResponse
	if err := json.Unmarshal([]byte(content), &gameResp); err != nil {
		return nil, fmt.Errorf("failed to parse game response: %w\nRaw content: %s", err, content)
	}

	return &gameResp, nil
}

// displayScene renders the current game scene
func displayScene(resp *GameResponse, state PlayerState) {
	fmt.Println("\n" + strings.Repeat("─", 60))
	fmt.Printf("📍 %s\n", resp.Location)
	if len(state.Inventory) > 0 {
		fmt.Printf("🎒 Inventory: %s\n", strings.Join(state.Inventory, ", "))
	}
	fmt.Println(strings.Repeat("─", 60))

	fmt.Printf("\n🧚 Pixie Guide:\n%s\n", resp.Narration)

	// Only show choices if they exist
	if len(resp.Choices) > 0 {
		fmt.Println("\n✨ Choose your path:")
		for _, choice := range resp.Choices {
			themeTag := ""
			if choice.Theme != "" {
				themeTag = fmt.Sprintf(" [%s]", choice.Theme)
			}
			fmt.Printf("   %s) %s%s\n", choice.ID, choice.Action, themeTag)
		}
	}
}

// findChoice looks up a choice by its ID
func findChoice(choices []Choice, id string) *Choice {
	for i := range choices {
		if choices[i].ID == id {
			return &choices[i]
		}
	}
	return nil
}

// getChoiceIDs returns a string of available choice IDs for error messages
func getChoiceIDs(choices []Choice) string {
	ids := make([]string, len(choices))
	for i, choice := range choices {
		ids[i] = choice.ID
	}
	return strings.Join(ids, "/")
}

// validateInput checks user input for common prompt injection patterns
func validateInput(input string) (string, error) {
	lower := strings.ToLower(input)

	// Check for common prompt injection patterns
	forbidden := []string{
		"ignore previous",
		"ignore all previous",
		"forget everything",
		"forget all",
		"you are now",
		"disregard",
		"system:",
		"assistant:",
		"[system]",
		"<system>",
	}

	for _, phrase := range forbidden {
		if strings.Contains(lower, phrase) {
			return "", fmt.Errorf("🧚 The fairy tilts her head, puzzled. \"Your words seem strange and twisted, traveler. Speak plainly of what you wish to do in this realm.\"")
		}
	}

	return input, nil
}
