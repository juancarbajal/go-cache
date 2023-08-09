package main
import( 
	"github.com/juancarbajal/go-cache/pkg/cache"
)

func main(){
	c := TCacheFactory.Create('file');
	errAdd := c.Add("f1", "hola esta es una prueba de cache ")
	if errAdd == nil {
		fmt.Println("ADD - OK")
	} else {
		fmt.Println("ADD - Error")
	}
	val, errFind := c.Find("f1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)

}
