import sys

with open("pkg/config/config.go", "r") as f:
    code = f.read()

target1 = """type GLMSearchConfig struct {
	Enabled  bool `json:"enabled"       env:"PICOCLAW_TOOLS_WEB_GLM_ENABLED"`
	apiKey   string
	secDirty bool
	BaseURL  string `json:"base_url"      env:"PICOCLAW_TOOLS_WEB_GLM_BASE_URL"`"""

replacement1 = """type GLMSearchConfig struct {
	Enabled  bool   `json:"enabled"       env:"PICOCLAW_TOOLS_WEB_GLM_ENABLED"`
	apiKey   string
	secDirty bool
	BaseURL  string `json:"base_url"      env:"PICOCLAW_TOOLS_WEB_GLM_BASE_URL"`"""

new_code = code.replace(target1, replacement1)

with open("pkg/config/config.go", "w") as f:
    f.write(new_code)
