package graphic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// UniformDataType type
type UniformDataType int32

// ShaderUniformDataType enumeration
const (
	ShaderUniformFloat UniformDataType = iota
	ShaderUniformVec2
	ShaderUniformVec3
	ShaderUniformVec4
	ShaderUniformInt
	ShaderUniformIvec2
	ShaderUniformIvec3
	ShaderUniformIvec4
	ShaderUniformSampler2d
)

type Shader struct {
	uniformLocations map[string]int32
	Shader           rl.Shader
}

func InitShader(shaderPath string) *Shader {
	s := new(Shader)
	s.Shader = rl.LoadShader("", shaderPath)
	s.uniformLocations = make(map[string]int32)
	return s

}

func (s *Shader) SetValueFromUniformName(uniformName string, value any, uniformType UniformDataType) {

	if location, ok := s.uniformLocations[uniformName]; ok {
		s.setValue(location, value, uniformType)

	} else {
		s.uniformLocations[uniformName] = rl.GetShaderLocation(s.Shader, uniformName)
		s.setValue(s.uniformLocations[uniformName], value, uniformType)

	}

}
func (s *Shader) setValue(uniformLocation int32, value any, uniformType UniformDataType) {
	switch value.(type) {
	case float32:
		rl.SetShaderValue(s.Shader, uniformLocation, []float32{value.(float32)}, rl.ShaderUniformDataType(uniformType))
	case []float32:
		rl.SetShaderValue(s.Shader, uniformLocation, value.([]float32), rl.ShaderUniformDataType(uniformType))
	default:
		panic("Invalid value type")
	}

}

func (s *Shader) Begin() {
	rl.BeginShaderMode(s.Shader)
}

func (s *Shader) End() {
	rl.EndShaderMode()
}
