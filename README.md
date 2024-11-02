## Scaffold

### A very small, and not very in-depth, tool to create a go project.

### Installation: 
```bash
go install github.com/justinbather/scaffold@latest
```

### Usage:
```bash
scaffold github.com/your-github-username/your-project-name
```
> Right now this only supports the above project name format. I am lazy and since I only create projects like this idc.

> Under the hood it parses the "your-project-name", creates a folder for it, enters it, runs git init, inits a go module with the package name, and creates a simple main file with main func.
