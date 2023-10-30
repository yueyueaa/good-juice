# good-juice
Qiniu's code competition

开发流程：

1. 从主仓库 fork 一份到个人仓库
2. git clone 个人仓库的ssh链接, 添加远程主仓库

```bash
git clone git@github.com:xxxxx/good-juice.git 
git remote add upstream git@github.com:yueyueaa/good-juice.git # 添加远程主仓库
git remote -v # 查看远程关联仓库
```

1. 创建新分支，切换到新分支
2. 需求开发、编写UT，本地测试
3. git add .  git commit -m “xxx” git  push origin branch_name
4. 提交PR，branch_name → main
5. 至少1人review后合入master分支

注意：每次开发新功能前先在本地切换到main分支，拉取远程主仓库main分支最新代码，避免冲突。