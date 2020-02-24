/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// repo patch parameter
type RepoPatchParam struct {
	// 用户授权码
	AccessToken string `json:"access_token,omitempty"`
	// 仓库名称
	Name string `json:"name,omitempty"`
	// 仓库描述
	Description string `json:"description,omitempty"`
	// 主页(eg: https://gitee.com)
	Homepage string `json:"homepage,omitempty"`
	// 允许提Issue与否。默认: 允许(true)
	HasIssues string `json:"has_issues,omitempty"`
	// 提供Wiki与否。默认: 提供(true)
	HasWiki string `json:"has_wiki,omitempty"`
	// 仓库公开或私有。
	Private string `json:"private,omitempty"`
	// 更新仓库路径
	Path string `json:"path,omitempty"`
	// 更新默认分支
	DefaultBranch string `json:"default_branch,omitempty"`
}
