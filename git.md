


          
# Git 使用指南

## 目录
1. [Git 基础配置](#git-基础配置)
2. [Git 基本操作](#git-基本操作)
3. [Git 分支管理](#git-分支管理)
4. [GitHub 仓库管理](#github-仓库管理)
5. [版本控制与历史查看](#版本控制与历史查看)
6. [常见问题解决](#常见问题解决)
7. [快捷操作](#快捷操作)

---

## Git 基础配置

### 配置用户信息
```bash
# 配置全局用户名和邮箱
git config --global user.name "你的GitHub用户名"
git config --global user.email "你的GitHub注册邮箱"

# 查看当前配置
git config --list
```

**重要说明：**
- 用户名应使用 GitHub 的用户名
- 邮箱应使用 GitHub 注册时的邮箱
- 这样配置可以确保提交记录正确关联到你的 GitHub 账户

### 初始化仓库
```bash
# 在项目目录中初始化 Git 仓库
git init

# 查看仓库状态
git status
```

---

## Git 基本操作

### 添加和提交文件
```bash
# 添加所有文件到暂存区
git add .

# 添加特定文件
git add filename.txt

# 提交更改
git commit -m "提交信息"

# 查看提交历史
git log
git log --oneline  # 简洁格式
```

### 常见错误解决
**问题：** `nothing added to commit but untracked files present`

**原因：** 文件未添加到暂存区

**解决方案：**
```bash
# 方案1：添加所有文件
git add .
git commit -m "first commit"

# 方案2：选择性添加
git add 文件名
git commit -m "提交信息"

# 方案3：先查看状态再决定
git status
```

### 连接远程仓库
```bash
# 添加远程仓库
git remote add origin https://github.com/用户名/仓库名.git

# 推送到远程仓库
git push -u origin main

# 后续推送
git push
```

---

## Git 分支管理

### 分支的作用
- **隔离开发**：不同功能在不同分支开发，互不影响
- **并行开发**：多人可以同时在不同分支工作
- **版本管理**：便于管理不同版本和功能
- **安全性**：保护主分支稳定性

### 常见分支类型
| 分支类型 | 用途 | 示例名称 |
|---------|------|----------|
| 主分支 | 稳定的生产代码 | `main`, `master` |
| 开发分支 | 日常开发 | `develop`, `dev` |
| 功能分支 | 新功能开发 | `feature/login`, `feature/payment` |
| 修复分支 | Bug修复 | `hotfix/bug-123`, `fix/login-error` |
| 发布分支 | 版本发布准备 | `release/v1.0.0` |

### 分支操作命令
```bash
# 查看所有分支
git branch
git branch -a  # 包括远程分支

# 创建新分支
git branch 分支名

# 切换分支
git checkout 分支名

# 创建并切换到新分支
git checkout -b 分支名

# 合并分支
git checkout main  # 切换到主分支
git merge 功能分支名  # 合并功能分支

# 删除分支
git branch -d 分支名  # 删除已合并的分支
git branch -D 分支名  # 强制删除分支
```

### 分支合并详解

#### 合并类型
1. **Fast-Forward 合并**（当主分支无新提交时）
```bash
git checkout main
git merge feature-branch
```

2. **Three-Way 合并**（需要创建合并提交）
```bash
git checkout main
git merge feature-branch
# 会创建一个新的合并提交
```

#### 合并后的同步
- **功能分支的所有修改都会同步到主分支**
- **包括代码变更、提交历史和文件操作**
- **合并后可以安全删除功能分支**

#### 合并验证
```bash
# 查看合并后的历史
git log --graph --oneline

# 查看文件差异
git diff HEAD~1  # 与上一次提交比较

# 查看分支合并情况
git show --stat
```

### 实际工作流程示例
```bash
# 1. 从主分支创建功能分支
git checkout main
git pull origin main  # 确保主分支是最新的
git checkout -b feature/user-login

# 2. 在功能分支开发
# ... 编写代码 ...
git add .
git commit -m "实现用户登录功能"

# 3. 推送功能分支
git push origin feature/user-login

# 4. 合并到主分支
git checkout main
git pull origin main  # 确保主分支最新
git merge feature/user-login

# 5. 推送合并结果
git push origin main

# 6. 清理功能分支
git branch -d feature/user-login
git push origin --delete feature/user-login
```

---

## GitHub 仓库管理

### 创建仓库配置建议

**基于你的 review-service 项目：**
- **仓库名称：** `review-service`
- **描述：** "A review service built with Go and Kratos framework"
- **可见性：** 选择 Private（私有）
- **模板：** 不选择任何模板
- **README：** 不勾选（项目中已存在）
- **.gitignore：** 不勾选（项目中已存在完善的 .gitignore）
- **许可证：** 可选择 "No license" 或 "MIT License"

### 仓库可见性管理

#### 设置为私有仓库
1. 进入仓库设置页面：`https://github.com/用户名/仓库名/settings`
2. 滚动到页面底部找到 "Danger Zone"
3. 点击 "Change repository visibility"
4. 选择 "Make private" 并确认

**注意事项：**
- GitHub 免费账户可能有私有仓库数量限制
- 私有仓库只有你和授权用户可以访问
- 可以随时在公开和私有之间切换

#### 删除仓库
1. 进入仓库设置页面
2. 滚动到页面底部的 "Danger Zone"
3. 点击 "Delete this repository"
4. 输入仓库名确认删除

**警告：** 删除操作不可逆，请谨慎操作！

#### 替代方案
- **归档仓库：** 保留但标记为只读
- **重命名仓库：** 改变仓库名称
- **设为私有：** 隐藏但保留访问权限

### 查看 GitHub 账户信息

#### 查看登录邮箱
1. **设置页面：** https://github.com/settings/emails
2. **个人资料：** https://github.com/settings/profile（如果设置了公开）
3. **账户安全：** https://github.com/settings/security

**注意事项：**
- 主邮箱用于重要通知
- 可以添加多个备用邮箱
- 注意邮箱验证状态
- 考虑隐私设置

---

## 版本控制与历史查看

### 获取历史版本

#### 克隆整个仓库后切换版本
```bash
# 1. 克隆仓库
git clone https://github.com/username/repository.git
cd repository

# 2. 查看提交历史
git log --oneline

# 3. 切换到特定版本
git checkout <commit-hash>

# 4. 基于历史版本创建新分支
git checkout -b new-branch-name <commit-hash>
```

#### 克隆特定分支或标签
```bash
# 克隆特定分支
git clone -b branch-name https://github.com/username/repository.git

# 克隆特定标签版本
git clone -b v1.0.0 https://github.com/username/repository.git

# 浅克隆（只获取最新版本）
git clone --depth 1 https://github.com/username/repository.git
```

### GitHub 网页查看提交历史

#### 查看提交历史
1. 进入仓库主页
2. 点击提交数量（如 "23 commits"）
3. 浏览所有提交记录

#### 查看单个提交详情
1. 点击具体的提交记录
2. 查看文件变更统计
3. 查看代码差异对比

#### 高级查看功能
- **分屏视图 vs 统一视图**
- **按文件筛选变更**
- **忽略空白字符差异**
- **查看原始差异**
- **行级注释和讨论**

### 命令行查看历史
```bash
# 查看提交历史
git log
git log --oneline  # 简洁格式
git log --graph    # 图形化显示

# 查看特定文件的历史
git log -- filename.txt

# 查看两个版本之间的差异
git diff commit1..commit2

# 查看特定提交的详情
git show <commit-hash>
```

---

## 常见问题解决

### 1. 提交时出现 "nothing to commit"
**原因：** 文件未添加到暂存区
```bash
git add .
git commit -m "提交信息"
```

### 2. 推送时出现权限错误
**原因：** 认证问题或仓库权限不足
```bash
# 检查远程仓库地址
git remote -v

# 重新设置远程仓库
git remote set-url origin https://github.com/用户名/仓库名.git
```

### 3. 合并冲突
**解决步骤：**
```bash
# 1. 查看冲突文件
git status

# 2. 手动编辑冲突文件，解决冲突标记
# <<<<<<< HEAD
# 当前分支的内容
# =======
# 合并分支的内容
# >>>>>>> branch-name

# 3. 添加解决后的文件
git add 冲突文件名

# 4. 完成合并
git commit -m "解决合并冲突"
```

### 4. 撤销操作
```bash
# 撤销工作区修改
git checkout -- filename.txt

# 撤销暂存区文件
git reset HEAD filename.txt

# 撤销最后一次提交（保留修改）
git reset --soft HEAD~1

# 撤销最后一次提交（丢弃修改）
git reset --hard HEAD~1
```

---

## 最佳实践建议

### 提交信息规范
```bash
# 好的提交信息示例
git commit -m "feat: 添加用户登录功能"
git commit -m "fix: 修复登录页面样式问题"
git commit -m "docs: 更新README文档"
```

### 分支命名规范
- `feature/功能名称` - 新功能开发
- `fix/问题描述` - Bug修复
- `hotfix/紧急修复` - 紧急修复
- `release/版本号` - 版本发布

### 工作流程建议
1. **经常提交**：小步快跑，频繁提交
2. **有意义的提交信息**：清楚描述本次修改
3. **使用分支**：不要直接在主分支开发
4. **合并前测试**：确保代码质量
5. **定期同步**：及时拉取远程更新

---

## 快捷操作
```bash
# 进入项目目录
cd e:\777\2025summer\review_service\review_service

# 检查当前状态
git status

# 添加所有文件
git add .

# 提交代码
git commit -m "commit X, XXXXXX "

# 推送到远程仓库
git push -u origin main
```



## 总结

Git 是强大的版本控制工具，掌握基本操作后可以：
- 安全地管理代码版本
- 高效地进行团队协作
- 轻松地回溯和恢复历史版本
- 灵活地管理不同功能分支

记住：**每个提交都是项目的一个快照，Git 保存了完整的项目历史，任何版本都可以被恢复和访问。**
        