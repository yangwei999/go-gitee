/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 取消用户测试 Pull Request
type PullRequest struct {
	Id                int32       `json:"id,omitempty"`
	Url               string      `json:"url,omitempty"`
	HtmlUrl           string      `json:"html_url,omitempty"`
	DiffUrl           string      `json:"diff_url,omitempty"`
	PatchUrl          string      `json:"patch_url,omitempty"`
	IssueUrl          string      `json:"issue_url,omitempty"`
	CommitsUrl        string      `json:"commits_url,omitempty"`
	ReviewCommentsUrl string      `json:"review_comments_url,omitempty"`
	ReviewCommentUrl  string      `json:"review_comment_url,omitempty"`
	CommentsUrl       string      `json:"comments_url,omitempty"`
	StatusesUrl       string      `json:"statuses_url,omitempty"`
	AssigneesNumber   int32       `json:"assignees_number,omitempty"`
	TestersNumber     int32       `json:"testers_number,omitempty"`
	Number            int32       `json:"number,omitempty"`
	State             string      `json:"state,omitempty"`
	Title             string      `json:"title,omitempty"`
	Body              string      `json:"body,omitempty"`
	BodyHtml          string      `json:"body_html,omitempty"`
	Assignees         []UserBasic `json:"assignees,omitempty"`
	Testers           []UserBasic `json:"testers,omitempty"`
	Milestone         *Milestone  `json:"milestone,omitempty"`
	Labels            []Label     `json:"labels,omitempty"`
	Locked            bool        `json:"locked,omitempty"`
	CreatedAt         string      `json:"created_at,omitempty"`
	UpdatedAt         string      `json:"updated_at,omitempty"`
	ClosedAt          string      `json:"closed_at,omitempty"`
	MergedAt          string      `json:"merged_at,omitempty"`
	Mergeable         bool        `json:"mergeable,omitempty"`
	Head              *BasicInfo  `json:"head,omitempty"`
	Base              *BasicInfo  `json:"base,omitempty"`
	Links             interface{} `json:"_links,omitempty"`
	User              *UserBasic  `json:"user,omitempty"`
	Comments          int32       `json:"comments,omitempty"`
	Commits           int32       `json:"commits,omitempty"`
	Additions         int32       `json:"additions,omitempty"`
	Deletions         int32       `json:"deletions,omitempty"`
	ChangedFiles      int32       `json:"changed_files,omitempty"`
}
