package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"log"
	"math/rand"
)

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "vicmanbrile", // TODO: Your Battlesnake username
		Color:      "#991f00",     // TODO: Personalize
		Head:       "default",     // TODO: Personalize
		Tail:       "default",     // TODO: Personalize
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// Step 0: Don't let your Battlesnake move back in on it's own neck
	Head := state.You.Body[0] // Coordinates of your head
	Neck := state.You.Body[1]

	if Head.X > Neck.X {
		possibleMoves["right"] = true
		possibleMoves["left"] = false
	} else {
		possibleMoves["right"] = false
		possibleMoves["left"] = true
	}

	if Head.Y > Neck.Y {
		possibleMoves["up"] = true
		possibleMoves["down"] = false
	} else {
		possibleMoves["up"] = false
		possibleMoves["down"] = true
	}

	// TODO: Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	if possibleMoves["right"] == true && Head.X == boardWidth {
		possibleMoves["right"] = false
	} else if possibleMoves["left"] == true && Head.X == 1 {
		possibleMoves["left"] = false
	}

	if possibleMoves["up"] == true && Head.Y == boardHeight {
		possibleMoves["up"] = false
	} else if possibleMoves["down"] == true && Head.Y == 1 {
		possibleMoves["down"] = false
	}

	// TODO: Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
	// mybody := state.You.Body

	/*
		for _, snake := range state.Board.Snakes {
			for _, part := range snake.Body {
				if positionX(part, myHead) {
					possibleMoves["right"] = false
				} else {
					possibleMoves["left"] = false
				}

				if positionY(part, myHead) {
					possibleMoves["up"] = false
				} else {
					possibleMoves["down"] = false
				}
			}
		}

	*/

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
