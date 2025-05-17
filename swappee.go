package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	hook "github.com/robotn/gohook"
)

func main() {
	hook.Register(hook.KeyDown, []string{"\"", "cmd", "shift"}, func(e hook.Event) {
		togglePressAndHold(true)
	})
	hook.Register(hook.KeyDown, []string{"\\", "cmd", "shift"}, func(e hook.Event) {
		togglePressAndHold(false)
	})

	s := hook.Start()
	<-hook.Process(s)
}

func togglePressAndHold(toggle bool) {
	toggleStr := strconv.FormatBool(toggle)
	cmd := exec.Command("bash", "-c", "defaults write -g ApplePressAndHoldEnabled -bool "+toggleStr)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Toggled to %s\n", toggleStr)
}
