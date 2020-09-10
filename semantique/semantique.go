package semantique

import (
	parser "Compilateur/parser"
	"fmt"
	"log"
	"sync"
)

type Symbol struct {
	Identifiant string
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

func Sem(N parser.Noeud) parser.Noeud {
	switch N.TypeDeNoeud {
	default:
		for i, Fils := range N.Fils {
			N.Fils[i] = Sem(Fils)
		}
		break

	case parser.NoeudBlock:
		DebutBlock()
		for i, Fils := range N.Fils {
			N.Fils[i] = Sem(Fils)
		}
		FinBlock()
		break

	case parser.NoeudDec:
		_, err := Declarer(N.ValeurString)
		if err != nil {
			log.Fatal(" Erreur : Declarer")
			break
		}
		// S.Type = "variable"
		// S.Slot = NbSlot
		NbSlot++
		break

	case parser.NoeudRef:
		S, err := Acceder(N.ValeurString)
		if err != nil {
			log.Fatal(" Erreur : Acceder")
			break
		}
		if S.Type != "variable" {
			log.Fatal(" Erreur semantique : variable attendue, ", S.Type, " trouvé.")
			break
		} else {
			N.Slot = S.Slot
		}
		break

	}
	return N
}

func DebutBlock() {
	NouvelleHashMap := make(map[string]Symbol)
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
	if contains {
		log.Fatal(" Erreur " + ident + " déjà sur la pile")
	} else {
		S := Symbol{ident, "variable", NbSlot}
		top[ident] = S
		return S, nil
	}
	return Symbol{"", "", 0}, fmt.Errorf("Already exists")

}

func Acceder(ident string) (Symbol, error) {
	for pile.Size() > 0 {
		hash, err := pile.Front()
		if err != nil || len(hash) == 0 {
			log.Fatal(" Erreur :" + " Acceder -> Variable " + ident + " non initialisée")
		}
		_, contains := hash[ident]
		if contains {
			return hash[ident], nil
		}
	}
	return Symbol{"", "", 0}, fmt.Errorf("Not found")
}

func NouveauSymbole() Symbol {
	return Symbol{}
}
