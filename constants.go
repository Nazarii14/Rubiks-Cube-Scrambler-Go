package main

const U = "u"
const U2 = "u2"
const U_PRIME = "u_prime"

const D = "d"
const D2 = "d2"
const D_PRIME = "d_prime"

const R = "r"
const R2 = "r2"
const R_PRIME = "r_prime"

const L = "l"
const L2 = "l2"
const L_PRIME = "l_prime"

const F = "f"
const F2 = "f2"
const F_PRIME = "f_prime"

const B = "b"
const B2 = "b2"
const B_PRIME = "b_prime"

const M = "m"
const M2 = "m2"
const M_PRIME = "m_prime"

const E = "e"
const E2 = "e2"
const E_PRIME = "e_prime"

const S = "s"
const S2 = "s2"
const S_PRIME = "s_prime"

const X = "x"
const X2 = "x2"
const X_PRIME = "x_prime"

const Y = "y"
const Y2 = "y2"
const Y_PRIME = "y_prime"

const Z = "z"
const Z2 = "z2"
const Z_PRIME = "z_prime"

// MOVES NEEDED TO SCRAMBLE CUBE
var ScrambleMoves = []string{
	U, U2, U_PRIME, D, D2, D_PRIME,
	R, R2, R_PRIME, L, L2, L_PRIME,
	F, F2, F_PRIME, B, B2, B_PRIME,
}

var CHECKER1 = map[string]int{
	U: 1, U_PRIME: 1, U2: 1,
	D: 2, D_PRIME: 2, D2: 2,
	R: 3, R_PRIME: 3, R2: 3,
	L: 4, L_PRIME: 4, L2: 4,
	F: 5, F_PRIME: 5, F2: 5,
	B: 6, B_PRIME: 6, B2: 6,
}

var CHECKER2 = map[string]int{
	U: 1, U_PRIME: 1, U2: 1,
	D: 1, D_PRIME: 1, D2: 1,
	R: 2, R_PRIME: 2, R2: 2,
	L: 2, L_PRIME: 2, L2: 2,
	F: 3, F_PRIME: 3, F2: 3,
	B: 3, B_PRIME: 3, B2: 3,
}

var NOT_U_AND_D = []string{R, R_PRIME, R2, L, L_PRIME, L2, F, F_PRIME, F2, B, B_PRIME, B2}
var NOT_R_AND_L = []string{U, U_PRIME, U2, D, D_PRIME, D2, F, F_PRIME, F2, B, B_PRIME, B2}
var NOT_F_AND_B = []string{U, U_PRIME, U2, D, D_PRIME, D2, R, R_PRIME, R2, L, L_PRIME, L2}

var NOT_U_AND_D_AND_R_AND_L = []string{F, F_PRIME, F2, B, B_PRIME, B2}
var NOT_R_AND_L_AND_F_AND_B = []string{U, U_PRIME, U2, D, D_PRIME, D2}
var NOT_F_AND_B_AND_U_AND_D = []string{R, R_PRIME, R2, L, L_PRIME, L2}
