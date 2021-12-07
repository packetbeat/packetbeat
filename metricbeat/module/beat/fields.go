// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package beat

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "beat", asset.ModuleFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded zlib format compressed contents of module/beat.
func AssetBeat() string {
	return "eJzsXU1v5LgRvftXED5nBOxhLz4kQZAE2EsCBAvkEAQGrWZ3E5ZIDUl54v31gaRutT5YrJJEqnt3x5cZjEfvPRbJYvGzvrB38fnC3gR3T4w56Qrxwp7/Irh7fmLsIGxuZOWkVi/sj0+MMdb8ipX6UBfiiTEjCsGteGEn/sSYFc5JdbIv7D/P1hbPf2DPZ+eq5/82vztr415zrY7y9MKOvLDN90cpioN9aZG/MMVL0Wmxr9ZxZ9t/Z8x9Vg2F0XV1+Zfhd8NveVV+scJ8CNP/yve5D2IIUxmdC2u1Gf0WQoLQhoi24ipzhit71KbkjUXt7D9fCXghue+3FXfnzj5Za56MV+VrV9ysl5yRqK66hDGzUoZLGirttMS+QuDFpBe1VZ/BTL0Wx/N3Z3gudlKE8F11HQ0vd5IUoLqqwZtnfFkUzqu+g8j1QapT9+k++iicM325rpXbVx5MeVX3wQt5aK28p/1orB6N+9kQJR31D543//NX6zEHZXgovznS9UjecyiM6kO7/76nMpT2Pj5+JvEhPf1Q5cP6e6/Ix/P6Q5kP7fsBoeERwOra5KLkPg+/3f+33MIXjccseF+GLMR3n5540/aw/dAj8fF64U3kQ/dBr8xwDyyFMzJP0v32HRq7gqwaFYPKIAtgVmDTRQGoeDQjLDYEzdFMnOWdJC7p8L/JyqL6s7tWV9CnjPrUbAxcv853EMeCO5EoPlBOKPelEOrkzkn808Uc2aUYGYGSUsOxRYVGiNMvMlV4tpv5mzI8lu0vikKGr1Wuy8oIa8Xh114Bw7I8VkVMlIUqxAg+d20szlKV/MXn4+IVs9OegTz7GfyiBAlLpbVSnb6MmwoYKKzf5rmqIhP2NTbdFWObxjkjvtbCugQlvPwRJhi4hLw2RiiXfePSZWXcTbXLHwSWm11spZVNFQCkauy90Tv11Ag3eXCr35PGjNNit4XKAqT9PnOei8p5x7nU6lDq/UJuv0C87RDWQGI0nv0N0BUsuHYyWG/QuuDmBI2jSRWi3JOJ7F1Eoty90+VOFLKUd6lsnPwq82st6rsYMkzcd5dC+yP35PoQ5n6/T5s3eTgIaE0lqUacfBaT5J/3cT44/W3Kxmt31kb+cp+KJ/FfxUrlhFG8uIdQlHu8nnQPiQjzYAH9rO9S2R7mWzhVBgN+WFNIRxh1zO2N2h9xejSUS5wk1cqKhCL88FgQFok8XHJokrJ9Kqa0K/VBHiXouLZNyEaF7OJqjDL9LNEjCg/2wdlUdC0AE3HOFl0OyIdOg7Y3TyR0jlhWSqxOHMHjqwoSUsLy+JJgtp37LzpXJYU08QUF6Igz5/iagoSDYPqDy4K/FfsJwzipU6f4ysKMpGg0vqgA3WBgP+pa7acpSHjznK2D/VoLA83qUnhPhJS6/hJfWpiRtKKRwJXCdAsm3EmcxJx0eNvqacq+5K7UWdsN07X2rzOIvjPy0j8no2PPIPpB5LPaij2D6Ku5nk07lmLPIPrQThg7P1e0FH6KMj5fZze1CPtpnSijzaXzqs5ybbwH7xfOIjtlGYx45Sw0P2Q/RCME4EZsP8Zlm8ONyxaZzoM34lPaTFsEizLTef7hOZ4jHZaoUextAkPuH5OS++powJ6U/Edv1y7k22TAYJv6diUrUUiVaLu8kEK5iAd5L6XPrqqzEMNoepnx/D3miD9TgrH0aw4fAplcxth55bmTHwlWwWfFbotjM4Svn8kaXVUpdjogXRhhPz3jsthTF8LXy5KFE2ZXYRhj7zjqt0La857acMrb6r8Dp2kJhIXpbosmLsXeGSTKT9cvCteuquFtgS1ev+PfwcXBDp1tN2dnoJt/C5Ht5XWnmkg+9427/Aze0o2uCqPbbSgYy6IOBIe6KmTOk5wsA5RdGXGbpR6lxsKIY5TTuuQqobcdq8LodvK2U1EQ2fAoePb2Cdfyhhj0ogWhGCmJfQnSJ4Vw3dGzBBRPBAh+Zf9mpBOJKwXjGGtJXC1BksHpqaP0XXHcHhl0r1ZlplYKu0W5qpid9CD+RIp13KSYoF6UBOBnQnSVUscUHVh9ZpsWE6Q66kxUZ1EKw4vXwMkP8hIXDjkirysnSxHlFgKC16++c3UoAkujW/pLu8GSnbmJuHBx0Zsh2GMJVh8jbi2NJYDY/VSlirlreiX3ot4W4cvh83dzyi21espflfhfRHte1WYh5EHRtPl85UWhfbf4t0ro4DMYfqIjFKZF0EEIzGxEt9vT+1D7ARZqPBEWXKs6+bwbvHTE1kfVnUnabaEAfF/Ko80qYaQ+ZHWCee1ADI1pqOtrrR3fQxZK1G8GAo5sLCfUMhihdQwZO4uFOBnZDAwxRfcPFMZ+znE22rlCHEhfxdW5jHuuGHwUJa1Y5FmUvKp5nju4667wpDc1Ieipgtbdw1bariNIMB7aUpgjjDzgv8RWkee2YxUUmqGk2vJT9On2XFKIJjxORriV323iOpm/xyxj0yUx5JkEaBoUQwSMPVp+yz54EfNkb6MBAx4LSFATGPBUQIJ6wKH7c1BWmBRGQHAn9AlMMEJ+YpMVje718msDmfdr6PXyu57FEwnP4omEZ/FEwrN4Iu1ZPBE+i9e0L+t4Wc3axBS7w33+c//F8wxs2CLZqsblr8KIrQBsvWsWA7nJz9KJ3NUGPl0UXo1r9GRBnKFyj2UWEYEYgQ64iCAIrtMsNlUFd0f/0T9cO/M1obYo2mZB5HD/jMDuA76tBFf1xmusIoNAxiTK8x71YpI5yGRzYGNRQJQJzcbCgCjjMycbPRgE0rtZYT5kHu8i8fZNDJFdNPnmbRF8i4co6GngPrmSbQjoiaTmtvdktYFjspHROqx38flNjzYuvIjdT5sv56e/ZvPRfRywrIVuvpuD00Z7DHuKMs72c/0JXgXYrzvMTTf+NNjEsY+xRuv7fuSqo1kgfJ++0J6N56EF4GE+VISo27SeMC+sYlqG1R8DRzJoH3vmEfSPt7SfvKitE2Y+C6F8HDWmzrVyXCoBXV+j2SJRqPkujBLhvTxIGcMb2DIQUthLAaIEsKHqL7niJ1EK5TKh+Nv8QF0H8aZ1IfiUJjA8ND8/WZYL5QwvBjTsQvOndQGlVE6cZs9UIkL+UZdvwjB9vOBb5isrOdoM2ROLIf11iunnpWjUd+CstuLA3j7bIdgrorvrkUZDi83ehFQnvxBoX299hGE3RRi3u6/R3ByU/S+MGkIeGY6WCZCFx/QhIpTDD9eLaZ7qxjdyQbFs3HaQtFILAYOZoBZi0apmISg1u8E6UNpDlyRM+sv+a2E3ih1VE5iMjX1v+iuwCMnL1iB+70yP3plCea1Y1K4UznFFkMy+N4SEDQFMr8SitoKkbgFFi3nYDK1OYiEYYYlnIeS8lfxeLQNG+v7sPyxKnA9nAsLxMQ42WZRBs6KwZU0mkjcBUvGw30v5kYw47PdiBzARDYsbpsNJaQhiWfxyL8wKgyAP91QS+Szs2eIFZSdlZlmAF0isgpcdKz+LV/dsOvLuOugG86UQ1Q/xCHlOVqBGDnCIIXVMO0cqARsd2ETTgaxAJeTvWIFKSbixAhZLkLECEk1qsQKTkodijVRS3ogVwMQsDyuQCSkZVqCiORRWYKJpDxDMm0MMZypYBJMwVVjUkXwzBJSBYAHEPvnG8Gf9CZJZmmiCkAuNhEPOXUbbeArHdTEqhZyDilR8choZEhotp9PeTYWYhYa2hE9NU0ZCwx9JXwhITwpFgiOmxyFhIU+JL0QjPAG+EJGeNo3Wfmn5w6jNhJ6YaoZ4M9lRP/m+WHMqYNPJs01n3n6dx9ainhn8bnza9/0a0d5PcAej3ULzNFmYsMefD7rGMmOgbzhTMLZDAO9js6iBFGwuklZGM9tyrHVQg8YZraEj95wpq7PYJWUCRviSMRlgQymwK74EiOAdXfr3CxUkGfpJb4qhYwH6NhhYGO+TPZu8eRJvTHmuhrrmgD8xQ0RCFg0oh7BDT9jEc874kyxLJm0LXnpZBbv20MagDcIvq9A6FOFRFIKQ7umOJP0hQtNb8pYK4XAP/RUU0BcBb9Ot90bwq3Hk2gPfe6MjQC+1ERD8T7ghZvS/27jeisArhZTgPPQSI/l74BlFxArQmAjSIrcY2jsFHaiXz9SqJTxpo2snFVABy4lvV1Bu0OzyECyTCr7LsTShCqLj720zYbkuS62Y04wXRUs+LejDJ2ZBvRk1rwq+ir1XwgDS6/y09V3i6/W0FVTas+40LGrGEBIaPckHbb2TlJaDuChOed49HAGle+ia8sK1PWsTfqqV8Dg1BQR7WNqP0Y9r6xKFIG6y+flnd9sOCu4pdgbWBhkpxCOJbH5+/qwG1wNTujKioL81LNeT/BBb3DwqyD4/UTgbjdSdsRoBSn8rxOEUL8VKCrkddbiKB6vvxLQryQyLKog6lm1rAouzs9zRrpaSIybm0L5NbOQsLimUOq1ZQ07tWhvG/lVKf274mJpYNhwACO+GC4s5CvxL8ENwaYw6AoQWI1jSmm+ZQ8ZafMY0hcpGX4j/KrFNr5K61v/dkPwmqr0xlwscG3mAmm9rdCog9Jf/BwAA////wsMd"
}
