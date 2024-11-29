#include <GL/glew.h>
#include <SDL2/SDL.h>
#include <SDL2/SDL_opengl.h>

#include <fstream>
#include <iostream>
#include <signal.h>
#include <sstream>
#include <string>


#define TEST_OPENGL_ERROR()\
  do {\
    GLenum err = glGetError();\
    if (err != GL_NO_ERROR) {\
      std::cerr << "ERROR! At " << __FILE__ << ":" << __LINE__ << std::endl;\
      raise(SIGTRAP);\
    }\
  } while(0)

#define ASSERT(x) if (!(x)) raise(SIGTRAP);

#define WIN_W 1920
#define WIN_H 1080


void log_error(std::string func_name) {
    std::cerr << "Error in " << func_name << ": " << SDL_GetError() << std::endl;
}

struct ShaderProgramSource {
    std::string VertexSource;
    std::string FragmentSource;
};

static ShaderProgramSource ParseShaders(const std::string& filepath) {
    std::ifstream stream(filepath);

    enum class ShaderType {
        NONE = -1,
        VERTEX = 0,
        FRAGMENT = 1,
    };

    std::string line;
    std::stringstream ss[2];
    ShaderType type = ShaderType::NONE;

    while (getline(stream, line)) {
        if (line.find("#shader") != std::string::npos) {
            if (line.find("vertex") != std::string::npos) {
                type = ShaderType::VERTEX;
            } else if (line.find("fragment") != std::string::npos) {
                type = ShaderType::FRAGMENT;
            }
        } else {
            ss[(int)type] << line << '\n';
        }
    }

    return { ss[0].str(), ss[1].str() };
}

static unsigned int CompileShader(unsigned int type, const std::string& source) {
    unsigned int id = glCreateShader(type);
    const char *src = source.c_str();
    glShaderSource(id, 1, &src, nullptr); TEST_OPENGL_ERROR();
    glCompileShader(id); TEST_OPENGL_ERROR();

    int result;
    glGetShaderiv(id, GL_COMPILE_STATUS, &result); TEST_OPENGL_ERROR();
    if (result == GL_FALSE) {
        int length;
        glGetShaderiv(id, GL_INFO_LOG_LENGTH, &length); TEST_OPENGL_ERROR();
        char *message = (char*)alloca(length * sizeof(char));
        glGetShaderInfoLog(id, length, &length, message); TEST_OPENGL_ERROR();
        std::cout
            << "Failed to compile "
            << (type == GL_VERTEX_SHADER ? "vertex" : "fragment")
            << " shader!" << std::endl;
        std::cout << message << std::endl;

        glDeleteShader(id); TEST_OPENGL_ERROR();
        return 0;
    }

    return id;
}

static int CreateProgram(
        const std::string& vertexShader,
        const std::string& fragmentShader
) {
    unsigned int program = glCreateProgram(); TEST_OPENGL_ERROR();
    unsigned int vs = CompileShader(GL_VERTEX_SHADER, vertexShader);
    unsigned int fs = CompileShader(GL_FRAGMENT_SHADER, fragmentShader);

    glAttachShader(program, vs);    TEST_OPENGL_ERROR();
    glAttachShader(program, fs);    TEST_OPENGL_ERROR();
    glLinkProgram(program);         TEST_OPENGL_ERROR();
    glValidateProgram(program);     TEST_OPENGL_ERROR();

    glDeleteShader(vs); TEST_OPENGL_ERROR();
    glDeleteShader(fs); TEST_OPENGL_ERROR();

    return program;
}

