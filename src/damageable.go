package main

// Damageable représente toute cible pouvant recevoir des dégâts.
// Implémenté par *Character ET *Monster.
type Damageable interface {
	GetName() string
	GetHP() int
	SetHP(hp int)
	GetMaxHP() int
	IsDead() bool
}

// ==================== Implémentation pour Character ====================

func (c *Character) GetName() string { return c.Name }
func (c *Character) GetHP() int      { return c.CurrentHP }
func (c *Character) SetHP(hp int)    { c.CurrentHP = hp }
func (c *Character) GetMaxHP() int   { return c.MaxHP }

// ⚠️ IsDead() est déjà définie ailleurs pour Character → on ne la redéfinit pas ici.

// ==================== Implémentation pour Monster ====================

func (m *Monster) GetName() string { return m.Name }
func (m *Monster) GetHP() int      { return m.CurrentHP }
func (m *Monster) SetHP(hp int)    { m.CurrentHP = hp }
func (m *Monster) GetMaxHP() int   { return m.MaxHP }
func (m *Monster) IsDead() bool    { return m.CurrentHP <= 0 }
