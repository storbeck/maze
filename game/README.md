# Maze Game

This is a simple maze game implemented using Node.js and HTML5 Canvas. The game supports two players: Player 1 uses the arrow keys, and Player 2 uses the WASD keys. The objective is to navigate through the maze and reach the exit. When a player reaches the exit, a "YOU WIN" message is displayed along with a confetti effect indicating the winner.



https://github.com/storbeck/maze/assets/449874/fd502851-761f-4b7e-9aef-69408bde36e9



## Features
- Generate a new maze of specified size.
- Two-player support with different controls and colors.
- Confetti effect and winning message when a player reaches the exit.

## Requirements
- Node.js and npm installed.
- A maze generation script (`mazegen`) located in the `bin` directory in the parent directory.

## Installation

1. Install dependencies.
   ```bash
   cd game
   npm install
   ```

3. Ensure the `mazegen` script is executable and located in the `bin` directory in the parent directory.

## Usage

1. Start the Node.js server.
   ```bash
   cd game
   node server.js
   ```

2. Open your web browser and go to `http://localhost:3000?size=WIDTHxHEIGHT`, replacing `WIDTH` and `HEIGHT` with the desired dimensions of the maze (e.g., `http://localhost:3000?size=20x20`).

## Controls

- **Player 1 (Blue)**
  - Arrow Up: Move up
  - Arrow Down: Move down
  - Arrow Left: Move left
  - Arrow Right: Move right

- **Player 2 (Orange)**
  - W: Move up
  - S: Move down
  - A: Move left
  - D: Move right

## How It Works

- The server generates a maze using the `mazegen` script and serves it to the client.
- The client renders the maze on an HTML5 canvas and allows two players to navigate it.
- When a player reaches the exit, a "YOU WIN" message is displayed along with a confetti effect.

## Example

Open your browser and navigate to:

```
http://localhost:3000?size=20x20
```

You should see a maze where two players can navigate using their respective controls.

## License

This project is licensed under the MIT License.
