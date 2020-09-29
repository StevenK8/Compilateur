package gencode

import (
	"fmt"
	"log"
	"sync"

	"github.com/StevenK8/Compilateur/parser"
)

type customQueue struct {
	stack [][2]string
	lock  sync.RWMutex
}

var pile customQueue
var lblIncrementName = 1
var ListOfAssembleurInstructions []string

func (c *customQueue) Push(name [2]string) {
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

func (c *customQueue) Front() ([2]string, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[len-1], nil
	}
	var empty [2]string
	return empty, fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return len(c.stack)
}

func (c *customQueue) Empty() bool {
	return len(c.stack) == 0
}

func Gencode(Node parser.Noeud) {
	switch Node.TypeDeNoeud {

	case parser.NoeudConst:
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "push "+fmt.Sprint(Node.ValeurEntiere))
		break

	case parser.NoeudAdd:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "add")
		break

	case parser.NoeudSub:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "sub")
		break

	case parser.NoeudSubUnaire:
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "push 0")
		Gencode(Node.Fils[0])
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "sub")
		break

	case parser.NoeudMult:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "mul")
		break

	case parser.NoeudDiv:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "div")
		break

	case parser.NoeudMod:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "mod")
		break

	case parser.NoeudDebug:
		Gencode(Node.Fils[0])
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "dbg")
		break

	case parser.NoeudBlock:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		break

	case parser.NoeudDrop:
		Gencode(Node.Fils[0])
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "drop")
		break

	case parser.NoeudTest:
		label1 := "if" + fmt.Sprint(lblIncrementName)
		lblIncrementName++
		label2 := "if" + fmt.Sprint(lblIncrementName)
		lblIncrementName++
		Gencode(Node.Fils[0])
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "jumpf "+label1)
		Gencode(Node.Fils[1])
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "jump "+label2, "."+label1)
		if len(Node.Fils) == 3 {
			Gencode(Node.Fils[2])
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "."+label2)
		break

	case parser.NoeudAffect:
		Gencode(Node.Fils[1])

		slot := Node.Fils[0].Slot
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "dup", "set "+fmt.Sprint(slot))
		break

	case parser.NoeudRef:
		slot := Node.Slot
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "get "+fmt.Sprint(slot))
		break

	case parser.NoeudLoop:
		l1 := "loop" + fmt.Sprint(lblIncrementName)
		lblIncrementName++
		l2 := "loop" + fmt.Sprint(lblIncrementName)
		lblIncrementName++
		pile.Push([2]string{l1, l2})
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "."+l1)
		Gencode(Node.Fils[0])
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "jump "+l1, "."+l2)
		pile.Pop()
		break

	case parser.NoeudBreak:
		l, err := pile.Front()
		if err != nil {
			log.Fatal(" Erreur : NoeudBreak")
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "jump "+l[1])
		break

	case parser.NoeudContinue:
		l, err := pile.Front()
		if err != nil {
			log.Fatal(" Erreur : NoeudContinue")
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "jump "+l[0])
		break

	case parser.NoeudFonction:
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "."+Node.ValeurString, "resn "+fmt.Sprint(Node.Slot-(len(Node.Fils)-1)))
		Gencode(Node.Fils[len(Node.Fils)-1])
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "push 0", "ret")
		break

	case parser.NoeudAppel:
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "prep "+Node.ValeurString)
		for _, n := range Node.Fils {
			Gencode(n)
		}
		ListOfAssembleurInstructions = append(ListOfAssembleurInstructions, "call "+fmt.Sprint(len(Node.Fils)))
		break

	}

}

func Gen(Node parser.Noeud) []string {
	Gencode(Node)
	return ListOfAssembleurInstructions
}
