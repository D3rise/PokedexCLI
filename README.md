# Pokedex CLI

```
 _____       _            _            _____ _      _____
|  __ \     | |          | |          / ____| |    |_   _|
| |__) |___ | | _____  __| | _____  _| |    | |      | |
|  ___/ _ \| |/ / _ \/ _` |/ _ \ \/ / |    | |      | |
| |  | (_) |   <  __/ (_| |  __/>  <| |____| |____ _| |_
|_|   \___/|_|\_\___|\__,_|\___/_/\_\\_____|______|_____|
```

A command-line interface Pokedex application that allows you to explore the world of Pokemon, catch them, and view their details.

## Features

- Browse Pokemon locations
- Explore areas to find Pokemon
- Catch Pokemon with a chance-based system
- View your collection of caught Pokemon
- Inspect detailed information about caught Pokemon
- Caching system to reduce API calls

## Architecture

The application is built with Go and follows a modular architecture:

- **Main Package**: Entry point for the application, initializes the command registry and starts the REPL.
- **REPL**: Provides a Read-Eval-Print Loop for interacting with the application through the command line.
- **Commands**: Each command is implemented as a separate module with a consistent interface.
- **PokeAPI Client**: Handles communication with the [PokeAPI](https://pokeapi.co/) to retrieve Pokemon data.
- **Cache**: Implements a simple in-memory cache with expiration to reduce API calls.
- **Context**: Provides a shared context for storing and retrieving data across commands.
- **Pokedex**: Manages the user's collection of caught Pokemon.

## Installation

### Prerequisites

- Go 1.16 or higher

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/D3rise/pokedexcli.git
   ```

2. Navigate to the project directory:
   ```bash
   cd pokedexcli
   ```

3. Build the application:
   ```bash
   go build
   ```

4. Run the application:
   ```bash
   ./pokedexcli
   ```

## Usage

Once you start the application, you'll be presented with a prompt:

```
Welcome to the Pokedex!
Pokedex >
```

You can enter commands at this prompt to interact with the application.

### Example Workflow

1. View available locations:
   ```
   Pokedex > map
   ```

2. Explore a location to find Pokemon:
   ```
   Pokedex > explore canalave-city-area
   ```

3. Catch a Pokemon:
   ```
   Pokedex > catch pikachu
   ```

4. View your Pokedex:
   ```
   Pokedex > pokedex
   ```

5. Inspect a caught Pokemon:
   ```
   Pokedex > inspect pikachu
   ```

## Available Commands

- **help**: Displays a help message with all available commands
- **exit**: Exits the application
- **map**: Displays a list of location areas (20 at a time)
- **mapb**: Displays the previous list of location areas
- **explore <areaName>**: Explores a location area for Pokemon
- **catch <pokemonName>**: Attempts to catch a Pokemon
- **inspect <pokemonName>**: Displays detailed information about a caught Pokemon
- **pokedex**: Displays a list of all caught Pokemon

## API Usage

This application uses the [PokeAPI](https://pokeapi.co/) to retrieve Pokemon data. The API is free to use, but please be respectful of the rate limits. The application implements caching to reduce the number of API calls.

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgements

- [PokeAPI](https://pokeapi.co/) for providing the Pokemon data
- The Pokemon franchise for the inspiration