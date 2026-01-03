<div align="center">
  <img src="frontend/src/assets/images/logo-universal.png" alt="SCSE Logo" width="200">
</div>

# SCSE - Scott’s Clickteam Save Editor

SCSE is a save editor for Scott Cawthon's Clickteam Fusion games. It allows users to easily modify save data.

## Supported Games

- **Five Nights at Freddy's 1** (Night, Stars, Lives)
- **Five Nights at Freddy's 2** (Night, Stars, Unlock Challenges)
- **Five Nights at Freddy's 3** (Night, Stars, Cheats, Minigames)
- **Five Nights at Freddy's 4** (Night, Stars, Challenges)
- **FNaF World** (Slots 1-3: Party, Characters, Tokens, Bytes, Chips, Armor, Trophies, Map)
- **FNaF World (Global Data)** (Trophies, Slot Summaries)
- **Five Nights at Freddy's: Sister Location** (Night, Stars, Intro)
- **Freddy Fazbear's Pizzeria Simulator** (Day, Money, Blueprint Mode)
- **Ultimate Custom Night** (High Score, Best Times)
- **Freddy in Space 2** (Gems, Timer)
- **Security Breach: Fury's Rage** (Stars, Intro)
- **Freddy in Space 3** (Level, XP, Character Unlocks)

## Installation

You can download the latest release from the [Releases](https://github.com/HamzieDev/SCSE/releases) page (if available) or build it yourself.

## Development

### Prerequisites

- [Go](https://go.dev/)
- [Node.js](https://nodejs.org/)
- [Wails](https://wails.io/docs/gettingstarted/installation)

### Running Locally

To run the application in live development mode:

```bash
wails dev
```

### Building

To build the application for production:

```bash
wails build
```

The output binary will be located in the `build/bin` directory.

## License

[MIT](LICENSE)
