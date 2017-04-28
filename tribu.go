package main

import (
    "fmt"
    "strconv"
    "math/rand"
)

type People struct {
    Sex bool //true if female, false if male
    Name string
    Living bool
    Birthdate int
    Deathdate int
    Attractivity int //0 if really ugly, 10 if really beautiful
    Mother *People
    Children []*People
}

type Event struct {
    StartYear int
    EndYear int
    Influence float64
}

func eventInfluence(current_year int, events []Event) float64 {
    influence := 0.0
    nbInfluence := 0

    for _, event := range events {
        if current_year >= event.StartYear && current_year <= event.EndYear {
            influence += event.Influence
            nbInfluence++
        }
    }

    return influence / float64(nbInfluence)
}

func boundProbability(p float64) float64 {
    if p < 0 {
        return 0
    } else if p > 1 {
        return 1
    } else {
        return p
    }
}

func dies(current_year int, person People) bool {
    //params
    life_expectancy := 45
    event_factor := 0.1
    death_factor := 0.5

    death_events := []Event{
        Event{StartYear:720, EndYear:740, Influence:0.3},
        Event{StartYear:858, EndYear:859, Influence:0.2},
        Event{StartYear:885, EndYear:885, Influence:0.4},
        Event{StartYear:901, EndYear:901, Influence:0.2},
        Event{StartYear:904, EndYear:904, Influence:0.4},
        Event{StartYear:908, EndYear:967, Influence:0.6},
        Event{StartYear:967, EndYear:983, Influence:-0.1},
        Event{StartYear:997, EndYear:1010, Influence:0.2},
        Event{StartYear:1010, EndYear:1017, Influence:0.6},
    }

    //algo
    age := current_year - person.Birthdate

    var ageFactor float64
    if age <= 3 {
        ageFactor = 0.4
    } else if age <= 10 {
        ageFactor = 0.2
    } else if age <= 15 {
        ageFactor = 0.05
    } else if age <= life_expectancy {
        ageFactor = 0.15
    } else {
        ageFactor = 0.20 + float64(age-life_expectancy)*0.05
    }

    mortality_probability := ageFactor + event_factor * eventInfluence(current_year, death_events);

    return rand.Float64() < death_factor * boundProbability(mortality_probability)
}

func impregnates(current_year int, person People) bool {
    //params
    natural_fertility := 1.0
    event_factor := 0.4
    uglyness_factor := 0.8

    birth_events := []Event{
        Event{StartYear:858, EndYear:859, Influence:-0.75},
        Event{StartYear:885, EndYear:885, Influence:-0.6},
        Event{StartYear:901, EndYear:901, Influence:0.3},
        Event{StartYear:904, EndYear:904, Influence:-0.6},
        Event{StartYear:904, EndYear:906, Influence:-0.1},
        Event{StartYear:967, EndYear:983, Influence:-0.1},
        Event{StartYear:997, EndYear:1005, Influence:0.2},
        Event{StartYear:1005, EndYear:1010, Influence:-0.2},
        Event{StartYear:1010, EndYear:1017, Influence:-0.5},
    }

    //algo
    if !person.Sex{
        return false
    }

    age := current_year - person.Birthdate
    var ageFactor float64
    if age < 15 {
        ageFactor = 0
    } else if age < 25 {
        ageFactor = 1.5
    } else if age < 35 {
        ageFactor = 1
    } else if age < 40 {
        ageFactor = 0.5
    } else if age < 45 {
        ageFactor = 0.2
    } else {
        ageFactor = 0
    }

    birth_probability := natural_fertility * ageFactor + ((1.0/uglyness_factor) * float64(person.Attractivity-5)/10.0) + event_factor * eventInfluence(current_year, birth_events)

    return rand.Float64() < boundProbability(birth_probability)
}

func name(person People) string {
    var filiation_name string
    if person.Sex {
        filiation_name = " ferch " //daughter of
    } else {
        filiation_name = " ap " //son of
    }

    name := person.Name
    if person.Mother != nil {
        name = name + filiation_name + person.Mother.Name
    }

    return name
}

func info(person People) string {
    sexinfo := "female"
    if !person.Sex {
        sexinfo = "male"
    }

    birthinfo := "born on year " + strconv.Itoa(person.Birthdate)
    deathinfo := "dead on year " + strconv.Itoa(person.Deathdate)

    var attractivityinfo string
    switch person.Attractivity {
    case 0,1:
        attractivityinfo = "really ugly"
    case 2,3:
        attractivityinfo = "ugly"
    case 4:
        attractivityinfo = "negatively neutral"
    case 5:
        attractivityinfo = "neutral"
    case 6:
        attractivityinfo = "positively neutral"
    case 7,8:
        attractivityinfo = "beautiful"
    case 9,10:
        attractivityinfo = "really beautiful"
    }

    infostring := name(person) + " (" + sexinfo + "), " + birthinfo
    if !person.Living {
        infostring += ", " + deathinfo
    }

    infostring += " (" + attractivityinfo + ", " + strconv.Itoa(len(person.Children)) + " children)"

    return infostring
}

