## What is iam-spy

Wouldn't it be nice if we can answer questions like "Which users in my AWS account have access to a particular service?

Which policy grants access for user X to service Y?

With current AWS IAM API offerings there is no easy way to get these answers in a straight forward way. This tool aims to help in answering such questions without making multiple queries.

It analyses an aws service and lists all users, roles, gtoups or policies that have access to a selected service.

### Credentials

This tool looks for AWS credential in obvious places like environment variables, .aws folder etc. These credentials should have permissions to get a list of IAM entities( users, roles, policies).

### Get it

For Linux
```
wget https://github.com/bit-cloner/iam-spy/releases/download/0.9/iam-spy
chmod +x ./iam-spy
./iam-spy

```
One Liner

```
wget https://github.com/bit-cloner/iam-spy/releases/download/0.9/iam-spy && sudo chmod +x ./iam-spy && ./iam-spy
```
