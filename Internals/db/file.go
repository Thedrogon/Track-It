// To IMPLEMENT FILE BASE DB
package db

import (
	"fmt"
	"log"

	readers "github.com/Thedrogon/GophersCSV/readers"
	writers "github.com/Thedrogon/GophersCSV/writers"
)

func Error_check() {

}

func take_input() (string , string , []string){
	var problem_id int;
	fmt.Scanf("enter Problem ID : %v",problem_id )

	var title string;
	fmt.Scanf("Enter Title : %v",title)

	var tags []string;
	fmt.Scanf("Enter tags : %v",tags) //check for error

	return string(problem_id),title,tags
	
}

func Init_file_db() {

	var num_records int //taking number of records to be written or read

	var operation string
	fmt.Println("The operations are : 1--> Read , 2--> Write")
	fmt.Scanf("Enter the required operation : %v",&operation)

	switch operation {
	case "Read":
		readers.Read_csv("./data/my_200.csv")
	case "Write":
		writers.Create_csv("./data/my_200.csv",[]string{"problem_ID","Title","Tags"}) //creating csv file

		fmt.Scanf("Enter no. of records to be taken: %v",num_records)
		var records [][]string;

		err := writers.Add_data("./data/my_200.csv",records)

		if err != nil{
			log.Fatalln("Error while Adding data.")
		}
	default :   
		log.Fatalln("Not a valid case")
	}

}
