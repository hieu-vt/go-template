## Go template from (hieu-vt)

### Create new go template

```
npx degit hieu-vt/go-template/template <project-name>
```

### **Replace module name**

- 1:
  Open your project in GoLand.
- 2:
  Press **_"Shift" + "Ctrl" + "R"_** (on Windows/Linux) or **_"Shift" + "Command" + "R"_** (on macOS) to open the "Replace in Path" dialog
- 3:
  In the "Text to find" field, enter **_go-template_**.
- 4:
  In the "Replace with" field, enter **_<project-name>_**
- 5:
  Click **_"Replace all"_** to replace all occurrences of **_go-template_** with **_<project-name>_**.
- 6:
  Open a terminal in your project directory.
- 7:
  Run the command `go mod edit -module <project-name>` to update your module to use **_<project-name>_** as the module name.
- 8:
  Run the command `go mod tidy` to clean up the module and update its dependencies.
