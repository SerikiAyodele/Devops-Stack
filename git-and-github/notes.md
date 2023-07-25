==========
BASICS ...
===========
* Git is a version control system, a VCS is a software that tracks and manages chnages to a file over time
* Git is the VCS that runs locally on your machine 
* Github is a servuce that hosts git repositories in the cloud and makesnit easy to collaborate with other people

BASIC COMMANDS
==============
* ls                      #list the contents or file of a folder                 
* mkdir plants            #creates a directory called plants
* cd                      #change directory
* touch                   #create a new file
* rm                      #remove a file
* rm -rf                  #remove a directory

NOTES
=====
* A git repo is a workspace that tracks and manages files within a folder
* Git status: gives information on the current status of a git repository and it's content 
* Git init: to create a new git repository- creates a hidden .git folder
  * rm -rf .git: to remove a .git folder

COMMITS
=======
* Each commit is like a snapshot or a checkpoint at a particular time, of your repo
* Working directory: where we actually work on our project
* Staging area: Where we add our changes to, before we make a commit
* Atomic commit: when possible, a commit should encompass a single feature, change or fix.
  This makes it easy to roll back changes later
* git log --abbrev-commit: will give you shorter info on the logs
* git log --pretty=oneline: is easier to read

  AMENDING COMMITS
  ================
  * Allows us to update just the previous commit using the --amend option
  * you can use this to chnage the comment meassage or add a new file to the previous commit 
  
  IGNORING FILES 
  ==============
  * Files to ignore include:
      - Secrets, api keys, credentials etc
      - Operating sytem files
      - Log files
      - Dependencies and packages
  * .DS_Store: will ignore files named .DS_Store
    folderName/: will ignore an entire directory
    *.log: will ignore any files with the .log extension