package main

import (
	"log"
	"os/exec"
	"runtime"
	"testing"

	"github.com/StevenK8/Compilateur/gencode"
)

//L'exécution doit se faire sur linux ou WSL (./msm non reconnu sur windows)
func execute(fileName string) (string, error) {
	/// Check if machine is linux or windows
	if runtime.GOOS == "windows" {
		var err error
		out, err := exec.Command("powershell", "./MSM.exe", fileName).Output()
		return string(out), err
	} else {
		print("Linux execute :")
		var err error
		out, err := exec.Command("./msm", fileName).Output()
		return string(out), err
	}
}

func createFileAndExecute(data []byte, file string) string {
	compileRuntime()
	compile(data, false)
	gencode.AddToList([]string{".start", "prep main", "call 0", "halt"})
	writeOutput("test/"+file, "")
	gencode.Clear()
	res, err := execute("test/" + file + ".out")
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func assertEquals(t *testing.T, s1 string, s2 string) {
	if s1 != s2 {
		t.Errorf("Failed '%q', '%q' isn't equal to '%q'", t.Name(), s1, s2)
	}

}

const EOF_CONSTANT string = "\r\n"

func TestMult(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 5;
			int b;
			b = 3;

			print(a*b);
			return 0;
		}`)

	expectedval := "15"

	assertEquals(t, createFileAndExecute(data, "mult"), expectedval)
}

func TestAdd(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 5;
			int b;
			b = 3;

			print(a+b);
			return 0;
		}`)

	expectedval := "8"

	assertEquals(t, createFileAndExecute(data, "add"), expectedval)
}

func TestSub(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 5;
			int b;
			b = 6;

			print(a-b);
			return 0;
		}`)

	expectedval := "-1"

	assertEquals(t, createFileAndExecute(data, "sub"), expectedval)
}

func TestDiv(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 30;
			int b;
			b = 5;

			print(a/b);
			return 0;
		}`)

	expectedval := "6"

	assertEquals(t, createFileAndExecute(data, "div"), expectedval)
}

func TestMod(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 10;
			int b;
			b = 3;

			print(a%b);
			return 0;
		}`)

	expectedval := "1"

	assertEquals(t, createFileAndExecute(data, "mod"), expectedval)
}

func TestBoucleFunc(t *testing.T) {

	data := []byte(`
		int boucleFunction(int a){
			while (a<5){
				a = a+1;
				print(a);
			}
			return a;
		}

		int main(){
			int a;
			a = 0;

			return boucleFunction(a);
		}
		`)

	expectedval := "12345"

	assertEquals(t, createFileAndExecute(data, "bouclefunc"), expectedval)
}

func TestPtr(t *testing.T) {

	data := []byte(`
		int increment(int a){
			*(a) = *(a)+20;
			return 0;
		}

		int main(){
			int t;
			t = malloc(1);
			*(t+0) = 5;
			increment(t);

			print(*(t));
			return 0;
		}
		`)

	expectedval := "25"

	assertEquals(t, createFileAndExecute(data, "ptr"), expectedval)
}

func TestTableau(t *testing.T) {

	data := []byte(`
		int main(){
			int t;
			int lenTab;
			
			lenTab = 5;
			t = malloc(5);

			int i;
			for(i=0;i<lenTab;i=i+1){
				*(t+i) = i;
			}

			for(i=0;i<lenTab;i=i+1){
				print(*(t+i));
			}

			return 0;
		}
		`)

	expectedval := "01234"

	assertEquals(t, createFileAndExecute(data, "tab"), expectedval)
}

func TestFor(t *testing.T) {

	data := []byte(`
		int main(){
			int i;
			for(i=1;i<10;i=i+1){
				print(i);
			}

			return 0;
		}
		`)

	expectedval := "123456789"

	assertEquals(t, createFileAndExecute(data, "for"), expectedval)
}

func TestPrint(t *testing.T) {

	data := []byte(`
		int main(){
			int n;

			n = 5;

			print(n);

			return 0;
		}
		`)

	expectedval := "5"

	assertEquals(t, createFileAndExecute(data, "print"), expectedval)
}
