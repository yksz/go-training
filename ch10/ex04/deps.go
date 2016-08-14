package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("usage: %s <package>\n", os.Args[0])
		os.Exit(1)
	}
	target := os.Args[1]
	if !exists(target) {
		log.Fatalf("No such package: %s\n", target)
	}
	pkgs, err := packages()
	if err != nil {
		log.Fatal(err)
	}
	for _, pkg := range pkgs {
		if contains(target, pkg) {
			fmt.Println(pkg.ImportPath)
		}
	}
}

type Package struct {
	ImportPath string
	Deps       []string
}

func exists(pkgName string) bool {
	out, err := exec.Command("go", "list", pkgName).Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(out)) == pkgName
}

func packages() ([]*Package, error) {
	cmd := exec.Command("go", "list", "-f", "{{.ImportPath}}:{{join .Deps \",\"}}", "...")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	var pkgs []*Package
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		index := strings.Index(line, ":")
		if index == -1 {
			continue
		}
		key := line[:index]
		val := strings.Split(line[index+1:], ",")
		pkgs = append(pkgs, &Package{ImportPath: key, Deps: val})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return pkgs, nil
}

func contains(name string, pkg *Package) bool {
	for _, dep := range pkg.Deps {
		if dep == name {
			return true
		}
	}
	return false
}
