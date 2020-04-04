package model

type Specialist struct {
	ID        int
	FirstName string
	LastName  string
	Ability   []SpecialistAbility
}

func (specialist *Specialist) AddAbility(abilityType Procedure) {
	ability := SpecialistAbility{AbilityType: abilityType}
	specialist.Ability = append(specialist.Ability, ability)
}

type SpecialistAbility struct {
	ID           int
	AbilityType  Procedure
	SpecialistID int
}
