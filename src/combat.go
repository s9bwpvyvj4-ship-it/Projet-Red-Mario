package main

import (
	"fmt"
	"strings"
	"time"
)

// ==================== UTILITAIRES VISUELS ====================

func HPBar(current, max int, width int) string {
	if max <= 0 {
		return strings.Repeat("░", width)
	}
	filled := current * width / max
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}

	var color string
	ratio := float64(current) / float64(max)
	switch {
	case ratio > 0.6:
		color = Green
	case ratio > 0.3:
		color = Yellow
	default:
		color = Red
	}

	bar := color + strings.Repeat("█", filled) + Reset +
		White + strings.Repeat("░", width-filled) + Reset
	return "[" + bar + "]"
}

func ManaBar(current, max int, width int) string {
	if max <= 0 {
		return strings.Repeat("░", width)
	}
	filled := current * width / max
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	bar := Blue + strings.Repeat("█", filled) + Reset +
		White + strings.Repeat("░", width-filled) + Reset
	return "[" + bar + "]"
}

func FlashDamage() {
	for i := 0; i < 3; i++ {
		fmt.Print("\r" + BgRed + "  💥 DÉGÂTS 💥  " + Reset)
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\r                  ")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

func AnimateAttack(attacker, target string) {
	fmt.Printf("\n%s%s%s ", Bold, attacker, Reset)
	for i := 0; i < 5; i++ {
		fmt.Print(Yellow + "»" + Reset)
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Print(" 💥 ")
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("%s%s%s\n", Bold, target, Reset)
	time.Sleep(300 * time.Millisecond)
}

// ==================== AFFICHAGE DU COMBAT ====================

func DisplayMonster(m *Monster) {
	switch m.Name {
	case "Goomba d'entraînement":
		fmt.Println(Red + GoombaArt + Reset)
	case "Thwomp":
		fmt.Println(Blue + ThwompArt + Reset)
	case "Bowser":
		fmt.Println(Red + BowserArt + Reset)
	}
}

func DisplayCombatUI(player *Character, monster *Monster, turn int) {
	Clear()

	// ===== Bandeau du tour (centré, dynamique) =====
	titreTour := fmt.Sprintf("⚔  TOUR %d ⚔", turn)
	contenuTour := "  " + Bold + Yellow + titreTour + Reset + "  "
	largeurTour := visibleLength(contenuTour)
	margeTour := strings.Repeat(" ", 10)

	fmt.Println(Cyan + "╔" + strings.Repeat("═", largeurTour+40) + "╗" + Reset)
	fmt.Println(Cyan + "║" + Reset + margeTour + contenuTour + strings.Repeat(" ", largeurTour+40-visibleLength(contenuTour)-10) + Cyan + "║" + Reset)
	fmt.Println(Cyan + "╚" + strings.Repeat("═", largeurTour+40) + "╝" + Reset)
	fmt.Println()

	// ===== Cadre dynamique autour du nom du monstre =====
	nomMonstre := "  " + Bold + monster.Name + Reset + "  "
	largeurNom := visibleLength(nomMonstre)
	bordureNom := strings.Repeat("═", largeurNom)
	marge := strings.Repeat(" ", 20)

	fmt.Println(marge + Red + "╔" + bordureNom + "╗" + Reset)
	fmt.Println(marge + Red + "║" + Reset + nomMonstre + Red + "║" + Reset)
	fmt.Println(marge + Red + "╚" + bordureNom + "╝" + Reset)

	DisplayMonster(monster)

	monsterHP := HPBar(monster.CurrentHP, monster.MaxHP, 20)
	fmt.Printf("       %s%s%s  %s  %s%d/%d%s\n\n",
		Bold, monster.Name, Reset,
		monsterHP,
		Red, monster.CurrentHP, monster.MaxHP, Reset)

	fmt.Println(White + "──────────────────────────────────────────────────────────" + Reset)
	fmt.Println()

	DisplayCharacterArt(player.Class)

	playerHP := HPBar(player.CurrentHP, player.MaxHP, 20)
	playerMana := ManaBar(player.Mana, player.ManaMax, 20)

	fmt.Printf("       %s%s%s  %s  %s%d/%d PV%s\n",
		Bold, player.Name, Reset,
		playerHP,
		Green, player.CurrentHP, player.MaxHP, Reset)
	fmt.Printf("       %sMana%s   %s  %s%d/%d%s\n",
		Blue, Reset,
		playerMana,
		Blue, player.Mana, player.ManaMax, Reset)
	fmt.Printf("       %sPièces : %d%s   %sExp : %d/%d%s\n\n",
		Yellow, player.Gold, Reset,
		Purple, player.Exp, player.ExpMax, Reset)

	fmt.Println(White + "──────────────────────────────────────────────────────────" + Reset)
}

// ==================== COMBAT PRINCIPAL ====================

func TrainingFight(player *Character, monster Monster) {
	PlayBattleMusic()

	defer func() {
		StopBattleMusic()
		PlayBackgroundMusic()
	}()

	Clear()

	DisplayMonster(&monster)
	SlowPrint(Yellow+"⚡ Un "+monster.Name+" apparaît !", 30*time.Millisecond)
	fmt.Print(White + "Préparation au combat" + Reset)
	LoadingDots()
	time.Sleep(500 * time.Millisecond)
	Pause()

	turn := 1
	playerFirst := player.Initiative >= monster.Initiative

	Clear()
	if playerFirst {
		SlowPrint(Green+"⚡ Vous êtes plus rapide que le monstre !", 25*time.Millisecond)
	} else {
		SlowPrint(Purple+"⚡ Le monstre est plus rapide que vous !", 25*time.Millisecond)
	}
	time.Sleep(1 * time.Second)

	for player.CurrentHP > 0 && monster.CurrentHP > 0 {
		DisplayCombatUI(player, &monster, turn)
		time.Sleep(400 * time.Millisecond)

		if playerFirst {
			player.CharacterTurn(&monster)
			if monster.CurrentHP <= 0 {
				break
			}
			fmt.Println()
			fmt.Print(Yellow + "Le monstre riposte" + Reset)
			LoadingDots()
			time.Sleep(400 * time.Millisecond)
			DisplayCombatUI(player, &monster, turn)
			time.Sleep(300 * time.Millisecond)
			monster.Pattern(turn, player)
		} else {
			monster.Pattern(turn, player)
			if player.CurrentHP <= 0 {
				break
			}
			fmt.Println()
			fmt.Print(Green + "À vous de jouer" + Reset)
			LoadingDots()
			time.Sleep(400 * time.Millisecond)
			DisplayCombatUI(player, &monster, turn)
			time.Sleep(300 * time.Millisecond)
			player.CharacterTurn(&monster)
		}

		turn++
		Pause()
	}

	// ==================== FIN DU COMBAT ====================
	Clear()

	if player.CurrentHP <= 0 {
		fmt.Println(Red + GameOverArt + Reset)
		time.Sleep(500 * time.Millisecond)
		SlowPrint(Red+"Vous vous effondrez...", 40*time.Millisecond)
		PlayDeathSound()
		time.Sleep(2 * time.Second)
		player.IsDead()
	} else {
		fmt.Println(Green + "╔══════════════════════════════════════╗" + Reset)
		fmt.Println(Green + "║          ★ ★ ★ VICTOIRE ! ★ ★ ★      ║" + Reset)
		fmt.Println(Green + "╚══════════════════════════════════════╝" + Reset)

		PlayVictorySound()
		time.Sleep(500 * time.Millisecond)
		fmt.Println()
		SlowPrint(Yellow+fmt.Sprintf("Vous gagnez %d EXP et %d pièces !",
			monster.ExpReward, monster.GoldReward), 25*time.Millisecond)
		player.GainExp(monster.ExpReward)
		player.Gold += monster.GoldReward
	}
	Pause()
}

// ==================== TOUR DU JOUEUR ====================

func (c *Character) CharacterTurn(monster *Monster) {
	for {
		fmt.Printf("  %s1.%s %s Attaquer (Saut)%s\n", Cyan, Reset, Bold, Reset)
		fmt.Printf("  %s2.%s Inventaire\n", Cyan, Reset)
		fmt.Printf("  %s3.%s Sorts (%sMana: %d/%d%s)\n", Cyan, Reset, Blue, c.Mana, c.ManaMax, Reset)

		choice := AskInt("\nVotre choix : ")

		switch choice {
		case 1:
			c.PlayerAttack(monster, "Saut", 5)
			return

		case 2:
			if c.AccessInventoryInCombat(monster) {
				return
			}
			Clear()
			DisplayCombatUI(c, monster, 0)
			continue

		case 3:
			if c.UseSpellInCombat(monster) {
				return
			}

		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

func (c *Character) PlayerAttack(monster *Monster, attackName string, damage int) {
	Clear()
	fmt.Println(Cyan + "═══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()

	AnimateAttack(c.Name, monster.Name)
	FlashDamage()

	monster.CurrentHP -= damage
	if monster.CurrentHP < 0 {
		monster.CurrentHP = 0
	}

	SlowPrint(Green+fmt.Sprintf("%s utilise %s et inflige %d dégâts !",
		c.Name, attackName, damage), 15*time.Millisecond)
	time.Sleep(400 * time.Millisecond)

	fmt.Println()
	fmt.Printf("  %s%s%s  %s  %s%d/%d%s\n",
		Bold, monster.Name, Reset,
		HPBar(monster.CurrentHP, monster.MaxHP, 25),
		Red, monster.CurrentHP, monster.MaxHP, Reset)
}

// ==================== SORTS EN COMBAT ====================

func (c *Character) UseSpellInCombat(monster *Monster) bool {
	Clear()
	fmt.Println(Cyan + "╔════════════════════════════════════════╗" + Reset)
	fmt.Println(Cyan + "║              📖 SORTS                  ║" + Reset)
	fmt.Println(Cyan + "╚════════════════════════════════════════╝" + Reset)
	fmt.Printf("  Mana : %s %s%d/%d%s\n\n",
		ManaBar(c.Mana, c.ManaMax, 15), Blue, c.Mana, c.ManaMax, Reset)

	for i, s := range c.Skills {
		cost := spellManaCost(s)
		dmg := spellDamage(s)
		color := Green
		if c.Mana < cost {
			color = Red
		}

		if s == "Étoile" {
			fmt.Printf("  %s%d.%s %-15s %s(SOIN +40 PV, %d mana)%s\n",
				Cyan, i+1, Reset, s, color, cost, Reset)
		} else {
			fmt.Printf("  %s%d.%s %-15s %s(%d dégâts, %d mana)%s\n",
				Cyan, i+1, Reset, s, color, dmg, cost, Reset)
		}
	}
	fmt.Printf("  %s0.%s Retour\n", Yellow, Reset)

	choice := AskInt("\nChoix : ")
	if choice == 0 {
		return false
	}
	if choice < 1 || choice > len(c.Skills) {
		fmt.Println(Red + "Choix invalide." + Reset)
		time.Sleep(500 * time.Millisecond)
		return false
	}

	spell := c.Skills[choice-1]
	cost := spellManaCost(spell)

	if c.Mana < cost {
		fmt.Println(Red + "❌ Mana insuffisant !" + Reset)
		time.Sleep(800 * time.Millisecond)
		return false
	}

	c.Mana -= cost

	if spell == "Étoile" {
		return c.CastHealSpell()
	}

	damage := spellDamage(spell)

	Clear()
	fmt.Println(Cyan + "═══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()

	if spell == "Boule de Feu" {
		fmt.Printf("%s%s%s lance une ", Bold, c.Name, Reset)
		time.Sleep(200 * time.Millisecond)
		fmt.Print(Red + "🔥 " + Yellow + "BOULE " + Red + "DE " + Yellow + "FEU " + Red + "🔥" + Reset)
		time.Sleep(400 * time.Millisecond)
		fmt.Println()
		for i := 0; i < 5; i++ {
			fmt.Print(Yellow + "»" + Red + "»" + Yellow + "»" + Reset + " ")
			time.Sleep(80 * time.Millisecond)
		}
		fmt.Print(" 💥💥💥")
		fmt.Println()
	} else {
		AnimateAttack(c.Name, monster.Name)
	}

	FlashDamage()

	monster.CurrentHP -= damage
	if monster.CurrentHP < 0 {
		monster.CurrentHP = 0
	}

	time.Sleep(300 * time.Millisecond)
	SlowPrint(Green+fmt.Sprintf("%s lance %s et inflige %d dégâts !",
		c.Name, spell, damage), 15*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	fmt.Println()
	fmt.Printf("  %s%s%s  %s  %s%d/%d%s\n",
		Bold, monster.Name, Reset,
		HPBar(monster.CurrentHP, monster.MaxHP, 25),
		Red, monster.CurrentHP, monster.MaxHP, Reset)
	fmt.Printf("  %sMana restant : %s %s%d/%d%s\n",
		Blue, ManaBar(c.Mana, c.ManaMax, 15), Blue, c.Mana, c.ManaMax, Reset)
	return true
}

// ==================== SORT DE SOIN ÉTOILE ====================

func (c *Character) CastHealSpell() bool {
	Clear()
	fmt.Println(Cyan + "═══════════════════════════════════════════════════════════" + Reset)
	fmt.Println()

	fmt.Printf("%s%s%s invoque ", Bold, c.Name, Reset)
	time.Sleep(200 * time.Millisecond)
	fmt.Print(Yellow + "✨ ÉTOILE ✨" + Reset)
	time.Sleep(400 * time.Millisecond)
	fmt.Println()

	for i := 0; i < 5; i++ {
		fmt.Print(Yellow + "★ " + Reset)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()

	before := c.CurrentHP
	c.CurrentHP += 40
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	healed := c.CurrentHP - before

	time.Sleep(300 * time.Millisecond)
	SlowPrint(Green+fmt.Sprintf("Vous récupérez %d PV !", healed), 15*time.Millisecond)
	time.Sleep(300 * time.Millisecond)

	fmt.Println()
	fmt.Printf("  %s%s%s  %s  %s%d/%d%s\n",
		Bold, c.Name, Reset,
		HPBar(c.CurrentHP, c.MaxHP, 25),
		Green, c.CurrentHP, c.MaxHP, Reset)
	fmt.Printf("  %sMana restant : %s %s%d/%d%s\n",
		Blue, ManaBar(c.Mana, c.ManaMax, 15), Blue, c.Mana, c.ManaMax, Reset)
	return true
}

// ==================== COÛTS ET DÉGÂTS DES SORTS ====================

func spellManaCost(spell string) int {
	switch spell {
	case "coup de poing":
		return 0
	case "Boule de Feu":
		return 15
	case "Étoile":
		return 20
	}
	return 0
}

func spellDamage(spell string) int {
	switch spell {
	case "coup de poing":
		return 8
	case "Boule de Feu":
		return 18
	case "Étoile":
		return 0
	}
	return 10
}