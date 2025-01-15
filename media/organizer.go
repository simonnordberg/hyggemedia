package media

type MediaOrganizer interface {
	Organize(title, srcDir, destDir string, dryRun, move bool) error
}
