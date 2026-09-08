> 🌐 本文档由 [ollama/ollama](https://github.com/ollama/ollama) 翻译,英文原版见原项目。

# 参与贡献 Ollama

感谢你有兴趣为 Ollama 做贡献!下面是一些帮助你上手 contributor(贡献者)工作的指引。

## 环境准备

请参阅[开发文档](./docs/development.md),了解如何在本地构建和运行 Ollama。

### 受欢迎的 issue

* [Bug](https://github.com/ollama/ollama/issues?q=is%3Aissue+is%3Aopen+label%3Abug):Ollama 停止工作或出现意外错误的问题。
* [性能](https://github.com/ollama/ollama/issues?q=is%3Aissue+is%3Aopen+label%3Aperformance):让 Ollama 在模型推理、下载或上传方面更快的问题。
* [安全](https://github.com/ollama/ollama/blob/main/SECURITY.md):可能引发安全漏洞的问题。如 [SECURITY.md](https://github.com/ollama/ollama/blob/main/SECURITY.md) 所述,请勿公开披露安全漏洞。

### 较难审查的贡献

* 新功能:新功能(例如 API 字段、环境变量)会增加 Ollama 的表面积,长期维护难度更大,因为将来一旦移除它们,可能会破坏现有用户的使用。
* 重构:大型代码改进固然重要,但审查和合并起来可能更困难、更耗时。
* 文档:补充或修正缺失文档的小幅更新很有帮助,但大规模的文档新增在长期维护上会很吃力。

### 可能不会被接受的改动

* 破坏 Ollama API(包括 OpenAI 兼容 API)向后兼容性的改动
* 给用户体验带来明显阻力的改动
* 给维护者和贡献者带来巨大未来维护负担的改动

## 提议一个(非琐碎的)改动

> 所谓"非琐碎",指的是不属于 bug 修复或小幅文档更新的改动。如果你不确定,请到我们的 [Discord 服务器](https://discord.gg/ollama)上询问。

在发起非琐碎的 Pull Request 之前,请先开一个 issue 讨论该改动并听取维护者的反馈。这能帮助我们理解改动的背景、以及它如何契合 Ollama 的路线图,避免重复劳动,也免得你把时间花在一个我们可能无法接受的改动上。

提案小技巧:

* 说明你要解决的问题,而不是你想做的事情。
* 说明这个改动为什么重要。
* 说明这个改动将如何被使用。
* 说明这个改动将如何被测试。

另外,加分项:提供一份你预期在该改动被接受后会出现的文档草稿。

## Pull Request

**Commit 提交信息**

标题格式应为:

    <包名>: <简短描述>

包名指受影响最大的 Go 包。如果改动不涉及 Go 代码,则改用目录名。针对根目录下某个知名单文件的改动,可以使用文件名。

简短描述应以小写字母开头,并能接续下面这句话:

      "This changes Ollama to..."(这个改动让 Ollama ……)

示例:

      llm/backend/mlx: support the llama architecture
      CONTRIBUTING: provide clarity on good commit messages, and bad

反面示例:

      feat: add more emoji
      fix: was not using famous web framework
      chore: generify code

**测试**

请附带测试。尽量测试行为,而不是测试实现细节。

**新增依赖**

依赖应当克制地添加。如果你要引入新依赖,请说明它的必要性,以及你尝试过哪些不依赖它也能行得通的替代方案。

## 需要帮助?

有任何问题,欢迎到我们的 [Discord 服务器](https://discord.gg/ollama)上联系我们。
