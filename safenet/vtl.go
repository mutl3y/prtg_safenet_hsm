package safenet

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/PaesslerAG/go-prtg-sensor-api"
	"os/exec"
	"strings"
	"time"
)

type vtl struct {
	dir, app string
}

func NewVtl(dir, app string) *vtl {
	v := &vtl{}
	v.dir = dir
	v.app = app
	return v
}

type Vtl interface {
	Verify() (resT time.Duration, err error)
}

func (v *vtl) Verify(s string) (rtnErr error) {

	start := time.Now()
	_, ss, err := v.runVtl("verify")
	if err != nil {
		return err
	}
	runTime := time.Since(start)
	r := new(prtg.SensorResponse)

	if len(ss) < 1 {
		r.SensorResult.Error = 1
		r.SensorResult.Text = fmt.Sprintf("no slots returned ")
	}

	r.Channels = append(r.Channels, prtg.SensorChannel{
		Name:  fmt.Sprintf("slots"),
		Value: float64(len(ss)),
		Unit:  prtg.UnitOne,
	})

	r.Channels = append(r.Channels, prtg.SensorChannel{
		Name:  fmt.Sprintf("response time"),
		Value: runTime.Round(time.Microsecond).Seconds() * 1000,
		Unit:  prtg.UnitTimeResponse,
	})

	// check serial if specified
	if s != "" {
		if s == ss[0][1] {
			r.Channels = append(r.Channels, prtg.SensorChannel{Name: "serial match", ValueLookup: "true"})
		} else {
			r.Channels = append(r.Channels, prtg.SensorChannel{
				Name:        "serial match",
				ValueLookup: "false",
				Warning:     1,
			})
			txt := "serial number mismatch want %v got %v"
			r.SensorResult.Text = fmt.Sprintf(txt, s, ss[0][1])
			rtnErr = fmt.Errorf(txt, s, ss[0][1])
		}
	}

	fmt.Println(r.String())

	return
}

func (v *vtl) runVtl(args string) (headers []string, data [][]string, err error) {
	sh := make([]string, 0, 10)
	ss := make([][]string, 0, 10)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	cmd := exec.CommandContext(ctx, v.app, args)
	cmd.Dir = v.dir
	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(combinedOutput))
	scanner.Split(bufio.ScanLines)
	runOnce := true
	for scanner.Scan() {
		str, scanErr := vtlCleanup(scanner.Text())
		if scanErr != nil {
			fmt.Printf("Error parsing %v\n", scanErr)
		}
		//fmt.Printf("str %v\n",str)
		if str != "" {
			// capture headers
			if runOnce {
				sh = strings.Fields(str)
				runOnce = false
			} else {
				ss = append(ss, strings.Fields(str))
			}

		}

	}

	return sh, ss, nil
}

// strips comments and headers
func vtlCleanup(b string) (string, error) {
	switch {
	case strings.HasPrefix(b, "="):
		fallthrough
	case strings.Contains(b, "Luna"):
		return "", nil
	}
	return strings.ReplaceAll(b, "#", ""), nil
}
