# cmake_minimum_required(VERSION <specify CMake version here>)
project(demo1)

set(CMAKE_CXX_STANDARD 14)

add_executable(demo1 main.cpp test1.cpp main.cpp test2.cpp test3.cpp)