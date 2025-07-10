# DecoyOps

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Linux-lightgrey.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D%201.16-00ADD8.svg)

DecoyOps is a powerful cross-platform system reconnaissance and analysis tool written in Go. It provides detailed insights into system processes, file systems, and network services.

## 🚀 Features

- **Process Discovery**: Enumerate and analyze running processes
- **File System Enumeration**: Scan and catalog file system contents
- **Network Service Discovery**: Identify open ports and running services

## 📋 Requirements

- Go 1.16 or higher
- Windows or Linux operating system
- Administrative privileges (for some features)

## ⚡ Quick Start

1. **Clone the repository**
```bash
git clone https://github.com/zeusnotfound04/DecoyOps.git
cd DecoyOps
```

2. **Build the project**

For Windows:
```cmd
go build -o DecoyOps.exe .\cmd\main.go
```

For Linux:
```bash
GOOS=linux GOARCH=amd64 go build -o DecoyOps ./cmd/main.go
```

## 🛠️ Usage

### Available Commands

1. **Process Discovery**
```cmd
DecoyOps.exe --process-discovery
```
- Enumerates all running processes
- Shows detailed process information (CPU, Memory, Status)
- Identifies system and user processes

2. **File Enumeration**
```cmd
DecoyOps.exe --file-enum
```
- Lists all files and directories
- Shows file permissions and sizes
- Generates detailed file system report

3. **Network Scanning**
```cmd
DecoyOps.exe --Network-scan
```
- Scans for open ports
- Identifies running services
- Displays network configuration

### Help Command
```cmd
DecoyOps.exe --help
```

## 📁 Project Structure

```
DecoyOps/
├── cmd/
│   └── main.go           # Main application entry point
├── internal/
│   ├── executor/         # Core functionality implementations
│   ├── logger/          # Logging mechanism
│   └── utils/           # Utility functions
├── output/              # Output files directory
└── techniques/          # Technique definitions
```

## 📝 Output

All scan results are saved in:
- `output/output.json`: Detailed JSON-formatted results

## ⚠️ Notes

- Some features require administrative privileges
- Network scanning is limited to localhost by default
- Use responsibly and in accordance with your organization's security policies

## 🔒 Security Considerations

- The tool performs active system reconnaissance
- Some operations may trigger security tools or EDR solutions
- Always obtain proper authorization before use

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit pull requests.

1. Fork the project
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## ✨ Acknowledgments

- Go community for excellent networking and system libraries
- Contributors and testers

---
Created with by ⚡ ZeusNotfound ⚡
