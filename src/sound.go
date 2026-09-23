package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

<<<<<<< HEAD
func PlaySound(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
=======
// ==================== LECTURE AUDIO ====================

func PlaySound(filePath string) {
	if _, err := os.Stat(filePath); err != nil {
>>>>>>> 7287e4d6f999db582656947acc8a74bd9b791839
		fallbackBeep()
		return
	}

	switch runtime.GOOS {
	case "darwin":
<<<<<<< HEAD
		cmd = exec.Command("afplay", filePath)
=======
		exec.Command("afplay", filePath).Start()
>>>>>>> 7287e4d6f999db582656947acc8a74bd9b791839

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
<<<<<<< HEAD

	_ = cmd.Start()
}

=======
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

>>>>>>> 7287e4d6f999db582656947acc8a74bd9b791839
func fallbackBeep() {
	fmt.Print("\a")
}


<<<<<<< HEAD
var musicCmd *exec.Cmd 
=======
var musicCmd *exec.Cmd
>>>>>>> 7287e4d6f999db582656947acc8a74bd9b791839

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


<<<<<<< HEAD
var battleCmd *exec.Cmd 
=======
var battleCmd *exec.Cmd
>>>>>>> 7287e4d6f999db582656947acc8a74bd9b791839

func PlayBattleMusic() {
	StopBackgroundMusic()

	filePath := "sounds/battle.mp3"

<<<<<<< HEAD
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
=======
	if _, err := os.Stat(filePath); err != nil {
>>>>>>> 7287e4d6f999db582656947acc8a74bd9b791839
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

<<<<<<< HEAD
=======
// ==================== SONS DU JEU ====================

>>>>>>> 7287e4d6f999db582656947acc8a74bd9b791839
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
