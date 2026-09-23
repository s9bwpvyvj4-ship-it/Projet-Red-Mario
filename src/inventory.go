package main

import (
	"fmt"
	"strings"
	"time"
)


func (c *Character) AddInventory(item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println(Red + "❌ Inventaire plein !" + Reset)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func (c *Character) RemoveInventory(item string) bool {
	for i, v := range c.Inventory {
		if v == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Character) CountItem(item string) int {
	count := 0
	for _, v := range c.Inventory {
		if v == item {
			count++
		}
	}
	return count
}

func (c *Character) HasItem(item string) bool {
	return c.CountItem(item) > 0
}

func (c *Character) CanAddItem() bool {
	return len(c.Inventory) < c.MaxInventory
}

func (c *Character) InventoryFullPercent() int {
	if c.MaxInventory == 0 {
		return 100
	}
	return len(c.Inventory) * 100 / c.MaxInventory
}

func (c *Character) ListItemsByPrefix(prefix string) []string {
	var result []string
	for _, item := range c.Inventory {
		if strings.HasPrefix(item, prefix) {
			result = append(result, item)
		}
	}
	return result
}


func (c *Character) AccessInventory(target Damageable) {
	for {
		Clear()
		PrintTitle("INVENTAIRE")
		fmt.Printf("  Capacité : %s%d / %d%s\n\n",
			Yellow, len(c.Inventory), c.MaxInventory, Reset)

		if len(c.Inventory) == 0 {
			fmt.Println(Red + "  Inventaire vide." + Reset)
		} else {
			for i, item := range c.Inventory {
				icon := itemIcon(item)
				fmt.Printf("  %s%d.%s %s %s\n", Cyan, i+1, Reset, icon, item)
			}
		}
		fmt.Println()
		fmt.Printf("  %s0.%s Retour\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
			continue
		}

		item := c.Inventory[choice-1]
		c.UseItem(item, target)
		Pause()
	}
}

func (c *Character) AccessInventoryInCombat(opponent Damageable) bool {
	for {
		Clear()
		PrintTitle("INVENTAIRE (combat)")
		fmt.Printf("  Capacité : %s%d / %d%s\n\n",
			Yellow, len(c.Inventory), c.MaxInventory, Reset)

		if len(c.Inventory) == 0 {
			fmt.Println(Red + "  Inventaire vide." + Reset)
		} else {
			for i, item := range c.Inventory {
				icon := itemIcon(item)
				fmt.Printf("  %s%d.%s %s %s\n", Cyan, i+1, Reset, icon, item)
			}
		}
		fmt.Println()
		fmt.Printf("  %s0.%s Retour (sans utiliser d'item)\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")
		if choice == 0 {
			return false
		}
		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
			continue
		}

		item := c.Inventory[choice-1]
		c.UseItem(item, opponent)
		Pause()
		return true
	}
}

func itemIcon(item string) string {
	switch {
	case strings.Contains(item, "Champignon Super"):
		return "🍄"
	case strings.Contains(item, "Champignon Poison"):
		return "☠️"
	case strings.Contains(item, "Potion de Mana"):
		return "🔵"
	case strings.Contains(item, "Fleur de Feu"):
		return "🔥"
	case strings.Contains(item, "Casquette"):
		return "🧢"
	case strings.Contains(item, "Salopette"):
		return "👕"
	case strings.Contains(item, "Bottes"):
		return "👢"
	case strings.Contains(item, "Fourrure"):
		return "🐺"
	case strings.Contains(item, "Peau de Troll"):
		return "🧟"
	case strings.Contains(item, "Cuir"):
		return "🐗"
	case strings.Contains(item, "Plume"):
		return "🪶"
	default:
		return "📦"
	}
}


func (c *Character) UseItem(item string, target Damageable) {
	switch item {
	case "Champignon Super":
		c.TakePot()
	case "Champignon Poison":
		c.PoisonPot(target)
	case "Fleur de Feu":
		c.SpellBook()
	case "Potion de Mana":
		c.DrinkManaPot()
	case "Casquette Mario", "Salopette Mario", "Bottes Kuribo":
		c.EquipItem(item)
	default:
		fmt.Printf(Yellow+"%s ne peut pas être utilisé ici.\n"+Reset, item)
	}
}


func (c *Character) TakePot() {
	if !c.RemoveInventory("Champignon Super") {
		fmt.Println(Red + "Pas de Champignon Super dans l'inventaire." + Reset)
		return
	}
	fmt.Println(Green + "🍄 Vous utilisez Champignon Super !" + Reset)
	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf(Green+"PV : %d / %d\n"+Reset, c.CurrentHP, c.MaxHP)
}

func (c *Character) PoisonPot(target Damageable) {
	if !c.RemoveInventory("Champignon Poison") {
		fmt.Println(Red + "Pas de Champignon Poison dans l'inventaire." + Reset)
		return
	}

	if target == nil {
		fmt.Println(Yellow + "☠ Vous lancez le Champignon Poison... mais personne en face !" + Reset)
		return
	}

	fmt.Println(Purple + "☠ Vous lancez Champignon Poison sur " + target.GetName() + " !" + Reset)

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		hp := target.GetHP() - 10
		if hp < 0 {
			hp = 0
		}
		target.SetHP(hp)
		fmt.Printf(Red+"Poison ! %s PV : %d / %d\n"+Reset,
			target.GetName(), hp, target.GetMaxHP())
		if target.IsDead() {
			return
		}
	}
}

