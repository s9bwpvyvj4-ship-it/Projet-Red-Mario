package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// ==================== LECTURE AUDIO ====================

func PlaySound(filePath string) {
	if _, err := os.Stat(filePath); err != nil {
		fallbackBeep()
		return
	}

	switch runtime.GOOS {
	case "darwin":
		exec.Command("afplay", filePath).Start()

	case "linux":
		if _, err := exec.LookPath("mpg123"); err == nil {
			exec.Command("mpg123", "-q", filePath).Start()
		} else {
			fallbackBeep()
		}

	case "windows":
		playMP3(filePath)

	default:
		fallbackBeep()
	}
}

// playMP3 utilise PowerShell pour lire un MP3 avec Windows
func playMP3(filePath string) {
	script := fmt.Sprintf(`
Add-Type -AssemblyName presentationCore
$player = New-Object System.Windows.Media.MediaPlayer
$player.Open([Uri]::new((Resolve-Path '%s').Path))
$player.Play()
Start-Sleep -Seconds 3
$player.Close()
`, filePath)

	exec.Command(
		"powershell",
		"-NoProfile",
		"-Command",
		script,
	).Start()
}

func fallbackBeep() {
	fmt.Print("\a")
}

// ==================== MUSIQUE DE FOND ====================

var musicCmd *exec.Cmd

func PlayBackgroundMusic() {
	StopBackgroundMusic()

	filePath := "sounds/theme.mp3"

	if _, err := os.Stat(filePath); err != nil {
		return
	}

	switch runtime.GOOS {
	case "windows":
		musicCmd = exec.Command(
			"powershell",
			"-NoProfile",
			"-Command",
			fmt.Sprintf(`
Add-Type -AssemblyName presentationCore
$player = New-Object System.Windows.Media.MediaPlayer
$player.Open([Uri]::new((Resolve-Path '%s').Path))
$player.Play()
while ($true) {
	Start-Sleep -Seconds 1
}
`, filePath),
		)

	case "darwin":
		musicCmd = exec.Command("afplay", filePath)

	case "linux":
		if _, err := exec.LookPath("mpg123"); err == nil {
			musicCmd = exec.Command("mpg123", "-q", filePath)
		} else {
			return
		}

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

// ==================== MUSIQUE DE COMBAT ====================

var battleCmd *exec.Cmd

func PlayBattleMusic() {
	StopBackgroundMusic()

	filePath := "sounds/battle.mp3"

	if _, err := os.Stat(filePath); err != nil {
		PlayBackgroundMusic()
		return
	}

	switch runtime.GOOS {
	case "windows":
		battleCmd = exec.Command(
			"powershell",
			"-NoProfile",
			"-Command",
			fmt.Sprintf(`
Add-Type -AssemblyName presentationCore
$player = New-Object System.Windows.Media.MediaPlayer
$player.Open([Uri]::new((Resolve-Path '%s').Path))
$player.Play()
while ($true) {
	Start-Sleep -Seconds 1
}
`, filePath),
		)

	case "darwin":
		battleCmd = exec.Command("afplay", filePath)

	case "linux":
		if _, err := exec.LookPath("mpg123"); err == nil {
			battleCmd = exec.Command("mpg123", "-q", filePath)
		} else {
			return
		}

	default:
		return
	}

	_ = battleCmd.Start()
}

func StopBattleMusic() {
	if battleCmd != nil && battleCmd.Process != nil {
		_ = battleCmd.Process.Kill()
		battleCmd = nil
	}
}

// ==================== SONS DU JEU ====================

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

// ==================== BIPS DE SECOURS ====================

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
