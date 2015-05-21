package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type Tshark struct {
	Protocol  string
	Interface string
}

// ValidateInstallation Will figure out if tshark is available and if so will print version
func (t Tshark) ValidateInstallation() error {

	if out, err := exec.Command("tshark", "-v").Output(); err != nil {
		return err
	} else {
		vline := strings.SplitAfter(string(out), "\n")[0]
		Debug("Tshark info: %s", strings.Replace(vline, "\"", "", -1))
	}

	return nil
}

// Capture -
func (t Tshark) Capture(packets chan []byte) error {

	args := []string{
		"-i", fmt.Sprintf("%s", t.Interface),
		"-Y", fmt.Sprintf("%s", t.Protocol),
		"-O", fmt.Sprintf("%s", t.Protocol),
		"-V", "-x", "-l", "-c", "1",
	}

	for {
		var pdu []byte

		cmd := exec.Command("tshark", args...)

		stdout, err := cmd.StdoutPipe()

		if err != nil {
			return err
		}

		if err := cmd.Start(); err != nil {
			return err
		}

		scanner := bufio.NewScanner(stdout)

		go func() {
			for scanner.Scan() {
				for _, b := range []byte(scanner.Text()) {
					pdu = append(pdu, b)
				}
			}
		}()

		if err := cmd.Wait(); err != nil {
			return err
		}

		if len(pdu) > 0 {
			packets <- pdu
		}
	}

	return nil
}

func NewTshark(iface string) (Tshark, error) {
	tshark := Tshark{
		Protocol:  "smpp",
		Interface: iface,
	}

	if err := tshark.ValidateInstallation(); err != nil {
		return tshark, err
	}

	return tshark, nil
}
