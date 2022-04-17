// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package airflow

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "airflow", asset.ModuleFieldsPri, AssetAirflow); err != nil {
		panic(err)
	}
}

// AssetAirflow returns asset data.
// This is the base64 encoded zlib format compressed contents of module/airflow.
func AssetAirflow() string {
	return "eJzMl8GK4zAMQO/5CtHLwMAM9BCG5LCwXxKUSA1q7djYSmf694vTdslCtk0hB+dUJFnvKQjqfMCJLzWghINx3wWAihquYff7GtkVAMSxC+JV3FDDrwIA4JYF62g0XAAENoyRa2hZsQA4CBuK9VT8AQNankPSoxfPNfTBjf4WmZ+Zn3t7/9zbJqDy29/c/bxrj9zpLHwNNNcsubGd9BayjUXvZehvpbv33axuYeT7cx99byEpgYrlEMGyBumW3Mv83Mu17vsM5fflSvnOjYPmZD4JcYhLshZ/clK1+PP8BVvGIbv1SFIrFySVZqe+wpokN+9ktMJc8tKWFc7+q8zJ+asEz6HjQcWs2HBfZWVfvWpfNVVW/tVn9eoEefm/ZB+ViM85DRAVB8JAQHwWTJXPpzijGbP6h+px7HnhEkDYNwcxC64nvny7QK9hCHtI7dKrQcLpWr5AFNqQJ/SAdnRtk35uxDu6dur8gOg8B1QXtsTeez5je+fMltzU7xkzKuoYNwJemz2gKcbTduuTuv1nf/79uPwTAAD//9XSEKo="
}
