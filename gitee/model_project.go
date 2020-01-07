/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 搜索仓库
type Project struct {
	Id                  int32      `json:"id,omitempty"`
	FullName            string     `json:"full_name,omitempty"`
	HumanName           string     `json:"human_name,omitempty"`
	Url                 string     `json:"url,omitempty"`
	Namespace           *Namespace `json:"namespace,omitempty"`
	Path                string     `json:"path,omitempty"`
	Name                string     `json:"name,omitempty"`
	Owner               *UserBasic `json:"owner,omitempty"`
	Description         string     `json:"description,omitempty"`
	Private             bool       `json:"private,omitempty"`
	Public              bool       `json:"public,omitempty"`
	Internal            bool       `json:"internal,omitempty"`
	Fork                bool       `json:"fork,omitempty"`
	HtmlUrl             string     `json:"html_url,omitempty"`
	SshUrl              string     `json:"ssh_url,omitempty"`
	ForksUrl            string     `json:"forks_url,omitempty"`
	KeysUrl             string     `json:"keys_url,omitempty"`
	CollaboratorsUrl    string     `json:"collaborators_url,omitempty"`
	HooksUrl            string     `json:"hooks_url,omitempty"`
	BranchesUrl         string     `json:"branches_url,omitempty"`
	TagsUrl             string     `json:"tags_url,omitempty"`
	BlobsUrl            string     `json:"blobs_url,omitempty"`
	StargazersUrl       string     `json:"stargazers_url,omitempty"`
	ContributorsUrl     string     `json:"contributors_url,omitempty"`
	CommitsUrl          string     `json:"commits_url,omitempty"`
	CommentsUrl         string     `json:"comments_url,omitempty"`
	IssueCommentUrl     string     `json:"issue_comment_url,omitempty"`
	IssuesUrl           string     `json:"issues_url,omitempty"`
	PullsUrl            string     `json:"pulls_url,omitempty"`
	MilestonesUrl       string     `json:"milestones_url,omitempty"`
	NotificationsUrl    string     `json:"notifications_url,omitempty"`
	LabelsUrl           string     `json:"labels_url,omitempty"`
	ReleasesUrl         string     `json:"releases_url,omitempty"`
	Recommend           bool       `json:"recommend,omitempty"`
	Homepage            string     `json:"homepage,omitempty"`
	Language            string     `json:"language,omitempty"`
	ForksCount          int32      `json:"forks_count,omitempty"`
	StargazersCount     int32      `json:"stargazers_count,omitempty"`
	WatchersCount       int32      `json:"watchers_count,omitempty"`
	DefaultBranch       string     `json:"default_branch,omitempty"`
	OpenIssuesCount     int32      `json:"open_issues_count,omitempty"`
	HasIssues           bool       `json:"has_issues,omitempty"`
	HasWiki             bool       `json:"has_wiki,omitempty"`
	PullRequestsEnabled bool       `json:"pull_requests_enabled,omitempty"`
	HasPage             bool       `json:"has_page,omitempty"`
	License             string     `json:"license,omitempty"`
	Outsourced          bool       `json:"outsourced,omitempty"`
	ProjectCreator      string     `json:"project_creator,omitempty"`
	Members             []string   `json:"members,omitempty"`
	PushedAt            string     `json:"pushed_at,omitempty"`
	CreatedAt           string     `json:"created_at,omitempty"`
	UpdatedAt           string     `json:"updated_at,omitempty"`
	Parent              *Project   `json:"parent,omitempty"`
	Paas                string     `json:"paas,omitempty"`
	Stared              string     `json:"stared,omitempty"`
	Watched             string     `json:"watched,omitempty"`
	Permission          string     `json:"permission,omitempty"`
	Relation            string     `json:"relation,omitempty"`
}
