# KSP2 Science Report

A small terminal (CLI) app to get an overview from KSP2 save file of what science one have done and where.

## How to use

Locate the ksp2sr.exe in any location of your wish. Ensure that whatever folder the ksp2sr.exe file is in, there is also a folder named ´Saves´. This is where you´ll need copy-paste your save file to be interrogated into.

When a save file have been copied to the ´Saves´ folder, open the ksp2sr.exe app by either navigating to where the file is located and typing in ´ksp2sr.exe´ or by double clicking on the file from File-Explorer. Hell, you can in File-Explorer make a convenient little desktop shortcut to place wherever you want, and double click on that.

This (if not launching from the CLI) will open a CLI window, and present you with the menu. If more than one save is in the ´Saves´ directory, you will be promted to select a save for interrogation.

Press ´1´ through N to select relevant save. Press ´Q´ to quit.

Once done, and assuming no errors, you will be given a list of Bodies (Planets) for which Science reports may be available. Select which one by pressing ´1´ through N. Press ´B´ to go back to file selection. Press ´Q´ to quit.

If selecting a planet with one or more moons, you´ll be prompted to select the planet itself (by pressing ´1´) or one of the moons (pressing ´2´ through N).

*Coming Soon:*
>When viewing the science for a planet or moon, press ´P´ to export to PDF. A PDF report will be created in the folder ´Report´ where the ksp2sr.exe file is located. 

Hit Return to go back to planet/moon selection.

## Run and build from source

To build your own executable simply run you will need Go (aka Golang) 1.20.3 or newer installed.

<https://go.dev/>

This project will run and compile for all popular OS and CPUs without any exernal dependencies.
At first run, any needed third party libraries will be downloaded automatically, or you can trigger this manually by running:

```go mod tidy```

To run the project from source, use:

```go run *.go```

To build the project into a single static executable use:

```go build -o <file name here>```

If on - or compiling for - Windows, it might be prudent to add a ´.exe´ suffix to the filename. More reasonable OS do not require this.

#### Compile App for Windows if on Mac or Linux: 

```GOOS=windows GOARCH=amd64 go build -o ksp2sr.exe```

## Releases

- Version format: [major release].[new feature(s)].[bugfix patch-version]
- Date format: yyyy-mm-dd

### v.1.0.0: 2023-12-28

- First release

## License: MIT

Full license text found in `LICENCE.md` file

## Copyright © 2023 Roy Dybing

ʕ◔ϖ◔ʔ