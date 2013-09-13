package main

import (
	"flag"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	externalOnly := flag.Bool("x", false, "List only the external dependencies")
	flag.Parse()

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	pkg, err := build.ImportDir(workingDir, build.AllowBinary)
	if err != nil {
		log.Fatal(err)
	}

	var lockFile string
	for _, pkgName := range pkg.Imports {
		pkgDir := packageDir(pkgName, build.Default.SrcDirs())
		pkgHash := packageHash(pkgDir)

		if !*externalOnly || (*externalOnly && packageExternal(pkgName)) {
			lockFile += pkgHash + " " + pkgName + "\n"
		}
	}
	ioutil.WriteFile(".dondur.lock", []byte(lockFile), os.ModePerm)

	print(lockFile)
}

func packageHash(pkgDir string) string {
	gitHashCmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
	gitHashCmd.Dir = pkgDir
	gitHash, err := gitHashCmd.Output()
	if err == nil {
		return strings.Trim(string(gitHash), "\n")
	}

	hgHashCmd := exec.Command("hg", "--debug", "id", "-i")
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
