package main 
import "testing"
func TestCleanUp(t *testing.T){
	t.Cleanup(func() {
		println("1 cleanup.")
	})
	t.Cleanup(func() {
		println("2 cleanup.")
	})
	t.Cleanup(func() {
		println("3 cleanup.")
	})
	t.Log("body")
}

func TestCleanUp2(t *testing.T){
	func(){
		t.Cleanup(func() {
		println("1 cleanup.")
		})
	}()
	func(){
		defer println("defer")
	}()
	t.Log("body")
}

