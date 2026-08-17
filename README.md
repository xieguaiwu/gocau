# gocau · 终端计算器 CLI

极简终端计算器命令行工具，提供算术平均、加权平均与语言学 Fitness Metric 计算。零第三方依赖（Go 标准库）。

## 子命令

| 命令 | 说明 |
|---|---|
| `gocau avg` | 输入若干数值，计算算术平均值 |
| `gocau wavg` | 输入数值与对应权重，计算加权平均值 |
| `gocau fitness_metric` | 输入多轮违规数/超集参数/优雅度，计算语言学 Fitness Metric |

## 安装

### 方式一：源码构建

```bash
go build -o gocau .
./gocau avg
```

### 方式二：GitHub Release

从 [Releases](https://github.com/xieguaiwu/gocau/releases) 下载 `gocau-<version>-linux-amd64.tar.gz`：

```bash
tar xzf gocau-*-linux-amd64.tar.gz
sudo install -Dm755 gocau /usr/local/bin/gocau
```

### 方式三：COPR（Fedora）

```bash
sudo dnf copr enable xieguaiwu/gocau
sudo dnf install gocau
```

## 使用示例

```bash
$ gocau avg
Enter number of values: 3
Enter all the numbers:
1 2 3
Average = 2.0000

$ gocau wavg
Enter number of values: 2
Enter numbers:
80 90
Enter corresponding weights:
2 1
Weighted Average = 86.6667
```

## 许可证

MIT
