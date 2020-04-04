package model

type Specialist struct {
	ID        int
	FirstName string
	LastName  string
	Ability   []SpecialistAbility
}

func (specialist *Specialist) AddAbility(abilityType Procedure) {

	if specialist.HasAbility(abilityType) {
		return
	}

	ability := SpecialistAbility{AbilityType: abilityType}
	specialist.Ability = append(specialist.Ability, ability)
}

func (specialist Specialist) HasAbility(abilityType Procedure) bool {
	for _, value := range specialist.Ability {
		if value.AbilityType == abilityType {
			return true
		}
	}

	return false
}

type SpecialistAbility struct {
	ID           int
	AbilityType  Procedure
	SpecialistID int
}
