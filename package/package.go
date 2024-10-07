package cmimalloc

import (
	denv "github.com/jurgen-kluft/ccode/denv"
	cunittest "github.com/jurgen-kluft/cunittest/package"
)

// GetPackage returns the package object of 'cmimalloc'
func GetPackage() *denv.Package {
	// Dependencies
	cunittestpkg := cunittest.GetPackage()

	// The main (cmimalloc) package
	mainpkg := denv.NewPackage("cmimalloc")
	mainpkg.AddPackage(cunittestpkg)

	// 'cmimalloc' library
	mainlib := denv.SetupDefaultCppLibProjectWithLibs("cmimalloc", "github.com\\jurgen-kluft\\cmimalloc", getPlatformLibs())

	// 'cmimalloc' unittest project
	maintest := denv.SetupDefaultCppTestProject("cmimalloc"+"_test", "github.com\\jurgen-kluft\\cmimalloc")
	maintest.Dependencies = append(maintest.Dependencies, cunittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}

func getPlatformLibs() []*denv.Lib {
	if denv.IsWindows() {
		winLibs := []*denv.Lib{}
		return winLibs
	}
	if denv.IsMacOS() {
		macLibs := []*denv.Lib{}
		return macLibs
	}
	return []*denv.Lib{}
}
