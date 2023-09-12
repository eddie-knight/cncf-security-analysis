package main

import (
	"fmt"
	"time"
)

// struct for results.json data
type CheckResults struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Score        ProjectScore `json:"score"`
	Rating       string       `json:"rating"`
	Category     string       `json:"category"`
	LogoURL      string       `json:"logo_url"`
	Maturity     string       `json:"maturity"`
	Snapshots    []string     `json:"snapshots"`
	Foundation   string       `json:"foundation"`
	UpdatedAt    int64        `json:"updated_at"`
	AcceptedAt   int64        `json:"accepted_at"`
	Description  string       `json:"description"`
	DisplayName  string       `json:"display_name"`
	Repositories []Repository `json:"repositories"`
}

type Repository struct {
	URL          string          `json:"url"`
	Name         string          `json:"name"`
	Score        RepositoryScore `json:"score"`
	Digest       string          `json:"digest"`
	Report       Report          `json:"report"`
	CheckSets    []string        `json:"check_sets"`
	RepositoryID string          `json:"repository_id"`
}

type Report struct {
	Data      Data     `json:"data"`
	ReportID  string   `json:"report_id"`
	CheckSets []string `json:"check_sets"`
	UpdatedAt int64    `json:"updated_at"`
}

type ProjectScore struct {
	Legal         float64 `json:"legal"`
	Global        float64 `json:"global"`
	License       float64 `json:"license"`
	Security      float64 `json:"security"`
	Documentation float64 `json:"documentation"`
	GlobalWeight  int     `json:"global_weight"`
	BestPractices float64 `json:"best_practices"`
}

type RepositoryScore struct {
	Legal               float64 `json:"legal"`
	Global              float64 `json:"global"`
	License             float64 `json:"license"`
	Security            float64 `json:"security"`
	LegalWeight         int     `json:"legal_weight"`
	Documentation       float64 `json:"documentation"`
	GlobalWeight        int     `json:"global_weight"`
	BestPractices       float64 `json:"best_practices"`
	LicenseWeight       int     `json:"license_weight"`
	SecurityWeight      int     `json:"security_weight"`
	DocumentationWeight int     `json:"documentation_weight"`
	BestPracticesWeight int     `json:"best_practices_weight"`
}

type Data struct {
	Legal         Legal         `json:"legal"`
	License       License       `json:"license"`
	Security      Security      `json:"security"`
	Documentation Documentation `json:"documentation"`
	BestPractices BestPractices `json:"best_practices"`
}

type Legal struct {
	TrademarkDisclaimer TrademarkDisclaimer `json:"trademark_disclaimer"`
}

type TrademarkDisclaimer struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type License struct {
	LicenseSPDXID   LicenseSPDXID   `json:"license_spdx_id"`
	LicenseApproved LicenseApproved `json:"license_approved"`
	LicenseScanning LicenseScanning `json:"license_scanning"`
}

type LicenseSPDXID struct {
	Value  string `json:"value"`
	Exempt bool   `json:"exempt"`
	Failed bool   `json:"failed"`
	Passed bool   `json:"passed"`
}

type LicenseApproved struct {
	Value  bool `json:"value"`
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type LicenseScanning struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type Security struct {
	SBOM                 SBOM                 `json:"sbom"`
	Maintained           Maintained           `json:"maintained"`
	CodeReview           CodeReview           `json:"code_review"`
	SecurityPolicy       SecurityPolicy       `json:"security_policy"`
	SignedReleases       SignedReleases       `json:"signed_releases"`
	BinaryArtifacts      BinaryArtifacts      `json:"binary_artifacts"`
	TokenPermissions     TokenPermissions     `json:"token_permissions"`
	DangerousWorkflow    DangerousWorkflow    `json:"dangerous_workflow"`
	DependencyUpdateTool DependencyUpdateTool `json:"dependency_update_tool"`
}

type SBOM struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type Maintained struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
	// Details string `json:"details"` // Omitted due to length
}

type CodeReview struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
	// Details string `json:"details"` // Omitted due to length
}

type SecurityPolicy struct {
	URL    string `json:"url"`
	Exempt bool   `json:"exempt"`
	Failed bool   `json:"failed"`
	Passed bool   `json:"passed"`
	// Details string `json:"details"` // Omitted due to length
}

type SignedReleases struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
	// Details string `json:"details"` // Omitted due to length
}

type BinaryArtifacts struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
	// Details string `json:"details"` // Omitted due to length
}

type TokenPermissions struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
	// Details string `json:"details"` // Omitted due to length
}

type DangerousWorkflow struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
	// Details string `json:"details"` // Omitted due to length
}

type DependencyUpdateTool struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
	// Details string `json:"details"` // Omitted due to length
}

type Documentation struct {
	Readme        Readme        `json:"readme"`
	Roadmap       Roadmap       `json:"roadmap"`
	Website       Website       `json:"website"`
	Adopters      Adopters      `json:"adopters"`
	Changelog     Changelog     `json:"changelog"`
	Governance    Governance    `json:"governance"`
	Maintainers   Maintainers   `json:"maintainers"`
	Contributing  Contributing  `json:"contributing"`
	CodeOfConduct CodeOfConduct `json:"code_of_conduct"`
}

