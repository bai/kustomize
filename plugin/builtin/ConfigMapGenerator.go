// Code generated by pluginator on ConfigMapGenerator; DO NOT EDIT.
package builtin

import (
	"sigs.k8s.io/kustomize/pkg/ifc"
	"sigs.k8s.io/kustomize/pkg/resmap"
	"sigs.k8s.io/kustomize/pkg/types"
	"sigs.k8s.io/yaml"
)

type ConfigMapGeneratorPlugin struct {
	ldr ifc.Loader
	rf  *resmap.Factory
	types.GeneratorOptions
	types.ConfigMapArgs
}

func NewConfigMapGeneratorPlugin() *ConfigMapGeneratorPlugin {
	return &ConfigMapGeneratorPlugin{}
}

func (p *ConfigMapGeneratorPlugin) Config(
	ldr ifc.Loader, rf *resmap.Factory, config []byte) (err error) {
	p.GeneratorOptions = types.GeneratorOptions{}
	p.ConfigMapArgs = types.ConfigMapArgs{}
	err = yaml.Unmarshal(config, p)
	p.ldr = ldr
	p.rf = rf
	return
}

func (p *ConfigMapGeneratorPlugin) Generate() (resmap.ResMap, error) {
	return p.rf.FromConfigMapArgs(p.ldr, &p.GeneratorOptions, p.ConfigMapArgs)
}
