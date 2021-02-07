# git-release-materials

## synopsis

Prepare release materials that are managed by Git.  

* Create a list of files that have changed.
* Collect sources before and after modification.
* Collect sources to be released.

## usage

This tool uses the git diff command internally.

`commit1` is the start of the commit.  `commit2` is the end of the commit.  

### Create a list of files that have changed.

```shell
git-release-materials list commit1 commit2
```

### Collect sources before and after modification.

```shell
git-release-materials before-after commit1 commit2
```

### Collect sources to be released.

```shell
git-release-materials release commit1 commit2
```

### Options

#### -d, --workdir

You can specify the working directory in which to run this program. It will be an error if the .git directory does not exist in this directory.

#### -o, --output

You can specify the directory where the materials will be placed. If the directory does not exist, it will be created.