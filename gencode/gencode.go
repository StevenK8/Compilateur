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
var listOfAssembleurInstructions []string

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
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "push "+fmt.Sprint(Node.ValeurEntiere))

	case parser.NoeudAdd:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "add")

	case parser.NoeudSub:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "sub")

	case parser.NoeudSubUnaire:
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "push 0")
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "sub")

	case parser.NoeudMult:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "mul")

	case parser.NoeudDiv:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "div")

	case parser.NoeudMod:
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "mod")

	case parser.NoeudDebug:
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "send")

	case parser.NoeudReturn:
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "ret")

	case parser.NoeudBlock:
		for _, n := range Node.Fils {
			Gencode(n)
		}

	case parser.NoeudDrop:
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "drop")

	case parser.NoeudTest:
		label1 := "if" + fmt.Sprint(lblIncrementName)
		lblIncrementName++
		label2 := "if" + fmt.Sprint(lblIncrementName)
		lblIncrementName++
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "jumpf "+label1)
		Gencode(Node.Fils[1])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "jump "+label2, "."+label1)
		if len(Node.Fils) == 3 {
			Gencode(Node.Fils[2])
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "."+label2)

	case parser.NoeudAffect:
		Gencode(Node.Fils[1])

		slot := Node.Fils[0].Slot
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "dup", "set "+fmt.Sprint(slot))

	case parser.NoeudRef:
		slot := Node.Slot
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "get "+fmt.Sprint(slot))

	case parser.NoeudLoop:
		l1 := "loop" + fmt.Sprint(lblIncrementName)
		lblIncrementName++
		l2 := "loop" + fmt.Sprint(lblIncrementName)
		lblIncrementName++
		pile.Push([2]string{l1, l2})
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "."+l1)
		Gencode(Node.Fils[0])
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "jump "+l1, "."+l2)
		pile.Pop()

	case parser.NoeudBreak:
		l, err := pile.Front()
		if err != nil {
			log.Fatal(" Erreur : NoeudBreak")
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "jump "+l[1])

	case parser.NoeudContinue:
		l, err := pile.Front()
		if err != nil {
			log.Fatal(" Erreur : NoeudContinue")
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "jump "+l[0])

	case parser.NoeudFonction:
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "."+Node.ValeurString, "resn "+fmt.Sprint(Node.Slot-(len(Node.Fils)-1)))
		Gencode(Node.Fils[len(Node.Fils)-1])
		// listOfAssembleurInstructions = append(listOfAssembleurInstructions, "push 0", "ret")

	case parser.NoeudAppel:
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "prep "+Node.ValeurString)
		for _, n := range Node.Fils {
			Gencode(n)
		}
		listOfAssembleurInstructions = append(listOfAssembleurInstructions, "call "+fmt.Sprint(len(Node.Fils)))

	}

}

func Gen(Node parser.Noeud) []string {
	Gencode(Node)
	return listOfAssembleurInstructions
}
