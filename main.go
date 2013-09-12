package main

import (
	"flag"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	externalOnly := flag.Bool("x", false, "List only the external dependencies")
	flag.Parse()

	workingDir, _ := os.Getwd()
	pkg, _ := build.ImportDir(workingDir, build.AllowBinary)

	var lockFile string
	for _, pkgName := range pkg.Imports {
		pkgHash := packageHash(pkgName, build.Default.SrcDirs())

		if !*externalOnly || (*externalOnly && packageExternal(pkgName)) {
			lockFile += pkgName + " - " + pkgHash + "\n"
		}
	}
	ioutil.WriteFile(".dondur.lock", []byte(lockFile), os.ModePerm)

	print(lockFile)
}

func packageHash(pkgName string, srcDirs []string) string {
	pkgDir := packageDir(pkgName, srcDirs)

	gitHashCmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
	gitHashCmd.Dir = pkgDir
	gitHash, err := gitHashCmd.Output()
	if err == nil {
		return strings.Trim(string(gitHash), "\n")
	}

	hgHashCmd := exec.Command("hg", "id", "-i")
	hgHashCmd.Dir = pkgDir
	hgHash, err := hgHashCmd.Output()
	if err == nil {
		return strings.TrimSpace(string(hgHash))
	}

	return "?"
}

func packageDir(pkgName string, srcDirs []string) string {
	for _, srcDir := range srcDirs {
		if packageInDir(pkgName, srcDir) {
			return srcDir + "/" + pkgName
		}
	}
	return ""
}

func packageExternal(pkgName string) bool {
	pkgNameParts := strings.Split(pkgName, "/")
	if strings.Contains(pkgNameParts[0], ".") {
		return true
	}
	return false
}

func packageInDir(pkgName string, srcDir string) bool {
	_, err := build.ImportDir(srcDir+"/"+pkgName, build.AllowBinary)
	if err != nil {
		return false
	}
	return true
}
