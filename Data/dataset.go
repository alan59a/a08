package Data

// Returens the size of the Dataset
func (d *Dataset) Size() int { return d.size }

// Sets the Root path for data storage
func (d *Dataset) SetRoot(path string) { d.root = path }

// Sets the download link
func (d *Dataset) SetLink(url string) { d.link = url }

// Switches the Dataset group
func (d *Dataset) ChangeGroup() { d.train = !d.train }

// Switches the command for download
func (d *Dataset) ChangeDownload() { d.download = !d.download }

// Prints a summary of Dataset
func (d *Dataset) Summary() {
	println("SORRY! not implemented yet")
}
