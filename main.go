package main

import (
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	workingDir, _ := os.Getwd()
	pkg, _ := build.ImportDir(workingDir, build.AllowBinary)

	var lockFile string
	for _, dependency := range pkg.Imports {
		packageHash := packageHash(dependency, build.Default.SrcDirs())
		lockFile += dependency + " - " + packageHash + "\n"
	}
	ioutil.WriteFile(".dondur.lock", []byte(lockFile), os.ModePerm)

	print(lockFile)
}

func packageHash(pkgName string, srcDirs []string) string {
	packageDir := packageDir(pkgName, srcDirs)

	gitHashCmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
	gitHashCmd.Dir = packageDir
	gitHash, err := gitHashCmd.Output()
	if err == nil {
		return strings.Trim(string(gitHash), "\n")
	}

	hgHashCmd := exec.Command("hg", "id", "-i")
	hgHashCmd.Dir = packageDir
	hgHash, err := hgHashCmd.Output()
	if err == nil {
		return strings.Trim(string(hgHash), "\n")
	}

	return "?"
}

func packageDir(pkg string, srcDirs []string) string {
	for _, srcDir := range srcDirs {
		if packageInDir(pkg, srcDir) {
			return srcDir + "/" + pkg
		}
	}
	return ""
}

func packageInDir(pkgName string, srcDir string) bool {
	pkg, _ := build.ImportDir(srcDir+"/"+pkgName, build.AllowBinary)
	if pkg.Name == "" {
		return false
	}
	return true
}