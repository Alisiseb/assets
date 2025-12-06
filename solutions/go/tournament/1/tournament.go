package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type stats struct {
	MP int
	W  int
	D  int
	L  int
	P  int
}

func Tally(reader io.Reader, writer io.Writer) error {
	Team := make(map[string]*stats)
	Newline := bufio.NewScanner(reader)
	for Newline.Scan() {
		line := Newline.Text()
		if line == "\n" || line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		part := strings.Split(line, ";")
		if len(part) != 3 {
			return errors.New("invalid file")
		}
		t1 := ensureTeamMapHaveAmount(Team, part[0])
		t2 := ensureTeamMapHaveAmount(Team, part[1])
		switch part[2] {
		case "win":
			win(t1, t2)
		case "loss":
			win(t2, t1)
		case "draw":
			draw(t1, t2)
		default:
			return errors.New("invalid result")
		}
	}
	list := make([]string, 0, len(Team))
	for k := range Team {
		list = append(list, k)
	}
	sort.SliceStable(list, func(i, j int) bool {
		if Team[list[i]].P == Team[list[j]].P {
			return list[i] < list[j]
		}
		return Team[list[i]].P > Team[list[j]].P
	})
	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, v := range list {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", v, Team[v].MP, Team[v].W, Team[v].D, Team[v].L, Team[v].P)
	}
	return nil
}
func win(winner, loser *stats) {
	winner.P = winner.P + 3
	winner.W++
	winner.MP++
	loser.L++
	loser.MP++
}

func draw(teamName1, teamName2 *stats) {
	teamName1.D++
	teamName2.D++
	teamName1.MP++
	teamName2.MP++
	teamName1.P++
	teamName2.P++
}
func ensureTeamMapHaveAmount(t map[string]*stats, name string) *stats {
	if t[name] == nil {
		t[name] = &stats{}
	}
	return t[name]
}
