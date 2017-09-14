# Tools

## Env
Go library to load environment properties

Is really simple to use.


### Examples

#### Installation

Clone in your go path

```
git clone https://github.com/basset-la/tools
```

#### Importing

```go
import "github.com/basset-la/tools"
```

#### Create properties

Create a folder in your project

```
mkdir YOUR_PROJECT_PATH/properties
```

Create prop files

```
touch YOUR_PROJECT_PATH/properties/development.yaml
```

Put some props

```yaml
---
app:
  appPath: /test
```

#### Create a struct to access your properties easily and define a global var

```go
// global var
var props Properties

// Properties : Struct used to map env.yaml properties file
type Properties struct {
	App struct {
		Path string `yaml:"appPath"`
	}
}
```

#### Load your props

In your main func call to LoadProperties func telling the name of the properties folder and passing the struct to map the yaml

```go
env.LoadProperties("properties", props)
```

#### Running your project

When you run your project don't forget to pass the -e parameter