#pragma once

#include <GL/glew.h>

#define TEST_OPENGL_ERROR()\
  do {\
    GLenum err = glGetError();\
    if (err != GL_NO_ERROR) {\
      std::cerr << "ERROR! At " << __FILE__ << ":" << __LINE__ << std::endl;\
      raise(SIGTRAP);\
    }\
  } while(0)

#define ASSERT(x) if (!(x)) raise(SIGTRAP);
