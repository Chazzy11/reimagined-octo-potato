# 🧚 Dark Labyrinth

A dark, philosophical text-based adventure game powered by OpenAI's API. Explore enchanted realms where beauty meets shadow, and every choice carries weight.

## Description

This CLI game creates an interactive fairy tale experience inspired by stories of Studio Ghibli or those found in *Women Who Run With the Wolves*—beautiful but not sanitised, where magic is real but dangerous, and moral choices are rarely simple. Powered by GPT-4o-mini, each playthrough generates a unique story that responds to your actions.

## Features

- 🌙 Dark, atmospheric storytelling with philosophical depth
- 🎭 Morally complex characters and meaningful choices
- 🔄 Hybrid gameplay: Free-form exploration mixed with multiple-choice decisions at pivotal moments
- 📦 State tracking: Automatic inventory and location management
- ✨ Poetic, evocative language
- 🔮 Unique story every time you play

## Requirements

- Go 1.21 or higher
- OpenAI API key

## Installation

1. Clone this repository:
```bash
git clone https://github.com/Chazzy11/reimagined-octo-potato.git
cd reimagined-octo-potato
```

2. Install dependencies

```bash
go mod download
```

3. Get an OpenAI API key from https://platform.openai.com/api-keys

4. Set your API key as an environment variable:
```bash
export OPENAI_API_KEY="your-api-key-here"
```

5. Build the game:

```bash
make build
```

## Usage

Run the game:

```bash
make run
```

The game intelligently switches between two modes:
**Free-form Exploration** (most of the time):

```bash
📍 The Whispering Woods
────────────────────────────────
🧚 Pixie Guide: You find yourself in a misty glade...

✨ Your action: examine the ancient tree
```
Type your actions naturally, and the story will respond.

**Multiple Choice Decisions** (at critical moments):

```bash
📍 The Crossroads
🎒 Inventory: silver thread, broken compass
────────────────────────────────
🧚 Pixie Guide: Three paths stretch before you...

✨ Choose your path:
   a) The moonlit path [courage]
   b) The shadowed path [wisdom]
   c) Wait until dawn [patience]

💫 Your choice: a
```

 Type `quit` to exit.

## Technical Details

The game uses a structured JSON response system where the AI returns:

```bash
{
  "narration": "Story text...",
  "scene_type": "exploration|decision_point|revelation|danger|conversation",
  "inventory": ["item1", "item2"],
  "location": "Current Location",
  "choices": [] // Optional: present only at critical moments
}
```
This enables:

- State persistence across the story
- Dynamic switching between exploration and constrained choices
- Inventory and location tracking

## Cost

This game uses OpenAI's GPT-4o-mini model. Typical costs:

- Per game session (10-20 exchanges): ~$0.01-0.03
- An hour of gameplay: ~$0.10-0.20

## Development

Format code:

```bash
make fmt
```

Run tests:

```bash
make test
```

## License

MIT License - see LICENSE file for details

## Contributing

This is a personal project, but feel free to fork and create your own variations!

## Acknowledgements
- Built with the go-openai library
- Powered by OpenAI's GPT-4o-mini model