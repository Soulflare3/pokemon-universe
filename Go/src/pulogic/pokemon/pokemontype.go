package pokemon

var typeList map[int]*PokemonType = make(map[int]*PokemonType)

var (
	T_NORMAL = NewPokemonType(0)
	T_FIRE = NewPokemonType(1)
	T_WATER = NewPokemonType(2)
	T_ELECTRIC = NewPokemonType(3)
	T_GRASS = NewPokemonType(4)
	T_ICE = NewPokemonType(5)
	T_FIGHTING = NewPokemonType(6)
	T_POISON = NewPokemonType(7)
	T_GROUND = NewPokemonType(8)
	T_FLYING = NewPokemonType(9)
	T_PSYCHIC = NewPokemonType(10)
	T_BUG = NewPokemonType(11)
	T_ROCK = NewPokemonType(12)
	T_GHOST = NewPokemonType(13)
	T_DRAGON = NewPokemonType(14)
	T_DARK = NewPokemonType(15)
	T_STEEL = NewPokemonType(16)
	T_TYPELESS = NewPokemonType(17)
)

func GetTypesString() [18]string {
	return [18]string{	"Water",
			        	"Electric",
				        "Grass",
				        "Ice",
				        "Fighting",
				        "Poison",
				        "Ground",
				        "Flying",
				        "Psychic",
				        "Bug",
				        "Rock",
				        "Ghost",
				        "Dragon",
				        "Dark",
				        "Steel",
				        "Typeless" }
}

func GetTypesSpecial() [18]bool {
	return [18]bool{false,
			        true,
			        true,
			        true,
			        true,
			        true,
			        false,
			        false,
			        false,
			        false,
			        true,
			        false,
			        false,
			        false,
			        true,
			        true,
			        false,
			        false }
}

func GetTypesMultiplier() [18][18]float32 {
	return [18][18]float32{{ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.5, 0, 1, 1, 0.5, 1 },
							{ 1, 0.5, 0.5, 1, 2, 2, 1, 1, 1, 1, 1, 2, 0.5, 1, 0.5, 1, 2, 1 },
							{ 1, 2, 0.5, 1, 0.5, 1, 1, 1, 2, 1, 1, 1, 2, 1, 0.5, 1, 1, 1 },
							{ 1, 1, 2, 0.5, 0.5, 1, 1, 1, 0, 2, 1, 1, 1, 1, 0.5, 1, 1, 1 },
							{ 1, 0.5, 2, 1, 0.5, 1, 1, 0.5, 2, 0.5, 1, 0.5, 2, 1, 0.5, 1, 0.5, 1 },
							{ 1, 0.5, 0.5, 1, 2, 0.5, 1, 1, 2, 2, 1, 1, 1, 1, 2, 1, 0.5, 1 },
							{ 2, 1, 1, 1, 1, 2, 1, 0.5, 1, 0.5, 0.5, 0.5, 2, 0, 1, 2, 2, 1 },
							{ 1, 1, 1, 1, 2, 1, 1, 0.5, 0.5, 1, 1, 1, 0.5, 0.5, 1, 1, 0, 1 },
							{ 1, 2, 1, 2, 0.5, 1, 1, 2, 1, 0, 1, 0.5, 2, 1, 1, 1, 2, 1 },
							{ 1, 1, 1, 0.5, 2, 1, 2, 1, 1, 1, 1, 2, 0.5, 1, 1, 1, 0.5, 1 },
							{ 1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 0.5, 1, 1, 1, 1, 0, 0.5, 1 },
							{ 1, 0.5, 1, 1, 2, 1, 0.5, 0.5, 1, 0.5, 2, 1, 1, 0.5, 1, 2, 0.5, 1 },
							{ 1, 2, 1, 1, 1, 2, 0.5, 1, 0.5, 2, 1, 2, 1, 1, 1, 1, 0.5, 1 },
							{ 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 2, 1, 0.5, 0.5, 1 },
							{ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 0.5, 1 },
							{ 1, 1, 1, 1, 1, 1, 0.5, 1, 1, 1, 2, 1, 1, 2, 1, 0.5, 0.5, 1 },
							{ 1, 0.5, 0.5, 0.5, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 0.5, 1 },
					        { 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 } }
}

func GetTypes() map[int]*PokemonType {
	return typeList
}

type PokemonType struct {
	Type	int
}

func NewPokemonType(_i int) *PokemonType {
	pokemonType := &PokemonType{ Type: _i }
	
	typeList[_i] = pokemonType
	
	return pokemonType
}

func (p *PokemonType) IsSpecial() bool {
	return GetTypesSpecial()[p.Type]
}

func (p *PokemonType) GetMultiplier(_type *PokemonType) float32 {
	return GetTypesMultiplier()[p.Type][_type.Type]
}

func (p *PokemonType) Equals(_type *PokemonType) bool {
	return p.Type == _type.Type
}

func (p *PokemonType) String() string {
	return GetTypesString()[p.Type]
}