package toggle

import "fmt"

type FeatureRepository interface{}
type Feature interface{}

func main() {
	usesrRepository := NewInMemoryUserRepository()
	repository := NewInMemoryFeatureRepository()
	features := NewFeatures(repository, usersRepository)

	features.RegisterFeature(Feature{
		Namespace: "my",
		Name:      "feature-one",
	})

	if feature.isActive("my.feature-one") {
		fmt.Println("my.feature is enabled")
	}
}
