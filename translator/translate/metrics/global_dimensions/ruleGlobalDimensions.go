// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package global_dimensions

import (
	parent "github.com/aws/amazon-cloudwatch-agent/translator/translate/metrics"
)

type globalDimensions struct {
}

const SectionKey = "global_dimensions"

func (ad *globalDimensions) ApplyRule(input interface{}) (returnKey string, returnVal interface{}) {
	im := input.(map[string]interface{})

	dimensions := map[string]interface{}{}

	if _, ok := im[SectionKey]; !ok {
		returnKey = ""
		returnVal = ""
	} else {
		for key, val := range im[SectionKey].(map[string]interface{}) {
			if key != "" && val != "" {
				dimensions[key] = val
			}
		}

		returnKey = "outputs"
		returnVal = map[string]interface{}{
			SectionKey: dimensions,
		}
	}
	return
}

func init() {
	gd := new(globalDimensions)
	parent.RegisterRule(SectionKey, gd)
}
