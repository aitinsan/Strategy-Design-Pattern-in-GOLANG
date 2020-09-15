package main

import "fmt"

type King struct {
	Character
	Name string
	Race string
}
func (king *King)setKingWeapon(behavior WeaponBehavior){
	king.weapon = behavior
}
type Quenn struct {
	Character
	Name string
	Race string
}
func (quenn *Quenn)setQueenWeapon(behavior WeaponBehavior){
	quenn.weapon = behavior
}
type Knight struct {
	Character
	Name string
	Race string
}
func (knight *Knight)setKnightWeapon(behavior WeaponBehavior){
	knight.weapon = behavior
}
type Troll struct {
	Character
	Name string
	Race string
}
func (troll *Troll)setTrollWeapon(behavior WeaponBehavior){
	troll.weapon = behavior
}
func RunFantasyGame(){
	Hero := Hero()
	King1 := King{
		Hero,
		"Wraith King",
		"dead People",
	}
	fmt.Println("First character - "+King1.Name)
	King1.Damage()
	King1.setKingWeapon(SwordBehavior{})
	King1.Damage()

	Quenn1 := Quenn{
		Hero,
		"Queen of pain",
		"People or koza",
	}
	fmt.Println(" ")
	fmt.Println("Second character - "+Quenn1.Name)
	Quenn1.Damage()
	Quenn1.setQueenWeapon(KnifeBehavior{})
	Quenn1.Damage()
	 Knight1:= Knight{
		Hero,
		"Dragon Knight",
		"People and dragon",
	}
	fmt.Println(" ")
	fmt.Println("Third character - "+Knight1.Name)
	Knight1.Damage()
	Knight1.setKnightWeapon(SwordBehavior{})
	Knight1.Damage()
	Troll1:= Troll{
		Hero,
			"Troll Warlock",
			"Troll",
	}
	fmt.Println(" ")
	fmt.Println("Fourth character - "+Troll1.Name)
	Troll1.Damage()
	Troll1.setTrollWeapon(AxeBehavior{})
	Troll1.Damage()



}
