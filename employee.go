/*
The program should implement the following menu:
          1. Add Employee(EmplD, EmpName, EmpSalary, EmpExperience)
          2.Fetch Employee with highest salary.
          3.Fetch 5 employees with the least experience.
          4.Update employee details of an employee with a given Id.
          5. Delete the employee with the given Id.
*/

package main
import "os"
import "fmt"
import "sort"

type Employee struct{  //Structure declaration for Employee
EmpID int
EmpName string
EmpSalary int
EmpExperience int
}

var employeeSlice = []Employee{}  //declare slice to store Employee structure variables


//function to add employee details
func addEmployee(){   
	
	var e Employee
	var flag bool

	fmt.Println("enter employee id: ")
	fmt.Scanln(&e.EmpID)

	flag = checkEmpID(e.EmpID) //use checkEmpID() function to check if Employee already exists
	if(flag){  
		fmt.Println("Id already exists")
		return
	}

	fmt.Println("enter employee name: ")
	fmt.Scanln(&e.EmpName)

	fmt.Println("enter employee salary: ")
	fmt.Scanln(&e.EmpSalary)

	fmt.Println("enter employee experience: ")
	fmt.Scanln(&e.EmpExperience)

	employeeSlice = append(employeeSlice, e)
}


//function to display employee with highest Salary
func highestSalary(){    

    sort.Slice(employeeSlice, func(p, q int) bool {  // sort slice by salary in descending order
      return employeeSlice[p].EmpSalary > employeeSlice[q].EmpSalary })
     
    fmt.Println(employeeSlice[0])  //print the highest salary employee
}


//function to display 5 employees with least experience
func leastExperience(){  
   
    sort.Slice(employeeSlice, func(p, q int) bool {
      return employeeSlice[p].EmpExperience < employeeSlice[q].EmpExperience })

    for i := 0; i<5 && i<len(employeeSlice); i++{
        if i == 5{    //limit the display to 5 employees
            break
        }
        fmt.Println("\n",i+1 , employeeSlice[i])
    }
}


//function to update employee details
func updateDetails(){  

	var id int
	var flag bool
	var existFlag bool
	var i int
	fmt.Println("enter the employee id you want to change: ")
	fmt.Scanln(&id)

	for i := range employeeSlice{  //find employee from slice
		if employeeSlice[i].EmpID == id{
			flag = true
			break;
		}
	}

	if(flag==true){ //change data if employee is found

		var choice int

		fmt.Println("1. Change employee id\n 2. Change employee name\n 3. Change employee salary\n 4. Change employee experience")
		fmt.Scanln(&choice)

		switch choice{
			case 1:
			var tempID int  //temporary variable to check if ID already exists before updating

			fmt.Println("enter new employee id: ")
			fmt.Scanln(&tempID)
			existFlag = checkEmpID(tempID) //use checkEmpID() function to check if Employee id already exists

			if(existFlag){
			fmt.Println("Id already exists")
			return // exit function if id already exist

			} else{
				employeeSlice[i].EmpID = tempID
			}

			case 2:
			fmt.Println("enter new name: ")
			fmt.Scanln(&employeeSlice[i].EmpName)

			case 3:
			fmt.Println("enter new salary: ")
			fmt.Scanln(&employeeSlice[i].EmpSalary)

			case 4:
			fmt.Println("enter new experience: ")
			fmt.Scanln(&employeeSlice[i].EmpExperience)
		}

		fmt.Println("Employee details changed!!")
		fmt.Println(employeeSlice[i])  //display new employee details
		
	} else{
		fmt.Println("employee doesnt exist")
	}
}


//function to delete an employee
func deleteEmployee(){ 
	var id int
	var flag bool
	var i int
	fmt.Println("enter the employee id you want to delete: ")
	fmt.Scanln(&id)

	for i := range employeeSlice{  //find employee from slice
		if employeeSlice[i].EmpID == id{
			flag = true
			break;
		}
	}
	if(flag==true){ //delete employee if found

		//Remove the element at index i from employeeSlice.
		employeeSlice[i] = employeeSlice[len(employeeSlice)-1] // Copy last element to index i.
		employeeSlice = employeeSlice[:len(employeeSlice)-1]   //Truncate slice.
		
		fmt.Println("Employee deleted!!")
		fmt.Println(employeeSlice)  //display remaining slice
		
	} else{
		fmt.Println("employee doesnt exist")
	}
}


//function to check if Employee ID already exist
func checkEmpID(id int) bool{
    var flag bool
    for i := range employeeSlice {
        if employeeSlice[i].EmpID == id {
            flag= true
        }
    }
    return flag
}


func main(){
	for{
		fmt.Println("-------------------\n1.Add Employee \n2.Fetch Employee with highest salary. \n3.Fetch five employees with the least experience.\n4.Update employee details of an employee with a given Id.\n5.Delete the employee with the given Id.\n6.Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice{
			case 1:
			addEmployee() //add employee details

			case 2:
			highestSalary() //display employee with highest Salary

			case 3:
			leastExperience() // display 5 employees with least experience

			case 4:
			updateDetails() //update employee details

			case 5:
			deleteEmployee() //delete an employee

			case 6:
			os.Exit(0)
		}
	}
}


/* OUTPUT
-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit

1
enter employee id: 
001
enter employee name: 
nevin
enter employee salary: 
15000
enter employee experience: 
1

-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit

1
enter employee id: 
002
enter employee name: 
sidarth
enter employee salary: 
20000
enter employee experience: 
2
-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit

1
enter employee id: 
003
enter employee name: 
akshay
enter employee salary: 
23000
enter employee experience: 
5

-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit

2
{3 akshay 23000 5}

-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit
3

 1 {1 nevin 15000 1}

 2 {2 sidarth 20000 2}

 3 {3 akshay 23000 5}
-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit

4
enter the employee id you want to change: 
001
 1. Change employee id
 2. Change employee name
 3. Change employee salary
 4. Change employee experience
2
enter new name: 
paul      
Employee details changed!!
{1 paul 15000 1}

-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit

5
enter the employee id you want to delete: 
003
Employee deleted!!
[{3 akshay 23000 5} {2 sidarth 20000 2}]

-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit

3

 1 {2 sidarth 20000 2}

 2 {3 akshay 23000 5}

-------------------
1.Add Employee 
2.Fetch Employee with highest salary. 
3.Fetch five employees with the least experience.
4.Update employee details of an employee with a given Id.
5.Delete the employee with the given Id.
6.Exit
6

*/