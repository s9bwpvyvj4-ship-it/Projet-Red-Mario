package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func PlaySound(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fallbackBeep()
		return
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("afplay", filePath)

	case "linux":
		if _, err := exec.LookPath("mpg123"); err == nil {
			cmd = exec.Command("mpg123", "-q", filePath)
		} else if _, err := exec.LookPath("aplay"); err == nil {
			cmd = exec.Command("aplay", filePath)
		} else if _, err := exec.LookPath("paplay"); err == nil {
			cmd = exec.Command("paplay", filePath)
		} else {
			fallbackBeep()
			return
		}

	case "windows":
		cmd = exec.Command("powershell", "-c",
			fmt.Sprintf("(New-Object Media.SoundPlayer '%s').PlaySync()", filePath))

	default:
		fallbackBeep()
		return
	}

	_ = cmd.Start()
}

func fallbackBeep() {
	fmt.Print("\a")
}

var musicCmd *exec.Cmd

func PlayBackgroundMusic() {
	StopBackgroundMusic()

	filePath := "sounds/theme.mp3"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return
	}

	switch runtime.GOOS {
	case "darwin":
		musicCmd = exec.Command("afplay", filePath)

	case "linux":
		if _, err := exec.LookPath("mpg123"); err == nil {
			musicCmd = exec.Command("mpg123", "-q", filePath)
		} else if _, err := exec.LookPath("aplay"); err == nil {
			musicCmd = exec.Command("aplay", filePath)
		} else {
			return
		}

	case "windows":
		musicCmd = exec.Command("powershell", "-c",
			fmt.Sprintf("(New-Object Media.SoundPlayer '%s').PlayLooping()", filePath))

	default:
		return
	}

	_ = musicCmd.Start()
}

func StopBackgroundMusic() {
	if musicCmd != nil && musicCmd.Process != nil {
		_ = musicCmd.Process.Kill()
		musicCmd = nil
	}
}


var battleCmd *exec.Cmd

func PlayBattleMusic() {
	StopBackgroundMusic()

	filePath := "sounds/battle.mp3"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		PlayBackgroundMusic()
		return
	}

	switch runtime.GOOS {
	case "darwin":
		battleCmd = exec.Command("afplay", filePath)

	case "linux":
		if _, err := exec.LookPath("mpg123"); err == nil {
			battleCmd = exec.Command("mpg123", "-q", filePath)
		} else if _, err := exec.LookPath("aplay"); err == nil {
			battleCmd = exec.Command("aplay", filePath)
		} else {
			return
		}

	case "windows":
		battleCmd = exec.Command("powershell", "-c",
			fmt.Sprintf("(New-Object Media.SoundPlayer '%s').PlayLooping()", filePath))

	default:
		return
	}

	_ = battleCmd.Start()
}

// StopBattleMusic arrête la musique de combat
func StopBattleMusic() {
	if battleCmd != nil && battleCmd.Process != nil {
		_ = battleCmd.Process.Kill()
		battleCmd = nil
	}
}

func PlayDeathSound() {
	PlaySound("sounds/mario-death.mp3")
}

func PlayVictorySound() {
	PlaySound("sounds/victory.mp3")
}

func PlayCoinSound() {
	PlaySound("sounds/coin.mp3")
}

func PlayHammerSound() {
	PlaySound("sounds/hammer.mp3")
}

func PlayQuitSound() {
	PlaySound("sounds/quit.mp3")
}


func BeepShort() {
	fmt.Print("\a")
}

func BeepDouble() {
	fmt.Print("\a")
	time.Sleep(100 * time.Millisecond)
	fmt.Print("\a")
}

func BeepTriple() {
	fmt.Print("\a")
	time.Sleep(200 * time.Millisecond)
	fmt.Print("\a")
	time.Sleep(300 * time.Millisecond)
	fmt.Print("\a")
}
