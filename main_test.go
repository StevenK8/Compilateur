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

			debug a*b;
			return 0;
		}`)

	expectedval := "15" + EOF_CONSTANT

	assertEquals(t, createFileAndExecute(data, "mult"), expectedval)
}

func TestAdd(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 5;
			int b;
			b = 3;

			debug a+b;
			return 0;
		}`)

	expectedval := "8" + EOF_CONSTANT

	assertEquals(t, createFileAndExecute(data, "add"), expectedval)
}

func TestSub(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 5;
			int b;
			b = 6;

			debug a-b;
			return 0;
		}`)

	expectedval := "-1" + EOF_CONSTANT

	assertEquals(t, createFileAndExecute(data, "sub"), expectedval)
}

func TestDiv(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 30;
			int b;
			b = 5;

			debug a/b;
			return 0;
		}`)

	expectedval := "6" + EOF_CONSTANT

	assertEquals(t, createFileAndExecute(data, "div"), expectedval)
}

func TestMod(t *testing.T) {

	data := []byte(`
		int main(){
			int a;
			a = 10;
			int b;
			b = 3;

			debug a%b;
			return 0;
		}`)

	expectedval := "1" + EOF_CONSTANT

	assertEquals(t, createFileAndExecute(data, "mod"), expectedval)
}

func TestBoucleFunc(t *testing.T) {

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

	expectedval := "1\r\n2\r\n3\r\n4\r\n5" + EOF_CONSTANT

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

			debug *(t);
			return 0;
		}
		`)

	expectedval := "25" + EOF_CONSTANT

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
				debug *(t+i);
			}

			return 0;
		}
		`)

	expectedval := "0\r\n1\r\n2\r\n3\r\n4" + EOF_CONSTANT

	assertEquals(t, createFileAndExecute(data, "tab"), expectedval)
}

func TestFor(t *testing.T) {

	data := []byte(`
		int main(){
			int i;
			for(i=1;i<10;i=i+1){
				debug i;
			}

			return 0;
		}
		`)

	expectedval := "1\r\n2\r\n3\r\n4\r\n5\r\n6\r\n7\r\n8\r\n9" + EOF_CONSTANT

	assertEquals(t, createFileAndExecute(data, "for"), expectedval)
}
