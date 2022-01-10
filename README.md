# model-identity-go

The main goal with this repo is to share common data transfer object / structs definitions for other repos of the fuel traceability project.

_technology_

![](https://img.shields.io/badge/Tech-Golang-informational?style=flat&logo=Go&logoColor=white)
![](https://img.shields.io/badge/Tech-Json-informational?style=flat&logo=JSON&logoColor=white)
![](https://img.shields.io/badge/Tech-Blockchain-informational?style=flat&logo=Blockchain.com&logoColor=white)
![](https://img.shields.io/badge/Tech-Hyperledger-informational?style=flat&logo=Hyperledger&logoColor=white)

_misc_

![](https://img.shields.io/badge/build-passing-brightgreen?style=flat)
![](https://img.shields.io/badge/release-v0.0.0-inactive?style=flat)
![](https://img.shields.io/badge/coverage-90%25-green?style=flat)
![](https://img.shields.io/badge/reposize-0MB-inactive?style=flat)


## 📑 Including "model-identity-go" as go module into another project's source code.
Since the "model-identity-go" module is in the "**ic-matcom**" private github repository, several previous steps
must be carried out.

#### 1. Set GOPRIVATE in a environment like below: 

```bash
go env -w GOPRIVATE=github.com/ic-matcom
```
❗ Multiple values are separated by a comma.

#### 2. Passing repository credentials to Go module
```bash
git config --global url."https://username:accessToken_or_Password@github.com".insteadOf "https://github.com"
```
❗ If your user has any @ change to %40. This type of special strings must be url encoded.

#### 3. Finally, we import the module
```bash
go get github.com/ic-matcom/model-identity-go
```
