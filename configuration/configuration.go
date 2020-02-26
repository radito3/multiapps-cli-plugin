package configuration

import (
	"fmt"
	"github.com/cloudfoundry-incubator/multiapps-cli-plugin/configuration/properties"
	"os"
)

const unknownError = "An unknown error occurred during the parsing of the environment variable \"%s\". Please report this! Value type: %T"

type Snapshot struct {
	backendURL          string
	uploadChunkSizeInMB uint64
}

func NewSnapshot() Snapshot {
	return Snapshot{
		backendURL:          getBackendURLFromEnvironment(),
		uploadChunkSizeInMB: getUploadChunkSizeInMBFromEnvironment(),
	}
}

func (c Snapshot) GetBackendURL() string {
	return c.backendURL
}

func (c Snapshot) GetUploadChunkSizeInMB() uint64 {
	return c.uploadChunkSizeInMB
}

func getBackendURLFromEnvironment() string {
	return getStringProperty(properties.BackendURL)
}

func getUploadChunkSizeInMBFromEnvironment() uint64 {
	return getUint64Property(properties.UploadChunkSizeInMB)
}

func getStringProperty(property properties.ConfigurableProperty) string {
	uncastedValue := getPropertyOrDefault(property)
	value, ok := uncastedValue.(string)
	if !ok {
		panic(fmt.Sprintf(unknownError, property.Name, uncastedValue))
	}
	return value
}

func getUint64Property(property properties.ConfigurableProperty) uint64 {
	uncastedValue := getPropertyOrDefault(property)
	value, ok := uncastedValue.(uint64)
	if !ok {
		panic(fmt.Sprintf(unknownError, property.Name, uncastedValue))
	}
	return value
}

func getPropertyOrDefault(property properties.ConfigurableProperty) interface{} {
	value := getPropertyWithNameOrDefaultIfInvalid(property, property.Name)
	if value != nil {
		return value
	}
	for _, deprecatedName := range property.DeprecatedNames {
		value := getPropertyWithNameOrDefaultIfInvalid(property, deprecatedName)
		if value != nil {
			fmt.Printf("Attention: You're using a deprecated environment variable \"%s\". Use \"%s\" instead.\n\n", deprecatedName, property.Name)
			return value
		}
	}
	return property.DefaultValue
}

func getPropertyWithNameOrDefaultIfInvalid(property properties.ConfigurableProperty, name string) interface{} {
	propertyValue, err := getPropertyWithName(name, property.Parser)
	if err != nil {
		propertyValue = os.Getenv(name)
		fmt.Printf(property.ParsingFailureMessage, propertyValue, name, property.DefaultValue)
		return property.DefaultValue
	}
	if propertyValue != nil {
		fmt.Printf(property.ParsingSuccessMessage, propertyValue, name)
		return propertyValue
	}
	return nil
}

func getPropertyWithName(name string, parser properties.Parser) (interface{}, error) {
	propertyValue, isSet := os.LookupEnv(name)
	if isSet {
		return parser.Parse(propertyValue)
	}
	return nil, nil
}
