#ThreatCmd
---

[![codebeat badge](https://codebeat.co/badges/022d0311-1c2a-4594-ab59-091d052c61b9)](https://codebeat.co/projects/github-com-jheise-threatcmd)

CLI for ThreatCrowd.org

Available Functions

- email - query for the given email address

- ip - query for the given ip address

- domain - query for the given domain

- av - query for antivirus data

- file - query for md5 hash

---

Example

```
threatcmd ip 4.2.2.1

threatcmd domain google.com

threatcmd email fake@mail.com

threatcmd av plugx

threatcmd file 45b7fb10a4f9aebe85f2c537b33cc27c
```
