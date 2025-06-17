cmake_minimum_required(VERSION 3.15)
project(AUTH-SERVICE VERSION 1.0.0
                     LANGUAGES CXX)
add_executable(${PROJECT_NAME} dev/main.cpp)
target_compile_options(${PROJECT_NAME} PUBLIC -std=c++20)
target_include_directories(${PROJECT_NAME} PUBLIC include)