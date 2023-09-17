package main

import "fmt"

func main() {
	// New CSVContent to read for project-security-scores.csv and fix the CSV contents
	// csvContent := CSVContent{}
	// csvContent.ReadCSV("project-security-scores.csv")
	// csvContent.FixBooleanValues()

	// *** API Calls ***
	projects := Projects{}
	projects.Init()
	projects.WriteOverviewData()
	projects.WriteSecurityData()
	fmt.Printf("Completed %d requests\n", REQUEST_COUNTER)

	// // print all unique project names
	// for _, project := range repos.GetProjects() {
	// 	fmt.Println(project)
	// }

	// // print all projects with the "code-lite" check set
	// for _, project := range repos.GetProjectsWithCheckSet("code-lite") {
	// 	fmt.Println(project.RepositoryURL)
	// }

	// print all possible check sets
	// for _, checkSet := range repos.GetCheckSets() {
	// 	fmt.Println(checkSet)
	// }

	// test getting project data
	// data := projects.GetProjectData("cloudevents")
	// jsonData, _ := json.MarshalIndent(data, "", "  ")
	// ioutil.WriteFile("cloudevents.json", jsonData, 0644)
}
