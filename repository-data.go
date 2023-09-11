package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Projects struct {
	metadata []ProjectMetadata
	projects []ProjectData
}

type ProjectData struct {
	ProjectName            string
	CheckResultEndpoints   []string
	HistoricalCheckResults []CheckResults
}

type ProjectMetadata struct {
	Foundation      string
	Project         string
	RepositoryURL   string
	CheckSets       []string
	Adopters        bool
	AnnualReview    bool
	Changelog       bool
	CodeOfConduct   bool
	Contributing    bool
	Governance      bool
	Maintainers     bool
	Readme          bool
	Roadmap         bool
	SummaryTable    bool
	Website         bool
	LicenseApproved bool
	LicenseScanning bool
	LicenseSPDXID   string
	Analytics       string
	// everything after this is bool
	ArtifactHubBadge          bool
	CLA                       bool
	CommunityMeeting          bool
	DCO                       bool
	GitHubDiscussions         bool
	OpenSSFBestPracticesBadge bool
	OpenSSFScorecardBadge     bool
	RecentRelease             bool
	SlackPresence             bool
	BinaryArtifacts           bool
	CodeReview                bool
	DangerousWorkflow         bool
	DependencyUpdateTool      bool
	Maintained                bool
	SBOM                      bool
	SecurityPolicy            bool
	SignedReleases            bool
	TokenPermissions          bool
	TrademarkDisclaimer       bool
}

// GetRepositories returns a list of repositories
func (p *Projects) GetRepoURLs() []string {
	var names []string
	for _, repo := range p.metadata {
		names = append(names, repo.Project)
	}
	return names
}

// GetProjects returns a list of unique project names from the repositories— omitting duplicates
func (p *Projects) GetProjects() []string {
	var projects []string
	for _, repo := range p.metadata {
		if !contains(projects, repo.Project) {
			projects = append(projects, repo.Project)
		}
	}
	return projects
}

// GetProjectsWithCheckSet returns a list of projects matching a given check set
func (p *Projects) GetProjectsWithCheckSet(checkSet string) []ProjectMetadata {
	var projects []ProjectMetadata
	for _, repo := range p.metadata {
		if contains(repo.CheckSets, checkSet) {
			projects = append(projects, repo)
		}
	}
	return projects
}

// GetProjectsByFoundation returns a list of projects matching a given foundation
func (p *Projects) GetProjectsByFoundation(foundation string) []ProjectMetadata {
	var projects []ProjectMetadata
	for _, repo := range p.metadata {
		if repo.Foundation == foundation {
			projects = append(projects, repo)
		}
	}
	return projects
}

// GetProjectsByName returns a list of projects matching a given project name
func (p *Projects) GetProjectsByName(project string) []ProjectMetadata {
	var projects []ProjectMetadata
	for _, repo := range p.metadata {
		if repo.Project == project {
			projects = append(projects, repo)
		}
	}
	return projects
}

// GetProjectsByFoundationAndCheckSet returns a list of projects matching a given foundation and check set
func (p *Projects) GetProjectsByFoundationAndCheckSet(foundation string, checkSet string) []ProjectMetadata {
	var projects []ProjectMetadata
	for _, repo := range p.metadata {
		if repo.Foundation == foundation && contains(repo.CheckSets, checkSet) {
			projects = append(projects, repo)
		}
	}
	return projects
}

// GetCheckSets returns a list of unique check sets from the repositories— omitting duplicates
func (p *Projects) GetCheckSets() []string {
	var checkSets []string
	for _, repo := range p.metadata {
		for _, checkSet := range repo.CheckSets {
			if !contains(checkSets, checkSet) {
				checkSets = append(checkSets, checkSet)
			}
		}
	}
	return checkSets
}

// CountAPIEndpoints will count all API endpoints for all projects. Should total over 2600.
func (p *Projects) CountAPIEndpoints() int {
	var count int
	for _, project := range p.projects {
		count += len(p.GetAPIEndpointsByProject(project.ProjectName))
	}
	return count
}

// GetAPIEndpointsByProject returns a list of API Endpoints matching a given a project name
func (p *Projects) GetAPIEndpointsByProject(project string) (endpoints []string) {
	// Given a date and project name, return a list of API Endpoints in the following format: https://clomonitor.io/api/projects/<FOUNDATION>/<PROJECT>/snapshots/<DATE>
	for _, repo := range p.metadata {
		if repo.Project == project {
			for _, date := range getLastDaysOfMonths() {
				endpoints = append(endpoints, fmt.Sprintf("https://clomonitor.io/api/projects/"+repo.Foundation+"/"+repo.Project+"/snapshots/"+date))
			}
		}
	}
	return
}

