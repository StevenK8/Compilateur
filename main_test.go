package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/StevenK8/Compilateur/gencode"
)

//L'exécution doit se faire sur linux ou WSL (./msm non reconnu sur windows)
func execute(fileName string) (string, error) {
	var err error
	out, err := exec.Command("./msm", fileName).Output()
	return string(out[:len(out)-1]), err
}

func TestMult(t *testing.T) {
	compileRuntime()

	data := []byte(`
		int main(){
			int a;
			a = 5;
			int b;
			b = 3;

			debug a*b;
			return 0;
		}`)

	expectedval := 15
	compile(data)

	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})

	writeOutput("testmult")
	result, err := execute("testmult.out")

	val, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if val != expectedval || err != nil {
		t.Errorf("Multiplication incorrecte, reçu %d au lieu de %d.", val, expectedval)
	}
}

func TestAdd(t *testing.T) {
	compileRuntime()

	data := []byte(`
		int main(){
			int a;
			a = 5;
			int b;
			b = 3;

			debug a+b;
			return 0;
		}`)

	expectedval := 8
	compile(data)

	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})

	writeOutput("testmult")
	result, err := execute("testmult.out")

	val, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if val != expectedval || err != nil {
		t.Errorf("Addition incorrecte, reçu %d au lieu de %d.", val, expectedval)
	}
}

func TestSub(t *testing.T) {
	compileRuntime()

	data := []byte(`
		int main(){
			int a;
			a = 5;
			int b;
			b = 6;

			debug a-b;
			return 0;
		}`)

	expectedval := -1
	compile(data)

	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})

	writeOutput("testmult")
	result, err := execute("testmult.out")

	val, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if val != expectedval || err != nil {
		t.Errorf("Soustraction incorrecte, reçu %d au lieu de %d.", val, expectedval)
	}
}

func TestDiv(t *testing.T) {
	compileRuntime()

	data := []byte(`
		int main(){
			int a;
			a = 30;
			int b;
			b = 5;

			debug a/b;
			return 0;
		}`)

	expectedval := 6
	compile(data)

	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})

	writeOutput("testmult")
	result, err := execute("testmult.out")

	val, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if val != expectedval || err != nil {
		t.Errorf("Division incorrecte, reçu %d au lieu de %d.", val, expectedval)
	}
}

func TestMod(t *testing.T) {
	compileRuntime()

	data := []byte(`
		int main(){
			int a;
			a = 10;
			int b;
			b = 3;

			debug a%b;
			return 0;
		}`)

	expectedval := 1
	compile(data)

	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})

	writeOutput("testmult")
	result, err := execute("testmult.out")

	val, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if val != expectedval || err != nil {
		t.Errorf("Modulo incorrect, reçu %d au lieu de %d.", val, expectedval)
	}
}

func TestBoucleFunc(t *testing.T) {
	compileRuntime()

	data := []byte(`
		int boucleFunction(int a){
			while (a<5){
				a = a+1;
				debug a;
			}
			return a;
		}

		int main(){
			int a;
			a = 0;

			return boucleFunction(a);
		}
		`)

	expectedval := `1
2
3
4
5`

	compile(data)

	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})

	writeOutput("testmult")
	result, err := execute("testmult.out")

	if result != expectedval || err != nil {
		t.Errorf("BoucleFunc incorrecte, reçu %s au lieu de %s.", result, expectedval)
	}
}
