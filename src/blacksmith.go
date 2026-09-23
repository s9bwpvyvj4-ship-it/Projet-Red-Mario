package main

import "fmt"


type Recipe struct {
	Name      string
	Cost      int
	Materials map[string]int
}


var recipes = []Recipe{
	{"Casquette Mario", 5, map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}},
	{"Salopette Mario", 5, map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}},
	{"Bottes Kuribo", 5, map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}},
}


func (c *Character) Blacksmith() {
	for {
		Clear()

		fmt.Println(Yellow + BlacksmithArt + Reset)

		PrintTitle("🔨 FORGERON")

		fmt.Printf("  %sVotre bourse : %d pièces%s\n", Yellow, c.Gold, Reset)
		fmt.Printf("  %sInventaire  : %d / %d%s\n\n",
			Cyan, len(c.Inventory), c.MaxInventory, Reset)

		for i, r := range recipes {
			canCraft := c.canCraft(r)

			titleColor := Green
			if !canCraft {
				titleColor = Red
			}

			fmt.Printf("  %s%d.%s %s%s%s (%d pièces)\n",
				Cyan, i+1, Reset, titleColor, r.Name, Reset, r.Cost)

			for mat, qty := range r.Materials {
				have := c.CountItem(mat)
				matColor := Green
				if have < qty {
					matColor = Red
				}
				fmt.Printf("       %s- %s x%d (vous : %d)%s\n",
					matColor, mat, qty, have, Reset)
			}
			fmt.Println()
		}
		fmt.Printf("  %s0.%s Retour\n", Yellow, Reset)

		choice := AskInt("\nChoix : ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(recipes) {
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
			continue
		}

		c.CraftItem(recipes[choice-1])
		Pause()
	}
}


func (c *Character) CraftItem(r Recipe) {
	if c.Gold < r.Cost {
		fmt.Println(Red + "❌ Pas assez de pièces !" + Reset)
		return
	}

	for mat, qty := range r.Materials {
		if c.CountItem(mat) < qty {
			fmt.Printf(Red+"❌ Ressource manquante : %s x%d\n"+Reset, mat, qty)
			return
		}
	}

	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println(Red + "❌ Inventaire plein !" + Reset)
		return
	}

	c.Gold -= r.Cost
	for mat, qty := range r.Materials {
		for i := 0; i < qty; i++ {
			c.RemoveInventory(mat)
		}
	}
	c.AddInventory(r.Name)

	PlayHammerSound()

	fmt.Println()
	fmt.Println(Green + "╔════════════════════════════════════════╗" + Reset)
	fmt.Println(Green + "║       ⚒  FABRICATION RÉUSSIE ! ⚒	     ║" + Reset)
	fmt.Println(Green + "╚════════════════════════════════════════╝" + Reset)
	fmt.Printf(Green+"✓ Vous fabriquez : %s\n"+Reset, r.Name)
	fmt.Printf(Yellow+"  -%d pièces (reste : %d)\n"+Reset, r.Cost, c.Gold)
}


func (c *Character) canCraft(r Recipe) bool {
	if c.Gold < r.Cost {
		return false
	}
	for mat, qty := range r.Materials {
		if c.CountItem(mat) < qty {
			return false
		}
	}
	if len(c.Inventory) >= c.MaxInventory {
		return false
	}
	return true
}
