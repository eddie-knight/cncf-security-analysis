package main

func main() {
	repos := Projects{}
	repos.Init()

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

	// print project data for argo
	// data := repos.GetProjectData("argo")
	// jsonData, _ := json.MarshalIndent(data, "", "  ")
	// fmt.Println(string(jsonData))
}
