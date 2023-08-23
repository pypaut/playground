#include <GL/glew.h>
#include <GLFW/glfw3.h>

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

int main(void) {
    GLFWwindow* window;

    /* Initialize GLFW */
    if (!glfwInit()) {
        std::cerr << "glfwInit() error!" << std::endl;
        return -1;
    }

    /* Create a windowed mode window and its OpenGL context */
    window = glfwCreateWindow(640, 480, "Hello World", NULL, NULL);
    TEST_OPENGL_ERROR();
    if (!window)
    {
        glfwTerminate();
        return -1;
    }

    /* Make the window's context current */
    glfwMakeContextCurrent(window); TEST_OPENGL_ERROR();

    /* Linked to the refresh rate */
    glfwSwapInterval(1); TEST_OPENGL_ERROR();

    /* Initialize GLEW */
    if (glewInit() != GLEW_OK) {
        std::cerr << "glewInit() error!" << std::endl;
    }

    /* Print version */
    std::cout << "OpenGL version: " << glGetString(GL_VERSION) << std::endl;

    /* Data buffer */
    float positions[8] = {
        -0.5f, -0.5f,
         0.5f,  0.5f,
         0.5f, -0.5f,
        -0.5f,  0.5f,
    };

    unsigned int indices[] = {
        0, 1, 2,
        0, 1, 3,
    };

    /* Setup data buffer to send input to shaders */
    unsigned int buffer;
    glGenBuffers(1, &buffer); TEST_OPENGL_ERROR();
    glBindBuffer(GL_ARRAY_BUFFER, buffer); TEST_OPENGL_ERROR();
    glBufferData(GL_ARRAY_BUFFER, 8 * sizeof(float), positions, GL_STATIC_DRAW);
    TEST_OPENGL_ERROR();

    glEnableVertexAttribArray(0); TEST_OPENGL_ERROR();
    glVertexAttribPointer(0, 2, GL_FLOAT, GL_FALSE, sizeof(float) * 2, 0);
    TEST_OPENGL_ERROR();

    unsigned int ibo;
    glGenBuffers(1, &ibo); TEST_OPENGL_ERROR();
    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, ibo); TEST_OPENGL_ERROR();
    glBufferData(
        GL_ELEMENT_ARRAY_BUFFER,
        6 * sizeof(unsigned int), 
        indices,
        GL_STATIC_DRAW
    ); TEST_OPENGL_ERROR();

    ShaderProgramSource source = ParseShaders("res/shaders/basic.shader");
    unsigned int program = CreateProgram(source.VertexSource, source.FragmentSource);
    glUseProgram(program); TEST_OPENGL_ERROR();

    int location = glGetUniformLocation(program, "u_Color"); TEST_OPENGL_ERROR();
    ASSERT(location != -1);
    glUniform4f(location, 0.2f, 0.3f, 0.3f, 0.5f); TEST_OPENGL_ERROR();

    float r = 0.0f;
    float g = 0.5f;
    float b = 1.0f;
    float r_increment = 0.03f;
    float g_increment = 0.05f;
    float b_increment = 0.07f;

    /* Loop until the user closes the window */
    while (!glfwWindowShouldClose(window))
    {
        /* Render here */
        glClear(GL_COLOR_BUFFER_BIT); TEST_OPENGL_ERROR();

        glUniform4f(location, r, g, b, 0.5f); TEST_OPENGL_ERROR();
        glDrawElements(GL_TRIANGLES, 6, GL_UNSIGNED_INT, nullptr);
        TEST_OPENGL_ERROR();

        if (r < 0.0f || r > 1.0f)
            r_increment *= -1;
        r += r_increment;

        if (g < 0.0f || g > 1.0f)
            g_increment *= -1;
        g += g_increment;

        if (b < 0.0f || b > 1.0f)
            b_increment *= -1;
        b += b_increment;

        /* Swap front and back buffers */
        glfwSwapBuffers(window); TEST_OPENGL_ERROR();

        /* Poll for and process events */
        glfwPollEvents(); TEST_OPENGL_ERROR();
    }

    glDeleteProgram(program); TEST_OPENGL_ERROR();

    glfwTerminate(); TEST_OPENGL_ERROR();
    return 0;
}
