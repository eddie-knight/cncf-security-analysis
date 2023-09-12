package main

func main() {
	projects := Projects{}
	projects.Init()
	projects.WriteOverviewData()
	projects.WriteSecurityData()
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
	// data := projects.GetProjectData("akri")
	// jsonData, _ := json.MarshalIndent(data, "", "  ")
	// ioutil.WriteFile("akri.json", jsonData, 0644)
}
