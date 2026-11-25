<p align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:111827,100:2563eb&height=160&section=header&text=Go%20Channel%20Select%20Lab&fontSize=30&fontColor=ffffff" alt="Go Channel Select Lab header"/></p>
<p align="center"><img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=18&pause=1000&color=60A5FA&center=true&vCenter=true&width=760&lines=Explore%20buffered%20and%20unbuffered%20channels%2C%20select%20cases%2C%20default%20branches%2C%20and%20closed-channel%20receives." alt="Animated tagline"/></p>
<p align="center"><a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-2563eb.svg" alt="MIT"/></a> <img src="https://img.shields.io/badge/language-Go-0ea5e9.svg" alt="Go"/> <img src="https://img.shields.io/github/stars/CyberLoop-dev/go-channel-select-lab?style=flat" alt="Stars"/> <img src="https://img.shields.io/github/last-commit/CyberLoop-dev/go-channel-select-lab?style=flat" alt="Last commit"/></p>
<p align="center"><img src="assets/mark.svg" width="110" alt="Custom project mark"/></p>

## 🧭 Overview
Explore buffered and unbuffered channels, select cases, default branches, and closed-channel receives. The examples use the Go standard library and short, bounded demonstrations.

## ✨ Features
- 📬 Compare channel capacity, queued values, and blocking sends
- 🔀 Observe select choosing among ready cases and a timeout
- 🚪 Inspect the value and ok result received from a closed channel

## 🎬 Demo
Run the CLI and compare the printed event order and context/channel state. No hosted demo is claimed.

## 🚀 Quick Start
```bash
go run .
```
Use `go run . -scenario buffered`, `unbuffered`, `select`, `closed`, or `all`.

## 🧰 Tech Stack
- 🧩 Go standard library
- 🔒 No network access or external services
- 📦 No third-party dependencies

## 🗂️ Folder Structure
```text
go-channel-select-lab/
├── assets/mark.svg
├── main.go
├── main_test.go
├── go.mod
├── .gitignore
├── LICENSE
└── README.md
```

## 🛣️ Roadmap
- [ ] Add a JSON output mode
- [ ] Add more guided scenarios

## 🤝 Contributing
Issues and pull requests are welcome. Keep examples deterministic where possible and explain any timing-dependent behavior.

## 📄 License
Released under the [MIT License](LICENSE).

<p align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:2563eb,100:111827&height=105&section=footer" alt="Animated wave footer"/></p>