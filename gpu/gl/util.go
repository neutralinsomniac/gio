// SPDX-License-Identifier: Unlicense OR MIT

package gl

import (
	"errors"
	"fmt"
	"strings"
)

func CreateProgram(ctx Functions, vsSrc, fsSrc string, attribs []string) (Program, error) {
	vs, err := createShader(ctx, VERTEX_SHADER, vsSrc)
	if err != nil {
		return Program{}, err
	}
	defer ctx.DeleteShader(vs)
	fs, err := createShader(ctx, FRAGMENT_SHADER, fsSrc)
	if err != nil {
		return Program{}, err
	}
	defer ctx.DeleteShader(fs)
	prog := ctx.CreateProgram()
	if !prog.valid() {
		return Program{}, errors.New("glCreateProgram failed")
	}
	ctx.AttachShader(prog, vs)
	ctx.AttachShader(prog, fs)
	for i, a := range attribs {
		ctx.BindAttribLocation(prog, Attrib(i), a)
	}
	ctx.LinkProgram(prog)
	if ctx.GetProgrami(prog, LINK_STATUS) == 0 {
		log := ctx.GetProgramInfoLog(prog)
		ctx.DeleteProgram(prog)
		return Program{}, fmt.Errorf("program link failed: %s", strings.TrimSpace(log))
	}
	return prog, nil
}

func GetUniformLocation(ctx Functions, prog Program, name string) Uniform {
	loc := ctx.GetUniformLocation(prog, name)
	if !loc.valid() {
		panic(fmt.Errorf("uniform %s not found", name))
	}
	return loc
}

func createShader(ctx Functions, typ Enum, src string) (Shader, error) {
	sh := ctx.CreateShader(typ)
	if !sh.valid() {
		return Shader{}, errors.New("glCreateShader failed")
	}
	ctx.ShaderSource(sh, src)
	ctx.CompileShader(sh)
	if ctx.GetShaderi(sh, COMPILE_STATUS) == 0 {
		log := ctx.GetShaderInfoLog(sh)
		ctx.DeleteShader(sh)
		return Shader{}, fmt.Errorf("shader compilation failed: %s", strings.TrimSpace(log))
	}
	return sh, nil
}

func ParseGLVersion(glVer string) ([2]int, error) {
	var ver [2]int
	if _, err := fmt.Sscanf(glVer, "OpenGL ES %d.%d", &ver[0], &ver[1]); err == nil {
		return ver, nil
	} else if _, err := fmt.Sscanf(glVer, "WebGL %d.%d", &ver[0], &ver[1]); err == nil {
		// WebGL major version v corresponds to OpenGL ES version v + 1
		ver[0]++
		return ver, nil
	} else if _, err := fmt.Sscanf(glVer, "%d.%d", &ver[0], &ver[1]); err == nil {
		return ver, nil
	}
	return ver, fmt.Errorf("failed to parse OpenGL ES version (%s)", glVer)
}
