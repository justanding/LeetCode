package predictPartyVictory_test

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	t.Log(predictPartyVictory("RD"))
	t.Log(predictPartyVictory("RDD"))
	t.Log(predictPartyVictory("DDRRR"))
}

/*
 * https://leetcode-cn.com/problems/dota2-senate/
 */
func predictPartyVictory(senate string) string {
	R, D := true, true
	person := 0
	s := []byte(senate)
	l := len(s)
	for R && D {
		R, D = false, false
		for i := 0; i < l; i++ {
			if s[i] == 'R' {
				R = true
				if person < 0 {
					s[i] = 'O'
				}
				person++
			} else if s[i] == 'D' {
				D = true
				if person > 0 {
					s[i] = 'O'
				}
				person--
			}
		}
	}

	if person > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