type Readme struct {
	URL    string `json:"url"`
	Exempt bool   `json:"exempt"`
	Failed bool   `json:"failed"`
	Passed bool   `json:"passed"`
}

type Roadmap struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type Website struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type Adopters struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type Changelog struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type Governance struct {
	URL    string `json:"url"`
	Exempt bool   `json:"exempt"`
	Failed bool   `json:"failed"`
	Passed bool   `json:"passed"`
}

type Maintainers struct {
	URL    string `json:"url"`
	Exempt bool   `json:"exempt"`
	Failed bool   `json:"failed"`
	Passed bool   `json:"passed"`
}

type Contributing struct {
	URL    string `json:"url"`
	Exempt bool   `json:"exempt"`
	Failed bool   `json:"failed"`
	Passed bool   `json:"passed"`
}

type CodeOfConduct struct {
	URL    string `json:"url"`
	Exempt bool   `json:"exempt"`
	Failed bool   `json:"failed"`
	Passed bool   `json:"passed"`
}

type BestPractices struct {
	CLA                  CLA                  `json:"cla"`
	DCO                  DCO                  `json:"dco"`
	Analytics            Analytics            `json:"analytics"`
	OpenSSFBadge         OpenSSFBadge         `json:"openssf_badge"`
	RecentRelease        RecentRelease        `json:"recent_release"`
	SlackPresence        SlackPresence        `json:"slack_presence"`
	ArtifactHubBadge     ArtifactHubBadge     `json:"artifacthub_badge"`
	CommunityMeeting     CommunityMeeting     `json:"community_meeting"`
	GitHubDiscussions    GitHubDiscussions    `json:"github_discussions"`
	OpenSSFBestPractices OpenSSFBestPractices `json:"open_ssf_best_practices"`
}

type CLA struct {
	Exempt          bool   `json:"exempt"`
	Failed          bool   `json:"failed"`
	Passed          bool   `json:"passed"`
	ExemptionReason string `json:"exemption_reason"`
}

type DCO struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type Analytics struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type OpenSSFBadge struct {
	URL    string `json:"url"`
	Exempt bool   `json:"exempt"`
	Failed bool   `json:"failed"`
	Passed bool   `json:"passed"`
}

type RecentRelease struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type SlackPresence struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type ArtifactHubBadge struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type CommunityMeeting struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type GitHubDiscussions struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

type OpenSSFBestPractices struct {
	Exempt bool `json:"exempt"`
	Failed bool `json:"failed"`
	Passed bool `json:"passed"`
}

// ToCSV converts the ProjectScore struct to a CSV string using fmt.Printf
func (ps *ProjectScore) ToCSV() string {
	return fmt.Sprintf("%f,%f,%f,%f,%f,%f",
		ps.Global,
		ps.BestPractices,
		ps.Documentation,
		ps.Legal,
		ps.License,
		ps.Security,
	)
}

func (cr *CheckResults) SummaryScoreCSV() string {
	return fmt.Sprintf("%s,%s,%s,%s\n",
		cr.Name,
		time.Unix(cr.UpdatedAt, 0).Format("2006-01"),
		cr.Rating,
		cr.Score.ToCSV(),
	)
}

func (s *Security) GetSecurityResults() string {
	return fmt.Sprintf("%t,%t,%t,%t,%t,%t,%t,%t,%t",
		s.SBOM.Passed,
		s.Maintained.Passed,
		s.CodeReview.Passed,
		s.SecurityPolicy.Passed,
		s.SignedReleases.Passed,
		s.BinaryArtifacts.Passed,
		s.TokenPermissions.Passed,
		s.DangerousWorkflow.Passed,
		s.DependencyUpdateTool.Passed,
	)
}

func (cr *CheckResults) GetSecurityCSV() (csv string) {
	for _, repo := range cr.Repositories {
		if !contains(repo.CheckSets, "code") && !contains(repo.CheckSets, "code-lite") {
			continue
		}
		csv += fmt.Sprintf("%s,%s,%f,%s,%s,%s,%s,%t,%s\n",
			repo.URL,
			time.Unix(cr.UpdatedAt, 0).Format("2006-01"),
			repo.Score.Security,
			cr.Name,
			cr.Foundation,
			cr.Maturity,
			repo.CheckSets,
			IsSlam22Participant(cr.Foundation),
			repo.Report.Data.Security.GetSecurityResults(),
		)
	}
	return csv
}

// GetSecScoreCSVHeaders returns the CSV headers for GetSecurityCSV, which includes GetSecurityResults as well.
func GetSecScoreCSVHeaders() (csv string) {
	return "URL,Date,Security Score,Project,Foundation,Maturity,CheckSets,Slam23 Participant,SBOM,Maintained,Code Review,Security Policy,Signed Releases,Binary Artifacts,Token Permissions,Dangerous Workflow,Dependency Update Tool\n"
}

func IsSlam22Participant(projectName string) bool {
	participants := []string{
		"argo",
		"artifact-hub",
		"chaos-mesh",
		"cloudevents",
		"cortex",
		"flagger",
		"flux",
		"k8gb",
		"kubewarden",
		"open-feature",
		"pixie",
	}
	// if projectName is in participants, return true
	for _, participant := range participants {
		if projectName == participant {
			return true
		}
	}
	return false
}
