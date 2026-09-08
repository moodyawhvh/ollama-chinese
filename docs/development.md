> 🌐 本文档由 [ollama/ollama](https://github.com/ollama/ollama) 翻译,英文原版见原项目。

# 开发

安装前置依赖:

- [Go](https://go.dev/doc/install)
- [CMake](https://cmake.org/download/) 3.24 或更新
- C/C++ 编译器:macOS 用 Clang,Windows 用 Visual Studio 2022 C++ 工具,Linux 用 GCC/Clang
- 建议把 [Ninja](https://github.com/ninja-build/ninja/releases) 加入 `PATH`,Windows 上尤其推荐

针对已有原生产物做纯 Go 迭代时,从仓库根目录运行 Ollama:

```shell
go run . serve
```

> [!NOTE]
> Ollama 包含用 CGO 编译的原生代码。这些数据结构偶尔会变化,CGO 可能因此不同步,导致意外崩溃。可以先运行 `go clean -cache` 强制完整重建原生代码。

## 原生构建模型

全新 checkout 或修改原生代码后,从仓库根目录构建。在 macOS arm64 上,这会构建 Metal 推理;在所有其他平台上,这会构建仅 CPU 推理。它会在仓库根目录构建 Go 二进制,并把原生运行时产物安装到 `build/lib/ollama` 下。

```shell
cmake -B build .
cmake --build build --parallel 8
./ollama serve
```

安装到标准前缀布局:

```shell
cmake --install build --prefix /path/to/install
```

在 macOS arm64 以外的所有平台上,构建 GPU 后端需要显式选择后端:

```shell
cmake -B build . -DOLLAMA_LLAMA_BACKENDS="cuda_v13;vulkan"
cmake --build build --parallel 8
```

支持的后端取值有 `cuda_v12`、`cuda_v13`、`rocm_v7_1`、`rocm_v7_2`、`vulkan`、`cuda_jetpack5` 和 `cuda_jetpack6`。

使用标准 CMake 架构覆盖参数,把 GPU 构建收窄到本地硬件:

```shell
# CUDA
cmake -B build . -DOLLAMA_LLAMA_BACKENDS=cuda_v13 -DCMAKE_CUDA_ARCHITECTURES=native

# ROCm / HIP
cmake -B build . -DOLLAMA_LLAMA_BACKENDS=rocm_v7_2 -DCMAKE_HIP_ARCHITECTURES=gfx1100
```

你可以在 configure 阶段设置 `GGML_*` 值来调整 GGML 构建选项。例如,本地调试时禁用 CUDA flash attention 内核:

```shell
cmake -B build . -DOLLAMA_LLAMA_BACKENDS=cuda_v12 -DGGML_CUDA_FA=OFF
```

## macOS (Apple Silicon)

额外前置依赖:

MLX Metal 需要 Metal 工具链。先安装 [Xcode](https://developer.apple.com/xcode/),然后:

```shell
xcodebuild -downloadComponent MetalToolchain
```

## Windows

额外前置依赖:

- [Visual Studio 2022](https://visualstudio.microsoft.com/downloads/),包含 Native Desktop 工作负载
- (可选)AMD GPU 支持
    - [ROCm](https://rocm.docs.amd.com/en/latest/)
- (可选)NVIDIA GPU 支持
    - [CUDA SDK](https://developer.nvidia.com/cuda-downloads?target_os=Windows&target_arch=x86_64&target_type=exe_network)
- (可选)Vulkan GPU 支持
    - [Vulkan SDK](https://vulkan.lunarg.com/sdk/home) - 对 AMD/Intel GPU 有用
- (可选)MLX 引擎支持
    - [CUDA 13+ SDK](https://developer.nvidia.com/cuda-downloads)
    - [cuDNN 9+](https://developer.nvidia.com/cudnn)

使用 Ninja 构建时,请从 Developer PowerShell/Command Prompt 或其他 Visual Studio 编译器可用的 shell 中运行 CMake。

> 构建 Vulkan 需要设置 VULKAN_SDK 环境变量:
> 
> PowerShell
> ```powershell
> $env:VULKAN_SDK="C:\VulkanSDK\<version>"
> ```
> CMD
> ```cmd
> set VULKAN_SDK=C:\VulkanSDK\<version>
> ```

## Windows (ARM)

Windows ARM 目前不支持额外的加速库。

## Linux

额外前置依赖:

- (可选)AMD GPU 支持
    - [ROCm](https://rocm.docs.amd.com/projects/install-on-linux/en/latest/install/quick-start.html)
- (可选)NVIDIA GPU 支持
    - [CUDA SDK](https://developer.nvidia.com/cuda-downloads)
- (可选)Vulkan GPU 支持
    - [Vulkan SDK](https://vulkan.lunarg.com/sdk/home) - 对 AMD/Intel GPU 有用
    - 或通过包管理器安装:`sudo apt install vulkan-sdk`(Ubuntu/Debian)或 `sudo dnf install vulkan-sdk`(Fedora/CentOS)
- (可选)MLX 引擎支持
    - [CUDA 13+ SDK](https://developer.nvidia.com/cuda-downloads)
    - [cuDNN 9+](https://developer.nvidia.com/cudnn)
    - OpenBLAS/LAPACK:`sudo apt install libopenblas-dev liblapack-dev liblapacke-dev`(Ubuntu/Debian)
> [!IMPORTANT]
> 运行 CMake 前,确保前置依赖都在 `PATH` 中。

## MLX 引擎(可选)

MLX 引擎用于运行基于 safetensor 的模型。在 macOS arm64 上 MLX 默认启用;在其他平台上,通过 `OLLAMA_MLX_BACKENDS` 选择 MLX 后端。

### CUDA

需要 CUDA 13+ 和 [cuDNN](https://developer.nvidia.com/cudnn) 9+。

```shell
cmake -B build . -DOLLAMA_MLX_BACKENDS=cuda_v13
cmake --build build --parallel 8
```

### 本地 MLX 源码覆盖

要基于本地 checkout 的 MLX 和/或 MLX-C 构建(开发时有用),在运行 CMake 前设置环境变量:

```shell
export OLLAMA_MLX_SOURCE=/path/to/mlx
export OLLAMA_MLX_C_SOURCE=/path/to/mlx-c
```

macOS arm64 上:

```shell
OLLAMA_MLX_SOURCE=../mlx OLLAMA_MLX_C_SOURCE=../mlx-c cmake -B build .
cmake --build build --parallel 8
```

CUDA 上:

```powershell
$env:OLLAMA_MLX_SOURCE="../mlx"
$env:OLLAMA_MLX_C_SOURCE="../mlx-c"
cmake -B build . -DOLLAMA_MLX_BACKENDS=cuda_v13
cmake --build build --parallel 8
```

## Docker

```shell
docker build .
```

### ROCm

```shell
docker build --build-arg FLAVOR=rocm .
```

## 运行测试

使用 `go test` 运行测试:

```shell
go test ./...
```

## 运行库探测

Ollama 会在已安装布局和本地开发布局中查找原生辅助二进制和加速库:

* `../lib/ollama` - 标准安装(`ollama` 位于 `bin/` 下)
* `./lib/ollama` - Windows 发布版风格的产物和本地 dist 输出
* `.` - macOS 发布产物(辅助程序与 `ollama` 同目录)
* `build/lib/ollama` 和 `dist/<platform>/lib/ollama` - 本地开发构建

如果找不到这些库,Ollama 将无法以任何加速库运行。
