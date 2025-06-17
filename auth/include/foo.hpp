#pragma once

#include <iostream>

namespace auth::sample
{

auto printer() noexcept -> void
{
  std::cout << "Hello, world!" << std::endl;
}

} /* namespace auth::sample */