package scaffold

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const repoTemplate = `
package repositories

import (
	"${moduleName}/models"
	"gorm.io/gorm"
)

const ${modelCamel}PageQuery = "-- Write your query here --"

type I${modelName}Repo interface {
	IBaseCrudRepo[models.${modelName}, models.${modelName}DTO, models.${modelName}Page]
}

type ${modelName}Repo struct {
	BaseCrudRepo[models.${modelName}, models.${modelName}DTO, models.${modelName}Page]
}

func New${modelName}Repo(db *gorm.DB) *${modelName}Repo {
	return &${modelName}Repo{
		BaseCrudRepo: BaseCrudRepo[models.${modelName}, models.${modelName}DTO, models.${modelName}Page]{
			db:        db,
			pageQuery: ${modelCamel}PageQuery,
		},
	}
}
`

func ScaffoldRepo(modelName string, modelCamel string) {
	filename := fmt.Sprintf("repositories/%s.repo.go", ToSeparateByCharLowered(modelName, '_'))
	log.Printf("Generating repo file: %s", filename)

	newRepo := repoTemplate
	newRepo = strings.ReplaceAll(newRepo, "${modelCamel}", modelCamel)
	newRepo = strings.ReplaceAll(newRepo, "${modelName}", modelName)
	newRepo = strings.ReplaceAll(newRepo, "${moduleName}", moduleName)

	if err := os.WriteFile(filename, []byte(newRepo), 0644); err != nil {
		panic("Failed writing repo file")
	}
}
