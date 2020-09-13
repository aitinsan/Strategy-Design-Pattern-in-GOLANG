package main
type Character struct {
	weapon WeaponBehavior


}
func (character Character)Damage(){
	character.weapon.damage()
}
func Hero()Character{
	return Character{
		weapon:HasWeaponBehavior{},

	}
}

