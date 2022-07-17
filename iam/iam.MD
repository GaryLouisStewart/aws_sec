# Creation of IAM Resources on AWS

The following script allows us to create a set of IAM users from a text file that is saved in the repository. This is particularly useful when intially setting up your aws account for defining some default users for IAC tooling/administration
In order to create an IAM user please run the following script. 

Creating multiple users can be done by doing the following...


defining a file, there is already a demo file called aws-users.txt

```bash
# run the following command..
# usage: go run iam_createuser.go
go run iam_createuser.go aws-users.txt

# output is similar to below

go run iam_createuser.go aws-users.txt
Success {
  User: {
    Arn: "arn:aws:iam::<accountID>:user/gary-test",
    CreateDate: 2022-07-17 19:14:47 +0000 UTC,
    Path: "/",
    UserId: "AIDAWPVARBQGAN7NI7ABY",
    UserName: "gary-test"
  }
}
Success {
  User: {
    Arn: "arn:aws:iam::<accountID>:user/synaptik-dev",
    CreateDate: 2022-07-17 19:14:48 +0000 UTC,
    Path: "/",
    UserId: "AIDAWPVARBQGLXR5CR5A5",
    UserName: "synaptik-dev"
  }
}
Success {
  User: {
    Arn: "arn:aws:iam::<accountID>:user/synaptik-stage",
    CreateDate: 2022-07-17 19:14:48 +0000 UTC,
    Path: "/",
    UserId: "AIDAWPVARBQGIDMKMVEM2",
    UserName: "synaptik-stage"
  }
}
Success {
  User: {
    Arn: "arn:aws:iam::<accountID>:user/synaptik-prod",
    CreateDate: 2022-07-17 19:14:48 +0000 UTC,
    Path: "/",
    UserId: "AIDAWPVARBQGKRRYZWDSO",
    UserName: "synaptik-prod"
  }
}
Success {
  User: {
    Arn: "arn:aws:iam::<accountID>:user/synaptik-admin",
    CreateDate: 2022-07-17 19:14:49 +0000 UTC,
    Path: "/",
    UserId: "AIDAWPVARBQGAQEXYGIWZ",
    UserName: "synaptik-admin"
  }
}
Success {
  User: {
    Arn: "arn:aws:iam::<accountID>:user/synaptik-security",
    CreateDate: 2022-07-17 19:14:49 +0000 UTC,
    Path: "/",
    UserId: "AIDAWPVARBQGKT77SBI3X",
    UserName: "synaptik-security"
  }
}
```

This will create the users that you have defined in a text file, by default the users will have no permissions to do anything straight away
