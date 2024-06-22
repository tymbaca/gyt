package main

func CatFile(repo Repo, mode string, shaStr string) {
	sha, err := NewSHAFromString(shaStr)
	if err != nil {
		panic(err)
	}

	o, err := ReadObject(repo, sha)
	if err != nil {
		panic(err)
	}

	switch mode {
	case "-t":
		Log(o.Kind)
	case "-s":
		Log(o.Size)
	case "-p":
		Logf("%s", o.Body)
	default:
		panic("pass cat-file option")
	}
}
