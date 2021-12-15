package main

import "fmt"

func writeEmployeeInfo(employeeId string) error {
	file := openFile("employees.txt")

	// fetches info about the employee from external api
	info, err := getInfo(employeeId)
	// if request was unsuccessful
	if err != nil {
		return err
	}

	// if request was successful, info is written to file
	_, err = writeFile(file, info)
	// if write operation was unsuccessful
	if err != nil {
		closeFile(file)
		return err
	}

	// if write operation was successful, upload file to server
	err = uploadFile(file)

	closeFile(file)
	return err
}

func main() {
  cleanFile("employees.txt")
	for _, employeeId := range employees {
		if err := writeEmployeeInfo(employeeId); err != nil {
			fmt.Println(employeeId, err)
		}
	}
}
