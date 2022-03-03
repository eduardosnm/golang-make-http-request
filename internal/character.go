package internal

type Character struct {
	Id        string
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
}

func (c *Character) ToArray() []string {
	return []string{c.Name, c.Height, c.Mass, c.HairColor, c.SkinColor, c.EyeColor, c.BirthYear, c.Gender}
}

type CharacterRepo interface {
	GetCharacter(id string) (*Character, error)
	AddCharacter(character *Character) error
}
