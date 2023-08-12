## --- Day 21: Dirac Dice ---

Part 1 is pretty easy, so I won't go into any details, and go straight to Part 2.

Each turn results in 3 rolls and 27 forks. Below is how sum of three rolls distributed among the forks. E.g the sum of 5 happens in 6 different forks: `1,2,2`,`2,1,2`,`2,2,1`,`1,1,3`,`1,3,1`,`3,1,1`.

| Score | Forks |
| ----- | ----- |
| 3     | 1     |
| 4     | 3     |
| 5     | 6     |
| 6     | 7     |
| 7     | 6     |
| 8     | 3     |
| 9     | 1     |
| Total | 27    |

Assume that you are in a certain point in the game where your can arrive in `X` forks and it's your time to roll. You roll, and in the resulting 27 forks you win in `Y` forks, and you do not win in `Z=27-Y` forks. Now in total *on the current turn* you win in `X*Y` forks and you do not win in `X*Z` forks.

The next major insight, is that the order of the rolls between the players does not impact the number of the forks, because for the operation of multiplication (as seen in the previous paragraph), the order of operands do not matter. So if the first player did not win on a turn in `A` forks out of 27, and the second player in `B` forks out of 27 (remember that these are different 27 forks for the players as they roll separately and each roll forks), and on the next turn, the first player did not win in `C` forks, and the second player did not win in `D` forks, then the second player did not win on that two turns in `A*B*C*D` forks in total. This means, that the players can roll separately, multiply the fork numbers independently, and then in the end, if we multiply the number of forks from first and second player together we will still get the right number of forks.

Note how we say "does not win" instead of "loses". These are two different things. You lose on a turn if you opponent wins this turn. But if it is unknown if the opponent won, it also unknow if you lost, so we just say that you did not win in this case.

Lastly, we can represent the universe splitting as a tree graph for each player separately. The first level of the tree will be before turn one. The second level will have one node per each possible rolled turn score (3-9). The node can keep the current number of forks, current score, position on the board and the turn number. When score of a node is 21 or more, it becomes a leaf node and does not produce the next level. Now all we need to do is to walk this tree (e.g. depth first) and for each turn collect the number of forks that win and that do not win from each node.

Here is an example for the player that starts at position 3:

| Turn | Win      | No Win  |
| ---- | -------- | ------- |
| 1    |          | 27      |
| 2    |          | 729     |
| 3    | 5401     | 14282   |
| 4    | 267290   | 118324  |
| 5    | 2573275  | 621473  |
| 6    | 15358599 | 1421172 |
| 7    | 37322440 | 1049204 |
| 8    | 28186346 | 142162  |
| 9    | 3837558  | 816     |
| 10   | 22032    |         |

On turn 1 it is not possible to get over 21 yet, and neither on turn 2. On turn 3, 5401 forks leads to a win, and those forks will not fork themselves anymore since they are leaves now. 14282 forks which did not win will fork 27 times each, and so on. Eventually the score will become that big that any three rolls will get it to 21 or over (on step 9 in the case above), all the remaining forks will be in the win column.

Once we calculated those numbers for both players, we are ready to get the final counts:

```go
	var firstWinsCount, secondWinsCount int64
	// a bit of cheating here, we know that 10 is always the max number of turns from debugging
	// otherwise we should have really found the max keys in the maps first and used that
	for i := 3; i <= 10; i++ {
		// first player wins on this turn in the number of his forks that he wins in
		// multiplied by the number of forks the second player did not win on previous turn
		firstWinsCount += winFirst[i] * noWinSecond[i-1]
		// second player wins on this turn in the number of his forks that he wins in
		// multiplied by the number of forks the first player did not win on this turn
		secondWinsCount += winSecond[i] * noWinFirst[i]
	}
```
A few things to note. The cheat I mention in the code below is easy enough to fix in a few lines of code, but I decided not to for clarity. No one can win on turns 1 and 2, so the loop start from 3.
