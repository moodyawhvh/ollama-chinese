<div align="center">

# Ollama 中文翻译版

**[中文版] Ollama — 在你自己的电脑上一键运行 Llama、Qwen、DeepSeek 等开源大模型的本地运行框架**

[![原项目](https://img.shields.io/badge/原项目-ollama--ollama-blue?style=flat-square&logo=github)](https://github.com/ollama/ollama)
[![中文文档](https://img.shields.io/badge/中文文档-README.zh--CN.md-orange?style=flat-square)](README.zh-CN.md)
[![GitHub Stars](https://img.shields.io/github/stars/ollama/ollama?style=flat-square&label=原项目Stars)](https://github.com/ollama/ollama/stargazers)
[![微信联系](https://img.shields.io/badge/微信-uaycar-brightgreen?style=flat-square&logo=wechat)](#)

</div>

---

> 这是 [ollama/ollama](https://github.com/ollama/ollama) 的中文翻译版本。
> 完整源代码请访问原项目:https://github.com/ollama/ollama

**代部署 / 定制服务 / 技术咨询 请添加微信：uaycar**

---

## 📖 项目简介

Ollama 是一个开源的本地大模型运行框架，让你在自己的电脑上就能跑起 Llama、Qwen、DeepSeek、Gemma 等开源大模型，数据全程不出本机。它把模型下载、量化推理和 API 服务压缩成一条命令的事：安装之后 `ollama run <模型名>` 即可开聊，再通过内置 REST API 和官方 Python / JavaScript SDK 把模型能力接进你自己的应用。它还与 Claude Code、Codex、Copilot 等主流编码工具深度集成，配合 OpenClaw 更能变身覆盖 WhatsApp、Telegram、Slack、Discord 的个人 AI 助手。

## ✨ 主要特性

- 🖥️ **跨平台**：macOS、Windows、Linux 一键安装，官方 Docker 镜像开箱即用
- ⚡ **一条命令跑模型**：`ollama run <模型名>` 自动下载并进入对话
- 📚 **模型库丰富**：Llama、Qwen、DeepSeek、Gemma 等开源模型齐聚 [ollama.com/library](https://ollama.com/library)
- 🔌 **内置 REST API**：本地 `localhost:11434` 直接调用，官方 Python / JavaScript SDK 同步支持
- 🛠️ **编码工具集成**：一条命令接入 Claude Code、Codex、Copilot CLI、OpenCode 等编码智能体
- 🤖 **AI 助手**：配合 OpenClaw 可打通 WhatsApp、Telegram、Slack、Discord 等平台
- 🧱 **成熟推理后端**：基于 Georgi Gerganov 发起的 [llama.cpp](https://github.com/ggml-org/llama.cpp) 项目
- 🌍 **社区生态庞大**：Open WebUI、Dify、AnythingLLM 等上百个社区集成可选

## 📁 文件说明

| 文件 | 说明 |
|:-----|:-----|
| README.md | 本文件（中文简介） |
| README.zh-CN.md | 详细中文文档（完整汉化） |

## 🚀 快速开始

**1. 安装**

macOS / Linux：

```shell
curl -fsSL https://ollama.com/install.sh | sh
```

Windows（PowerShell）：

```shell
irm https://ollama.com/install.ps1 | iex
```

Docker：直接使用 Docker Hub 上的官方镜像 [`ollama/ollama`](https://hub.docker.com/r/ollama/ollama)。

**2. 启动**

```
ollama
```

启动后会引导你运行模型，或把 Ollama 接入 Claude Code、OpenClaw、OpenCode、Codex、Copilot 等现有工具。

**3. 与模型对话**

```
ollama run gemma4
```

**4. 调用 REST API**

```
curl http://localhost:11434/api/chat -d '{
  "model": "gemma4",
  "messages": [{
    "role": "user",
    "content": "Why is the sky blue?"
  }],
  "stream": false
}'
```

**5. 在代码中使用（Python / JavaScript）**

```
pip install ollama
```

```
npm i ollama
```

完整源代码与最新版本请访问原项目：https://github.com/ollama/ollama

## 📞 联系方式

**代部署 / 定制服务 / 技术咨询 请添加微信：uaycar**

---

本项目为 [ollama/ollama](https://github.com/ollama/ollama) 的中文翻译版本，所有代码版权归原项目作者所有，遵循原项目许可证（MIT）。

**代部署 / 定制服务 / 技术咨询 请添加微信：uaycar**

**如果觉得有用，请给原项目点个 Star！** ⭐
