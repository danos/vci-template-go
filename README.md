# vci-template-go

This repository contains a minimal VCI component written in Go. This
package may be copied and modified to serve as a template for creating
a new VCI component in Go.

## Steps for creating a new component
 - Replace "vci-template-go" with the desired name of the component in
   every file in this repo (don't forget the debian/ directory).
 - Replace "yang/vci-template-go-v1.yang" with your the desired YANG defintion.
 - Create your component logic in the library portion of the new repo
   (modify "vci-template-go.go" in the top level of the repo).
 - Rename "cmd/vci-template-go" to the desired name of your component.
 - Ensure the "debian/*.install" files are pointed at the appropriate new files.

## Additional documentation
For full VCI documentation see the following pages on the DANOS community wiki.
 - [VCI overview](https://danosproject.atlassian.net/wiki/spaces/DAN/pages/684654694/VCI+API)
 - [Go VCI API](https://danosproject.atlassian.net/wiki/spaces/DAN/pages/684621933/Go+VCI+API)
 - [godoc](https://pkg.go.dev/github.com/danos/vci)

## License
The repository is licensed MIT-0 which means you may copy and modify
the files with no requirements (including no need for attribution).