func birth(current_year int, mother *People) People {
    names := []string{"Aubrée", "Alan", "Alann", "Albin", "Allan", "Alain", "Alderic", "Amael", "Annick", "Annaïg", "Arwen", "Armel", "Armelle", "Arthur", "Audran", "Alistair", "Alon", "Arzel", "Abriel", "Aelwenn", "Aeryn", "Brendan", "Brayan", "Briac", "Brian", "Brice", "Bryan", "Bryce", "Brieuc", "Byron", "Cédric", "Corantin", "Corentin", "Corentine", "Delwyn", "Dilwen", "Douglas", "Deirdre", "Dyclan", "Duncan", "Donovan", "Declan", "Donald", "Edern", "Enora", "Erin", "Erwan", "Erwann", "Ewarn", "Emeline", "Ewen", "Elouan", "Evène", "Ewan", "Fiacre", "Gaël", "Ganaël", "Gauvain", "Gildas", "Gladys", "Glenn", "Goulven", "Goulwen", "Gwenael", "Gwenaelle", "Gwendoline", "Gwenola", "Gwladys", "Gurvan", "Guenièvre", "Guénolé", "Guerdiale", "Gwendal", "Gwidel", "Hervé", "Herlé", "Jennifer", "Judicaël", "Jaouen", "Julven", "Janig", "Kelvin", "Ken", "Kenelm", "Kenan", "Kenny", "Kerian", "Kevin", "Kilian", "Killian", "Killyan", "Kilyan", "Klervi", "Keith", "Kenya", "Kendall", "Logan", "Loan", "Loann", "Lohan", "Luan", "Lénaig", "Maclou", "Maelle", "Magloire", "Malo", "Malou", "Melwyn", "Malvina", "Metig", "Maë", "Maël", "Maëlig", "Morgane", "Morgan", "Maudan", "Maïwenn", "Nelly", "Nigel", "Nimué", "Nolan", "Nolann", "Nolwenn", "Nominoë", "Pierrick", "Renan", "Ronan", "Riwanon", "Riwann", "Rivoal", "Rivoual", "Rian", "Ryan", "Riwalenn", "Rozenn", "Roween", "Rowen", "Rowena", "Sleipnir", "Solen", "Siobhan", "Seylan", "Sterenn", "Sloan", "Soizik", "Tangi", "Tanguy", "Tristan", "Tugdual", "Tual", "Thurien", "Yann", "Yannick", "Ylan", "Almedea", "Mamaeth", "Serwyl", "Llywelyn", "Ogg", "Cunedda", "Mair", "Brithyll", "Senana", "Maches", "Glawdys"}

    sex := rand.Float64() < 0.5
    name := names[rand.Intn(len(names))]
    motherAttractivity := boundProbability(rand.NormFloat64() * 0.5 + 0.5)
    if mother != nil {
        motherAttractivity = float64(mother.Attractivity)/10.0
    }
    attractivity := int(10 * (rand.NormFloat64() * 0.2 + motherAttractivity))
    if attractivity < 0 {
        attractivity = 0
    } else if attractivity > 10 {
        attractivity = 10
    }

    return People{Sex:sex, Name:name, Living:true, Birthdate:current_year, Attractivity:attractivity, Mother:mother}
}

func main() {
    rand.Seed(42)

    creation_year := 720
    current_year := 1017

    tribe_attractivity := 0.5

    //tribe input
    tribe_creator := People{Sex:true, Name:"Almedea", Living:true, Birthdate:687, Attractivity:5}
    tribe := []People{tribe_creator}

    //Almedea's children
    for year := 702; year < 720; year++ {
        if year % 2 == 0 || year == 701 {
            child := birth(year, &tribe_creator)
            tribe_creator.Children = append(tribe_creator.Children, &child)
            tribe = append(tribe, child)
        }
    }

    //persons who joined the tribe
    for i := 0; i < 10; i++ {
        birth_year := creation_year - (rand.Intn(10) + 20)
        person := birth(birth_year, nil)
        tribe = append(tribe, person)
    }

    fmt.Println("Tribe in 720")
    for _, person := range tribe {
        fmt.Println(info(person))
    }

    for year := creation_year; year < current_year; year++ {
        fmt.Println("year ", year, ":")

        dyingPersons := []int{}
        for i, person := range tribe {
            if dies(year, person) {
                person.Living = false
                person.Deathdate = year
                fmt.Println(name(person), "dies at age", person.Deathdate - person.Birthdate)
                dyingPersons = append(dyingPersons, i)
            } else if impregnates(year, person) {
                child := birth(year, &person)
                fmt.Println("new birth:", info(child))
                person.Children = append(person.Children, &child)
                tribe = append(tribe, child)
            }
        }
        for _, i := range dyingPersons {
            if i == 0 {
                tribe = tribe[1:]
            } else if i < len(tribe) {
                tribe = append(tribe[:i], tribe[i+1:]...)
            }
        }

        if rand.Float64() < tribe_attractivity {
            //joining
            number_of_people_joining_tribe := int(rand.NormFloat64() * 10 + 1)
            if number_of_people_joining_tribe < 1 {
                number_of_people_joining_tribe = 1
            }

            for i := 0; i < number_of_people_joining_tribe; i++ {
                person := birth(year - (rand.Intn(10) + 20), nil)
                fmt.Println("joining tribe:", info(person))
                tribe = append(tribe, person)
            }
        }
    }

    fmt.Println("Tribe of", len(tribe), "members")
}
