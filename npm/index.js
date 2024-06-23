// Copyright © 2023 Cowk8s Corp
// SPDX-License-Identifier: Apache-2.0

var binwrap = require("binwrap")
var path = require("path")

var packageInfo = require(path.join(__dirname, "..", "package.json"))
var version = packageInfo.version
var root = "https://github.com/cowk8s/cli/releases/download/v" + version

module.exports = binwrap({
  dirname: __dirname,
  binaries: ["cowk8s"],
  urls: {
    "linux-x64": root + "/cowk8s_" + version + "-linux_64bit.tar.gz",
    "linux-arm64": root + "/cowk8s_" + version + "-linux_arm64.tar.gz",
    "win32-x64": root + "/cowk8s_" + version + "-windows_64bit.zip",
    "darwin-x64": root + "/cowk8s_" + version + "-macOS_64bit.tar.gz",
    "darwin-arm64": root + "/cowk8s_" + version + "-macOS_arm64.tar.gz",
  },
})
