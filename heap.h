#include <stdexcept>

template <typename T, typename Compare = std::less<T>>
class Heap {
public:
    // Проверяет, пуста ли куча
    bool empty() const;

    // Добавляет элемент в кучу
    void push(const T& value);

    T pop();


};