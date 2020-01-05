# git commands

- add files to stage
```
git add <filename>
```

- add files recursively
```
git add . 
```

- add files recursively and also update any files that are updated, deleted etc
```
git add -A
```

- add files and also update the file index for git understand the changes
```
git add -u
```

- remove files
```
git rm <filename> 
```

- back changes to working dir
```
git reset HEAD <filename>
```

- discart changes in working dir
```
git checkout -- <filename>
```

- commit files to repository
```
git commit 
```

- commit inline
```
git commit -m "message"
```

```
git commit -am "message"
```

- move files
```
git mv <filename> <new-location>
```

- list tracking files
```
git ls-files
```

- see history

```
git log
```

- see details from commit a

```
git show <commit-hash>
```

- abrrev history
```
git log --abbrev-commit
```

- history branching graph with decoration

```
git log --oneline --graph --decorate
```

- see history range based on hash
```
git log <hash$1> <hash$2>
```

- see history based on days
```
git log --since="3 days ago"
```

- see history based on a file 
```
git log -- <filename>
```

- see history based on a file  changes
```
git log --follow -- <path/filename>
```

- git alias *(command without git)

```
git config --global alias.hist "log --all --graph --decorate --oneline" 
```

- path to change alias 
```
~/.gitcofig
```

- git ignore pattern examples
```
specific file: MyFile.ext
file pattern: *.ext
folder: my-folder/
```