int main() {
    SDL_Window *window = SDL_CreateWindow(
            "SDL & OpenGL!",
            SDL_WINDOWPOS_CENTERED,
            SDL_WINDOWPOS_CENTERED,
            WIN_W,
            WIN_H,
            SDL_WINDOW_OPENGL
    );
    if (!window) {
        log_error("SDL_CreateWindow");
    }

    SDL_GLContext context = SDL_GL_CreateContext(window);
    context = context;

    /* Initialize GLEW: the window must initialized by now, or SEGFAULT */
    if (glewInit() != GLEW_OK) {
        std::cerr << "glewInit() error!" << std::endl;
    }

    /* Print version */
    std::cout << "OpenGL version: " << glGetString(GL_VERSION) << std::endl;

    /* Setup data buffers to send input to shaders */

    /* Data buffers */

    /* Positions are x, y coordinates */
    float positions[16] = {
        // Player 1
        -0.90f, -0.2f, // x, y
        -0.85f,  0.2f,
        -0.90f,  0.2f,
        -0.85f, -0.2f,

        // Player 2
        0.90f, -0.2f, // x, y
        0.85f,  0.2f,
        0.90f,  0.2f,
        0.85f, -0.2f,
    };
    size_t nb_positions = 16;

    /* Indices will select which set of positions to draw triangles to */
    /* Here, we draw 2 triangles with 2 common vertices */
    unsigned int indices[] = {
        // Player 1
        0, 1, 2, // pos[0,1], pos[2,3], and pos[4,5] are the 3 points
        0, 1, 3, // pos[0,1], pos[2,3], and pos[6,7] are the 3 points

        // Player 2
        4, 5, 6,
        4, 5, 7,

    };
    size_t nb_indices = 12;

    unsigned int vao;
    glGenVertexArrays(1, &vao); TEST_OPENGL_ERROR();
    glBindVertexArray(vao); TEST_OPENGL_ERROR();

    unsigned int buffer;
    glGenBuffers(1, &buffer); TEST_OPENGL_ERROR();
    glBindBuffer(GL_ARRAY_BUFFER, buffer); TEST_OPENGL_ERROR();
    glBufferData(GL_ARRAY_BUFFER, nb_positions * sizeof(float), positions, GL_STATIC_DRAW);
    TEST_OPENGL_ERROR();

    glEnableVertexAttribArray(0); TEST_OPENGL_ERROR();
    glVertexAttribPointer(0, 2, GL_FLOAT, GL_FALSE, sizeof(float) * 2, 0);
    TEST_OPENGL_ERROR();

    unsigned int ibo;
    glGenBuffers(1, &ibo); TEST_OPENGL_ERROR();
    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, ibo); TEST_OPENGL_ERROR();
    glBufferData(
        GL_ELEMENT_ARRAY_BUFFER,
        nb_indices * sizeof(unsigned int),
        indices,
        GL_STATIC_DRAW
    ); TEST_OPENGL_ERROR();

    ShaderProgramSource source = ParseShaders("res/shaders/basic.shader");
    unsigned int program = CreateProgram(source.VertexSource, source.FragmentSource);
    glUseProgram(program); TEST_OPENGL_ERROR();

    int location = glGetUniformLocation(program, "u_Color"); TEST_OPENGL_ERROR();
    ASSERT(location != -1);
    glUniform4f(location, 0.2f, 0.3f, 0.3f, 0.5f); TEST_OPENGL_ERROR();

    glBindVertexArray(0); TEST_OPENGL_ERROR();
    glUseProgram(0); TEST_OPENGL_ERROR();
    glBindBuffer(GL_ARRAY_BUFFER, 0); TEST_OPENGL_ERROR();
    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, 0); TEST_OPENGL_ERROR();

    float r = 1.0f;
    float g = 1.0f;
    float b = 1.0f;

    /* Main loop */
    while (true) {
        /* Check for quit event */
        SDL_Event event;
        bool should_quit = false;
        while (SDL_PollEvent(&event)) {
            if (event.type == SDL_QUIT) {
                should_quit = true;
                break;
            }
        }
        if (should_quit) {
            break;
        }

        /* Render here */
        glClear(GL_COLOR_BUFFER_BIT); TEST_OPENGL_ERROR();
        glUseProgram(program); TEST_OPENGL_ERROR();
        glUniform4f(location, r, g, b, 0.5f); TEST_OPENGL_ERROR();

        glBindVertexArray(vao); TEST_OPENGL_ERROR();
        glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, ibo); TEST_OPENGL_ERROR();

        glDrawElements(GL_TRIANGLES, nb_indices, GL_UNSIGNED_INT, nullptr);
        TEST_OPENGL_ERROR();

        glViewport(0, 0, WIN_W, WIN_H); TEST_OPENGL_ERROR();
        glClearColor(0.f, 0.f, 0.f, 0.f); TEST_OPENGL_ERROR();

        SDL_GL_SwapWindow(window); TEST_OPENGL_ERROR();
    }

    SDL_Quit();
    return 0;
}
