#include <functional>  // Для std::less, std::greater

template <typename T, typename Compare = std::less<T>>
class BinaryTree {
public:
    void push(const T& value);

    bool pop(const T& value);

    
    bool search(const T& value) const;

 
};