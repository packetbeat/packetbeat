// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package okta

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "okta", asset.ModuleFieldsPri, AssetOkta); err != nil {
		panic(err)
	}
}

// AssetOkta returns asset data.
// This is the base64 encoded gzipped contents of module/okta.
func AssetOkta() string {
	return "eJzsWk1zqzgW3edX3HprtRfz0YsspopnlDw6tnEBjisrSsGyrQlGbkkkz/3rp8SHjUEGYpOpXjzvAtI5515JR1civ8EbPdwDf1PkDkAxFdN7cJ8C6w5gRWUk2F4xntzDf+4AAKZ8lcYU1lzAliSrmCUbkAep6A5ivpGwFnyXdR/dAawZjVfyPuv4GyRkR49E+qcOe3oPG8HTffHEQKh/DxnOObb+VfGrHGnKVseHx6AWC8euPJVbLtQ9BFsKacL+TCmwFU0UWzMqgK9BbWlGBhO+we80UaNK5wtC9e86wDwXb/TwwcVJeyMyqvuFunEzvgwTgvN3lSh1r1LHFTF1de8ZwTsVkvGkKf+58aKiveh1g/weCD0jkPSdCqYOzRD85ptKDGW/G4K4CAHTVCp4pcCTbJBs/H3xiMCZPbgIlpY3Q8AFYM9zvSsiXjG5j8kh3FEpycYw9ey8AUwbDSrxFyhQoNyQhk8g9YyQRIqLZlxW7XERTeFHBW3WN3PEK8IpoNSWKIipggNPQSouKLBkzcWOqMqk7WZqWioYfLLyqkzAmWFeSltHLABOw/AyyaM25pqZXc0dVPypm5XEioqEKBoOE7lV4hlMv1tNucD0X0OoKdejxutSUWqIYkYT1VwD4+w5XFwF5JWnKmPIES7wDbAKejNdvwr2huzXHrZPwi0FZw5ktRJUHh0ilzs6a2tgTyUVIdmcj0IHYWMQNAhkICYHMSkxZ6uZLzh7W4oW5CO8ILxt+naElWdSkI9jEKcYThFezGil3JTDSnL9iiLjmj5xvwr+IakYVkABashLfWwNM+wvngzjtVuaYfWZYia7o+8sGkzIRQ0FT6uUYdw/k9Hw/YuZOE7OVEV8Z6in3PxFW2FzbsK9nbOg/CJfNqBf7cWCEnl2HLhpdHK0MqENnUZ+mcZqOH6NVuevle334C/GY+z7CB4sZ7LwMAL/yZnPsY3AmkzcJQIbz14QjH9YkwmePWIEi9nTzF3OWqaZImJDDTt7UH9eKdZjJnO1WRv5ieK8pWeePyIEOfyqTn9Vp5+pTpUgiSSRMl4dBMaXVzikoDFRdFVl+wKn7GD5253hjDq/cq1Ub5sq3PVLjm9L/P0bgm9/uN+/tUydFX1NN2HEE0V/GlzQ1q9h3Hh95QabsUHB9kXb7EWOq6dOnqMVOd7N9hymXInud92Zwlg352VbuGbJhoq9YEMfKyrA5RRvqRSrFcGfKZWq7sk362l6csHUU08q2MDnrhwYFp7TIUFtBSUqlKnc00jRQROjkeGI3CEkFfGwOVh4kxZPIana6kGLsqV52Vyss3bDuMw59xdaTTfR1X5Ty99e8He2qh2VBzoSnmyzFk9J2thX3KfACq1F8APPAmdsBY47C+ee++zY2ENgjQPnGYe24+Fx4HovCCa2NUfwgG3sZY0R+O7YsSa6itdNjp3ba77zlEhFjZdRiaKbWqI601ELXEO3SokEzSyJxP+HkTmRdYzI2MO2HhFrUhkNz7cQ+C9TaxbgMYJH132cYAT2wkXwsvjuPOGXvqEOWb60hqlxmyEGcwT+1Ecwt3x/6Xr6yOf72MtnlLO0EOCp5UwQuHpu/gPBH8sAwVi3eNDTFCOYezj0f1getkP/ZTrFgeeMwyf8gooMThw8C0If+34GauNnZ4zDhe3Y7ZckUqafnXPmJZgjmWqF+qmwZ63w9XtxXbJ5H6zPnZuFVD+3tkg4fhj+mZ3w4lBSKbV/DHexVSCWqTllpuQ8tVgDSQ7tUylRVKzJkLd/BSCkkq5GQEebEQI3VTHnbwjc9ZpF9J+//xvBhwxEKtXlTV3SKBVMHS5v537RYpiNvOT7wi28jeL6zVtCr2K/z77EE77jqSz+h2T0eQNI0t3rhTt2017Za+1ZfgHbseq52JCE/UVqFw9tOenFX8XNx5l/JDJbcwZp5mQZZFeSVr/JgQ7D6pRtEK5JOrYVU4VzlQs42gUSqsCnQp/kYF6WEq2nXr4jbLAr5hytT9i6qvp5MPC+ch5TkvTnXW6p2lIBTAGTQCADBi4g4W0fHooDXtPgvMaLK69BGifYQS9ACnQELMn+1m30rlTePbN9GG0JG+harQDrPyqTmgzgr/+l0fld+I13JKxuK8Zv1Z3L9vStusPqmv+2BbfVOM68xGyUws//QvD8e4cgyVMRDVhz+Rledb51CNhQvhFkv2URiQ0VQw/KxwqCmbgalWkvaZ0qlUPO+f+kdeeqJl27m8a4kJHKoCiibtpYsnpTg3RS7blUWeJXNxPmUKChOmkjniZK3J7NHKaTbkN5zCNTgdFjR348dW6bXscJRnm45yxRd/8LAAD//zCDmHs="
}
