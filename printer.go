package ppanalysis

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	. "github.com/logrusorgru/aurora"
)

func Print(w io.Writer, body AnalysisBody) {
	for p, analysisres := range body {
		fmt.Fprint(w, Sprintf(Gray("package: %s\n"), p))
		for aname, mes := range analysisres {
			fmt.Fprintf(w, Sprintf(Gray("analyzer name: %s\n"), aname))

			for _, m := range mes {
				fmt.Fprintf(w, "location: %s\n", m.Posn)
				fmt.Fprintf(w, Sprintf("message: %s\n", m.Message))

				fileinfos := strings.Split(m.Posn, ":")
				func(fileinfos []string) {
					f, err := os.Open(fileinfos[0])
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()
					linenum, err := strconv.ParseInt(fileinfos[1], 10, 32)
					printSpecificLine(w, f, int(linenum))
				}(fileinfos)

				fmt.Fprint(w, "\n")
			}
		}

	}
}

func printSpecificLine(w io.Writer, r io.Reader, lnum int) {
	var line int
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line++
		if line == lnum-1 || line == lnum+1 || line == lnum-2 || line == lnum+2 {
			fmt.Fprint(w, Sprintf(Gray("%d| %s\n"), Gray(line), Gray(sc.Text())))
		}
		if line == lnum {
			fmt.Fprint(w, Sprintf(Red("%d| %s\n"), Red(line).Bold(), Red(sc.Text()).Bold()))
		}
	}
}
