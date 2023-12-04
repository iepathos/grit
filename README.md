Grit
----


command line tool for pull and checking out branches across multiple git repositories in parallel.  Intended to be useful managing cloning and pulling to sync changes in larger multi-repo projects.


If this error is seen during pull
```
2023/12/03 15:53:26 ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain
```

Then the user must run 
```
ssh-add
```

if their default ssh key ~/.ssh/id_rsa hasn't been added to their ssh-agent as this is what go-git relies on.




Roadmap
-----

have different repos output pulls to different lines in terminal so it's easier to track which one is at what state of pull or clone

add ability to pass in config to grit

add ability to pass in remote config to grit

see if we can fix needing to run ssh-add may be better solution that that