// SetProjectDataValues will set the values of the Projects object
func (p *Projects) SetProjectDataValues(projectName string) {
	p.projects = append(p.projects, ProjectData{
		ProjectName:            projectName,
		CheckResultEndpoints:   p.GetAPIEndpointsByProject(projectName),
		HistoricalCheckResults: p.GetHistoricalCheckResults(projectName),
	})
}

// GetHistoricalCheckResults returns a list of CheckResults objects matching the given project name
func (p *Projects) GetHistoricalCheckResults(projectName string) (checkResults []CheckResults) {
	for _, url := range p.GetAPIEndpointsByProject(projectName) {
		res := getHttpResponse(url)
		if res.Body != nil {
			defer res.Body.Close()
		}
		var data CheckResults
		json.NewDecoder(res.Body).Decode(&data)
		checkResults = append(checkResults, data)
	}
	return
}

// GetProjectData returns a ProjectData object matching a given project name
func (p *Projects) GetProjectData(name string) ProjectData {
	for _, project := range p.projects {
		if project.ProjectName == name {
			return project
		}
	}
	return ProjectData{}
}

// GetProjectMetadata returns a ProjectMetadata object matching a given project name
func (p *Projects) SetProjectMetadataValues() {
	url := "https://clomonitor.io/data/repositories.csv"
	res := getHttpResponse(url)
	if res.Body != nil {
		defer res.Body.Close()
	}
	reader := csv.NewReader(res.Body)
	records, _ := reader.ReadAll()

	// iterate through the records
	for i := 1; i < len(records); i++ {
		// record is a list of string values compatible with Repository struct values if converted
		record := records[i]

		// create new Repository object with record values
		projectMetadata := ProjectMetadata{
			Foundation:                record[0],
			Project:                   record[1],
			RepositoryURL:             record[2],
			CheckSets:                 convertStringToList(record[3]),
			Adopters:                  convertStringToBool(record[4]),
			AnnualReview:              convertStringToBool(record[5]),
			Changelog:                 convertStringToBool(record[6]),
			CodeOfConduct:             convertStringToBool(record[7]),
			Contributing:              convertStringToBool(record[8]),
			Governance:                convertStringToBool(record[9]),
			Maintainers:               convertStringToBool(record[10]),
			Readme:                    convertStringToBool(record[11]),
			Roadmap:                   convertStringToBool(record[12]),
			SummaryTable:              convertStringToBool(record[13]),
			Website:                   convertStringToBool(record[14]),
			LicenseApproved:           convertStringToBool(record[15]),
			LicenseScanning:           convertStringToBool(record[16]),
			LicenseSPDXID:             record[17],
			Analytics:                 record[18],
			ArtifactHubBadge:          convertStringToBool(record[19]),
			CLA:                       convertStringToBool(record[20]),
			CommunityMeeting:          convertStringToBool(record[21]),
			DCO:                       convertStringToBool(record[22]),
			GitHubDiscussions:         convertStringToBool(record[23]),
			OpenSSFBestPracticesBadge: convertStringToBool(record[24]),
			OpenSSFScorecardBadge:     convertStringToBool(record[25]),
			RecentRelease:             convertStringToBool(record[26]),
			SlackPresence:             convertStringToBool(record[27]),
			BinaryArtifacts:           convertStringToBool(record[28]),
			CodeReview:                convertStringToBool(record[29]),
			DangerousWorkflow:         convertStringToBool(record[30]),
			DependencyUpdateTool:      convertStringToBool(record[31]),
			Maintained:                convertStringToBool(record[32]),
			SBOM:                      convertStringToBool(record[33]),
			SecurityPolicy:            convertStringToBool(record[34]),
			SignedReleases:            convertStringToBool(record[35]),
			TokenPermissions:          convertStringToBool(record[36]),
			TrademarkDisclaimer:       convertStringToBool(record[37]),
		}
		// add record to list of records
		p.metadata = append(p.metadata, projectMetadata)
	}
}

// Init will initialize the Projects object
func (p *Projects) Init() {
	p.SetProjectMetadataValues()
	totalCount := len(p.metadata)
	for progress, project := range p.metadata {
		if progress > 3 {
			break
		}
		fmt.Printf("Progress: %d/%d\n", progress+1, totalCount)
		p.SetProjectDataValues(project.Project)
	}
	p.WriteAllProjectData()
}

// WriteAllProjectData will write all project data to a file in CSV format
func (p *Projects) WriteAllProjectData() {
	csv := "ProjectName, UpdatedAt, Global, BestPractices, Documentation, Legal, License, Security\n"
	for _, project := range p.projects {
		csv += project.CLOMonitorScoreCSV()
	}
	// fmt.Println(csv)
	writeToFile("all-projects.csv", csv)
}

func (p *ProjectData) CLOMonitorScoreCSV() (csv string) {
	for _, results := range p.HistoricalCheckResults {
		csv += results.SummaryScoreCSV()
	}
	return
}

func writeToFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(data)
}
