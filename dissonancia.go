package main
import(
	"fmt"
	"math"
	"strconv"
)
func dis(a string, b string) (float64, float64, float64){//formato da nota X(#)n
	var f_A0 float64=27.5;
	var r12_2 float64=math.Pow(2.0, 1.0/12.0);
	//fmt.Printf("R12_2:%f\n", r12_2);
	var fa, fb float64=1, 1;
	var da, db int=0, 0;
	if len(a)==3{da++;}
	if len(b)==3{db++;}
	var n int;
	if string(a[len(a)-1])!="0" && string(a[len(a)-1])!="1"{
		n, _=strconv.Atoi(string(a[len(a)-1]));
		da+=(n-1)*12;
	}
	if string(b[len(b)-1])!="0" && string(b[len(b)-1])!="1"{
		n, _=strconv.Atoi(string(b[len(b)-1]));
		db+=(n-1)*12;
	}
	//fmt.Printf("DA:%d\nDB:%d\n", da, db);
	switch string(a[0]){
	case "B":
		da+=2;
	case "C":
		da+=3;
	case "D":
		da+=5;
	case "E":
		da+=7;
	case "F":
		da+=8;
	case "G":
		da+=10;
	}
	switch string(b[0]){
	case "B":
		db+=2;
	case "C":
		db+=3;
	case "D":
		db+=5;
	case "E":
		db+=7;
	case "F":
		db+=8;
	case "G":
		db+=10;
	}
	//fmt.Printf("DA:%d\nDB:%d\n", da, db);
	fa=f_A0*math.Pow(r12_2, float64(da));
	fb=f_A0*math.Pow(r12_2, float64(db));
	return fa, fb, fa/fb;
}
func rac(n float64) (int, int){
	var np int=0;
	var dp int=1;

	var nm int=1;
	var dm int=1;
	for math.Abs(1-n/(float64(np)/float64(dp)))>0.1{//aproximação por porcentagem(se diferença<=10%, vá lá)
		if math.Abs(n-float64(np)/float64(dp))>math.Abs(n-float64(nm)/float64(dm)){
			np+=nm;
			dp+=dm;
		}else{
			nm+=np;
			dm+=dp;
		}
	}
	return np, dp;
}
func main(){
	var n1, n2 string;
	fmt.Println("Escreva duas notas no formato X(#)n [Não há bemois]:");
	fmt.Scanln(&n1);
	fmt.Scanln(&n2);
	fa, fb, r:=dis(n1, n2);
	fmt.Printf("Frequência 1: %f\nFrequência 2: %f\nRazão entre as frequências: %f\n", fa, fb, r);
	n, d:=rac(r);
	fmt.Printf("%d : %d\n", n, d);
}