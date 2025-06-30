#include <stdexcept>
#include <deque>  // или другая базовая структура

template <typename T>
class Queue {
public:
  
    bool empty() const;

    void push(const T& value);

    void push(T&& value);

    T pop();

};