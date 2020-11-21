package main

import (
	"os/exec"
	"testing"

	"github.com/StevenK8/Compilateur/gencode"
)

//L'exécution doit se faire sur linux ou WSL (./msm non reconnu sur windows)
func execute(fileName string) (string, error) {
	var err error
	out, err := exec.Command("./msm", fileName).Output()
	return string(out[:]), err
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

	expectedval := "15"

	compile(data, false)

	for _, gen := range gencode.GetGenList() {
		println(gen)
	}

	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})
	writeOutput("test", "")
	result, err := execute("test.out")

	if len(result) > 1 {
		result = result[:len(result)-1]
	} else {
		t.Errorf("Pas de résultat")
	}

	if result != expectedval || err != nil {
		t.Errorf("Multiplication incorrecte, reçu %s au lieu de %s.", result, expectedval)
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

	expectedval := "8"

	compile(data, false)
	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})
	writeOutput("test", "")

	result, err := execute("test.out")
	if len(result) > 1 {
		result = result[:len(result)-1]
	} else {
		t.Errorf("Pas de résultat")
	}

	if result != expectedval || err != nil {
		t.Errorf("Addition incorrecte, reçu %s au lieu de %s.", result, expectedval)
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

	expectedval := "-1"

	compile(data, false)
	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})
	writeOutput("test", "")

	result, err := execute("test.out")
	if len(result) > 1 {
		result = result[:len(result)-1]
	} else {
		t.Errorf("Pas de résultat")
	}

	if result != expectedval || err != nil {
		t.Errorf("Soustraction incorrecte, reçu %s au lieu de %s.", result, expectedval)
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

	expectedval := "6"

	compile(data, false)
	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})
	writeOutput("test", "")

	result, err := execute("test.out")
	if len(result) > 1 {
		result = result[:len(result)-1]
	} else {
		t.Errorf("Pas de résultat")
	}

	if result != expectedval || err != nil {
		t.Errorf("Division incorrecte, reçu %s au lieu de %s.", result, expectedval)
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

	expectedval := "1"

	compile(data, false)
	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})
	writeOutput("test", "")

	result, err := execute("test.out")
	if len(result) > 1 {
		result = result[:len(result)-1]
	} else {
		t.Errorf("Pas de résultat")
	}

	if result != expectedval || err != nil {
		t.Errorf("Modulo incorrecte, reçu %s au lieu de %s.", result, expectedval)
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

	compile(data, false)
	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})
	writeOutput("test", "")

	result, err := execute("test.out")
	if len(result) > 1 {
		result = result[:len(result)-1]
	} else {
		t.Errorf("Pas de résultat")
	}

	if result != expectedval || err != nil {
		t.Errorf("BoucleFunc incorrecte, reçu %s au lieu de %s.", result, expectedval)
	}
}
