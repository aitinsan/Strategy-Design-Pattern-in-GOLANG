package main
import "fmt"
type WeaponBehavior interface{
	damage()
}
type HasWeaponBehavior struct{}

func (HasWeaponBehavior)damage(){
	fmt.Println("I can Damage you with weapon")
}
type AxeBehavior struct{}

func (AxeBehavior)damage(){
	fmt.Println("My weapon is axe")
}
type SwordBehavior struct{}

func (SwordBehavior)damage(){
	fmt.Println("My weapon is sword")
}
type KnifeBehavior struct{}

func (KnifeBehavior)damage(){
	fmt.Println("My weapon is Knife")
}
type BowAndArrowBehavior struct{}

func (BowAndArrowBehavior)damage(){
	fmt.Println("My weapon is Bow and Arrow")
}

