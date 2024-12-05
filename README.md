# Containers Types

A Go module for handling images in Docker, Docker Compose, and Helm files. This module provides structures and utilities to manage image information across different types of files in a Kubernetes environment.

## Installation

```bash
go get github.com/Checkmarx/containers-types
```

## Usage

### Importing the Module

```go
import "github.com/Checkmarx/containers-types"
```

### File Structures

#### FileImages

```go
type FileImages struct {
    Dockerfile    []FilePath
    DockerCompose []FilePath
    Helm          []HelmChartInfo
}
```

#### FilePath

```go
type FilePath struct {
    FullPath     string
    RelativePath string
}
```

#### ImageModel

```go
type ImageModel struct {
    Name           string
    ImageLocations []ImageLocation
}
```

#### HelmChartInfo

```go
type HelmChartInfo struct {
    Directory     string     // Absolute path to the Helm chart directory
    ValuesFile    string     // Relative path to the values.yaml file
    TemplateFiles []FilePath // Relative paths to template files
}
```

#### ImageLocation

```go
type ImageLocation struct {
    Origin string
    Path   string
}
```

### Constants

```go
const (
    UserInput               = "UserInput"
    DockerFileOrigin        = "Dockerfile"
    DockerComposeFileOrigin = "DockerCompose"
    HelmFileOrigin          = "Helm"
    NoFilePath              = "NONE"
)
```

### Microservice

```go
type Microservice struct {
    Spec struct {
        Image struct {
            Registry string yaml:"registry"
            Name     string yaml:"name"
            Tag      string yaml:"tag"
        } yaml:"image"
    } yaml:"spec"
}
```

### Functions

#### ToImageModels

Converts a list of image names into a list of `ImageModel` structures, marking them as user input with no specific file paths.

```go
func ToImageModels(images []string) []ImageModel {
    var imageNames []ImageModel

    for _, image := range images {
        imageNames = append(imageNames, ImageModel{
            Name: image,
            ImageLocations: []ImageLocation{
                {
                    Origin: UserInput,
                    Path:   NoFilePath,
                },
            },
        })
    }

    return imageNames
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