func (c *Character) DrinkManaPot() {
	if !c.RemoveInventory("Potion de Mana") {
		fmt.Println(Red + "Pas de Potion de Mana dans l'inventaire." + Reset)
		return
	}
	fmt.Println(Blue + "🔵 Vous buvez une Potion de Mana !" + Reset)
	c.Mana += 30
	if c.Mana > c.ManaMax {
		c.Mana = c.ManaMax
	}
	fmt.Printf(Blue+"Mana : %d / %d\n"+Reset, c.Mana, c.ManaMax)
}


func (c *Character) SpellBook() {
	for _, s := range c.Skills {
		if s == "Boule de Feu" {
			fmt.Println(Yellow + "Vous connaissez déjà ce sort !" + Reset)
			return
		}
	}
	c.Skills = append(c.Skills, "Boule de Feu")
	c.RemoveInventory("Fleur de Feu")
	fmt.Println(Green + "🔥 Vous apprenez : Boule de Feu !" + Reset)
}


func (c *Character) EquipItem(item string) {
	var slot *string
	var bonus int

	switch item {
	case "Casquette Mario":
		slot = &c.Equipment.Head
		bonus = 10
	case "Salopette Mario":
		slot = &c.Equipment.Torso
		bonus = 25
	case "Bottes Kuribo":
		slot = &c.Equipment.Feet
		bonus = 15
	default:
		return
	}

	if *slot != "" {
		oldItem := *slot
		oldBonus := equipmentBonus(oldItem)
		c.MaxHP -= oldBonus
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		c.AddInventory(oldItem)
		fmt.Printf(Yellow+"↩ Vous rangez : %s\n"+Reset, oldItem)
	}

	*slot = item
	c.MaxHP += bonus
	c.RemoveInventory(item)
	fmt.Printf(Green+"✓ Vous équipez : %s (+%d PV max)\n"+Reset, item, bonus)
}

func equipmentBonus(item string) int {
	switch item {
	case "Casquette Mario":
		return 10
	case "Salopette Mario":
		return 25
	case "Bottes Kuribo":
		return 15
	}
	return 0
}


func (c *Character) UpgradeInventorySlot() bool {
	if c.InventoryUpgrades >= 3 {
		fmt.Println(Red + "Vous avez déjà utilisé toutes vos augmentations !" + Reset)
		return false
	}
	c.MaxInventory += 10
	c.InventoryUpgrades++
	fmt.Printf(Green+"✓ Capacité d'inventaire augmentée à %d !\n"+Reset, c.MaxInventory)
	return true
}


func (c *Character) DisplayEquipment() {
	fmt.Println(Cyan + "\n=== Équipement actuel ===" + Reset)
	fmt.Printf("  🧢 Tête  : %s\n", emptyOr(c.Equipment.Head, "—"))
	fmt.Printf("  👕 Torse : %s\n", emptyOr(c.Equipment.Torso, "—"))
	fmt.Printf("  👢 Pieds : %s\n", emptyOr(c.Equipment.Feet, "—"))
}


func hasPrefixAny(s string, prefixes []string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
