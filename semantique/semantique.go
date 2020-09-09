package semantique

import (
	parser "Compilateur/parser"
	"fmt"
	"log"
	"sync"
)

type Symbol struct {
	Identifiant int
	Type        string
	Slot        int
}

type customQueue struct {
	stack []map[string]Symbol
	lock  sync.RWMutex
}

func (c *customQueue) Push(name map[string]Symbol) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customQueue) Pop() error {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len-1]
		return nil
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (map[string]Symbol, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[len-1], nil
	}
	var empty map[string]Symbol
	return empty, fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return len(c.stack)
}

func (c *customQueue) Empty() bool {
	return len(c.stack) == 0
}

var NbSlot int
var pile customQueue

// var pile []map[string]Symbol

func Sem(N parser.Noeud) {
	switch N.TypeDeNoeud {
	default:
		for _, Fils := range N.Fils {
			Sem(Fils)
		}
		break

	case parser.NoeudBlock:
		DebutBlock()
		for _, Fils := range N.Fils {
			Sem(Fils)
		}
		FinBlock()
		break

	case parser.NoeudDec:
		S, err := Declarer(N.ValeurString)
		if err != nil {
			log.Fatal(" Erreur : Declarer")
			break
		}
		S.Type = "Variable"
		S.Slot = NbSlot
		NbSlot++
		break

	case parser.NoeudRef:
		S, err := Acceder(N.ValeurString)
		if err != nil {
			log.Fatal(" Erreur : Acceder")
			break
		}
		if S.Type != "variable" {
			log.Fatal(" Erreur : variable attendue, ", S.Type, " trouvé.")
			break
		} else {
			N.Slot = S.Slot
		}
		break

	}
}

func DebutBlock() {
	var NouvelleHashMap map[string]Symbol
	pile.Push(NouvelleHashMap)
}

func FinBlock() {
	pile.Pop()
}

func Declarer(ident string) (Symbol, error) {
	top, err := pile.Front()
	if err != nil {
		log.Fatal(" Erreur top")
	}
	_, contains := top[ident]
	if !contains {
		log.Fatal(" Erreur " + ident + " pas sur la pile")
	} else {
		S := NouveauSymbole()
		top[ident] = S
		return S, nil
	}
	return Symbol{0, "", 0}, fmt.Errorf("Already exists")

}

func Acceder(ident string) (Symbol, error) {
	for pile.Size() > 0 {
		hash, _ := pile.Front()
		_, contains := hash[ident]
		if contains {
			return hash[ident], nil
		}
	}
	return Symbol{0, "", 0}, fmt.Errorf("Not found")
}

func NouveauSymbole() Symbol {
	return Symbol{0, "", 0}
}
