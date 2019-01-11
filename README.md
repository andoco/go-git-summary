Command line utility to check the status of multiple Git repositories.

The intended use case is to ensure that no uncommitted or unpushed changes exist on the local machine.

## Features

- Checks for uncommitted changes in the work tree

## Building

```
go install
```

## Usage

```
gorepocheck $GOPATH/src
```

## Behaviours

Given a root path
Should find all Git repos under root
And check for outstanding changes

Given a repo path
Where repo has uncommited changes
Should print "DIRTY [commit]"

Given a repo path
Where repo has new files
Should print "DIRTY [commit]"

Given a repo path
Where repo has unpushed commits to tracked remote
Should print "DIRTY [push]"
