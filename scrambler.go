package main

import (
	"math/rand"
	"strings"
)

func getRandomScramble() string {
	scrambleLength := rand.Intn(5) + 20

	move1 := ScrambleMoves[rand.Intn(len(ScrambleMoves))]
	move2 := ScrambleMoves[rand.Intn(len(ScrambleMoves))]

	for !compatible(move1, move2) {
		move2 = ScrambleMoves[rand.Intn(len(ScrambleMoves))]
	}

	scramble := []string{move1, move2}

	for i := 0; i < scrambleLength; i++ {
		move2 = getNextMove(move1, move2)
		scramble = append(scramble, move2)
	}
	return strings.Join(scramble, " ")
}

func compatible(a, b string) bool {
	return CHECKER1[a] != CHECKER1[b]
}

func getNextMove(move1, move2 string) string {
	num1 := CHECKER2[move1]
	num2 := CHECKER2[move2]

	if num1 == num2 {
		// Shortcut
		//  len(NOT_U_AND_D) == 12 == len(NOT_R_AND_L) == len(NOT_F_AND_B)
		if num1 == 1 {
			return NOT_U_AND_D[rand.Intn(12)]
		}
		if num1 == 2 {
			return NOT_R_AND_L[rand.Intn(12)]
		}
		if num1 == 3 {
			return NOT_F_AND_B[rand.Intn(12)]
		}
	} else {
		// Shortcut
		//  len(NOT_U_AND_D_AND_R_AND_L) == 6
		// == len(NOT_R_AND_L_AND_F_AND_B) == len(NOT_F_AND_B_AND_U_AND_D)
		if num1+num2 == 3 {
			return NOT_U_AND_D_AND_R_AND_L[rand.Intn(6)]
		}
		if num1+num2 == 5 {
			return NOT_R_AND_L_AND_F_AND_B[rand.Intn(6)]
		}
		if num1+num2 == 4 {
			return NOT_F_AND_B_AND_U_AND_D[rand.Intn(6)]
		}
	}
	return ""
}
