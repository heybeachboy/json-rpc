package json_rpc

/*type CalculationsService struct {}

func (c * CalculationsService)Sum(args ...int)(int) {

	 total := 0

	 for _, val := range args {
	 	total += val
	 }

	 return  total
}

func (c * CalculationsService)Multiply(args ...int)(int) {
	 total := 1

	 for _,val := range args {
	 	total *= val
	 }

	 return total
}

func (c * CalculationsService)Divide(a,b int)(int) {

	 return a / b
}

func CalculationsMinus(a,b int)(int) {

	return a - b
}

type Student struct {
	 No int
	 Name string
	 Age  int
}


 /**
  * create a student
  */

/*
func CreateAStudent(id, age int,name string)(Student) {
	return Student{id,name,age}

} */




/*func main() {

	 fy := reflect.TypeOf(CalculationsMinus)
	 fv := reflect.ValueOf(CalculationsMinus)

	 fmt.Println("reflect type is : ",fy,":",fv.Kind(),"reflect value is :",fv)
	 param := make([]reflect.Value,2)
	 param[0] = reflect.ValueOf(20)
	 param[1] = reflect.ValueOf(10)
	 fmt.Println("result 1 is :",fv.Call(param)[0])

	 param[0] = reflect.ValueOf(30)
	 param[1] = reflect.ValueOf(10)
	 fmt.Println("result 2 is :",fv.Call(param)[0])

	 param[0] = reflect.ValueOf(40)
	 param[1] = reflect.ValueOf(10)
	 fmt.Println("result 3 is :",fv.Call(param)[0])

	 sv := reflect.ValueOf(CreateAStudent)

	 argument := make([]reflect.Value,3)
	 argument[0] = reflect.ValueOf(1)
	 argument[1] = reflect.ValueOf(25)
	 argument[2] = reflect.ValueOf("MR.zhou.lele")
	 st :=sv.Call(argument)
	 fmt.Println("student information:",st[0].Interface())

	 calv := reflect.ValueOf(new(CalculationsService))
	// calt := reflect.TypeOf(&CalculationsService{})
     //fmt.Println("calculations of type :",calt)
	 calParam := make([]reflect.Value,4)
	 calParam[0] =  reflect.ValueOf(10)
	 calParam[1] =  reflect.ValueOf(10)
	 calParam[2] =  reflect.ValueOf(10)
	 calParam[3] =  reflect.ValueOf(10)
	 fmt.Println("method name :",calv.MethodByName("Sum").Call(calParam)[0].Interface())
	 fmt.Println("method name :",calv.MethodByName("Multiply").Call(calParam)[0].Interface())
}*/
