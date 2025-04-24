/*
 * N-Queens Problem
 *
 * For an NxN chess board, 'safely' place a chess queen in every column and row such that none can attack another.
 * This solution is based Wirth Pascal solution, although a tad cleaner, thus easier to understand as it uses Go/C
 * style indexing and naming, and also prints the Queen using a Unicode 'rune' (which other languages do not handle natively).
 *
 * N rows by N columns are number left to right top to bottom 0 - 7
 *
 * There are 2N-1 diagonals (showing an 8x8)
 *  the upper-right to lower-left are numbered row + col that is:
 *    0   1   2   3   4   5   6   7
 *    1   2   3   4   5   6   7   8
 *    2   3   4   5   6   7   8   9
 *    3   4   5   6   7   8   9  10
 *    4   5   6   7   8   9  10  11
 *    5   6   7   8   9  10  11  12
 *    6   7   8   9  10  11  12  13
 *    7   8   9  10  11  12  13  14
 *
 *	the upper-left to lower-right are numbered N-1 + row - col
 *    7   6   5   4   3   2   1   0
 *    8   7   6   5   4   3   2   1
 *    9   8   7   6   5   4   3   2
 *   10   9   8   7   6   5   4   3
 *   11  10   9   8   7   6   5   4
 *   12  11  10   9   8   7   6   5
 *   13  12  11  10   9   8   7   6
 *   14  13  12  11  10   9   8   7
 */
