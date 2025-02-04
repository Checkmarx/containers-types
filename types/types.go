//go:build !coverage

package types

type FileImages struct {
	Dockerfile    []FilePath
	DockerCompose []FilePath
	Helm          []HelmChartInfo
}

type FilePath struct {
	FullPath     string
	RelativePath string
}

type ImageModel struct {
	Name           string
	ImageLocations []ImageLocation
}

type HelmChartInfo struct {
	Directory     string     // Absolute path to the Helm chart directory
	ValuesFile    string     // Relative path to the values.yaml file
	TemplateFiles []FilePath // Relative paths to template files
}

type ImageLocation struct {
	Origin     string
	Path       string
	FinalStage bool
}

const (
	UserInput               = "UserInput"
	DockerFileOrigin        = "Dockerfile"
	DockerComposeFileOrigin = "DockerCompose"
	HelmFileOrigin          = "Helm"
	NoFilePath              = "NONE"
)

type Microservice struct {
	Spec struct {
		Image struct {
			Registry string `yaml:"registry"`
			Name     string `yaml:"name"`
			Tag      string `yaml:"tag"`
		} `yaml:"image"`
	} `yaml:"spec"`
}

func ToImageModels(images []string) []ImageModel {
	var imageNames []ImageModel

	for _, image := range images {
		imageNames = append(imageNames, ImageModel{
			Name: image,
			ImageLocations: []ImageLocation{
				{
					Origin:     UserInput,
					Path:       NoFilePath,
					FinalStage: false,
				},
			},
		})
	}

	return imageNames
}
