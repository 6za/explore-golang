# Overview 

Exploring design of golang+cobra-cli+reausbale patterns of commands.



## How to create a project similar to this:

To save time, there is an image you can build so, you can start coding from zero to go-hero: 
```bash 
docker build working-shed/cobra_dev -t cobra_dev_v2

docker run \
  -v  go-pkg:/go/pkg \
  -v $SSH_CREDS:/home/developer/.ssh/ \
  -it --name dev-${BRANCH}  \
  --dns="1.0.0.1" --dns="208.67.222.222" --dns="8.8.8.8" \
  -v $(PWD)/go/cobra/scenario1:/home/developer/app \
  -e "TERM=xterm-256color" \
  -v $AWS_CREDS_PATH:/home/developer/.aws \
   cobra_dev_v2
```

Then on the container or your local machine: 
```bash 
go mod init github.com/6za/explore-golang/go/cobra/scenario1
cobra-cli init
go mod tidy
```


## Notes on the design

This is based on `kubefirst`comamnd design goals: 
- [ ] To have every family of commands organized in a sub-folder of `/cmd`
- [ ] Head of family command to `./cmd/root.go` command, to help on some better visibility of child commands. 
- [ ] Manually create commands from this template. 

At `./cmd/packagecmd/commandName.go`:
```golang

import "github.com/spf13/cobra"

var (
	someFlag             string
	...
	someOtherFlagFlag           bool
)

func NewCommand() *cobra.Command {

	myActionCmd := &cobra.Command{
		Use:     "action",
		Short:   "something short",
		Long:    "something long",
		PreRunE: validateAction, 
		RunE:    runAction,
		// PostRunE: runPostAction,
	}

	// todo review defaults and update descriptions
	myActionCmd.Flags().StringVar(&someFlag, "some-flag", "", "....")
	myActionCmd.Flags().BoolVar(&someOtherFlagFlag, "some-other-flag", false, "", "....")

	// wire up new commands
	myActionCmd.AddCommand(ChildCommand())

	return myActionCmd
}

func ChildCommand() *cobra.Command {
	childActionCmd := &cobra.Command{
		Use:   "child-action",
		Short:   "something short",
		Long:    "something long",
		RunE:  someChildActionCivo,
	}

	return childActionCmd
}

```

- [ ]  All flags are defined as global variables at the package level, and shared by all functions at this package.
- [ ] Function called will receive all flags values as variables, instead of a shared global mechanism or something else like it. 

## Notes to explore beyond the notes from the design of kubefirst

- [ ] Explore a better way to collect parameters in commons `structs` or `interfaces`.
- [ ] Explore a better [MVC style separation from FlagsUI](https://github.com/kubefirst/kubefirst/discussions/1032) from the implementation of behaviors. 
- [ ] Try to find a pattern to "actionController", to make commands code cleaner and more re-usable cross different flavors.
- [ ] Search for an foundation SDK design to be shared in other projects.




# Sample of execution

```bash 
 go run . alpha


2023-01-24T16:07:02Z INF cmd/alpha/alpha.go:16 > Alpha CMD executed: (someFlag:  )  go_version=go1.18.7 pid=881
2023-01-24T16:07:02Z INF internal/alphaController/controller.go:7 > Execute alpha steps:  go_version=go1.18.7 pid=881
```

# Running tests

```bash 
go test -v -short ./...
```
