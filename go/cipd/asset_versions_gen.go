// Code generated by "go run gen_versions.go"; DO NOT EDIT

package cipd

var PACKAGES = map[string]*Package{
	"infra/3pp/tools/cpython/linux-amd64": {
		Path:    "cipd_bin_packages/cpython",
		Name:    "infra/3pp/tools/cpython/linux-amd64",
		Version: "version:2@2.7.18.chromium.42",
	},
	"infra/3pp/tools/cpython/linux-arm64": {
		Path:    "cipd_bin_packages/cpython",
		Name:    "infra/3pp/tools/cpython/linux-arm64",
		Version: "version:2@2.7.18.chromium.42",
	},
	"infra/3pp/tools/cpython/linux-armv6l": {
		Path:    "cipd_bin_packages/cpython",
		Name:    "infra/3pp/tools/cpython/linux-armv6l",
		Version: "version:2@2.7.18.chromium.42",
	},
	"infra/3pp/tools/cpython/mac-amd64": {
		Path:    "cipd_bin_packages/cpython",
		Name:    "infra/3pp/tools/cpython/mac-amd64",
		Version: "version:2@2.7.18.chromium.42",
	},
	"infra/3pp/tools/cpython/windows-386": {
		Path:    "cipd_bin_packages/cpython",
		Name:    "infra/3pp/tools/cpython/windows-386",
		Version: "version:2@2.7.18.chromium.42",
	},
	"infra/3pp/tools/cpython/windows-amd64": {
		Path:    "cipd_bin_packages/cpython",
		Name:    "infra/3pp/tools/cpython/windows-amd64",
		Version: "version:2@2.7.18.chromium.42",
	},
	"infra/3pp/tools/cpython3/linux-amd64": {
		Path:    "cipd_bin_packages/cpython3",
		Name:    "infra/3pp/tools/cpython3/linux-amd64",
		Version: "version:2@3.8.10.chromium.19",
	},
	"infra/3pp/tools/cpython3/linux-arm64": {
		Path:    "cipd_bin_packages/cpython3",
		Name:    "infra/3pp/tools/cpython3/linux-arm64",
		Version: "version:2@3.8.10.chromium.19",
	},
	"infra/3pp/tools/cpython3/linux-armv6l": {
		Path:    "cipd_bin_packages/cpython3",
		Name:    "infra/3pp/tools/cpython3/linux-armv6l",
		Version: "version:2@3.8.10.chromium.19",
	},
	"infra/3pp/tools/cpython3/mac-amd64": {
		Path:    "cipd_bin_packages/cpython3",
		Name:    "infra/3pp/tools/cpython3/mac-amd64",
		Version: "version:2@3.8.10.chromium.19",
	},
	"infra/3pp/tools/cpython3/windows-386": {
		Path:    "cipd_bin_packages/cpython3",
		Name:    "infra/3pp/tools/cpython3/windows-386",
		Version: "version:2@3.8.10.chromium.19",
	},
	"infra/3pp/tools/cpython3/windows-amd64": {
		Path:    "cipd_bin_packages/cpython3",
		Name:    "infra/3pp/tools/cpython3/windows-amd64",
		Version: "version:2@3.8.10.chromium.19",
	},
	"infra/3pp/tools/git/linux-amd64": {
		Path:    "cipd_bin_packages",
		Name:    "infra/3pp/tools/git/linux-amd64",
		Version: "version:2@2.36.0-rc1.chromium.8",
	},
	"infra/3pp/tools/git/linux-arm64": {
		Path:    "cipd_bin_packages",
		Name:    "infra/3pp/tools/git/linux-arm64",
		Version: "version:2@2.36.0-rc1.chromium.8",
	},
	"infra/3pp/tools/git/linux-armv6l": {
		Path:    "cipd_bin_packages",
		Name:    "infra/3pp/tools/git/linux-armv6l",
		Version: "version:2@2.36.0-rc1.chromium.8",
	},
	"infra/3pp/tools/git/mac-amd64": {
		Path:    "cipd_bin_packages",
		Name:    "infra/3pp/tools/git/mac-amd64",
		Version: "version:2@2.36.0-rc1.chromium.8",
	},
	"infra/3pp/tools/git/windows-386": {
		Path:    "cipd_bin_packages",
		Name:    "infra/3pp/tools/git/windows-386",
		Version: "version:2@2.35.1.chromium.8",
	},
	"infra/3pp/tools/git/windows-amd64": {
		Path:    "cipd_bin_packages",
		Name:    "infra/3pp/tools/git/windows-amd64",
		Version: "version:2@2.35.1.chromium.8",
	},
	"infra/gsutil": {
		Path:    "cipd_bin_packages",
		Name:    "infra/gsutil",
		Version: "version:4.46",
	},
	"infra/tools/cipd/${os}-${arch}": {
		Path:    ".",
		Name:    "infra/tools/cipd/${os}-${arch}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/git/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/git/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/luci-auth/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/luci-auth/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/luci/git-credential-luci/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/luci/git-credential-luci/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/luci/isolate/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/luci/isolate/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/luci/isolated/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/luci/isolated/${platform}",
		Version: "git_revision:dc3a3dc4272aeef30698752d137ccd4f09526d69",
	},
	"infra/tools/luci/kitchen/${platform}": {
		Path:    ".",
		Name:    "infra/tools/luci/kitchen/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/luci/lucicfg/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/luci/lucicfg/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/luci/swarming/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/luci/swarming/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/luci/vpython-native/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/luci/vpython-native/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"infra/tools/luci/vpython/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "infra/tools/luci/vpython/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
	"skia/bots/bazel": {
		Path:    "bazel",
		Name:    "skia/bots/bazel",
		Version: "version:3",
	},
	"skia/bots/bazelisk": {
		Path:    "bazelisk",
		Name:    "skia/bots/bazelisk",
		Version: "version:0",
	},
	"skia/bots/cockroachdb": {
		Path:    "cockroachdb",
		Name:    "skia/bots/cockroachdb",
		Version: "version:4",
	},
	"skia/bots/gcloud_linux": {
		Path:    "gcloud_linux",
		Name:    "skia/bots/gcloud_linux",
		Version: "version:15",
	},
	"skia/bots/go": {
		Path:    "go",
		Name:    "skia/bots/go",
		Version: "version:16",
	},
	"skia/bots/go_win": {
		Path:    "go_win",
		Name:    "skia/bots/go_win",
		Version: "version:4",
	},
	"skia/bots/mockery": {
		Path:    "mockery",
		Name:    "skia/bots/mockery",
		Version: "version:2",
	},
	"skia/bots/node": {
		Path:    "node",
		Name:    "skia/bots/node",
		Version: "version:2",
	},
	"skia/bots/protoc": {
		Path:    "protoc",
		Name:    "skia/bots/protoc",
		Version: "version:0",
	},
	"skia/tools/goldctl/${platform}": {
		Path:    "cipd_bin_packages",
		Name:    "skia/tools/goldctl/${platform}",
		Version: "git_revision:55394369dcc8d5ca65b764dab10b355453186dcc",
	},
}
