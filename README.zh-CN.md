<div align="center">
  <a href="https://ollama.com">
    <img src="https://github.com/ollama/ollama/assets/3325447/0d0b44e2-8f4a-4e99-9b52-a5c1c741c8f7" alt="ollama" width="200"/>
  </a>

# Ollama 中文文档

**[中文翻译版] 用开源模型开始构建 —— 在本地一键运行 Llama、Qwen、DeepSeek 等开源大模型**

[![原项目](https://img.shields.io/badge/原项目-ollama--ollama-blue?style=flat-square&logo=github)](https://github.com/ollama/ollama)
[![微信联系](https://img.shields.io/badge/微信-uaycar-brightgreen?style=flat-square&logo=wechat)](#)

</div>

> 📌 本文档是 [ollama/ollama](https://github.com/ollama/ollama) 官方 README 的中文翻译版本，仅供学习参考，内容以英文原版为准。
> 完整源代码与最新版本请访问原项目：https://github.com/ollama/ollama

**代部署 / 定制服务 / 技术咨询 请添加微信：uaycar**

---

## 📖 项目简介

Ollama 让你可以用开源模型开始构建。它在本地托管并运行 Llama、Qwen、DeepSeek、Gemma 等开源大模型，一条命令即可下载运行，并通过 REST API 和 Python / JavaScript SDK 对外提供服务。它还能与 Claude Code、Codex、Copilot 等主流编码智能体一键集成，配合 OpenClaw 可以变身覆盖多个聊天平台的个人 AI 助手。

## 📥 下载安装

### macOS

```shell
curl -fsSL https://ollama.com/install.sh | sh
```

或者[手动下载](https://ollama.com/download/Ollama.dmg)安装包。

### Windows

```shell
irm https://ollama.com/install.ps1 | iex
```

或者[手动下载](https://ollama.com/download/OllamaSetup.exe)安装包。

### Linux

```shell
curl -fsSL https://ollama.com/install.sh | sh
```

[手动安装说明](https://docs.ollama.com/linux#manual-install)

### Docker

官方 [Ollama Docker 镜像](https://hub.docker.com/r/ollama/ollama) `ollama/ollama` 已发布在 Docker Hub。

### 官方库

- [ollama-python](https://github.com/ollama/ollama-python)（Python SDK）
- [ollama-js](https://github.com/ollama/ollama-js)（JavaScript SDK）

### 社区

- [Discord](https://discord.gg/ollama)
- [𝕏 (Twitter)](https://x.com/ollama)
- [Reddit](https://reddit.com/r/ollama)

## 🚀 快速开始

```
ollama
```

启动后，程序会引导你运行一个模型，或者把 Ollama 接入你现有的智能体与应用，例如 `Claude Code`、`OpenClaw`、`OpenCode`、`Codex`、`Copilot` 等。

### 编码

启动指定集成：

```
ollama launch claude
```

支持的集成包括 [Claude Code](https://docs.ollama.com/integrations/claude-code)、[Codex](https://docs.ollama.com/integrations/codex)、[Copilot CLI](https://docs.ollama.com/integrations/copilot-cli)、[DeepSeek Harness](https://docs.ollama.com/integrations/deepseek-harness)、[Droid](https://docs.ollama.com/integrations/droid) 和 [OpenCode](https://docs.ollama.com/integrations/opencode)。

### AI 助手

使用 [OpenClaw](https://docs.ollama.com/integrations/openclaw) 可以把 Ollama 变成覆盖 WhatsApp、Telegram、Slack、Discord 等平台的个人 AI 助手：

```
ollama launch openclaw
```

### 与模型对话

运行并与 [Gemma 4](https://ollama.com/library/gemma4) 对话：

```
ollama run gemma4
```

完整模型列表见 [ollama.com/library](https://ollama.com/library)。

更多细节请参阅[快速入门指南](https://docs.ollama.com/quickstart)。

## 📚 模型库

Ollama 的所有模型统一托管在 [ollama.com/library](https://ollama.com/library)，收录了 Llama、Qwen、DeepSeek、Gemma 等主流开源模型。在模型页找到名称后，执行 `ollama run <模型名>` 即可自动下载并进入对话，例如 `ollama run gemma4`。完整列表以[模型库](https://ollama.com/library)为准。

## 🔌 REST API

Ollama 内置一套用于运行和管理模型的 REST API：

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

全部端点请查看 [API 文档](https://docs.ollama.com/api)。

### Python

```
pip install ollama
```

```python
from ollama import chat

response = chat(model='gemma4', messages=[
  {
    'role': 'user',
    'content': 'Why is the sky blue?',
  },
])
print(response.message.content)
```

### JavaScript

```
npm i ollama
```

```javascript
import ollama from "ollama";

const response = await ollama.chat({
  model: "gemma4",
  messages: [{ role: "user", content: "Why is the sky blue?" }],
});
console.log(response.message.content);
```

## 🧱 支持的后端

- [llama.cpp](https://github.com/ggml-org/llama.cpp) —— 由 Georgi Gerganov 发起的项目。

## 📖 文档

- [CLI 参考](https://docs.ollama.com/cli)
- [REST API 参考](https://docs.ollama.com/api)
- [导入模型](https://docs.ollama.com/import)
- [Modelfile 参考](https://docs.ollama.com/modelfile)
- [从源码构建](https://github.com/ollama/ollama/blob/main/docs/development.md)

## ❓ 常见问题

**Ollama 支持哪些操作系统？**
支持 macOS、Windows 和 Linux，也可以直接用官方 Docker 镜像部署。

**在哪里查看可用的模型？**
完整模型列表见 [ollama.com/library](https://ollama.com/library)。

**如何把 Ollama 接入我现有的应用或智能体？**
通过本地 REST API（`http://localhost:11434`）或官方 Python / JavaScript SDK 接入；也可以直接使用 `ollama launch` 接入 Claude Code、Codex、Copilot 等现成集成。

**遇到问题在哪里反馈？**
请到原项目 [ollama/ollama](https://github.com/ollama/ollama) 提交 Issue，或加入官方 [Discord](https://discord.gg/ollama) 社区。

## 🌐 社区集成

> 想把你的项目加进来？欢迎提交 Pull Request。

### 聊天界面

#### Web

- [Open WebUI](https://github.com/open-webui/open-webui) - 可扩展、可自托管的 AI 界面
- [Onyx](https://github.com/onyx-dot-app/onyx) - 互联式 AI 工作区
- [LibreChat](https://github.com/danny-avila/LibreChat) - 增强版 ChatGPT 克隆，支持多提供商
- [Lobe Chat](https://github.com/lobehub/lobe-chat) - 带插件生态的现代聊天框架（[文档](https://lobehub.com/docs/self-hosting/examples/ollama)）
- [NextChat](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web) - 跨平台 ChatGPT UI（[文档](https://docs.nextchat.dev/models/ollama)）
- [Perplexica](https://github.com/ItzCrazyKns/Perplexica) - AI 驱动的搜索引擎，开源版 Perplexity 替代品
- [big-AGI](https://github.com/enricoros/big-AGI) - 面向专业人士的 AI 套件
- [Lollms WebUI](https://github.com/ParisNeo/lollms-webui) - 多模型 Web 界面
- [ChatOllama](https://github.com/sugarforever/chat-ollama) - 带知识库的聊天机器人
- [Bionic GPT](https://github.com/bionic-gpt/bionic-gpt) - 本地部署的 AI 平台
- [Chatbot UI](https://github.com/ivanfioravanti/chatbot-ollama) - ChatGPT 风格的 Web 界面
- [Hollama](https://github.com/fmaclen/hollama) - 极简 Web 界面
- [Chatbox](https://github.com/Bin-Huang/Chatbox) - 桌面与 Web AI 客户端
- [chat](https://github.com/swuecho/chat) - 面向团队的聊天 Web 应用
- [Ollama RAG Chatbot](https://github.com/datvodinh/rag-chatbot.git) - 用 RAG 与多个 PDF 对话
- [Tkinter 客户端](https://github.com/chyok/ollama-gui) - Python 桌面客户端

#### 桌面

- [Dify.AI](https://github.com/langgenius/dify) - LLM 应用开发平台
- [AnythingLLM](https://github.com/Mintplex-Labs/anything-llm) - Mac、Windows、Linux 全能 AI 应用
- [Maid](https://github.com/Mobile-Artificial-Intelligence/maid) - 跨平台移动与桌面客户端
- [Witsy](https://github.com/nbonamy/witsy) - Mac、Windows、Linux AI 桌面应用
- [Cherry Studio](https://github.com/kangfenmao/cherry-studio) - 多提供商桌面客户端
- [Ollama App](https://github.com/JHubi1/ollama-app) - 桌面与移动多平台客户端
- [PyGPT](https://github.com/szczyglis-dev/py-gpt) - Linux、Windows、Mac 的 AI 桌面助手
- [Alpaca](https://github.com/Jeffser/Alpaca) - Linux 和 macOS 的 GTK4 客户端
- [SwiftChat](https://github.com/aws-samples/swift-chat) - 跨平台客户端，含 iOS、Android 与 Apple Vision Pro
- [Enchanted](https://github.com/AugustDev/Enchanted) - 原生 macOS 与 iOS 客户端
- [RWKV-Runner](https://github.com/josStorer/RWKV-Runner) - 多模型桌面运行器
- [Ollama Grid Search](https://github.com/dezoito/ollama-grid-search) - 评估与比较模型
- [macai](https://github.com/Renset/macai) - 面向 Ollama 和 ChatGPT 的 macOS 客户端
- [AI Studio](https://github.com/MindWorkAI/AI-Studio) - 多提供商桌面 IDE
- [Reins](https://github.com/ibrahimcetin/reins) - 参数调优与推理模型支持
- [ConfiChat](https://github.com/1runeberg/confichat) - 注重隐私、支持可选加密
- [LLocal.in](https://github.com/kartikm7/llocal) - Electron 桌面客户端
- [MindMac](https://mindmac.app) - Mac 的 AI 聊天客户端
- [Msty](https://msty.app) - 多模型桌面客户端
- [BoltAI for Mac](https://boltai.com) - Mac 的 AI 聊天客户端
- [IntelliBar](https://intellibar.app/) - macOS 的 AI 助手
- [Kerlig AI](https://www.kerlig.com/) - macOS 的 AI 写作助手
- [Hillnote](https://hillnote.com) - Markdown 优先的 AI 工作区
- [Perfect Memory AI](https://www.perfectmemory.ai/) - 结合屏幕与会议记录的个性化效率 AI

#### 移动

- [Ollama Android Chat](https://github.com/sunshine0523/OllamaServer) - 在 Android 上一键使用 Ollama

> 上文中的 SwiftChat、Enchanted、Maid、Ollama App、Reins 和 ConfiChat 同样支持移动平台。

### 代码编辑器与开发工具

- [Cline](https://github.com/cline/cline) - VS Code 扩展，支持多文件/整仓编码
- [Continue](https://github.com/continuedev/continue) - 面向任意 IDE 的开源 AI 代码助手
- [Void](https://github.com/voideditor/void) - 开源 AI 代码编辑器，Cursor 替代品
- [Copilot for Obsidian](https://github.com/logancyang/obsidian-copilot) - Obsidian 的 AI 助手
- [twinny](https://github.com/rjmacarthy/twinny) - Copilot 与 Copilot 聊天替代品
- [gptel Emacs 客户端](https://github.com/karthink/gptel) - Emacs 的 LLM 客户端
- [Ollama Copilot](https://github.com/bernardo-bruning/ollama-copilot) - 把 Ollama 当作 GitHub Copilot 用
- [Obsidian Local GPT](https://github.com/pfrankov/obsidian-local-gpt) - Obsidian 的本地 AI
- [Ellama Emacs 客户端](https://github.com/s-kostyaev/ellama) - Emacs 的 LLM 工具
- [orbiton](https://github.com/xyproto/orbiton) - 零配置文本编辑器，带 Ollama 补全
- [AI ST Completion](https://github.com/yaroslavyaroslav/OpenAI-sublime-text) - Sublime Text 4 AI 助手
- [VT Code](https://github.com/vinhnx/vtcode) - 基于 Rust 的终端编码智能体，使用 Tree-sitter
- [QodeAssist](https://github.com/Palm1r/QodeAssist) - Qt Creator 的 AI 编码助手
- [AI Toolkit for VS Code](https://aka.ms/ai-tooklit/ollama-docs) - 微软官方 VS Code 扩展
- [Open Interpreter](https://docs.openinterpreter.com/language-model-setup/local-models/ollama) - 用自然语言操作计算机

### 库与 SDK

- [LiteLLM](https://github.com/BerriAI/litellm) - 统一 100+ LLM 提供商的 API
- [Semantic Kernel](https://github.com/microsoft/semantic-kernel/tree/main/python/semantic_kernel/connectors/ai/ollama) - 微软 AI 编排 SDK
- [LangChain4j](https://github.com/langchain4j/langchain4j) - Java 版 LangChain（[示例](https://github.com/langchain4j/langchain4j-examples/tree/main/ollama-examples/src/main/java)）
- [LangChainGo](https://github.com/tmc/langchaingo/) - Go 版 LangChain（[示例](https://github.com/tmc/langchaingo/tree/main/examples/ollama-completion-example)）
- [Spring AI](https://github.com/spring-projects/spring-ai) - Spring 框架的 AI 支持（[文档](https://docs.spring.io/spring-ai/reference/api/chat/ollama-chat.html)）
- [LangChain](https://python.langchain.com/docs/integrations/chat/ollama/) 与 [LangChain.js](https://js.langchain.com/docs/integrations/chat/ollama/)（[示例](https://js.langchain.com/docs/tutorials/local_rag/)）
- [Ollama for Ruby](https://github.com/crmne/ruby_llm) - Ruby LLM 库
- [any-llm](https://github.com/mozilla-ai/any-llm) - Mozilla 出品的统一 LLM 接口
- [OllamaSharp for .NET](https://github.com/awaescher/OllamaSharp) - .NET SDK
- [LangChainRust](https://github.com/Abraxas-365/langchain-rust) - Rust 版 LangChain（[示例](https://github.com/Abraxas-365/langchain-rust/blob/main/examples/llm_ollama.rs)）
- [Agents-Flex for Java](https://github.com/agents-flex/agents-flex) - Java 智能体框架（[示例](https://github.com/agents-flex/agents-flex/tree/main/agents-flex-llm/agents-flex-llm-ollama/src/test/java/com/agentsflex/llm/ollama)）
- [Elixir LangChain](https://github.com/brainlid/langchain) - Elixir 版 LangChain
- [Ollama-rs for Rust](https://github.com/pepperoni21/ollama-rs) - Rust SDK
- [LangChain for .NET](https://github.com/tryAGI/LangChain) - .NET 版 LangChain（[示例](https://github.com/tryAGI/LangChain/blob/main/examples/LangChain.Samples.OpenAI/Program.cs)）
- [chromem-go](https://github.com/philippgille/chromem-go) - 支持 Ollama 嵌入的 Go 向量数据库（[示例](https://github.com/philippgille/chromem-go/tree/v0.5.0/examples/rag-wikipedia-ollama)）
- [LangChainDart](https://github.com/davidmigloz/langchain_dart) - Dart 版 LangChain
- [LlmTornado](https://github.com/lofcz/llmtornado) - 多推理 API 的统一 C# 接口
- [Ollama4j for Java](https://github.com/ollama4j/ollama4j) - Java SDK
- [Ollama for Laravel](https://github.com/cloudstudio/ollama-laravel) - Laravel 集成
- [Ollama for Swift](https://github.com/mattt/ollama-swift) - Swift SDK
- [LlamaIndex](https://docs.llamaindex.ai/en/stable/examples/llm/ollama/) 与 [LlamaIndexTS](https://ts.llamaindex.ai/modules/llms/available_llms/ollama) - LLM 应用数据框架
- [Haystack](https://github.com/deepset-ai/haystack-integrations/blob/main/integrations/ollama.md) - AI 流水线框架
- [Firebase Genkit](https://firebase.google.com/docs/genkit/plugins/ollama) - Google AI 框架
- [Ollama-hpp for C++](https://github.com/jmont-dev/ollama-hpp) - C++ SDK
- [PromptingTools.jl](https://github.com/svilupp/PromptingTools.jl) - Julia LLM 工具包（[示例](https://svilupp.github.io/PromptingTools.jl/dev/examples/working_with_ollama)）
- [Ollama for R - rollama](https://github.com/JBGruber/rollama) - R SDK
- [Portkey](https://portkey.ai/docs/welcome/integration-guides/ollama) - AI 网关
- [Testcontainers](https://testcontainers.com/modules/ollama/) - 基于容器的测试
- [LLPhant](https://github.com/theodo-group/LLPhant?tab=readme-ov-file#ollama) - PHP AI 框架

### 框架与智能体

- [AutoGPT](https://github.com/Significant-Gravitas/AutoGPT/blob/master/docs/content/platform/ollama.md) - 自主 AI 智能体平台
- [crewAI](https://github.com/crewAIInc/crewAI) - 多智能体编排框架
- [Strands Agents](https://github.com/strands-agents/sdk-python) - AWS 出品的模型驱动智能体构建 SDK
- [Cheshire Cat](https://github.com/cheshire-cat-ai/core) - AI 助手框架
- [any-agent](https://github.com/mozilla-ai/any-agent) - Mozilla 出品的统一智能体框架接口
- [Stakpak](https://github.com/stakpak/agent) - 开源 DevOps 智能体
- [Hexabot](https://github.com/hexastack/hexabot) - 对话式 AI 构建器
- [Neuro SAN](https://github.com/cognizant-ai-lab/neuro-san-studio) - 多智能体编排（[文档](https://github.com/cognizant-ai-lab/neuro-san-studio/blob/main/docs/user_guide.md#ollama)）

### RAG 与知识库

- [RAGFlow](https://github.com/infiniflow/ragflow) - 基于深度文档理解的 RAG 引擎
- [R2R](https://github.com/SciPhi-AI/R2R) - 开源 RAG 引擎
- [MaxKB](https://github.com/1Panel-dev/MaxKB/) - 开箱即用的 RAG 聊天机器人
- [Minima](https://github.com/dmayboroda/minima) - 本地部署或完全离线的 RAG
- [Chipper](https://github.com/TilmanGriesel/chipper) - 结合 Haystack RAG 的 AI 界面
- [ARGO](https://github.com/xark-argo/argo) - Mac/Windows/Linux 上的 RAG 与深度研究
- [Archyve](https://github.com/nickthecook/archyve) - 支持文档库的 RAG 工具
- [Casibase](https://casibase.org) - 带 RAG 与 SSO 的 AI 知识库
- [BrainSoup](https://www.nurgo-software.com/products/brainsoup) - 原生客户端，支持 RAG 与多智能体自动化

### 机器人与消息

- [LangBot](https://github.com/RockChinQ/LangBot) - 多平台消息机器人，支持智能体与 RAG
- [AstrBot](https://github.com/Soulter/AstrBot/) - 多平台聊天机器人，支持 RAG 与插件
- [Discord-Ollama Chat Bot](https://github.com/kevinthedang/discord-ollama) - TypeScript 版 Discord 机器人
- [Ollama Telegram Bot](https://github.com/ruecat/ollama-telegram) - Telegram 机器人
- [LLM Telegram Bot](https://github.com/innightwolfsleep/llm_telegram_bot) - 角色扮演用 Telegram 机器人

### 终端与 CLI

- [aichat](https://github.com/sigoden/aichat) - 全能 LLM CLI，带 Shell 助手、RAG 与 AI 工具
- [oterm](https://github.com/ggozad/oterm) - Ollama 终端客户端
- [gollama](https://github.com/sammcj/gollama) - 基于 Go 的 Ollama 模型管理器
- [tlm](https://github.com/yusufcanb/tlm) - 本地 shell copilot
- [tenere](https://github.com/pythops/tenere) - LLM 的 TUI 界面
- [ParLlama](https://github.com/paulrobello/parllama) - Ollama 的 TUI 界面
- [llm-ollama](https://github.com/taketwo/llm-ollama) - [Datasette 的 LLM CLI](https://llm.datasette.io/en/stable/) 插件
- [ShellOracle](https://github.com/djcopley/ShellOracle) - Shell 命令建议
- [LLM-X](https://github.com/mrdjohnson/llm-x) - LLM 渐进式 Web 应用
- [cmdh](https://github.com/pgibler/cmdh) - 自然语言转 Shell 命令
- [VT](https://github.com/vinhnx/vt.ai) - 极简多模态 AI 聊天应用

### 生产力与应用

- [AppFlowy](https://github.com/AppFlowy-IO/AppFlowy) - AI 协作工作区，可自托管的 Notion 替代品
- [Screenpipe](https://github.com/mediar-ai/screenpipe) - 7×24 小时屏幕与麦克风录制，带 AI 搜索
- [Vibe](https://github.com/thewh1teagle/vibe) - 转写并分析会议内容
- [Page Assist](https://github.com/n4ze3m/page-assist) - AI 上网助手 Chrome 扩展
- [NativeMind](https://github.com/NativeMindBrowser/NativeMindExtension) - 设备端隐私浏览器 AI 助手
- [Ollama Fortress](https://github.com/ParisNeo/ollama_proxy_server) - Ollama 安全代理
- [1Panel](https://github.com/1Panel-dev/1Panel/) - 基于 Web 的 Linux 服务器管理面板
- [Writeopia](https://github.com/Writeopia/Writeopia) - 集成 Ollama 的文本编辑器
- [QA-Pilot](https://github.com/reid41/QA-Pilot) - GitHub 代码仓库理解
- [Raycast 扩展](https://github.com/MassimilianoPasquini97/raycast_ollama) - 在 Raycast 中使用 Ollama
- [Painting Droid](https://github.com/mateuszmigas/painting-droid) - 带 AI 集成的绘画应用
- [Serene Pub](https://github.com/doolijb/serene-pub) - AI 角色扮演应用
- [Mayan EDMS](https://gitlab.com/mayan-edms/mayan-edms) - 集成 Ollama 工作流的文档管理
- [TagSpaces](https://www.tagspaces.org) - 支持 [AI 标签](https://docs.tagspaces.org/ai/)的文件管理

### 可观测性与监控

- [Opik](https://www.comet.com/docs/opik/cookbook/ollama) - 调试、评估与监控 LLM 应用
- [OpenLIT](https://github.com/openlit/openlit) - OpenTelemetry 原生的 Ollama 与 GPU 监控
- [Lunary](https://lunary.ai/docs/integrations/ollama) - 带 PII 脱敏的 LLM 可观测性
- [Langfuse](https://langfuse.com/docs/integrations/ollama) - 开源 LLM 可观测性
- [HoneyHive](https://docs.honeyhive.ai/integrations/ollama) - 面向智能体的 AI 可观测性与评估
- [MLflow Tracing](https://mlflow.org/docs/latest/llms/tracing/index.html#automatic-tracing) - 开源 LLM 可观测性

### 数据库与嵌入

- [pgai](https://github.com/timescale/pgai) - 把 PostgreSQL 用作向量数据库（[指南](https://github.com/timescale/pgai/blob/main/docs/vectorizer-quick-start.md)）
- [MindsDB](https://github.com/mindsdb/mindsdb/blob/staging/mindsdb/integrations/handlers/ollama_handler/README.md) - 把 Ollama 与 200+ 数据平台相连
- [chromem-go](https://github.com/philippgille/chromem-go/blob/v0.5.0/embed_ollama.go) - Go 可嵌入向量数据库（[示例](https://github.com/philippgille/chromem-go/tree/v0.5.0/examples/rag-wikipedia-ollama)）
- [Kangaroo](https://github.com/dbkangaroo/kangaroo) - AI 驱动的 SQL 客户端

### 基础设施与部署

#### 云

- [Google Cloud](https://cloud.google.com/run/docs/tutorials/gpu-gemma2-with-ollama)
- [Fly.io](https://fly.io/docs/python/do-more/add-ollama/)
- [Koyeb](https://www.koyeb.com/deploy/ollama)
- [Harbor](https://github.com/av/harbor) - 以 Ollama 为默认后端的容器化 LLM 工具箱

#### 包管理器

- [Pacman](https://archlinux.org/packages/extra/x86_64/ollama/)
- [Homebrew](https://formulae.brew.sh/formula/ollama)
- [Nix 包](https://search.nixos.org/packages?show=ollama&from=0&size=50&sort=relevance&type=packages&query=ollama)
- [Helm Chart](https://artifacthub.io/packages/helm/ollama-helm/ollama)
- [Gentoo](https://github.com/gentoo/guru/tree/master/app-misc/ollama)
- [Flox](https://flox.dev/blog/ollama-part-one)
- [Guix channel](https://codeberg.org/tusharhero/ollama-guix)

---

**代部署 / 定制服务 / 技术咨询 请添加微信：uaycar**

> 本项目为 [ollama/ollama](https://github.com/ollama/ollama) 的中文翻译版本，所有代码版权归原项目作者所有，遵循原项目许可证（MIT）。

**如果觉得有用，请给原项目点个 Star！** ⭐
