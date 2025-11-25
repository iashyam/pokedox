# Pokedox

A command-line Pokédex application built in Go that allows you to explore Pokémon locations and catch Pokémon. This project demonstrates core Go concepts including HTTP requests, JSON parsing, caching, concurrency, and terminal input handling.

## Introduction

### What is Pokedox?

Pokedox is an interactive CLI (Command Line Interface) Pokédex application that lets you:
- Explore Pokémon location areas
- Navigate through paginated location data
- Catch and manage Pokémon
- Search for Pokémon information

The application demonstrates best practices in Go development, including:
- **Api Calls**: Fetches data from the PokéAPI
- **Caching System**: Implements a custom cache with automatic expiration (reaping) to minimize API calls
- **Concurrency**: Uses goroutines and mutexes for thread-safe operations
- **Terminal Input**: Supports both text commands and arrow key navigation
- **Clean Architecture**: Organized package structure with separation of concerns

**Key Features**:
- Custom cache with read-write mutexes for concurrent access
- Background goroutine for automatic cache expiration
- Error handling and graceful CLI navigation

## Basic Commands

### Available Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `help` | Display all available commands | `help` |
| `exit` | Exit the Pokedox application | `exit` |
| `map` | Display the next set of location areas | `map` |
| `mapb` | Display the previous set of location areas | `mapb`|
| `catch <pokemon>` | Attempt to catch a Pokémon | `catch pikachu` |
| `inspect <pokemon>` | View details of a caught Pokémon | `inspect pikachu` |
| `pokedex` | Display all caught Pokémon | `pokedex` |

### How Commands Work

#### Navigation Commands (`map` / `mapb`)

1. **`map`** - Fetches the next page of location areas from PokéAPI
   - First call loads the default page from `https://pokeapi.co/api/v2/location-area`
   - Results are cached to avoid duplicate API calls

2. **`mapb`** - Fetches the previous page of location areas
   - Navigates backward through paginated results

#### Catching and Inspecting Pokémon

3. **`catch <pokemon>`** - Attempt to capture a Pokémon
   - Randomly determines success (simulates catching mechanics)
   - Stores caught Pokémon in `Config.Pokedox` map
   - Prevents duplicate catches

4. **`inspect <pokemon>`** - View details of a caught Pokémon
   - Displays information only if the Pokémon has been caught
   - Shows stats, abilities, and moves

5. **`pokedex`** - Lists all Pokémon you've caught

### Caching System

The application uses a custom `PokeCache` implementation that:
- Stores API responses with timestamps
- Automatically expires cached entries after 10 minutes (configurable)
- Uses `sync.RWMutex` for thread-safe concurrent access
- Runs background cleanup with a goroutine to free up memory

#### Why Caching Matters

**Performance**: Caching significantly improves application performance by:
- **Reducing API Calls**: Eliminates redundant requests to PokéAPI
- **Faster Response Times**: Returns cached data instantly instead of waiting for network requests
- **Bandwidth Savings**: Minimizes data transfer and API rate limit consumption
- **Offline Capability**: Users can navigate previously visited locations without internet connectivity

**How It Works**:
1. When you request location data with `map` or `mapb`, the app checks the cache first
2. If the data exists and hasn't expired, it returns immediately from cache (cache hit)
3. If the data is missing or expired, it fetches from PokéAPI and stores it in cache for future use
4. Background cleanup automatically removes expired entries to prevent memory bloat

This implementation demonstrates real-world caching patterns used in production applications.

### Future Work

- Adding new commands (e.g., search, filter, stats)
- Improving the caching system
- Enhancing error handling and user feedback
- Optimizing API calls and performance
- Adding unit tests
- Improving documentation

## Acknowledgments

Special thanks to:
- **[Boot.dev](https://boot.dev)** - For providing excellent Go learning resources and project ideas.
- **[Lane Wagner](https://twitter.com/wagslane)** - For mentoring, guidance, and creating the educational platform that inspired this project.

This project was built as part of the Boot.dev Go curriculum and serves as a practical demonstration of Go programming concepts.# pokedox
