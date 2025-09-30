# 🧚 Dark Labyrinth

A dark, philosophical text-based adventure game powered by OpenAI's API. Explore enchanted realms where beauty meets shadow, and every choice carries weight.

## Description

This CLI game creates an interactive fairy tale experience inspired by stories of Studio Ghibli or those found in *Women Who Run With the Wolves*—beautiful but not sanitised, where magic is real but dangerous, and moral choices are rarely simple. Powered by GPT-4o-mini, each playthrough generates a unique story that responds to your actions.

## Features

- 🌙 Dark, atmospheric storytelling with philosophical depth
- 🎭 Morally complex characters and meaningful choices
- 📖 Dynamic narrative that remembers your actions
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

Example gameplay:

```bash
✨ You: ask the old woman why she helps me
🧚 Pixie Guide: She doesn't answer right away. Her fingers continue 
spinning thread from nothing, from air and shadow. "Help?" she finally 
says. "I give you choices, child. Whether they help or harm you depends 
entirely on what you choose to see." Her eyes reflect something ancient. 
"Now—will you take the thread, or leave empty-handed?"

✨ You: take the thread
🧚 Pixie Guide: It burns cold in your palm, alive with possibility...
```

Type your actions naturally, and the story will respond. Type `quit` to exit.

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