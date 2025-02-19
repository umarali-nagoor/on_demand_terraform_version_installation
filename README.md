# Scan the versions.tf to get the terraform version & use the same for terraform commands execution

1) go run main.go use version.tf

NOTE: 

Command "tfenv use", checks whether requested terraform version installed or not.
If its already installed it wil make use of it for further terraform command exceution
else it installs the requested terraform version first then it will start making use of it.

2) go run main.go list

It will display the list of terraform versions installed and points to current version as follows

1.5.6
  1.5.5
  1.4.6
  1.3.2
  1.2.5
  1.2.4
  1.2.1
* 1.2.0 (set by /Users/umarali/.tfenv/version)
  1.1.8

3) Show all the released terraform versions

tfenv list-remote
